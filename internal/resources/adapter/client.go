package adapter

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"zoomer/configs"
)

var (
	endpoint    = configs.NewConfig().MinIOEndpoint //minio.example.com:9000
	accessKey   = configs.NewConfig().MinIOAccessKey
	secretKey   = configs.NewConfig().MinIOSecretKey
	useSSL      = false
	contentType = "application/octet-stream" // "image/png"
	location    = "us-east-1"
)

func MinioClient() (c *minio.Client, err error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}
	return minioClient, err
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
		if errBucketExists != nil {
			logrus.Errorf("[UploadImage] check bucket exists error: %s", err)
			return err
		}
		if !exists {
			logrus.Errorf("[UploadImage] make bucket error: %s", err)
			return err
		}
	}
	return err
}

func UploadData(client *minio.Client, bucketName string, objectName string, data io.Reader) error {
	_, err := client.GetBucketPolicy(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	n, err := client.PutObject(context.Background(), bucketName, objectName, data, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully uploaded bytes: ", n)
	return err
}

func GetData(client *minio.Client, bucketName string, objectName string) (file io.Reader) {
	_, err := client.GetBucketPolicy(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	file, err = client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	return file
}

func GetDataList(client *minio.Client, bucketName string) (file []io.Reader) {
	_, err := client.GetBucketPolicy(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	objectCh := client.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{
		Recursive: true,
	})
	for object := range objectCh {
		file = append(file, GetData(client, bucketName, object.Key))
	}
	return file
}

func DeleteData(client *minio.Client, bucketName string, objectName string) error {
	_, err := client.GetBucketPolicy(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	err = client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully deleted ", objectName)
	return err
}

func DeleteDataList(client *minio.Client, bucketName string) error {
	_, err := client.GetBucketPolicy(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	objectCh := client.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{
		Recursive: true,
	})
	for object := range objectCh {
		err = client.RemoveObject(context.Background(), bucketName, object.Key, minio.RemoveObjectOptions{})
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("Successfully deleted ", object.Key)
	}
	return err
}
