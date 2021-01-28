package sdk

import (
	"fmt"
	"strings"
	"testing"
)

func TestClient_Auth(t *testing.T) {
	//c := NewClient("ak-7Qf3KXH8QZOrW8Tf", "WaYGi4cBsievlfZsNhE3fY40ZB9dI9L3", "http", "127.0.0.1:8081", 8081, 3)
	c := NewClient("ak-BAQEFAASCAmEwggJd", "YysbbB3iIwKhSkY9QwIDAQAB", "http", "127.0.0.1:8081", 8081, 3)

	// GET Request Test
	fmt.Println("GET /x-admin/ping")

	statusCode, resp := c.Get("/x-admin/ping", "", nil)

	fmt.Println("Status Code: " + fmt.Sprint(statusCode))
	fmt.Println(resp)

	fmt.Println(strings.Repeat("-", 50))

	// POST Request Test
	body := `{"echo":{"int":1,"str":"Hello World","unicode":"你好，世界！","none":null,"boolean":true}}`
	fmt.Println("GET /x-admin/echo")
	fmt.Println(body)

	statusCode, resp = c.Post("/x-admin/echo", "", body, nil)

	fmt.Println("Status Code: " + fmt.Sprint(statusCode))
	fmt.Println(resp)
}

func TestClient_Ak(t *testing.T) {
	//c := NewClient("ak-7Qf3KXH8QZOrW8Tf", "WaYGi4cBsievlfZsNhE3fY40ZB9dI9L3", "http", "127.0.0.1:8081", 8081, 3)
	c := NewClient("ak-BAQEFAASCAmEwggJd", "YysbbB3iIwKhSkY9QwIDAQAB", "http", "127.0.0.1:8081", 8081, 3)
	c = NewClient("ak-abc-123", "sk-abc-123", "http", "sx.api.mengtuiapp.com:18081", 18081, 3)

	// POST Request Test
	body := `{
	  "name": "name-test-2",
	  "value": "value-test-1"
	}`
	body = `{ "开发环境_kafka-67_密码": "name-test-1", "name": "dev_kafka_67_pwd", "value": "rN5JkN4yXrR", "allow_ip": "127.0.0.1,192.168.1.*,10.105.*" }`
	fmt.Println("GET /x-admin/echo")
	fmt.Println(body)

	statusCode, resp := c.Post("/x-admin/ak_config", "", body, map[string]string{"Content-Type": "application/json"})

	fmt.Println("Status Code: " + fmt.Sprint(statusCode))
	fmt.Println(resp)

	statusCode, resp = c.Get("/x-admin/ak_config", "ip=10.105.16.1&name=name-test-2", nil)
	fmt.Println("Status Code: " + fmt.Sprint(statusCode))
	fmt.Println(resp)

}
