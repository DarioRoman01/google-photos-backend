package database

import "github.com/DarioRoman01/photos/models"

// DatabaseRepository is an interface that defines the methods that a database must implement.
type DatabaseRepository interface {
	// InsertImage inserts an image into the database.
	InsertImage(image *models.Image) error
	// InsertFolder inserts a folder into the database.
	InsertFolder(folder *models.Folder) error
	// InsertUser inserts a user into the database.
	InsertUser(user *models.User) error
	// UpdateUserStatus updates the status of a user to verified.
	UpdateUserStatus(id string) error
	// GetImage retrieves an image from the database.
	GetImage(id string) (*models.Image, error)
	// GetFolder retrieves a folder from the database.
	GetFolder(id string) (*models.Folder, error)
	// GetImages retrieves all images from the database  from the given user.
	GetImages(userID, cursor string, limit int) ([]*models.Image, error, bool)
	// GetImagesByFolder retrieves all images from the database  from the given folder.
	GetImagesByFolder(folderID string) ([]*models.Image, error)
	// GetFolders retrieves all folders from the database from the given user.
	GetFolders(userID string) ([]*models.Folder, error)
	// CheckFolder checks if the folder exists if not exists its create a new folder with the given name
	CheckFolder(userID, foldeName string) (string, error)
	// GetUserByUsername retrieves a user from the database by username.
	GetUserByUsername(username string) (*models.User, error)
	// GetUserByEmail retrieves a user from the database by email.
	GetUserByEmail(email string) (*models.User, error)
	// GetUserByID retrieves a user from the database by ID.
	GetUserByID(id string) (*models.User, error)
	// DeleteImage deletes an image from the database.
	DeleteImage(id string) error
	// DeleteFolder deletes a folder from the database.
	DeleteFolder(id string) error
	// DeleteUser deletes a user from the database.
	DeleteUser(id string) error
	// UpdateImage updates the image with the given id only the folder and the url can be chage.
	UpdateImage(req *models.MoveFileRequest, userId string) error
}

var databaseRepository DatabaseRepository

func SetDatabaseRepository(repository DatabaseRepository) {
	databaseRepository = repository
}

func InsertImage(image *models.Image) error {
	return databaseRepository.InsertImage(image)
}

func InsertFolder(folder *models.Folder) error {
	return databaseRepository.InsertFolder(folder)
}

func InsertUser(user *models.User) error {
	return databaseRepository.InsertUser(user)
}

func GetImage(id string) (*models.Image, error) {
	return databaseRepository.GetImage(id)
}

func GetFolder(id string) (*models.Folder, error) {
	return databaseRepository.GetFolder(id)
}

func GetImagesGetImages(userID, cursor string, limit int) ([]*models.Image, error, bool) {
	return databaseRepository.GetImages(userID, cursor, limit)
}

func GetFolders(userID string) ([]*models.Folder, error) {
	return databaseRepository.GetFolders(userID)
}

func GetImagesByFolder(folderID string) ([]*models.Image, error) {
	return databaseRepository.GetImagesByFolder(folderID)
}

func UpdateUserStatus(id string) error {
	return databaseRepository.UpdateUserStatus(id)
}

func UpdateImage(req *models.MoveFileRequest, userId string) error {
	return databaseRepository.UpdateImage(req, userId)
}

func CheckFolder(userID, foldeName string) (string, error) {
	return databaseRepository.CheckFolder(userID, foldeName)
}

func GetUserByUsername(username string) (*models.User, error) {
	return databaseRepository.GetUserByUsername(username)
}

func GetUserByEmail(email string) (*models.User, error) {
	return databaseRepository.GetUserByEmail(email)
}

func GetUserByID(id string) (*models.User, error) {
	return databaseRepository.GetUserByID(id)
}

func DeleteImage(id string) error {
	return databaseRepository.DeleteImage(id)
}

func DeleteFolder(id string) error {
	return databaseRepository.DeleteFolder(id)
}

func DeleteUser(id string) error {
	return databaseRepository.DeleteUser(id)
}
