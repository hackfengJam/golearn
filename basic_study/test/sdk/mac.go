package sdk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"io/ioutil"
	"net/http"
)

const maxContentLength = 1024 * 1024

type Mac struct {
	AccessKey string `json:"access_key"`
	SecretKey []byte `json:"secret_key"`
	Host      string `json:"host"`
}

type SignArgs struct {
	Method      string
	Path        string
	Query       string
	ContentType string
	Body        []byte
}

func (m *Mac) SignRequest(sa *SignArgs) (token string, err error) {
	method, path, query, contentType, body := sa.Method, sa.Path, sa.Query, sa.ContentType, sa.Body

	h := hmac.New(sha1.New, m.SecretKey)

	io.WriteString(h, method+" "+path)
	if query != "" {
		io.WriteString(h, "?"+query)
	}
	io.WriteString(h, "\nHost: "+m.Host)

	if contentType != "" {
		io.WriteString(h, "\nContent-Type: "+contentType)
	}

	io.WriteString(h, "\n\n")

	if incBody(body, contentType) {
		h.Write(body)
	}

	sign := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return m.AccessKey + ":" + sign, nil

}

func incBody(body []byte, ctType string) bool {
	typeOk := ctType != "" && ctType != "application/octet-stream"
	lengthOk := len(body) > 0 && len(body) < maxContentLength
	return typeOk && lengthOk
}

type readCloser struct {
	io.Reader
	io.Closer
}

func GetBody(httpReq *http.Request) (body []byte) {
	if !(httpReq.ContentLength > 0 && httpReq.ContentLength < maxContentLength) {
		return
	}
	body, er := ioutil.ReadAll(httpReq.Body)
	if er != nil {
		body = []byte{}
	}
	httpReq.Body = &readCloser{bytes.NewReader(body), httpReq.Body}
	return
}
