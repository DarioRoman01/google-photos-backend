package bucket

import (
	"bytes"
	"fmt"
	"log"
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
	svc := s3.New(r.client)
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(fmt.Sprintf("%s/%s/%s", username, folder, filename)),
	})

	return err
}

// Upload uploads an image to the bucket and returns the URL of the image.
func (repository *S3BucketRepository) Upload(file *bytes.Buffer, fileName, username, folder string) (string, error) {
	uploader := s3manager.NewUploader(repository.client)

	// verify if the path is not empty
	var path string
	if folder == "" {
		path = fmt.Sprintf("%s/default/%s", username, fileName)
	} else {
		path = fmt.Sprintf("%s/%s/%s", username, folder, fileName)
	}

	r, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("S3_BUCKET")),
		Key:         aws.String(path),
		Body:        file,
		ContentType: aws.String("image/jpeg"),
	})

	if err != nil {
		log.Printf("Error uploading file: %s", err)
		return "", err
	}

	return r.Location, nil
}

// MoveFile copy a file from one folder to another and deletes the original file and returns the new file's URL.
func (r *S3BucketRepository) MoveFile(username, oldPath, newPath, filename string) (string, error) {
	bucketName := os.Getenv("S3_BUCKET")
	svc := s3.New(r.client)
	opt, err := svc.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String(bucketName),
		CopySource: aws.String(fmt.Sprintf("%s/%s/%s/%s", bucketName, username, oldPath, filename)),
		Key:        aws.String(fmt.Sprintf("%s/%s/%s", username, newPath, filename)),
	})

	if err != nil {
		log.Printf("Error copying file: %s", err)
		return "", err
	}

	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(fmt.Sprintf("%s/%s/%s", username, oldPath, filename)),
	})

	if err != nil {
		log.Printf("Error deleting file: %s", err)
		return "", err
	}

	return opt.CopyObjectResult.GoString(), nil
}

func (r *S3BucketRepository) DeleteFolder(username, folderName string) error {
	svc := s3.New(r.client)
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(fmt.Sprintf("%s/%s", username, folderName)),
	})

	return err
}
