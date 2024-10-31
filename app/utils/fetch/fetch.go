//nolint:all
package fetch

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Fetch HTTP 客户端
type Fetch struct {
	Cookie []*http.Cookie
	client *http.Client
}

// InitUnSafe 初始化一个 HTTP 客户端，该客户端跳过 TLS 证书验证 (不安全)
func (f *Fetch) InitUnSafe() {
	f.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       time.Second * 15,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
}

// Init 初始化一个安全的 HTTP 客户端，不跳过 TLS 验证
func (f *Fetch) Init() {
	f.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       time.Second * 15,
	}
}

// SkipTlsCheck 动态设置客户端跳过 TLS 证书验证
func (f *Fetch) SkipTlsCheck() {
	f.client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

// Get 发起一个 GET 请求，并返回响应的内容
func (f *Fetch) Get(url string) ([]byte, error) {
	response, err := f.GetRaw(url)
	if err != nil {
		return nil, err
	}
	s, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetRaw 发起一个 GET 请求，返回原始 HTTP 响应对象
func (f *Fetch) GetRaw(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for _, v := range f.Cookie {
		request.AddCookie(v)
	}
	response, err := f.client.Do(request)
	if err != nil {
		return nil, err
	}
	f.Cookie = cookieMerge(f.Cookie, response.Cookies())
	return response, err
}

// PostFormRaw 发起一个 POST 表单请求，返回原始 HTTP 响应
func (f *Fetch) PostFormRaw(url string, requestData url.Values) (*http.Response, error) {
	request, _ := http.NewRequest("POST", url, strings.NewReader(requestData.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for _, v := range f.Cookie {
		request.AddCookie(v)
	}
	return f.client.Do(request)
}

// PostForm 发起一个 POST 表单请求，并返回响应的内容
func (f *Fetch) PostForm(url string, requestData url.Values) ([]byte, error) {
	response, err := f.PostFormRaw(url, requestData)
	if err != nil {
		return nil, err
	}
	f.Cookie = cookieMerge(f.Cookie, response.Cookies())
	return io.ReadAll(response.Body)
}

// PostJsonFormRaw 发起一个 POST JSON 请求，返回原始 HTTP 响应对象
func (f *Fetch) PostJsonFormRaw(url string, requestData map[string]any) (*http.Response, error) {
	bytesData, _ := json.Marshal(requestData)
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	request.Header.Set("Content-Type", "application/json")
	for _, v := range f.Cookie {
		request.AddCookie(v)
	}
	return f.client.Do(request)
}

// PostJsonForm 发起一个 POST JSON 请求，并返回响应的内容
func (f *Fetch) PostJsonForm(url string, requestData map[string]any) ([]byte, error) {
	response, err := f.PostJsonFormRaw(url, requestData)
	if err != nil {
		return nil, err
	}
	f.Cookie = cookieMerge(f.Cookie, response.Cookies())
	return io.ReadAll(response.Body)
}

// cookieMerge 合并新的 Cookie，将已有的同名 Cookie 替换
func cookieMerge(cookieA []*http.Cookie, cookieB []*http.Cookie) []*http.Cookie {
	for _, v := range cookieB {
		for k, v2 := range cookieA {
			if v.Name == v2.Name {
				cookieA = append(cookieA[:k], cookieA[k+1:]...)
				break
			}
		}
	}
	cookieA = append(cookieA, cookieB...)
	return cookieA
}
