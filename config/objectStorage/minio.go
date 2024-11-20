package objectStorage

import (
	"4u-go/config/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"strings"
)

// MinioCreateTempDirServant 结构体定义
type MinioCreateTempDirServant struct {
	Client  *minio.Client
	Bucket  string
	Domain  string
	TempDir string
}

// MinioService 是全局 MinIO 服务实例
var MinioService *MinioCreateTempDirServant

// Init 创建并返回 MinIO 服务客户端实例
func Init() error {
	// 从配置中获取 MinIO 配置信息
	endPoint := config.Config.GetString("minio.endPoint")
	accessKey := config.Config.GetString("minio.accessKey")
	secretKey := config.Config.GetString("minio.secretKey")
	secure := config.Config.GetBool("minio.secure")
	bucket := config.Config.GetString("minio.bucket")
	domain := config.Config.GetString("minio.domain")
	tempDir := strings.Trim(config.Config.GetString("minio.tempDir"), " /") + "/"

	// 初始化 MinIO 客户端对象
	client, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})
	if err != nil {
		logrus.Fatalf("objectStorage.minio 创建客户端失败: %s", err)
	}

	MinioService = &MinioCreateTempDirServant{
		Client:  client,
		Bucket:  bucket,
		Domain:  domain,
		TempDir: tempDir,
	}

	return nil
}
