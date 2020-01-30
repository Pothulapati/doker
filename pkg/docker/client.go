package docker

import (
	"context"
	"encoding/json"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

// ListDockerImages Returns a json response of all the docker images present
func ListDockerImages(ctx context.Context, all bool, filters filters.Args) ([]byte, error) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		return nil, err
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{All: all, Filters: filters})
	if err != nil {
		return nil, err
	}

	return json.Marshal(images)

}

// DockerPruneImages just prunes all the images
func DockerPruneImages(ctx context.Context, pruneFilters filters.Args) ([]byte, error) {

	cli, err := client.NewClientWithOpts()
	if err != nil {
		return nil, err
	}

	report, err := cli.ImagesPrune(ctx, pruneFilters)
	if err != nil {
		return nil, err
	}

	return json.Marshal(report)
}

// GetDockerImages returns an Io Reader for the given image names
func GetDockerImages(ctx context.Context, imagIds []string) (io.ReadCloser, error) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		return nil, err
	}

	rc, err := cli.ImageSave(ctx, imagIds)
	if err != nil {
		return nil, err
	}

	return rc, nil
}

// LoadDockerImage returns an Io Reader for the given imageids
func LoadDockerImage(ctx context.Context, r io.Reader) ([]byte, error) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		return nil, err
	}

	resp, err := cli.ImageLoad(ctx, r, true)
	if err != nil {
		return nil, err
	}

	return json.Marshal(resp)
}
