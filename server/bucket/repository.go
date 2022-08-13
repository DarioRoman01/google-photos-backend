package bucket

import (
	"mime/multipart"
)

// BucketRepository is an interface for a repository that stores and retrieves images.
type BucketRepository interface {
	Delete(username, filename, folder string) error                     // Delete deletes an image from the bucket.
	Upload(file multipart.File, fileHeader *multipart.FileHeader) error // Upload uploads an image to the bucket.
}

var bucketRepository BucketRepository

func SetBucketRepository(repository BucketRepository) {
	bucketRepository = repository
}

func Delete(username, filename, folder string) error {
	return bucketRepository.Delete(username, filename, folder)
}

func Upload(file multipart.File, fileHeader *multipart.FileHeader, username, folder string) error {
	return bucketRepository.Upload(file, fileHeader)
}
