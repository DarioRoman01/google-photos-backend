package bucket

import (
	"bytes"
)

// BucketRepository is an interface for a repository that stores and retrieves images.
type BucketRepository interface {
	// Delete deletes an image from the bucket.
	Delete(key string) error
	// Upload uploads an image to the bucket.
	Upload(file *bytes.Buffer, fileName, username, folder string) (string, error)
	// MoveFile copy a file from one folder to another and deletes the original file and returns the new file's URL.
	MoveFile(username, oldPath, newPath, filename string) (string, error)
}

var bucketRepository BucketRepository

func SetBucketRepository(repository BucketRepository) {
	bucketRepository = repository
}

func Delete(key string) error {
	return bucketRepository.Delete(key)
}

func Upload(file *bytes.Buffer, fileName, username, folder string) (string, error) {
	return bucketRepository.Upload(file, fileName, username, folder)
}

func MoveFile(username, oldPath, newPath, filename string) (string, error) {
	return bucketRepository.MoveFile(username, oldPath, newPath, filename)
}
