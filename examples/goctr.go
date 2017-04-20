package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/containerd/containerd/api/types/mount"
	"github.com/kunalkushwaha/go-containerd/execution"
)

func getTempDir(id string) (string, error) {
	err := os.MkdirAll(filepath.Join(os.TempDir(), "ctr"), 0700)
	if err != nil {
		return "", err
	}
	tmpDir, err := ioutil.TempDir(filepath.Join(os.TempDir(), "ctr"), fmt.Sprintf("%s-", id))
	if err != nil {
		return "", err
	}
	return tmpDir, nil
}

func main() {
	exec, err := execution.GetExecutionService("")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	cconfig, err := ioutil.ReadFile(filepath.Join("/home/kunal/work/busybox-bundle", "config.json"))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	root := []*mount.Mount{
		{
			Type:    "overlay",
			Source:  "overlay",
			Target:  "/home/kunal/work/aufs-test",
			Options: []string{"lowerdir=/home/kunal/work/busybox-bundle/rootfs,upperdir=/home/kunal/work/goctr-test,workdir=/home/kunal/work/goctr-work"},
		},
	}

	tmpDir, err := getTempDir("goctr")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = exec.Create("goctr", false, root, string(cconfig), "linux", tmpDir+"stdin", tmpDir+"stdout", tmpDir+"stderr")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
