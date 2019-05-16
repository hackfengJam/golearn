package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
	)

	cmd = exec.Command("E:\\ProgramFiles\\Git\\bin\\bash.exe", "-c", "echo 1")

	err = cmd.Run()

	fmt.Println(err)
}
