### go-containerd
This is go bindings for [containerd GRPC APIs](https://github.com/containerd/containerd/tree/master/api), intended for building integration testing.

``` go
import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/containerd/containerd/api/types/mount"
	"github.com/kunalkushwaha/go-containerd/execution"
)

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
```


How to build example code.
```
$ go get github.com/kunalkushwaha/go-containerd
$ vndr
$ make
$ sudo ../bin/goctr
```

__Status__: Work In Progress.
- [ ]execution
 - [ ]Run ?
 - [x]Create
 - [x]Start
 - [x]Delete
 - [ ]Info
 - [ ]List
 - [ ]Kill
 - [ ]Events
 - [ ]Exec
 - [ ]Pty
 - [ ]CloseStdin
- [ ]images
- [ ]content
- [ ]shim
- [ ]rootfs

- _PRs & Issues are most welcome._
