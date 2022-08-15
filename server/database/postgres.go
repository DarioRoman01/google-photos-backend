package database

import (
	"database/sql"
	"log"

	"github.com/DarioRoman01/photos/models"
	"github.com/DarioRoman01/photos/utils"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// PostgresRepository is a repository that uses a Postgres database.
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository returns a new PostgresRepository.
func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	repo := &PostgresRepository{db: db}
	repo.Migrate()
	return repo, nil
}

// Migrate the database to the latest schema version.
func (r *PostgresRepository) Migrate() {
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			is_verified BOOLEAN NOT NULL DEFAULT FALSE
		);
	`)

	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

	_, err = r.db.Exec(`
		CREATE TABLE IF NOT EXISTS folders (
			id VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
		);
	`)

	if err != nil {
		log.Fatalf("Error creating folders table: %v", err)
	}

	_, err = r.db.Exec(`
		CREATE TABLE IF NOT EXISTS images (
			id VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			folder_id VARCHAR(36) NOT NULL,
			user_id VARCHAR(36) NOT NULL,
			url VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			FOREIGN KEY (folder_id) REFERENCES folders(id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);
	`)

	if err != nil {
		log.Fatalf("Error creating images table: %v", err)
	}
}

// InsertImage inserts an image into the database.
func (r *PostgresRepository) InsertImage(image *models.Image) error {
	_, err := r.db.Exec(
		"INSERT INTO images (id, name, url, user_id, folder_id) VALUES ($1, $2, $3, $4, $5)",
		image.ID, image.Name, image.URL, image.UserID, image.FolderID,
	)

	return err
}

func (r *PostgresRepository) UpdateUserStatus(id string) error {
	_, err := r.db.Exec("UPDATE users SET is_verified = $1 WHERE id = $2", true, id)
	return err
}

// InsertFolder inserts a folder into the database.
func (r *PostgresRepository) InsertFolder(folder *models.Folder) error {
	_, err := r.db.Exec(
		"INSERT INTO folders (id, name, user_id) VALUES ($1, $2, $3)",
		folder.ID, folder.Name, folder.UserID,
	)

	return err
}

// CheckFolder checks if the folder exists if not exists its create a new folder with the given name
func (r *PostgresRepository) CheckFolder(userID, foldeName string) (string, error) {
	row := r.db.QueryRow("SELECT id FROM folders WHERE name = $1 AND user_id = $2", foldeName, userID)
	var id string
	err := row.Scan(&id)
	if err != nil {
		id = uuid.NewString()
		_, err := r.db.Exec(
			"INSERT INTO folders (id, name, user_id) VALUES ($1, $2, $3)",
			id, foldeName, userID,
		)

		if err != nil {
			return "", err
		}
	}

	return id, nil
}

// InsertUser inserts a user into the database.
func (r *PostgresRepository) InsertUser(user *models.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4)",
		user.ID, user.Username, user.Email, user.Password,
	)

	return err
}

// GetUserByEmail returns a user with the given email.
func (r *PostgresRepository) GetUserByEmail(email string) (*models.User, error) {
	row := r.db.QueryRow("SELECT id, username, email, password, is_verified FROM users WHERE email = $1", email)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.IsVerified)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID returns a user with the given id.
func (r *PostgresRepository) GetUserByID(id string) (*models.User, error) {
	row := r.db.QueryRow("SELECT id, username, email, password, is_verified FROM users WHERE id = $1", id)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.IsVerified)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByUsername returns a user with the given username.
func (r *PostgresRepository) GetUserByUsername(username string) (*models.User, error) {
	row := r.db.QueryRow("SELECT id, username, email, password, is_verified FROM users WHERE username = $1", username)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.IsVerified)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetImage returns an image with the given id.
func (r *PostgresRepository) GetImage(id string) (*models.Image, error) {
	row := r.db.QueryRow("SELECT id, name, url, user_id, folder_id, created_at FROM images WHERE id = $1", id)
	image := &models.Image{}
	err := row.Scan(&image.ID, &image.Name, &image.URL, &image.UserID, &image.FolderID, &image.CreatedAt)
	if err != nil {
		return nil, err
	}

	return image, nil
}

// GetImages returns all images for the given user.
func (r *PostgresRepository) GetImages(userID, cursor string, limit int) ([]*models.Image, error, bool) {
	if limit > 50 || limit < 1 {
		limit = 50
	}

	var rows *sql.Rows
	var err error
	if cursor == "" {
		rows, err = r.db.Query(`
			SELECT id, name, url, user_id, folder_id, created_at FROM images WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2
		`, userID, limit)

	} else {
		rows, err = r.db.Query(`
		"SELECT id, name, url, user_id, folder_id, created_at FROM images WHERE user_id = $1 AND created_at < $2 ORDER BY created_at DESC LIMIT $3"
		`, userID, cursor, limit)
	}

	if err != nil {
		return nil, err, false
	}

	defer rows.Close()
	images := []*models.Image{}
	for rows.Next() {
		image := &models.Image{}
		err := rows.Scan(&image.ID, &image.Name, &image.URL, &image.UserID, &image.FolderID, &image.CreatedAt)
		if err != nil {
			return nil, err, false
		}

		images = append(images, image)
	}

	if len(images) == 0 {
		return nil, nil, false
	}

	if len(images) == limit {
		return images, nil, true
	}

	return images, nil, false
}

// GetImagesByFolder returns all images for the given folder.
func (r *PostgresRepository) GetImagesByFolder(folderID string) ([]*models.Image, error) {
	rows, err := r.db.Query("SELECT id, name, url, user_id, folder_id, created_at FROM images WHERE folder_id = $1", folderID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	images := []*models.Image{}
	for rows.Next() {
		image := &models.Image{}
		err := rows.Scan(&image.ID, &image.Name, &image.URL, &image.UserID, &image.FolderID)
		if err != nil {
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}

// GetFolder returns a folder with the given id.
func (r *PostgresRepository) GetFolder(id string) (*models.Folder, error) {
	row := r.db.QueryRow("SELECT id, name, user_id, created_at FROM folders WHERE id = $1", id)
	folder := &models.Folder{}
	err := row.Scan(&folder.ID, &folder.Name, &folder.UserID, &folder.CreatedAt)
	if err != nil {
		return nil, err
	}

	return folder, nil
}

// GetFolders returns all folders for the given user.
func (r *PostgresRepository) GetFolders(userID string) ([]*models.Folder, error) {
	rows, err := r.db.Query("SELECT id, name, user_id, created_at FROM folders WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	folders := []*models.Folder{}
	for rows.Next() {
		folder := &models.Folder{}
		err := rows.Scan(&folder.ID, &folder.Name, &folder.UserID, &folder.CreatedAt)
		if err != nil {
			return nil, err
		}

		folders = append(folders, folder)
	}

	return folders, nil
}

// GetImageURL returns the url for the given image id.
func (r *PostgresRepository) getImageUrl(id string) (string, error) {
	row := r.db.QueryRow("SELECT url FROM images WHERE id = $1", id)
	var url string
	err := row.Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

// UpdateImage updates the image with the given id only the folder and the url can be chage.
func (r *PostgresRepository) UpdateImage(req *models.MoveFileRequest, userId string) error {
	folderId, err := r.CheckFolder(userId, req.NewFolderName)
	if err != nil {
		return err
	}

	imgUrl, err := r.getImageUrl(req.FileID)
	if err != nil {
		return err
	}

	newUrl := utils.ChangeUrlPath(imgUrl, req.NewFolderName)
	_, err = r.db.Exec("UPDATE images SET folder_id = $1, url = $2 WHERE id = $3", folderId, newUrl, req.FileID)
	return err
}

// DeleteFolder deletes a folder with the given id.
func (r *PostgresRepository) DeleteFolder(id string) error {
	_, err := r.db.Exec("DELETE FROM folders WHERE id = $1", id)
	return err
}

// DeleteImage deletes an image with the given id.
func (r *PostgresRepository) DeleteImage(id string) error {
	_, err := r.db.Exec("DELETE FROM images WHERE id = $1", id)
	return err
}

// DeleteUser deletes a user with the given id.
func (r *PostgresRepository) DeleteUser(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
