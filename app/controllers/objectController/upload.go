package objectController

import (
	"bytes"
	"errors"
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
	fileData, err := objectService.ReadFileToBytes(data.File)
	if err != nil {
		apiException.AbortWithException(c, apiException.UploadFileError, err)
		return
	}

	// 获取文件信息
	contentType, fileExt, err := objectService.GetFileInfo(fileData, fileSize, uploadType)
	if errors.Is(err, objectService.ErrSizeExceeded) {
		apiException.AbortWithException(c, apiException.FileSizeExceedError, err)
		return
	}
	if errors.Is(err, objectService.ErrUnsupportedUploadType) {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	if errors.Is(err, objectService.ErrNotImage) {
		apiException.AbortWithException(c, apiException.FileNotImageError, err)
		return
	}
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	if uploadType == objectService.TypeImage {
		d, s, err := objectService.ConvertToWebP(fileData)
		if err != nil {
			zap.L().Error("转换图片到 WebP 失败", zap.Error(err))
		} else { // 若转换成功则替代原文件
			fileData = d
			fileSize = s
			fileExt = ".webp"
			contentType = "image/webp"
		}
	}

	// 上传文件
	objectKey := objectService.GenerateObjectKey(uploadType, fileExt)
	objectUrl, err := objectService.PutObject(objectKey, bytes.NewReader(fileData), fileSize, contentType)
	if err != nil {
		apiException.AbortWithException(c, apiException.UploadFileError, err)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"type": contentType,
		"url":  objectUrl,
	})
}
