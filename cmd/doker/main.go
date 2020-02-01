package main

import (
	"image-loader/cmd/doker/images"
	"os"

	"github.com/spf13/cobra"
)

var (
	// RootCmd is the main parent doker cmd
	RootCmd = &cobra.Command{
		Use:   "doker",
		Short: "A tool to get a docker-cli like experience for your kubernetes cluster",
		Long: `doker is a CLI tool to control or interact with the docker daemons in your kubernetes cluster.
It can help you manage images (i.e list, load, prune, etc) and containers on your Kubernetes nodes.`,
	}
)

func init() {
	RootCmd.AddCommand(images.NewImagesCmd())

	// Also add images command alias
	var subAlias = &cobra.Command{
		Use:   "images",
		Short: "alias for image list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return images.NewListCmd().RunE(cmd,args)
		},
	}

	RootCmd.AddCommand(subAlias)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
