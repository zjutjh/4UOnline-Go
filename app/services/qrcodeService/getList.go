package qrcodeService

import (
	"4u-go/app/models"
	dbUtils "4u-go/app/utils/database"
	"4u-go/config/database"
)

// GetList 获取权益码信息列表的筛选,搜索,分页
// revive:disable:flag-parameter
func GetList(
	collegeFilter []uint,
	feedbackFilter []uint,
	qrcodeStatus bool,
	keyword string,
	page int, pageSize int,
) (qrcodeList []models.Qrcode, total int64, err error) {
	query := database.DB.Model(models.Qrcode{})

	// 关键词搜索
	if len(keyword) > 0 {
		query = query.Where("ID = ? "+
			"OR department LIKE ? "+
			"OR location LIKE ? "+
			"OR description LIKE ?", keyword, "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 筛选`权益码状态`
	if qrcodeStatus {
		query = query.Where("status = ?", 1)
	}

	// 筛选`责任部门`
	if len(collegeFilter) > 0 {
		query = query.Scopes(dbUtils.Filter("college", collegeFilter))
	}

	// 筛选`反馈类型`
	if len(feedbackFilter) > 0 {
		query = query.Scopes(dbUtils.Filter("feedback_type", feedbackFilter))
	}

	// 分页查找
	err = query.Count(&total).
		Scopes(dbUtils.Paginate(page, pageSize)).
		Find(&qrcodeList).Error

	return qrcodeList, total, err
}
