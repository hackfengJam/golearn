package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type EmptyResp struct {
}
type Task interface {
	Run(ctx context.Context, argsJson json.RawMessage) (res *EmptyResp, err error)
}
type NotifyTask interface {
	Task
	Start(ctx context.Context, argsJson json.RawMessage) (res *EmptyResp, err error)
}

type callBackTask struct {
	x uint64
}

func (s *callBackTask) Run(ctx context.Context, argsJson json.RawMessage) (res *EmptyResp, err error) {
	s.x = 2
	return
}

func (s *callBackTask) Start(ctx context.Context, argsJson json.RawMessage) (res *EmptyResp, err error) {
	return
}

func F(t Task) {
	_, _ = t.Run(context.Background(), []byte(""))
	fmt.Printf("F(t Task) -> %v\n", t)
	return
}
func A() {
	a := callBackTask{}
	a.x = 1
	fmt.Printf("A() 1-> %v\n", a)
	F(&a)
	fmt.Printf("A() 2-> %v\n", a)
	return
}

type TaskFunC func(ctx context.Context, argsJson json.RawMessage)

func main() {
	A()
}
