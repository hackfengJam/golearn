package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"golearn/frame/go_zero/greet/internal/svc"
	"golearn/frame/go_zero/greet/internal/types"
)

type GreetLogic struct {
	ctx context.Context
	logx.Logger
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GreetLogic {
	return GreetLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
	}
	// TODO need set model here from svc
}

func (l *GreetLogic) Greet(req types.Request) error {
	return nil
}
