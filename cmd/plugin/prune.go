package main

import (
	"fmt"
	"image-loader/pkg/k8s"

	"github.com/spf13/cobra"
)

func newPruneCmd() *cobra.Command {

	var pruneCmd = &cobra.Command{
		Use:   "prune",
		Short: "prunes the images in all nodes",
		RunE: func(cmd *cobra.Command, args []string) error {

			api, err := k8s.NewKubernetesAPI(kubeconfigPath, kubeContext)
			if err != nil {
				return err
			}
			// send a request to each of them using the Kubernetes API
			pods, err := api.GetPodsWithDefaultLabels()
			if err != nil {
				return err
			}

			for _, pod := range pods.Items {
				fmt.Println(pod.Spec.NodeName)
				resp, err := api.SendPodGetRequest(pod.Name, k8s.ImagesNamespace, "prune")
				if err != nil {
					return err
				}
				fmt.Println(resp)
			}
			return nil
		},
		Example: `
images prune
`,
	}

	return pruneCmd
}
