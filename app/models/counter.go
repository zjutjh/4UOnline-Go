package models

// Counter 用于记录每天相关字段的增量
type Counter struct {
	ID    uint
	Day   int64  // 基于时间戳计算当前天数 `time.Now()/86400`
	Name  string // 统计的字段名
	Count int64  // 计数器
}
