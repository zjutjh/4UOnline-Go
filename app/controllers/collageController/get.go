package collageController

import (
	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/collageService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type getCollageResponse struct {
	CollageList []models.Collage `json:"collage_list"`
}

// GetCollageList 获取学院列表
func GetCollageList(c *gin.Context) {
	collageList, err := collageService.GetCollageList()
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, getCollageResponse{collageList})
}
