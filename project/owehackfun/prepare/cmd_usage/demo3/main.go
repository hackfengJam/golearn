package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	output []byte
	err    error
}

func main() {
	// 执行一个cmd，让它在一个协程里去执行，让它执行2秒：sleep 2;echo hello;

	// 1秒的时候，杀死cmd
	var (
		bashPath string

		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *result
		res        *result
	)
	bashPath = "/bin/bash"

	// 创建一个结果队列
	resultChan = make(chan *result, 1000)

	// content : chan byte
	// cancelFunc :  close(chan byte)

	ctx, cancelFunc = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err    error
		)

		cmd = exec.CommandContext(ctx, bashPath, "-c", "sleep 5;ls -l")
		// select { case <- ctx.Done(): }
		// kill pid 进程ID，杀死子进程

		// 执行任务，捕获输出
		output, err = cmd.CombinedOutput()

		// 把任务输出结果，传给main 协程
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()

	// 继续往下走
	time.Sleep(1 * time.Second)
	//time.Sleep(1 * time.Millisecond)

	// 取消上下文
	cancelFunc()

	// 在main 协程里，等待子协程的退出，并打印任务执行结果
	res = <-resultChan

	// 打印任务执行结果
	fmt.Println(res.err, string(res.output))
}
