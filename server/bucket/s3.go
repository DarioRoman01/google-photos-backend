package bucket

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// S3BucketRepository is an implementation of the BucketRepository interface that stores and retrieves images in an S3 bucket.
type S3BucketRepository struct {
	client *session.Session // client is the AWS SDK client.
}

// NewS3BucketRepository creates a new S3BucketRepository.
func NewS3BucketRepository() (*S3BucketRepository, error) {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("KEY_ID"), os.Getenv("ACCESS_KEY"), ""),
	})

	if err != nil {
		return nil, err
	}

	return &S3BucketRepository{client: session}, nil
}

// Delete deletes an image from the bucket.
func (r *S3BucketRepository) Delete(username, filename, folder string) error {
	batcher := s3manager.NewBatchDelete(r.client)

	// verify if the path is not empty
	var path string
	if folder == "" {
		path = username + "/" + "default" + "/" + filename
	} else {
		path = username + "/" + folder + "/" + filename
	}

	objs := []s3manager.BatchDeleteObject{{
		Object: &s3.DeleteObjectInput{
			Bucket: aws.String(os.Getenv("S3_BUCKET")),
			Key:    aws.String(path),
		},
	}}

	return batcher.Delete(aws.BackgroundContext(), &s3manager.DeleteObjectsIterator{Objects: objs})
}

// Upload uploads an image to the bucket and returns the URL of the image.
func (repository *S3BucketRepository) Upload(file multipart.File, fileHeader *multipart.FileHeader, username, folder string) (string, error) {
	uploader := s3manager.NewUploader(repository.client)

	// verify if the path is not empty
	var path string
	if folder == "" {
		path = fmt.Sprintf("%s/default/%s", username, fileHeader.Filename)
	} else {
		path = fmt.Sprintf("%s/%s/%s", username, folder, fileHeader.Filename)
	}

	r, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(path),
		Body:   file,
	})

	return r.Location, err
}
