package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/net/http2"
)

// 与网关之间的通讯
type GateConn struct {
	schema string
	client *http.Client // 内置长连接+并发连接数
}

type GatewayConfig struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`

	GatewayMaxConnection int `json:"gatewayMaxConnection"`
	GatewayTimeout       int `json:"gatewayTimeout"`     // 单位是毫秒
	GatewayIdleTimeout   int `json:"gatewayIdleTimeout"` // 单位是秒
	GatewayPushRetry     int `json:"gatewayPushRetry"`   // 超过重试次数后, 消息将被丢弃
}

func GetConfig() *GatewayConfig {
	return &GatewayConfig{
		Hostname:             "localhost",
		Port:                 8080,
		GatewayMaxConnection: 32,
		GatewayTimeout:       3000,
		GatewayIdleTimeout:   60,
		GatewayPushRetry:     3,
	}
}

func InitGateConn(gatewayConfig *GatewayConfig) (gateConn *GateConn, err error) {
	var (
		transport *http.Transport
	)

	gateConn = &GateConn{
		schema: "https://" + gatewayConfig.Hostname + ":" + strconv.Itoa(gatewayConfig.Port),
	}

	transport = &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, // 不校验服务端证书
		MaxIdleConns:        gatewayConfig.GatewayMaxConnection,
		MaxIdleConnsPerHost: gatewayConfig.GatewayMaxConnection,
		IdleConnTimeout:     time.Duration(gatewayConfig.GatewayIdleTimeout) * time.Second, // 连接空闲超时
	}
	// 启动HTTP/2协议
	http2.ConfigureTransport(transport)

	// HTTP/2 客户端
	gateConn.client = &http.Client{
		Transport: transport,
		Timeout:   time.Duration(gatewayConfig.GatewayTimeout) * time.Millisecond, // 请求超时
	}
	return
}

// 出于性能考虑, 消息数组在此前已经编码成json
func (gateConn *GateConn) PushAll(itemsJson []byte) (err error) {
	var (
		apiUrl string
		form   url.Values
		resp   *http.Response
		retry  int
	)

	apiUrl = gateConn.schema + "/push/all"

	form = url.Values{}
	form.Set("items", string(itemsJson))

	for retry = 0; retry < GetConfig().GatewayPushRetry; retry++ {
		if resp, err = gateConn.client.PostForm(apiUrl, form); err != nil {
			continue
		}
		resp.Body.Close()
		break
	}
	return
}

func main() {
	var (
		conn *GateConn
		err  error
	)
	if conn, err = InitGateConn(GetConfig()); err != nil {
		return
	}

	i := 0
	for {
		body := fmt.Sprintf(`[{"name":"%d"}, {"name":"%d"}, {"name":"%d"}]`, i, i+1, i+2)
		conn.PushAll([]byte(body))
		i++

		time.Sleep(1 * time.Second)
	}

}
