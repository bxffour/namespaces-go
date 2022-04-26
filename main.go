package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/docker/docker/pkg/reexec"
)

func init() {
	reexec.Register("nsInitialisation", nsInitialisation)
	if reexec.Init() {
		os.Exit(0)
	}
}

func nsInitialisation() {
	newroot := os.Args[1]

	if err := mountProc(newroot); err != nil {
		log.Fatalf("error mounting /proc - %s", err)
	}

	if err := pivot_root(newroot); err != nil {
		log.Fatalf("error running pivot_root - %s", err)
	}

	if err := syscall.Sethostname([]byte("sandbox")); err != nil {
		log.Fatalf("error setting hostname - %s", err)
	}

	nsRun()
}

func nsRun() {
	cmd := exec.Command("/bin/sh")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = []string{`PS1=\w \$ `}

	must(cmd.Run())
}

func main() {
	var rootfsPath = "/tmp/ns-process/rootfs"

	cmd := reexec.Command("nsInitialisation", rootfsPath)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		// flags for the clone command
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET |
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
