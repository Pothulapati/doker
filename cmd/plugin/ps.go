package main

import (
	"fmt"
	"image-loader/pkg/k8s"

	"github.com/spf13/cobra"
)

func newPsCmd() *cobra.Command {

	// checkCmd represent kubectl pg check.
	var psCmd = &cobra.Command{
		Use:   "ps",
		Short: "ps lists all containers in a running container",
		RunE: func(cmd *cobra.Command, args []string) error {

			api, err := k8s.NewKubernetesAPI(kubeconfigPath, kubeContext)
			if err != nil {
				return err
			}

			err = GetContainers(api)
			if err != nil {
				return err
			}
			return nil
		},
		Example: `
images ps
images ps -l
`,
	}
	return psCmd
}

// GetNodeImages talks to the Kubernetes API and gets all the images from all nodes
func GetContainers(k8sAPI *k8s.KubernetesAPI) error {
	// Talk to the Kubernetes API to get all endpoints of the service
	// send a request to each of them using the Kubernetes API
	pods, err := k8sAPI.GetPodsWithDefaultLabels()
	if err != nil {
		return err
	}

	for _, pod := range pods.Items {
		fmt.Println(pod.Spec.NodeName)
		resp, err := k8sAPI.SendPodGetRequestWithParams(pod.Name, k8s.ImagesNamespace, "ps", nil)
		if err != nil {
			return err
		}

		fmt.Println(resp)

	}

	return nil
}
