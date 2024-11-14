package collegeController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/services/collegeService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type deleteCollegeData struct {
	ID uint `json:"id" binding:"required"`
}

// DeleteCollege 删除学院
func DeleteCollege(c *gin.Context) {
	var data deleteCollegeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// 判断学院是否存在
	_, err = collegeService.GetCollegeById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	err = collegeService.DeleteCollegeById(data.ID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
