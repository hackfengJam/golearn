package svc

import "golearn/frame/go_zero/greet/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(config config.Config) *ServiceContext {
	return &ServiceContext{Config: config}
}
