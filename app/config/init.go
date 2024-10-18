package config

// initKey 是用于初始化状态的配置键。
const initKey = "initKey"

// SetInit 设置初始化状态为 "True"。
func SetInit() error {
	return setConfig(initKey, "True")
}

// ResetInit 设置初始化状态为 "False"。
func ResetInit() error {
	return setConfig(initKey, "False")
}

// GetInit 获取当前的初始化状态。
func GetInit() bool {
	return getConfig(initKey) == "True"
}
