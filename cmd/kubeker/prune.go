package main

import (
	"encoding/json"
	"fmt"
	"image-loader/pkg/k8s"
	"os"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"

	"github.com/docker/cli/cli/command"

	"github.com/julienschmidt/httprouter"

	"github.com/docker/cli/opts"
	"github.com/spf13/cobra"
)

const (
	allImageWarning = `WARNING! This will remove all images without at least one container associated to them.
Are you sure you want to continue?`
	danglingWarning = `WARNING! This will remove all dangling images.
Are you sure you want to continue?`
)

type pruneOptions struct {
	force  bool
	all    bool
	filter opts.FilterOpt
}

func newPruneCmd() *cobra.Command {

	var options pruneOptions

	var pruneCmd = &cobra.Command{
		Use:   "prune",
		Short: "prunes the images in all nodes",
		RunE: func(cmd *cobra.Command, args []string) error {

			warning := danglingWarning
			if options.all {
				warning = allImageWarning
			}

			if !options.force && !command.PromptForConfirmation(os.Stdin, os.Stdout, warning) {
				return nil
			}

			api, err := k8s.NewKubernetesAPI(kubeconfigPath, kubeContext)
			if err != nil {
				return err
			}
			// send a request to each of them using the Kubernetes API
			pods, err := api.GetPodsWithDefaultLabels()
			if err != nil {
				return err
			}

			// Decode all and filters params into http params
			var params httprouter.Params
			allParam := httprouter.Param{Key: "all", Value: strconv.FormatBool(options.all)}
			params = append(params, allParam)

			forceParam := httprouter.Param{Key: "force", Value: strconv.FormatBool(true)}
			params = append(params, forceParam)

			for _, pod := range pods.Items {
				fmt.Println(pod.Spec.NodeName)
				resp, err := api.SendPodGetRequestWithParams(pod.Name, k8s.KubekerNamespace, "prune", params)
				if err != nil {
					return err
				}
				// Conver the resposne it to pruneReport format
				reader := strings.NewReader(resp)
				var report types.ImagesPruneReport
				err = json.NewDecoder(reader).Decode(&report)
				if err != nil {
					return err
				}

				if len(report.ImagesDeleted) > 0 {
					var sb strings.Builder
					sb.WriteString("Deleted Images:\n")
					for _, st := range report.ImagesDeleted {
						if st.Untagged != "" {
							sb.WriteString("untagged: ")
							sb.WriteString(st.Untagged)
							sb.WriteByte('\n')
						} else {
							sb.WriteString("deleted: ")
							sb.WriteString(st.Deleted)
							sb.WriteByte('\n')
						}
					}
					fmt.Println(sb.String())
				}

			}
			return nil
		},
		Example: `
images prune
`,
	}

	flags := pruneCmd.Flags()
	flags.BoolVarP(&options.force, "force", "f", false, "Do not prompt for confirmation")
	flags.BoolVarP(&options.all, "all", "a", false, "Remove all unused images, not just dangling ones")
	flags.Var(&options.filter, "filter", "Provide filter values (e.g. 'until=<timestamp>')")

	return pruneCmd
}
