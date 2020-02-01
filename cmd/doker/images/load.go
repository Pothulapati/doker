package images

import (
	"context"
	"fmt"
	"image-loader/pkg/docker"
	"image-loader/pkg/k8s"
	"strings"

	"github.com/spf13/cobra"
)

func NewLoadCmd() *cobra.Command {

	// checkCmd represent kubectl pg check.
	var loadCmd = &cobra.Command{
		Use:   "load",
		Short: "loads an image onto your kubernetes cluster",
		RunE: func(cmd *cobra.Command, args []string) error {

			api, err := k8s.NewKubernetesAPI(kubeconfigPath, kubeContext)
			if err != nil {
				return err
			}

			pods, err := api.GetPodsWithDefaultLabels()
			if err != nil {
				return err
			}

			for _, pod := range pods.Items {

				fmt.Println("Loading images to ", strings.Trim(pod.Spec.NodeName, " "))
				// Talk to the docker client and save the image into a file
				rc, err := docker.GetDockerImages(context.Background(), args)
				if err != nil {
					return err
				}

				rs, err := api.SendMultiPartHttpRequest(pod.Name, k8s.DokerNamespace, "load", rc)
				if err != nil {
					return err
				}

				fmt.Println(string(rs))

			}

			return nil
		},
		Example: `
images load hello-world
images load hello-world-1 hello-world-2"
`,
	}

	return loadCmd
}
