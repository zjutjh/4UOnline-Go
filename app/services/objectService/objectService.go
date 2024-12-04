package objectService

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"strings"
	"time"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/dustin/go-humanize"
	"github.com/gabriel-vasile/mimetype"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

var (
	// ErrUnsupportedUploadType 不支持的上传类型
	ErrUnsupportedUploadType = errors.New("unsupported upload type")

	// ErrSizeExceeded 文件大小超限
	ErrSizeExceeded = errors.New("file size exceeded")

	// ErrNotImage 使用 image 类型上传非图片的文件
	ErrNotImage = errors.New("file isn't a image")
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
	fileData []byte,
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
	mimeType, mimeExt := getFileTypeAndExt(fileData)

	// 检查是否为图像类型
	if uploadType == TypeImage && !strings.HasPrefix(mimeType, "image") {
		return "", "", ErrNotImage
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
func getFileTypeAndExt(fileData []byte) (mimeType string, mimeExt string) {
	mime := mimetype.Detect(fileData)
	return mime.String(), mime.Extension()
}

// ConvertToWebP 将图片转换为 WebP 格式
func ConvertToWebP(fileData []byte) ([]byte, int64, error) {
	img, err := imaging.Decode(bytes.NewReader(fileData))
	if err != nil {
		return nil, 0, err
	}

	var buf bytes.Buffer
	err = webp.Encode(&buf, img, &webp.Options{Quality: 100})
	if err != nil {
		return nil, 0, err
	}
	return buf.Bytes(), int64(buf.Len()), nil
}

// ReadFileToBytes 读取文件数据
func ReadFileToBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			zap.L().Warn("文件关闭失败", zap.Error(err))
		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return data, nil
}
