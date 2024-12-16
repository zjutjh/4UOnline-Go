//nolint:all
package request

import (
	"crypto/tls"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

// Client 包装 Resty 客户端
type Client struct {
	*resty.Client
}

// New 初始化一个 Resty 客户端
func New() Client {
	s := Client{
		Client: resty.New().
			SetTimeout(5 * time.Second).
			SetJSONMarshaler(jsoniter.ConfigCompatibleWithStandardLibrary.Marshal).
			SetJSONUnmarshaler(jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal),
	}
	// 利用中间件实现请求日志
	s.OnAfterResponse(RestyLogMiddleware)
	return s
}

// NewUnSafe 初始化一个 Resty 客户端并跳过 TLS 证书验证
func NewUnSafe() Client {
	s := New()
	s.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	return s
}

// Request 获取一个新的请求实例
func (s Client) Request() *resty.Request {
	return s.R().EnableTrace()
}

// RestyLogMiddleware Resty日志中间件
func RestyLogMiddleware(_ *resty.Client, resp *resty.Response) error {
	if resp.IsError() {
		method := resp.Request.Method
		url := resp.Request.URL
		zap.L().Warn("请求出现错误",
			zap.String("method", method),
			zap.String("url", url),
			zap.Int64("time_spent(ms)", resp.Time().Milliseconds()),
		)
	}
	return nil
}
