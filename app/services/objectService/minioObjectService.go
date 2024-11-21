package objectService

import (
	"context"
	"io"
	"strings"

	"4u-go/config/objectStorage"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
)

// ms 是全局的 MinioService 实例
var ms = &objectStorage.MinioService

// PutObject 用于上传对象
func PutObject(objectKey string, reader io.Reader, size int64, contentType string) (string, error) {
	opts := minio.PutObjectOptions{ContentType: contentType}
	_, err := (*ms).Client.PutObject(context.Background(), (*ms).Bucket, objectKey, reader, size, opts)
	if err != nil {
		return "", err
	}
	return (*ms).Domain + (*ms).Bucket + "/" + objectKey, nil
}

// PutTemporaryObject 用于上传临时对象
func PutTemporaryObject(objectKey string, reader io.Reader, size int64, contentType string) (string, error) {
	return PutObject((*ms).TempDir+objectKey, reader, size, contentType)
}

// GetObjectKeyFromUrl 从 Url 中提取 ObjectKey
// 若该 Url 不是来自我们的 Minio, 则 ok 为 false
func GetObjectKeyFromUrl(fullUrl string) (objectKey string, ok bool) {
	objectKey = strings.TrimPrefix(fullUrl, (*ms).Domain+(*ms).Bucket+"/")
	if objectKey == fullUrl {
		return "", false
	}
	return objectKey, true
}

// DeleteObject 用于删除相应对象
func DeleteObject(objectKey string) error {
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
