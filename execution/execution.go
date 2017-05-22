package execution

import (
	gocontext "context"
	"fmt"

	"github.com/containerd/console"
	"github.com/containerd/containerd/api/services/execution"
	"github.com/containerd/containerd/api/types/mount"
	protobuf "github.com/gogo/protobuf/types"
	"github.com/kunalkushwaha/go-containerd/pkg"
	"github.com/opencontainers/runtime-spec/specs-go"
)

// ServiceClient exposes all execution APIs
type ServiceClient struct {
	exec execution.ContainerServiceClient
}

//GetExecutionService return localhost ServiceClient for execution
func GetExecutionService(containerdSocket string) (ServiceClient, error) {
	if containerdSocket == "" {
		containerdSocket = "/run/containerd/containerd.sock"
	}
	conn, err := pkg.GetGRPCConnection(containerdSocket)
	if err != nil {
		return ServiceClient{}, err
	}
	exec := ServiceClient{exec: execution.NewContainerServiceClient(conn)}
	return exec, nil
}

//Create container with params provided
func (client *ServiceClient) Create(id string, tty bool, rootfs []*mount.Mount, spec, containerRuntime, stdin, stdout, stderr string) (*execution.CreateResponse, error) {
	//TODO: sanity check for inputs params
	create := &execution.CreateRequest{
		ID: id,
		Spec: &protobuf.Any{
			TypeUrl: specs.Version,
			Value:   []byte(spec),
		},
		Rootfs:   rootfs,
		Runtime:  containerRuntime,
		Terminal: tty,
		Stdin:    stdin,
		Stdout:   stdout,
		Stderr:   stderr,
	}

	if create.Terminal {
		con := console.Current()
		defer con.Reset()
		if err := con.SetRaw(); err != nil {
			return nil, err
		}
	}
	_, err := pkg.PrepareStdio(create.Stdin, create.Stdout, create.Stderr, create.Terminal)
	if err != nil {
		return nil, err
	}

	response, err := client.exec.Create(gocontext.Background(), create)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)
	return response, nil
}

//Start a container with id
func (client *ServiceClient) Start(id string) error {
	_, err := client.exec.Start(gocontext.Background(), &execution.StartRequest{ID: id})
	if err != nil {
		return err
	}
	return nil
}

//Delete a container with id
func (client *ServiceClient) Delete(id string) (*execution.DeleteResponse, error) {
	response, err := client.exec.Delete(gocontext.Background(), &execution.DeleteRequest{ID: id})
	if err != nil {
		return nil, err
	}
	return response, nil
}
