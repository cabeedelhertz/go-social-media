package s3

import (
	"bytes"
	"context"
	"social"
	"social/internal/store"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var _ store.ImageStore = (*S3Client)(nil)

type S3Client struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	cfg           *social.Config
}

func New(cfg *social.Config, awsConfig aws.Config) *S3Client {
	s3Client := s3.NewFromConfig(awsConfig)
	presignClient := s3.NewPresignClient(s3Client)
	return &S3Client{
		client:        s3Client,
		presignClient: presignClient,
		cfg:           cfg,
	}
}

func (c *S3Client) UploadImage(ctx context.Context, key string, content []byte) error {
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.cfg.BucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(content),
	})
	return err
}

func (c *S3Client) GetImageURL(ctx context.Context, key string) (string, error) {
	presignInput := &s3.GetObjectInput{
		Bucket: aws.String(c.cfg.BucketName),
		Key:    aws.String(key),
	}
	presignedResult, err := c.presignClient.PresignGetObject(ctx, presignInput, func(po *s3.PresignOptions) {
		po.Expires = 15 * time.Minute
	})
	if err != nil {
		return "", err
	}

	return presignedResult.URL, nil
}

func (c *S3Client) DeleteImage(ctx context.Context, path string) error {
	return nil
}
