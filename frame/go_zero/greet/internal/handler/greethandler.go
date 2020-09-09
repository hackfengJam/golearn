package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"golearn/frame/go_zero/greet/internal/logic"
	"golearn/frame/go_zero/greet/internal/svc"
	"golearn/frame/go_zero/greet/internal/types"
)

func greetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGreetLogic(r.Context(), ctx)
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		err := l.Greet(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
