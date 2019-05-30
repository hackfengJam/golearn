package master

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

// 任务的 HTTP 接口
type ApiServer struct {
	httpServer *http.Server
}

var (
	// 单例对象
	G_apiServer *ApiServer
)

func handleJobSave(w http.ResponseWriter, r *http.Request) {
	// 任务保存到 ETCD 中
}

// 初始化服务
func InitApiServer() (err error) {

	var (
		mux        *http.ServeMux
		listener   net.Listener
		httpServer *http.Server
	)

	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("job/save", handleJobSave)

	// 启动 TCP 监听
	if listener, err = net.Listen("tcp", ":"+strconv.Itoa(G_config.ApiPort)); err != nil {
		return
	}

	// 创建一个 HTTP 服务
	httpServer = &http.Server{
		ReadTimeout:  time.Duration(G_config.ApiReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(G_config.ApiWriteTimeout) * time.Millisecond,
		Handler:      mux,
	}

	// 赋值单例
	G_apiServer = &ApiServer{
		httpServer: httpServer,
	}

	// 启动了服务端
	go httpServer.Serve(listener)

	return
}
