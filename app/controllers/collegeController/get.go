package collegeController

import (
	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/collegeService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type getCollegeResponse struct {
	CollegeList []models.College `json:"college_list"`
}

// GetCollegeList 获取学院列表
func GetCollegeList(c *gin.Context) {
	collegeList, err := collegeService.GetCollegeList()
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, getCollegeResponse{
		CollegeList: collegeList,
	})
}
