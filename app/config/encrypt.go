package config

// encryptKey 是用于加密的配置键。
const encryptKey = "encryptKey"

// SetEncryptKey 设置加密密钥的值
func SetEncryptKey(value string) error {
	return setConfig(encryptKey, value)
}

// GetEncryptKey 获取当前配置的加密密钥值
func GetEncryptKey() string {
	return getConfig(encryptKey)
}

// IsSetEncryptKey 检查是否设置了加密密钥
func IsSetEncryptKey() bool {
	return checkConfig(encryptKey)
}

// DelEncryptKey 删除加密密钥
func DelEncryptKey() error {
	return delConfig(encryptKey)
}
