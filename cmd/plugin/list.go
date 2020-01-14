package main

import (
	"encoding/json"
	"fmt"
	"image-loader/pkg/k8s"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"

	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
)

type Node struct {
	NodeName string
	Images   []types.ImageSummary
}

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
			nodes, err := GetNodeImages(*api)
			if err != nil {
				return err
			}

			for _, node := range nodes {
				println(node.NodeName)
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"REPOSITORY", "IMAGE ID"})
				for _, image := range node.Images {
					table.Append([]string{ArrayToString(image.RepoTags), image.ID[13:24]})
				}
				table.Render() // Send output
			}

			return nil
		},
		Example: `
images list
images list -l
`,
	}

	return listCmd
}

// GetNodeImages talks to the Kubernetes API and gets all the images from all nodes
func GetNodeImages(k8sAPI k8s.KubernetesAPI) ([]Node, error) {
	// Talk to the Kubernetes API to get all endpoints of the service
	// send a request to each of them using the Kubernetes API
	pods, err := k8sAPI.GetPodsWithDefaultLabels()
	if err != nil {
		return nil, err
	}

	var Nodes []Node
	for _, pod := range pods.Items {
		resp, err := k8sAPI.SendPodGetRequest(pod.Name, k8s.ImagesNamespace)
		if err != nil {
			return nil, err
		}
		fmt.Println()
		reader := strings.NewReader(resp)
		var imageSummaries []types.ImageSummary
		err = json.NewDecoder(reader).Decode(&imageSummaries)
		node := Node{
			NodeName: pod.Spec.NodeName,
			Images:   imageSummaries,
		}
		Nodes = append(Nodes, node)
	}
	return Nodes, err
}

func ArrayToString(arr []string) string {

	if len(arr) == 0 {
		return "-"
	}

	return arr[0]
}
