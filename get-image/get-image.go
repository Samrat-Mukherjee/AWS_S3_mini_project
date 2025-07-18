package getimage

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GeneratePresignedURL(client *s3.Client, bucket, key string) {
	// ✅ Use s3.NewPresignClient (part of main s3 package)
	presignClient := s3.NewPresignClient(client)

	req, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}, s3.WithPresignExpires(15*time.Minute)) // expires in 15 mins

	if err != nil {
		fmt.Println("❌ Error generating presigned URL:", err)
		return
	}

	fmt.Println("🔗 Pre-signed URL:", req.URL)
}
