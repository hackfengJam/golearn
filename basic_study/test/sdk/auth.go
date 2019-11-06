package sdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	prefixMt = "mm"
)

type Client struct {
	AKId     string
	AKSecret string
	Host     string
	Scheme   string
	Port     int
	Timeout  int
}

func NewClient(akId, akSecret, scheme, host string, port int, timeout int) *Client {
	c := &Client{
		AKId:     akId,
		AKSecret: akSecret,
		Host:     host,
		Port:     port,
		Timeout:  timeout,
		Scheme:   scheme,
	}

	return c
}

// mt :xxxx -> xxx
func (c *Client) Sign2Authorization(prefix string, sign string) string {
	return strings.Join([]string{
		prefix, sign,
	}, " ")
}

// xxxx -> mt :xxxxx
func (c *Client) Authorization2Sign(prefix string, authorization string) string {
	return strings.TrimLeft(authorization, prefix+" ")
}

func (c *Client) GetSign(headers map[string]string, method, path, query, body string) string {
	mac := &Mac{
		AccessKey: c.AKId,
		SecretKey: []byte(c.AKSecret),
		Host:      c.Host,
	}

	contentType, ok := headers["Content-Type"]
	if !ok {
		contentType = ""
	}

	sa := &SignArgs{
		strings.ToUpper(method),
		path,
		query,
		contentType,
		[]byte(body),
	}
	sign, _ := mac.SignRequest(sa)
	return sign
}

func (c *Client) VerifySign(sign string, headers map[string]string, method, path, query, body string) bool {
	expectedSign := c.GetSign(headers, method, path, query, body)
	fmt.Printf("VerifySign: sign: %v, expectedSign: %v\n", sign, expectedSign)
	return sign == expectedSign
}

func (c *Client) GetAuthHeader(headers map[string]string, method, path, query, body string) map[string]string {
	sign := c.GetSign(headers, method, path, query, body)

	authHeader := map[string]string{
		"Authorization": c.Sign2Authorization(prefixMt, sign),
	}
	return authHeader
}

func (c *Client) VerifyAuthHeader(headers map[string]string, method, path, query, body string) bool {
	authorization, ok := headers["Authorization"]
	if !ok {
		return false
	}
	fmt.Printf("VerifyAuthHeader: authorization: %v\n", authorization)
	sign := c.Authorization2Sign(prefixMt, authorization)
	return c.VerifySign(sign, headers, method, path, query, body)
}

func (c *Client) Run(method, path, query, body string, headers map[string]string) (int, string) {
	if headers == nil {
		headers = map[string]string{}
	}

	method = strings.ToUpper(method)
	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" {
		panic("Unsupported HTTP Method: " + method)
	}

	if len(query) > 0 {
		path = path + query
	}

	authHeaders := c.GetAuthHeader(headers, method, path, query, body)
	for k, v := range authHeaders {
		headers[k] = v
	}

	httpClient := &http.Client{}

	req, err := http.NewRequest(method, c.Scheme+"://"+c.Host+path, strings.NewReader(body))
	if err != nil {
		panic(err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	respStatusCode := resp.StatusCode
	respData := string(respBody)

	return respStatusCode, respData
}

func (c *Client) Get(path, query string, headers map[string]string) (int, string) {
	return c.Run("GET", path, query, "", headers)
}

func (c *Client) Post(path, query, body string, headers map[string]string) (int, string) {
	return c.Run("POST", path, query, body, headers)
}

func (c *Client) Put(path, query, body string, headers map[string]string) (int, string) {
	return c.Run("Put", path, query, body, headers)
}

func (c *Client) Delete(path, query string, headers map[string]string) (int, string) {
	return c.Run("DELETE", path, query, "", headers)
}
