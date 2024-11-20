package objectController

import (
	"4u-go/app/apiException"
	"4u-go/app/services/objectService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"strings"
)

type UploadFileData struct {
	UploadType string                `form:"type" binding:"required"`
	File       *multipart.FileHeader `form:"file" binding:"required"`
}

// UploadFile 上传文件
func UploadFile(c *gin.Context) {
	var data UploadFileData
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
	if err != nil {
		if strings.Contains(err.Error(), "file size exceeds the maximum limit of 100MB") {
			apiException.AbortWithException(c, apiException.FileSizeExceedError, err)
			return
		}
		apiException.AbortWithException(c, apiException.GetFileInfoError, err)
		return
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	// 文件重命名
	ossSavePath, err := objectService.SetFileName(uploadType, fileExt)
	if err != nil {
		apiException.AbortWithException(c, apiException.SetFileNameError, err)
		return
	}

	// 上传文件
	objectUrl, err := objectService.PutObject(ossSavePath, file, fileHeader.Size, contentType, true)
	if err != nil {
		apiException.AbortWithException(c, apiException.UploadFileError, err)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"type": contentType,
		"url":  objectUrl,
	})
}
