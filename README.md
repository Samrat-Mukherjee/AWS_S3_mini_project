# 🗂️ S3 (Simple Storage System) with Golang

This mini project demonstrates how to **upload files to AWS S3** and **generate pre-signed URLs** to securely access them — all using **Golang** and the **AWS SDK v2**.

---

## 📦 Key Features

- ✅ Connect securely to AWS S3 using SDK v2
- 📤 Upload files to S3 using `PutObject`
- 🔗 Generate pre-signed URLs (valid for temporary access)
- 🛡️ Follow production-level best practices

---

## 🔧 S3 Integration Overview

### 1. ✅ Create S3 Client

To communicate with AWS S3, create a client using the AWS SDK:

```go
cfg, err := config.LoadDefaultConfig(
	context.TODO(),
	config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(accessKey, secretKey, session),
	),
	config.WithRegion(region),
)

client := s3.NewFromConfig(cfg)

```


### 2. 📤 Upload File to S3
Use manager.NewUploader() to upload a file to your S3 bucket:

```go
uploader := manager.NewUploader(client)

_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
	Bucket: &bucket,
	Key:    &key,  // Object key (e.g., "uploads/file.pdf")
	Body:   file,  // The file (e.g., multipart.File)
})

```

### 3. 🔗 Generate Pre-signed URL
Generate a secure, temporary URL (e.g., valid for 15 minutes) to access the file without making the bucket public:

```go
presign := s3.NewPresignClient(client)

res, err := presign.PresignGetObject(context.TODO(), &s3.GetObjectInput{
	Bucket: &bucket,
	Key:    &key,
})
```



