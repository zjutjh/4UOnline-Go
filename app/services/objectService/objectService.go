package objectService

import (
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gabriel-vasile/mimetype"
	uuid "github.com/satori/go.uuid"
)

var (
	// ErrUnsupportedUploadType 不支持的上传类型
	ErrUnsupportedUploadType = errors.New("unsupported upload type")

	// ErrSizeExceeded 文件大小超限
	ErrSizeExceeded = errors.New("file size exceeded")
)

var uploadTypeLimits = map[string]int64{
	"public/image":      humanize.MByte * 10,
	"public/attachment": humanize.MByte * 100,
}

// GetFileInfo 获取文件基本信息
func GetFileInfo(
	file multipart.File,
	fileHeader *multipart.FileHeader,
	uploadType string,
) (
	contentType string,
	fileExt string,
	err error,
) {
	// 检查文件大小
	if err = checkFileSize(uploadType, fileHeader.Size); err != nil {
		return "", "", err
	}

	// 通过文件头获取类型和扩展名
	mimeType, mimeExt, err := getFileTypeAndExt(file)
	if err != nil {
		return "", "", err
	}

	return mimeType, mimeExt, nil
}

// GetObjectKey 通过 UUID 作为文件名并返回 ObjectKey
func GetObjectKey(uploadType string, fileExt string) string {
	return fmt.Sprintf("%s/%d/%s%s", uploadType, time.Now().Year(), uuid.NewV1().String(), fileExt)
}

// checkFileSize 检查文件大小
func checkFileSize(uploadType string, size int64) error {
	maxSize, ok := uploadTypeLimits[uploadType]
	if !ok {
		return ErrUnsupportedUploadType
	}
	if size > maxSize {
		return ErrSizeExceeded
	}
	return nil
}

// getFileTypeAndExt 根据文件头（Magic Number）判断文件类型和扩展名
func getFileTypeAndExt(file multipart.File) (mimeType string, mimeExt string, err error) {
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return "", "", err
	}
	return mime.String(), mime.Extension(), nil
}
