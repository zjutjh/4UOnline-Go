package websiteController

import (
	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/websiteService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type getWebsiteData struct {
	Type    uint `json:"type" binding:"required"`
	College uint `json:"college"`
}

type getWebsiteResponse struct {
	WebsiteList []websiteElement `json:"website_list"`
}

type getEditableWebsiteResponse struct {
	WebsiteList []models.Website `json:"website_list"`
}

type websiteElement struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Condition   string `json:"condition"`
	Editable    bool   `json:"editable"`
}

// GetWebsiteList 获取网站列表
func GetWebsiteList(c *gin.Context) {
	var data getWebsiteData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	list, err := websiteService.GetWebsiteList(data.Type, data.College)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	websiteList := make([]websiteElement, 0)
	for _, website := range list {
		websiteList = append(websiteList, websiteElement{
			ID:          website.ID,
			Title:       website.Title,
			Description: website.Description,
			URL:         website.URL,
			Condition:   website.Condition,
		})
	}

	utils.JsonSuccessResponse(c, getWebsiteResponse{
		WebsiteList: websiteList,
	})
}

// GetEditableWebsites 获取可管理的网站列表
func GetEditableWebsites(c *gin.Context) {
	list, err := websiteService.GetAllWebsites()
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	websiteList := list

	user := c.GetUint("user_id")
	adminType := c.GetUint("admin_type")
	if adminType != models.SuperAdmin {
		editableList := make([]models.Website, 0)
		for _, website := range list {
			// TODO: 根据管理员对应学院进行筛选
			if website.AuthorID == user {
				editableList = append(editableList, website)
			}
		}
		websiteList = editableList
	}

	utils.JsonSuccessResponse(c, getEditableWebsiteResponse{
		WebsiteList: websiteList,
	})
}
