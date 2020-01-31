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
		Use:   "kubeker",
		Short: "A tool to get a docker-cli like experience for your kubernetes cluster",
		Long: `kubeker is a CLI tool to control or interact with the docker daemons in your kubernetes cluster.
It can help you manage images (i.e list, load, prune, etc) and containers on your Kubernetes nodes.`,
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
	RootCmd.AddCommand(newLoadCmd())
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
