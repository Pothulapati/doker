package main

import (
	"encoding/json"
	"fmt"
	"image-loader/pkg/k8s"
	"os"
	"strings"

	"github.com/docker/cli/cli/command/formatter"
	"github.com/docker/cli/opts"

	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
)

type Node struct {
	NodeName string
	Images   []types.ImageSummary
}

type imagesOptions struct {
	matchName string

	quiet       bool
	all         bool
	noTrunc     bool
	showDigests bool
	format      string
	filter      opts.FilterOpt
}

func newListCmd() *cobra.Command {

	options := imagesOptions{filter: opts.NewFilterOpt()}

	// checkCmd represent kubectl pg check.
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "list gets all the images in the nodes present in the cluster.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				options.matchName = args[0]
			}

			api, err := k8s.NewKubernetesAPI(kubeconfigPath, kubeContext)
			if err != nil {
				return err
			}
			nodes, err := GetNodeImages(*api)
			if err != nil {
				return err
			}

			for _, node := range nodes {
				fmt.Printf("\n%s\n\n", node.NodeName)
				imageCtx := formatter.ImageContext{
					Context: formatter.Context{
						Output: os.Stdout,
						Format: formatter.NewImageFormat(options.format, options.quiet, options.showDigests),
						Trunc:  !options.noTrunc,
					},
					Digest: options.showDigests,
				}

				err := formatter.ImageWrite(imageCtx, node.Images)
				if err != nil {
					return err
				}
			}

			return nil
		},
		Example: `
images list
images list -l
`,
	}

	flags := listCmd.Flags()

	flags.BoolVarP(&options.quiet, "quiet", "q", false, "Only show numeric IDs")
	flags.BoolVarP(&options.all, "all", "a", false, "Show all images (default hides intermediate images)")
	flags.BoolVar(&options.noTrunc, "no-trunc", false, "Don't truncate output")
	flags.BoolVar(&options.showDigests, "digests", false, "Show digests")
	flags.StringVar(&options.format, "format", formatter.TableFormatKey, "Pretty-print images using a Go template")
	flags.VarP(&options.filter, "filter", "f", "Filter output based on conditions provided")

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
		resp, err := k8sAPI.SendPodGetRequest(pod.Name, k8s.ImagesNamespace, "list")
		if err != nil {
			return nil, err
		}
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
