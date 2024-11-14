package collageController

import (
	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/collageService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type createCollageData struct {
	Name string `json:"name"`
}

// CreateCollage 新建学院
func CreateCollage(c *gin.Context) {
	var data createCollageData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = collageService.CreateCollage(models.Collage{
		Name: data.Name,
	})
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
