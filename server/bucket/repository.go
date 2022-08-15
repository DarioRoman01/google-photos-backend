package bucket

import (
	"bytes"
)

// BucketRepository is an interface for a repository that stores and retrieves images.
type BucketRepository interface {
	// Delete deletes an image from the bucket.
	Delete(username, filename, folder string) error
	// Upload uploads an image to the bucket.
	Upload(file *bytes.Buffer, fileName, username, folder string) (string, error)
	// MoveFile copy a file from one folder to another and deletes the original file and returns the new file's URL.
	MoveFile(username, oldPath, newPath, filename string) (string, error)
	// DeleteFolder deletes a folder from the bucket.
	DeleteFolder(username, folder string) error
}

var bucketRepository BucketRepository

func SetBucketRepository(repository BucketRepository) {
	bucketRepository = repository
}

func Delete(username, filename, folder string) error {
	return bucketRepository.Delete(username, filename, folder)
}

func Upload(file *bytes.Buffer, fileName, username, folder string) (string, error) {
	return bucketRepository.Upload(file, fileName, username, folder)
}

func MoveFile(username, oldPath, newPath, filename string) (string, error) {
	return bucketRepository.MoveFile(username, oldPath, newPath, filename)
}

func DeleteFolder(username, folder string) error {
	return bucketRepository.DeleteFolder(username, folder)
}
