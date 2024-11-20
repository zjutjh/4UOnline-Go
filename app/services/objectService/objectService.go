package objectService

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"mime/multipart"
	"strings"
	"time"
)

// GetFileInfo 实现了封装文件基本信息的功能
func GetFileInfo(file multipart.File, fileHeader *multipart.FileHeader, uploadType string) (contentType string, fileExt string, xerr error) {
	if xerr = fileCheck(uploadType, fileHeader.Size); xerr != nil {
		return "", "", xerr
	}
	contentType = fileHeader.Header.Get("Content-Type")
	fileExt, xerr = getFileExt(contentType)
	if xerr != nil {
		return "", "", xerr
	}

	magicExt, xerr := getFileExtByMagic(file)
	if xerr != nil {
		return "", "", xerr
	}
	if fileExt != magicExt {
		return "", "", errors.New("mismatch between Content-Type and file content")
	}

	return contentType, fileExt, nil
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
	if len(randomString) == 0 {
		return "", errors.New("failed to encode random bytes to hex string")
	}

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
	withoutDomain := strings.TrimPrefix(fullUrl, domain)
	withoutBucket := strings.TrimPrefix(withoutDomain, bucket+"/")
	return withoutBucket
}

// fileCheck 根据上传类型和文件大小进行检查
func fileCheck(uploadType string, size int64) error {
	if uploadType != "public/image" &&
		uploadType != "public/attachment" {
		return errors.New("unsupported upload type")
	}
	if size > 1024*1024*100 {
		return errors.New("file size exceeds the maximum limit of 100MB")
	}
	return nil
}

// getFileExt 根据文件类型获取文件扩展名
func getFileExt(s string) (string, error) {
	switch s {
	case "image/png":
		return ".png", nil
	case "image/jpg":
		return ".jpg", nil
	case "image/jpeg":
		return ".jpg", nil
	case "application/zip",
		"application/x-zip",
		"application/octet-stream",
		"application/x-zip-compressed":
		return ".zip", nil
	default:
		return "", errors.New("unsupported file type")
	}
}

// getFileExtByMagic 根据文件头（Magic Number）判断文件类型
func getFileExtByMagic(file multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", errors.New("failed to read file header")
	}

	if bytes.HasPrefix(buffer, []byte{0x89, 0x50, 0x4E, 0x47}) {
		return ".png", nil
	} else if bytes.HasPrefix(buffer, []byte{0xFF, 0xD8, 0xFF}) {
		return ".jpg", nil
	} else if strings.HasPrefix(string(buffer[:4]), "PK") {
		return ".zip", nil
	}

	return "", errors.New("unsupported file type")
}
