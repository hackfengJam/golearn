package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"golearn/basic_study/test/http2/server/common"
	config2 "golearn/basic_study/test/http2/server/config"
	"net"
	"net/http"
	"os"
	"time"
)

type Service struct {
	server *http.Server
}

var (
	serviceMgr *Service
)

func InitService() (err error) {
	var (
		mux      *http.ServeMux
		server   *http.Server
		listener net.Listener
	)

	config := config2.GetConfig()

	mux = http.NewServeMux()
	mux.HandleFunc("/push/all", handlePushAll)

	// TLS证书解析验证
	if _, err = tls.LoadX509KeyPair(config.ServerPem, config.ServerKey); err != nil {
		return common.ErrCertInvalid
	}

	// HTTP/2 服务
	server = &http.Server{
		ReadTimeout:  time.Duration(config.ServiceReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(config.ServiceWriteTimeout) * time.Millisecond,
		Handler:      mux,
	}

	port := "8080"
	// 监听端口
	if listener, err = net.Listen("tcp", ":"+port); err != nil {
		return
	}

	// 赋值全局变量
	serviceMgr = &Service{server: server}

	// 拉起服务
	go serviceMgr.server.ServeTLS(listener, config.ServerPem, config.ServerKey)

	return err
}

func main() {
	var (
		err error
	)

	if err = InitService(); err != nil {
		goto ERR
	}

	for {
		time.Sleep(1 * time.Second)
	}
	os.Exit(0)

ERR:
	fmt.Fprintln(os.Stderr, err)
	os.Exit(-1)
}

func handlePushAll(resp http.ResponseWriter, req *http.Request) {
	var (
		err    error
		items  string
		msgArr []json.RawMessage
		msgIdx int
	)
	if err = req.ParseForm(); err != nil {
		return
	}

	items = req.PostForm.Get("items")
	if err = json.Unmarshal([]byte(items), &msgArr); err != nil {
		return
	}

	for msgIdx, _ = range msgArr {
		fmt.Println(fmt.Sprintf("msgIdx: %d, msg: %v", msgIdx, string(msgArr[msgIdx])))
	}
}
