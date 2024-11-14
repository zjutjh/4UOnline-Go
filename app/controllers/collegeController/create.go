package collegeController

import (
	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/collegeService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type createCollegeData struct {
	Name string `json:"name" binding:"required"`
}

// CreateCollege 新建学院
func CreateCollege(c *gin.Context) {
	var data createCollegeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = collegeService.SaveCollege(models.College{
		Name: data.Name,
	})
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
