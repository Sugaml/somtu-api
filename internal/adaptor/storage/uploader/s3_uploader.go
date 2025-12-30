package uploader

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Uploader struct {
	Client *s3.Client
	Bucket string
}

func NewS3Uploader() (*S3Uploader, error) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET_NAME") // Use a more semantic name than FS_LOCATION for S3

	if accessKey == "" || secretKey == "" || region == "" || bucket == "" {
		return nil, fmt.Errorf("missing one or more required AWS environment variables")
	}

	// Manually configure credentials
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	return &S3Uploader{
		Client: client,
		Bucket: bucket,
	}, nil
}

func (u *S3Uploader) UploadFile(file *FileDetails) (string, error) {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file.File)
	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("%s/%s/%s", file.FileType, file.EntityID, file.FileHeader.Filename)
	_, err = u.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(u.Bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String(file.FileHeader.Header.Get("Content-Type")),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", u.Bucket, key), nil
}

func (u *S3Uploader) GetFileURL(fileType FileType, entityID, fileName string) (string, error) {
	key := fmt.Sprintf("%s/%s/%s", fileType, entityID, fileName)
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", u.Bucket, key), nil
}
