package collegeController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/services/collegeService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type updateCollegeData struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// UpdateCollege 更新学院信息
func UpdateCollege(c *gin.Context) {
	var data updateCollegeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	college, err := collegeService.GetCollegeById(data.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		return
	}
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	{ // 更新学院信息
		college.Name = data.Name
	}

	err = collegeService.SaveCollege(college)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
