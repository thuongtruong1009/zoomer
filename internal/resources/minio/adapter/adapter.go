package adapter

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/thuongtruong1009/zoomer/configs"
	"io"
	"log"
)

type MinioAdapter struct {
	client     *minio.Client
	bucketName string
	cfg        *configs.Configuration
}

func NewMinioAdapter(c *minio.Client, bucketName string, cfg *configs.Configuration) ResourceAdapter {
	return &MinioAdapter{
		client:     c,
		bucketName: bucketName,
		cfg:        cfg,
	}
}

func (ma *MinioAdapter) UploadData(objectName string, data io.Reader) error {
	_, err := ma.client.GetBucketPolicy(context.Background(), ma.bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	// objectName := file.Filename
	// fileBuffer := buffer
	// contentType := file.Header["Content-Type"][0]
	// fileSize := file.Size
	n, err := ma.client.PutObject(context.Background(), ma.bucketName, objectName, data, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully uploaded bytes: ", n)
	return err
}

func (ma *MinioAdapter) GetData(objectName string) (file io.Reader) {
	_, err := ma.client.GetBucketPolicy(context.Background(), ma.bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	file, err = ma.client.GetObject(context.Background(), ma.bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	return file
}

func (ma *MinioAdapter) GetDataList() (file []io.Reader) {
	_, err := ma.client.GetBucketPolicy(context.Background(), ma.bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	objectCh := ma.client.ListObjects(context.Background(), ma.bucketName, minio.ListObjectsOptions{
		Recursive: true,
	})
	for object := range objectCh {
		file = append(file, ma.GetData(object.Key))
	}
	return file
}

func (ma *MinioAdapter) DeleteData(objectName string) error {
	_, err := ma.client.GetBucketPolicy(context.Background(), ma.bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	err = ma.client.RemoveObject(context.Background(), ma.bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully deleted ", objectName)
	return err
}

func (ma *MinioAdapter) DeleteDataList() error {
	_, err := ma.client.GetBucketPolicy(context.Background(), ma.bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	objectCh := ma.client.ListObjects(context.Background(), ma.bucketName, minio.ListObjectsOptions{
		Recursive: true,
	})
	for object := range objectCh {
		err = ma.client.RemoveObject(context.Background(), ma.bucketName, object.Key, minio.RemoveObjectOptions{})
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("Successfully deleted ", object.Key)
	}
	return err
}
