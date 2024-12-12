package objectService

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	_ "image/gif" // 注册解码器
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"time"

	"github.com/chai2010/webp"
	"github.com/dustin/go-humanize"
	"github.com/gabriel-vasile/mimetype"
	uuid "github.com/satori/go.uuid"
	_ "golang.org/x/image/bmp" // 注册解码器
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

var (
	// ErrUnsupportedUploadType 不支持的上传类型
	ErrUnsupportedUploadType = errors.New("unsupported upload type")

	// ErrSizeExceeded 文件大小超限
	ErrSizeExceeded = errors.New("file size exceeded")
)

const (
	// TypeImage 图片
	TypeImage = "image"

	// TypeAttachment 附件
	TypeAttachment = "attachment"
)

var uploadTypeLimits = map[string]int64{
	TypeImage:      humanize.MByte * 10,
	TypeAttachment: humanize.MByte * 100,
}

// GetFileInfo 获取文件基本信息
func GetFileInfo(
	file multipart.File,
	fileSize int64,
	uploadType string,
) (
	contentType string,
	fileExt string,
	err error,
) {
	// 检查文件大小
	if err = checkFileSize(uploadType, fileSize); err != nil {
		return "", "", err
	}

	// 通过文件头获取类型和扩展名
	mimeType, mimeExt, err := getFileTypeAndExt(file)
	if err != nil {
		return "", "", err
	}
	return mimeType, mimeExt, nil
}

// GenerateObjectKey 通过 UUID 作为文件名并生成 ObjectKey
func GenerateObjectKey(uploadType string, fileExt string) string {
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
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", "", err
	}
	return mime.String(), mime.Extension(), nil
}

// ConvertToWebP 将图片转换为 WebP 格式
func ConvertToWebP(file multipart.File) (io.Reader, int64, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, 0, err
	}

	var buf bytes.Buffer
	err = webp.Encode(&buf, img, &webp.Options{Quality: 100})
	if err != nil {
		return nil, 0, err
	}
	return bytes.NewReader(buf.Bytes()), int64(buf.Len()), nil
}
