package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("/bin/sh")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWUSER,
	}

	must(cmd.Run())
}

func must(err error) {
	panic(err)
}
