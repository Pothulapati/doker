package images

import "github.com/spf13/cobra"

var (
	apiAddr        string // An empty value means "use the Kubernetes configuration"
	kubeconfigPath string
	kubeContext    string
	impersonate    string
	verbose        bool
)

func NewImagesCmd() *cobra.Command {

	imagesCmd := &cobra.Command{
		Use:   "image",
		Short: "Manage Images in your kubernetes cluster.",
	}

	imagesCmd.AddCommand(NewLoadCmd())
	imagesCmd.AddCommand(newPruneCmd())
	imagesCmd.AddCommand(NewListCmd())

	return imagesCmd

}
