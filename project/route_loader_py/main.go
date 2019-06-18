package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"strconv"
	"time"
)

func templateHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/test.html")
	fmt.Println(t.Name())
	t.Execute(w, "Hello world")
}

func main() {
	var (
		mux        *http.ServeMux
		listener   net.Listener
		httpServer *http.Server
		//staticDir     http.Dir     // 静态文件根目录
		//staticHandler http.Handler // 静态文件的 HTTP 回调
		err error
	)

	mux = http.NewServeMux()
	mux.HandleFunc("/index", templateHandler)

	// 静态文件目录
	//staticDir = http.Dir("./templates")
	//staticHandler = http.FileServer(staticDir)
	//mux.Handle("/", http.StripPrefix("/", staticHandler))

	// 启动 TCP 监听
	if listener, err = net.Listen("tcp", ":"+strconv.Itoa(8080)); err != nil {
		return
	}

	// 创建一个 HTTP 服务
	httpServer = &http.Server{
		ReadTimeout:  time.Duration(5000) * time.Millisecond,
		WriteTimeout: time.Duration(5000) * time.Millisecond,
		Handler:      mux,
	}

	httpServer.Serve(listener)
}
