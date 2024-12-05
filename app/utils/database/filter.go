package database

import "gorm.io/gorm"

// Filter 自定义筛选插件
// Usage: db.Scopes(dbUtils.Filter("college", [1,2,3]))
// Desc: 筛选出"college"字段为1,2,3其中之一的字段
//
// Spec: 使用泛型以接收所有类型的切片
func Filter[T any](name string, choices []T) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(name+" IN (?)", choices)
	}
}
