package uploader

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewUploader(file io.Reader, client *s3.Client, bucket, objectkey string) {
	// _, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
	// 	Bucket: &bucket,
	// 	Key:    &objectKey,
	// 	Body:   file,
	// })

	uploader := manager.NewUploader(client)

	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &objectkey,
		Body:   file,
	})

	if err != nil {
		fmt.Printf("There is an error while uploading %v", err)
		return
	}
}
