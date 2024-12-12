package objectController

import (
	"errors"
	"image"
	"mime/multipart"

	"4u-go/app/apiException"
	"4u-go/app/services/objectService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type uploadFileData struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// UploadFile 上传文件
func UploadFile(c *gin.Context) {
	var data uploadFileData
	if err := c.ShouldBind(&data); err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

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
	if fileSize > objectService.ImageLimit {
		apiException.AbortWithException(c, apiException.FileSizeExceedError, nil)
		return
	}

	reader, size, err := objectService.ConvertToWebP(file)
	if errors.Is(err, image.ErrFormat) {
		apiException.AbortWithException(c, apiException.FileNotImageError, err)
		return
	}
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	contentType := "image/webp"

	// 上传文件
	objectKey := objectService.GenerateObjectKey("image", ".webp")
	objectUrl, err := objectService.PutObject(objectKey, reader, size, contentType)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"type": contentType,
		"url":  objectUrl,
	})
}
