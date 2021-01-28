// DO NOT EDIT, generated by goctl
package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest"
	"golearn/frame/go_zero/greet/internal/svc"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/greet/from/:name",
			Handler: greetHandler(serverCtx),
		},
	})
}
