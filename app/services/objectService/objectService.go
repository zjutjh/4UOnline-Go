package objectService

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gabriel-vasile/mimetype"
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

// SetFileName 将文件重命名
func SetFileName(uploadType string, fileExt string) (string, error) {
	now := time.Now()
	timestamp := now.UnixNano() / 1e6

	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", errors.New("failed to generate random bytes")
	}

	randomString := hex.EncodeToString(randomBytes)
	randomPath := fmt.Sprintf("%d%s", timestamp, randomString)
	if randomPath == "" {
		return "", errors.New("failed to generate random path")
	}

	ossSavePath := fmt.Sprintf("%s/%d/%s%s", uploadType, now.Year(), randomPath, fileExt)
	if ossSavePath == "" {
		return "", errors.New("failed to generate ossSavePath")
	}

	return ossSavePath, nil
}

// RemoveDomain 去除 URL 的域名部分
func RemoveDomain(fullUrl string, domain string, bucket string) string {
	return strings.TrimPrefix(strings.TrimPrefix(fullUrl, domain), bucket+"/")
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
