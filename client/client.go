package client

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewClient(accessKey, secretKey string) *s3.Client {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		), config.WithRegion("ap-south-1"))
	if err != nil {
		log.Fatalf("Unable to load SDK config: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	return client
}
