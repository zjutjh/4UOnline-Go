package collageController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/services/collageService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type updateCollageData struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// UpdateCollage 更新学院信息
func UpdateCollage(c *gin.Context) {
	var data updateCollageData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	collage, err := collageService.GetCollageById(data.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		return
	}
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	{ // 更新学院信息
		collage.Name = data.Name
	}

	err = collageService.SaveCollage(collage)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
