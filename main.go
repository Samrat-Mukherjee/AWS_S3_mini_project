package main

import (
	"fmt"
	"log"
	"os"
	"s3-aws/client"
	getimage "s3-aws/get-image"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//s3Bucket := os.Getenv("S3_BUCKET")
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	bucket := "samrat8-bucket"
	filepath := "image/1st-post.png"
	objectKey := "1st post.png"

	client := client.NewClient(accessKey, secretKey)

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("error in file path: %v", err)
	}
	defer file.Close()

	//uploader.NewUploader(file, client, bucket, objectKey)

	getimage.GeneratePresignedURL(client, bucket, objectKey)

	fmt.Println("Uploaded successfully...")

}
