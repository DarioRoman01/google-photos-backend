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
