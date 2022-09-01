package main

import (
	"os"
	"path/filepath"
	"syscall"
)

func pivot_root(newroot string) error {
	// pivot_root(2) - putold must be under or underneath newroot.
	putold := filepath.Join(newroot, "/.pivot_root")

	// bind mounting the newroot to itself - this is a slight hack
	// to work around the pivot_root requirement (the newroot must
	// be a path to a mount point `pivot_root(2)`)
	if err := syscall.Mount(
		newroot,
		newroot,
		"",
		syscall.MS_BIND|syscall.MS_REC,
		"",
	); err != nil {
		return err
	}

	// creating putold directory for the old rootdir to be pivoted to.
	// This directory is created at `/.pivot_root` to satisfy the
	// pivot_mount(2) rule that requires the putold directory to be
	// underneath the newroot directory.
	if err := os.MkdirAll(putold, 0700); err != nil {
		return err
	}

	// calling pivot_root
	if err := syscall.PivotRoot(newroot, putold); err != nil {
		return err
	}

	// changing the directory into the new root directory. This is done
	// because pivot_root(2) only changes the root/current workig dirs of
	// each process or thread in the same mountspace to new_root if they
	// point to the old root dir. However, It doesn't change the caller's
	// current working directory (unless it's in the old root dir), thus it
	// should be followed with a chdir("/") call.
	if err := os.Chdir("/"); err != nil {
		return err
	}

	// unmounting putold, which now lives at /.pivot
	putold = "/.pivot_root"
	if err := syscall.Unmount(putold, syscall.MNT_DETACH); err != nil {
		return err
	}

	// removing putold
	if err := os.RemoveAll(putold); err != nil {
		return err
	}

	return nil

}

// mountProc mounts the proc fs in the new root dir
func mountProc(newroot string) error {
	source := "proc"
	target := filepath.Join(newroot, "/proc")
	fstype := "proc"
	flags := 0
	data := ""

	os.MkdirAll(target, 0755)
	if err := syscall.Mount(
		source,
		target,
		fstype,
		uintptr(flags),
		data,
	); err != nil {
		return err
	}

	return nil
}

// mountSys mounts the sysfs in the new root dir
func mountSys(newroot string) error {
	source := ""
	target := filepath.Join(newroot, "/sys")
	fstype := "sysfs"
	flags := syscall.MS_NOSUID | syscall.MS_NOEXEC | syscall.MS_NODEV | syscall.MS_RDONLY
	data := ""

	os.MkdirAll(target, 0755)
	if err := syscall.Mount(
		source,
		target,
		fstype,
		uintptr(flags),
		data,
	); err != nil {
		return err
	}

	return nil
}
