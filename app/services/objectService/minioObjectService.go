package objectService

import (
	"4u-go/config/objectStorage"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"io"
)

// ms 是全局的 MinioService 实例
var ms = &objectStorage.MinioService

// PutObject 用于将对象上传到 MinIO 对象存储中
func PutObject(objectKey string, reader io.Reader, objectSize int64, contentType string, persistence bool) (string, error) {
	opts := minio.PutObjectOptions{ContentType: contentType}
	objectName := objectKey
	if !persistence {
		objectName = (*ms).TempDir + objectKey
	}

	_, err := (*ms).Client.PutObject(context.Background(), (*ms).Bucket, objectName, reader, objectSize, opts)
	if err != nil {
		return "", err
	}
	return (*ms).Domain + (*ms).Bucket + "/" + objectKey, nil
}

// DelObject 用于删除相应对象
func DelObject(objectKey string) error {
	err := (*ms).Client.RemoveObject(context.Background(), (*ms).Bucket, objectKey, minio.RemoveObjectOptions{ForceDelete: true})
	if err != nil {
		return errors.New("failed to delete object from bucket")
	}
	return nil
}
