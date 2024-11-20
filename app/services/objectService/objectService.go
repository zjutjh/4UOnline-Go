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

// 定义支持的文件扩展名
const (
	// PNGExt 表示 PNG 文件的扩展名
	PNGExt = ".png"
	// JPGExt 表示 JPG 文件的扩展名
	JPGExt = ".jpg"
	// ZIPExt 表示 ZIP 文件的扩展名
	ZIPExt = ".zip"
	// RARExt 表示 RAR 文件的扩展名
	RARExt = ".rar"
)

var uploadTypeLimits = map[string]int64{
	"public/image":      1024 * 1024 * 10,
	"public/attachment": 1024 * 1024 * 100,
}

var supportedFileTypes = map[string]string{
	"image/png":                    ".png",
	"image/jpeg":                   ".jpg",
	"image/jpg":                    ".jpg",
	"application/zip":              ".zip",
	"application/x-zip":            ".zip",
	"application/octet-stream":     ".zip",
	"application/x-zip-compressed": ".zip",
	"application/vnd.rar":          ".rar",
	"application/x-rar-compressed": ".rar",
}

var magicNumberMapping = map[string][]byte{
	PNGExt: {0x89, 0x50, 0x4E, 0x47},
	JPGExt: {0xFF, 0xD8, 0xFF},
	ZIPExt: []byte("PK"),
	RARExt: {0x52, 0x61, 0x72, 0x21, 0x1A},
}

// GetFileInfo 实现了封装文件基本信息的功能
func GetFileInfo(
	file multipart.File,
	fileHeader *multipart.FileHeader,
	uploadType string,
) (
	contentType string,
	fileExt string,
	xerr error,
) {
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

// fileCheck 根据上传类型和文件大小进行检查
func fileCheck(uploadType string, size int64) error {
	maxSize, ok := uploadTypeLimits[uploadType]
	if !ok {
		return errors.New("unsupported upload type")
	}
	if size > maxSize {
		return errors.New("file size exceeds the maximum limit")
	}
	return nil
}

// getFileExt 根据文件类型获取文件扩展名
func getFileExt(contentType string) (string, error) {
	if ext, ok := supportedFileTypes[contentType]; ok {
		return ext, nil
	}
	return "", errors.New("unsupported file type")
}

// getFileExtByMagic 根据文件头（Magic Number）判断文件类型
func getFileExtByMagic(file multipart.File) (string, error) {
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return "", errors.New("failed to read file header")
	}
	for ext, magic := range magicNumberMapping {
		if bytes.HasPrefix(buffer, magic) {
			return ext, nil
		}
	}
	return "", errors.New("unsupported file type")
}
