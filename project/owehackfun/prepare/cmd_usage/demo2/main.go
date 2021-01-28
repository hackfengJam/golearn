package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		bashPath string

		cmd    *exec.Cmd
		output []byte
		err    error
	)
	bashPath = "/bin/bash"

	// 生成Cmd
	cmd = exec.Command(bashPath, "-c", "sleep 5;ls -l")

	// 执行了命令，捕获了子进程的输出（pipe）
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	// 打印子进程的输出
	fmt.Println(output)
	fmt.Println(string(output))

}
