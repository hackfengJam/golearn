package main

import (
	"context"
	"encoding/json"
)

type Exception error
type TaskRetVal struct {
	IsSuccess bool
	NeedRetry bool
	RetVal    json.RawMessage
}
type Backend interface {
}
type Task interface {
	GetBackend(ctx context.Context) (res Backend, err error)

	Run(ctx context.Context, argsJson json.RawMessage) (res *TaskRetVal, err error)

	OnRetry(ctx context.Context, exc Exception, taskId uint64, argsJson json.RawMessage, eInfo string) (err error)
	OnFail(ctx context.Context, exc Exception, taskId uint64, argsJson json.RawMessage, eInfo string) (err error)
	OnSuccess(ctx context.Context, retVal *TaskRetVal, taskId uint64, argsJson json.RawMessage) (err error)
}

type BaseTask struct {
	backend Backend
}

func (s *BaseTask) GetBackend(ctx context.Context) (res Backend, err error) {
	return s.backend, nil
}

func (s *BaseTask) Run(ctx context.Context, argsJson json.RawMessage) (res *TaskRetVal, err error) {
	return
}

func (s *BaseTask) OnRetry(ctx context.Context, exc Exception, taskId uint64, argsJson json.RawMessage, eInfo string) (err error) {
	// do something
	return
}

func (s *BaseTask) OnFail(ctx context.Context, exc Exception, taskId uint64, argsJson json.RawMessage, eInfo string) (err error) {
	// do something
	return
}

func (s *BaseTask) OnSuccess(ctx context.Context, retVal *TaskRetVal, taskId uint64, argsJson json.RawMessage) (err error) {
	// do something
	return
}

func Run(ctx context.Context, argsJson json.RawMessage) (res *TaskRetVal, err error) {
	return
}

func main() {
	// var x Task
	// x = &BaseTask{}
	// x.Run = Run

}
