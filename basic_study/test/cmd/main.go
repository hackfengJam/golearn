package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

func Run(arg string) (string, error) {
	goos := runtime.GOOS
	var cmd *exec.Cmd
	switch goos {
	case "darwin", "linux":
		cmd = exec.Command("sh", "-c", arg)
	case "windows":
		cmd = exec.Command("cmd.exe", "/c", arg)
	default:
		return "", fmt.Errorf("unexpected os: %v", goos)
	}
	dtsout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	cmd.Stdout = dtsout
	cmd.Stderr = stderr
	err := cmd.Run()
	if err != nil {
		if stderr.Len() > 0 {
			return "", errors.New(stderr.String())
		}
		return "", err
	}

	return dtsout.String(), nil
}

func t() error {
	sh := fmt.Sprintf(`echo hello > /dev/null `)
	// stdout, err := execx.Run(sh)
	stdout, err := Run(sh)
	if err != nil {
		return err
	}

	if len(stdout) > 0 {
		fmt.Println(stdout)
	}

	return nil
}

func main() {
	err := t()
	if err != nil {
		panic(err)
	}
}
