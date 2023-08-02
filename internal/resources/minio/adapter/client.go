package adapter

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"github.com/thuongtruong1009/zoomer/configs"
	"log"
)

var (
	useSSL      = false
	contentType = "application/octet-stream"
	// contentType = "image/png"
	location = "us-east-1"
)

func RegisterMinioClient(cfg *configs.Configuration) (*minio.Client, error) {
	minioClient, err := minio.New(cfg.MinIOEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinIOAccessKey, cfg.MinIOSecretKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return minioClient, nil
}

func SetPermission(client *minio.Client, bucketName string) error {
	policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::` + bucketName + `/*"],"Sid": ""}]}`

	err := client.SetBucketPolicy(context.Background(), bucketName, policy)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func CreateBucket(client *minio.Client, bucketName string) error {
	ctx := context.Background()
	err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			logrus.Infof("We already own %s\n", bucketName)
		} else {
			logrus.Infof("Failed to create bucket %s\n", bucketName)
		}
	} else {
		logrus.Infof("Successfully created %s\n", bucketName)
	}
	return err
}
