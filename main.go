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
		// flags for the clone command
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWUSER,

		// this is to break the mountns' link to the initial ns.
		// this prevents the sandbox from populating the initial ns' /proc/mount file
		Unshareflags: syscall.CLONE_NEWNS,

		// mapping uid in the new ns to it's uid in the initial namespace. Unmapped users
		// are assigned the overflow uid (65534) so we're mapping the container id 0 (root) to
		// the corresponding uid in the initial namespace.
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},

		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}

	must(cmd.Run())
}

func must(err error) {
	panic(err)
}
