package sdk

import (
	"fmt"
	"strings"
	"testing"
)

func TestClient_Auth(t *testing.T) {
	c := NewClient("ak-7Qf3KXH8QZOrW8Tf", "WaYGi4cBsievlfZsNhE3fY40ZB9dI9L3", "http", "127.0.0.1:8081", 8081, 3)

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
