package objectService

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif" // 注册解码器
	_ "image/jpeg"
	_ "image/png"
	"io"
	"time"

	"4u-go/config/sdk"
	"github.com/chai2010/webp"
	"github.com/dustin/go-humanize"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	_ "golang.org/x/image/bmp" // 注册解码器
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

// ImageLimit 图片上传大小限制
const ImageLimit = humanize.MByte * 10

// GenerateObjectKey 通过 UUID 作为文件名并生成 ObjectKey
func GenerateObjectKey(uploadType string, fileExt string) string {
	return fmt.Sprintf("%s/%d/%s%s", uploadType, time.Now().Year(), uuid.NewV1().String(), fileExt)
}

// ConvertToWebP 将图片转换为 WebP 格式
func ConvertToWebP(reader io.Reader) (io.Reader, int64, error) {
	img, _, err := image.Decode(reader)
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

// DeleteObjectByUrlAsync 通过给定的 Url 异步删除对象
func DeleteObjectByUrlAsync(url string) {
	objectKey, ok := sdk.MinioService.GetObjectKeyFromUrl(url)
	if ok {
		go func(objectKey string) {
			err := sdk.MinioService.DeleteObject(objectKey)
			if err != nil {
				zap.L().Error("Minio 删除对象错误", zap.String("objectKey", objectKey), zap.Error(err))
			}
		}(objectKey)
	}
}
