// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (

	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/containerd/containerd/api/types/mount"
	"github.com/kunalkushwaha/go-containerd/execution"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates Container",

	Run: func(cmd *cobra.Command, args []string) {
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

		_, err = exec.Create("goctr", false, root, string(cconfig), "linux", tmpDir+"stdin", tmpDir+"stdout", tmpDir+"stderr")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	},
}

func init() {
	RootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}


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
