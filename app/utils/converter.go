package utils

import "strings"

// StringsToString 将字符串数组转换为字符串,以逗号分隔
func StringsToString(strSlice []string) string {
	var result string
	for i, str := range strSlice {
		if i != 0 {
			result += ","
		}
		result += str
	}
	return result
}

// StringToStrings 将字符串转换为字符串数组
func StringToStrings(str string) []string {
	return strings.Split(str, ",")
}
