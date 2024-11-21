package objectController

import (
	"errors"
	"io"
	"mime/multipart"

	"4u-go/app/apiException"
	"4u-go/app/services/objectService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
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
	fileHeader := data.File
	// 获取文件流
	file, err := data.File.Open()
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			apiException.AbortWithException(c, apiException.ServerError, err)
			return
		}
	}(file)

	// 获取文件信息
	contentType, fileExt, err := objectService.GetFileInfo(file, fileHeader, uploadType)
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
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	// 上传文件
	objectKey := objectService.GetObjectKey(uploadType, fileExt)
	objectUrl, err := objectService.PutObject(objectKey, file, fileHeader.Size, contentType)
	if err != nil {
		apiException.AbortWithException(c, apiException.UploadFileError, err)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"type": contentType,
		"url":  objectUrl,
	})
}
