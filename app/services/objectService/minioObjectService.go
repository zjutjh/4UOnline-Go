package objectService

import (
	"context"
	"io"

	"4u-go/config/objectStorage"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
)

// ms 是全局的 MinioService 实例
var ms = &objectStorage.MinioService

// PutPersistentObject 用于上传持久化对象
func PutPersistentObject(objectKey string, reader io.Reader, size int64, contentType string) (string, error) {
	opts := minio.PutObjectOptions{ContentType: contentType}
	_, err := (*ms).Client.PutObject(context.Background(), (*ms).Bucket, objectKey, reader, size, opts)
	if err != nil {
		return "", err
	}
	return (*ms).Domain + (*ms).Bucket + "/" + objectKey, nil
}

// PutTemporaryObject 用于上传临时对象
func PutTemporaryObject(objectKey string, reader io.Reader, size int64, contentType string) (string, error) {
	opts := minio.PutObjectOptions{ContentType: contentType}
	objectName := (*ms).TempDir + objectKey
	_, err := (*ms).Client.PutObject(context.Background(), (*ms).Bucket, objectName, reader, size, opts)
	if err != nil {
		return "", err
	}
	return (*ms).Domain + (*ms).Bucket + "/" + objectName, nil
}

// DelObject 用于删除相应对象
func DelObject(objectKey string) error {
	err := (*ms).Client.RemoveObject(
		context.Background(),
		(*ms).Bucket,
		objectKey,
		minio.RemoveObjectOptions{ForceDelete: true},
	)
	if err != nil {
		return errors.New("failed to delete object from bucket")
	}
	return nil
}
