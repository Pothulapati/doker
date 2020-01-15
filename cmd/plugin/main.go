package main

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	apiAddr        string // An empty value means "use the Kubernetes configuration"
	kubeconfigPath string
	kubeContext    string
	impersonate    string
	verbose        bool

	RootCmd = &cobra.Command{
		Use:   "images",
		Short: "A tool to manage Container images in a kubernetes cluster",
		Long: `images is a CLI tool to manage docker images in your kubernetes cluster.
It can help you in retrieving, removing, loading container images into your kubernetes node.`,
	}
)

func init() {
	RootCmd.PersistentFlags().StringVar(&kubeconfigPath, "kubeconfig", "", "Path to the kubeconfig file to use for CLI requests")
	RootCmd.PersistentFlags().StringVar(&kubeContext, "context", "", "Name of the kubeconfig context to use")
	RootCmd.PersistentFlags().StringVar(&impersonate, "as", "", "Username to impersonate for Kubernetes operations")
	RootCmd.PersistentFlags().StringVar(&apiAddr, "api-addr", "", "Override kubeconfig and communicate directly with the control plane at host:port (mostly for testing)")
	RootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Turn on debug logging")
	RootCmd.AddCommand(newListCmd())
	RootCmd.AddCommand(newPruneCmd())
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
