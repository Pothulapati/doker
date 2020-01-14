package main

import (
	"fmt"
	"image-loader/pkg/k8s"

	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const ()

func newListCmd() *cobra.Command {
	// checkCmd represent kubectl pg check.
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "list gets all the images in the nodes present in the cluster.",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := k8s.NewKubernetesAPI(kubeconfigPath, kubeContext)
			if err != nil {
				return err
			}
			images, err := listImages(*api)
			if err != nil {
				return err
			}
			jsonImages, err := yaml.Marshal(images)
			if err != nil {
				return err
			}
			fmt.Println("Images:")
			fmt.Print(string(jsonImages))
			return nil
		},
		Example: `
images list
images list -l
`,
	}

	return listCmd
}

// listImages talks to the Kubernetes API and gets all the images from all nodes
func listImages(k8sAPI k8s.KubernetesAPI) ([]types.ImageSummary, error) {

	// Talk to the Kubernetes API to get all endpoints of the service
	// send a request to each of them using the Kubernetes API
	pods, err := k8sAPI.GetPodsWithDefaultLabels()
	if err != nil {
		return nil, err
	}

	for _, pod := range pods.Items {
		resp, err := k8sAPI.SendPodGetRequest(pod.Name, k8s.ImagesNamespace)
		if err != nil {
			return nil, err
		}
		fmt.Println(pod.Name)
		fmt.Println(resp)
	}
	return nil, err
}
