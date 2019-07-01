package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd      *exec.Cmd
		err      error
		bashPath string
	)
	bashPath = "/bin/bash"

	cmd = exec.Command(bashPath, "-c", "echo 1")

	err = cmd.Run()

	fmt.Println(err)
}
