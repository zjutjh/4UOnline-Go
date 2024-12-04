//nolint:all
package objectController

import (
	"errors"
	"io"
	"mime/multipart"

	"4u-go/app/apiException"
	"4u-go/app/services/objectService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type uploadFileData struct {
	UploadType string                `form:"type" binding:"required"`
	File       *multipart.FileHeader `form:"file" binding:"required"`
}

// UploadFile 上传文件
func UploadFile(c *gin.Context) {
	var data uploadFileData
	if err := c.ShouldBind(&data); err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	uploadType := data.UploadType
	fileSize := data.File.Size
	file, err := data.File.Open()
	if err != nil {
		apiException.AbortWithException(c, apiException.UploadFileError, err)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			zap.L().Warn("文件关闭错误", zap.Error(err))
		}
	}(file)

	// 获取文件信息
	contentType, fileExt, err := objectService.GetFileInfo(file, fileSize, uploadType)
	if errors.Is(err, objectService.ErrSizeExceeded) {
		apiException.AbortWithException(c, apiException.FileSizeExceedError, err)
		return
	}
	if errors.Is(err, objectService.ErrUnsupportedUploadType) {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	var fileReader io.Reader = file
	if uploadType == objectService.TypeImage {
		reader, size, err := objectService.ConvertToWebP(file)
		if err != nil {
			if errors.Is(err, objectService.ErrNotImage) {
				apiException.AbortWithException(c, apiException.FileNotImageError, err)
				return
			}
			zap.L().Error("转换图片到 WebP 失败", zap.Error(err))
		} else { // 若转换成功则替代原文件
			fileReader = reader
			fileSize = size
			fileExt = ".webp"
			contentType = "image/webp"
		}
	}

	// 上传文件
	objectKey := objectService.GenerateObjectKey(uploadType, fileExt)
	objectUrl, err := objectService.PutObject(objectKey, fileReader, fileSize, contentType)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"type": contentType,
		"url":  objectUrl,
	})
}
