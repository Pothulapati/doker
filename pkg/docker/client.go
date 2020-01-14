package docker

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)


// Returns a json response of all the docker images present
func GetDockerImages(ctx context.Context) ([]byte, error) {

	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	images, err := cli.ImageList(ctx,types.ImageListOptions{All: false})
	if err !=nil{
		return nil, err
	}

	return json.Marshal(images)

}
