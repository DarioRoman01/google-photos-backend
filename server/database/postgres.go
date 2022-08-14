package database

import (
	"database/sql"
	"log"

	"github.com/DarioRoman01/photos/models"
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
			FOREIGN KEY (user_id) REFERENCES users (id)
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
			FOREIGN KEY (folder_id) REFERENCES folders(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
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
	row := r.db.QueryRow("SELECT id, name, url, user_id, folder_id FROM images WHERE id = $1", id)
	image := &models.Image{}
	err := row.Scan(&image.ID, &image.Name, &image.URL, &image.UserID, &image.FolderID)
	if err != nil {
		return nil, err
	}

	return image, nil
}

// GetImages returns all images for the given user.
func (r *PostgresRepository) GetImages(userID string) ([]*models.Image, error) {
	rows, err := r.db.Query("SELECT id, name, url, user_id, folder_id FROM images WHERE user_id = $1", userID)
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

// GetImagesByFolder returns all images for the given folder.
func (r *PostgresRepository) GetImagesByFolder(folderID string) ([]*models.Image, error) {
	rows, err := r.db.Query("SELECT id, name, url, user_id, folder_id FROM images WHERE folder_id = $1", folderID)
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
	row := r.db.QueryRow("SELECT id, name, user_id FROM folders WHERE id = $1", id)
	folder := &models.Folder{}
	err := row.Scan(&folder.ID, &folder.Name, &folder.UserID)
	if err != nil {
		return nil, err
	}

	return folder, nil
}

// GetFolders returns all folders for the given user.
func (r *PostgresRepository) GetFolders(userID string) ([]*models.Folder, error) {
	rows, err := r.db.Query("SELECT id, name, user_id FROM folders WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	folders := []*models.Folder{}
	for rows.Next() {
		folder := &models.Folder{}
		err := rows.Scan(&folder.ID, &folder.Name, &folder.UserID)
		if err != nil {
			return nil, err
		}

		folders = append(folders, folder)
	}

	return folders, nil
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
