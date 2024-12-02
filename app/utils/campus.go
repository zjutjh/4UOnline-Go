package utils

const (
	campusZH uint = 1 << iota
	campusPF
	campusMGS
)

// EncodeCampus 存储校区信息
func EncodeCampus(campus []uint) uint {
	var result uint
	for _, c := range campus {
		if c > 3 || c == 0 { // 拦截错误参数
			continue
		}
		result |= 1 << (c - 1)
	}
	return result
}

// DecodeCampus 提取校区信息
func DecodeCampus(campus uint) []uint {
	result := make([]uint, 0)
	if campus&campusZH != 0 {
		result = append(result, 1)
	}
	if campus&campusPF != 0 {
		result = append(result, 2)
	}
	if campus&campusMGS != 0 {
		result = append(result, 3)
	}
	return result
}
