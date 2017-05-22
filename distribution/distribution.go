package distribution

import (
	"fmt"
	"path/filepath"
	"time"

	contentapi "github.com/containerd/containerd/api/services/content"
	imagesapi "github.com/containerd/containerd/api/services/images"
	"github.com/containerd/containerd/content"
	"github.com/containerd/containerd/images"
	"github.com/containerd/containerd/remotes"
	contentservice "github.com/containerd/containerd/services/content"
	imagesservice "github.com/containerd/containerd/services/images"
	"github.com/kunalkushwaha/go-containerd/pkg"
)

// DistServiceClient exposes all execution APIs
type DistServiceClient struct {
	cs content.Store
	rs remotes.Resolver
	is images.Store
}

//GetDistributionService return localhost ServiceClient for execution
func GetDistributionService(containerdSocket string, root string) (DistServiceClient, error) {
	if containerdSocket == "" {
		containerdSocket = "/run/containerd/containerd.sock"
	}

	//TODO:
	// - Get the ContentService.
	if root == "" {
		root = "/var/lib/containerd"
	}
	// Check if rootPath exist
	rootPath := filepath.Join(root, "content")
	if !filepath.IsAbs(rootPath) {
		var err error
		_, err = filepath.Abs(rootPath)
		if err != nil {
			return DistServiceClient{}, err
		}
	}
	conn, err := pkg.ConnectGRPC(containerdSocket, 20*time.Second)
	if err != nil {
		return DistServiceClient{}, err
	}

	cs := contentservice.NewStoreFromClient(contentapi.NewContentClient(conn))

	// - Get ImageStore
	is := imagesservice.NewStoreFromClient(imagesapi.NewImagesClient(conn))

	// - Get Image Reslover
	//	rs := getResolver(ctx, resolverContext)

	dist := DistServiceClient{cs: cs, is: is}
	return dist, nil
}

//Pull the image
func (ds *DistServiceClient) Pull(ref string) error {
	fmt.Println(ref)
	return nil
}

//List images
func (ds *DistServiceClient) List(ref string) ([]images.Image, error) {
	return nil, nil
}

//Delete the image
func (ds *DistServiceClient) Delete(ref string) error {
	return nil
}
