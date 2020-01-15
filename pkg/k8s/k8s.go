package k8s

import (
	"fmt"

	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "k8s.io/api/core/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	ImagesLabel      = "app"
	ImagesLabelValue = "images"
	ImagesNamespace  = "default"
)

type KubernetesAPI struct {
	*rest.Config
	kubernetes.Interface
}

func NewKubernetesAPI(kubeConfigPath, kubeContext string) (*KubernetesAPI, error) {
	config, err := GetConfig(kubeConfigPath, kubeContext)
	if err != nil {
		return nil, fmt.Errorf("error configuring Kubernetes API clientset: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error configuring Kubernetes API clientset: %v", err)
	}

	return &KubernetesAPI{
		Config:    config,
		Interface: clientset,
	}, nil
}

func (k *KubernetesAPI) GetPodsWithDefaultLabels() (*v1.PodList, error) {
	pods, err := k.CoreV1().Pods(ImagesNamespace).List(v12.ListOptions{LabelSelector: fmt.Sprintf("%s=%s", ImagesLabel, ImagesLabelValue)})
	if err != nil {
		return nil, err
	}

	return pods, nil
}

func (k *KubernetesAPI) SendPodGetRequest(name, namespace, path string) (string, error) {
	req := k.CoreV1().RESTClient().Get().
		Resource("pods").
		Namespace(namespace).
		Name(name).
		SubResource("proxy").Suffix(fmt.Sprintf("/%s", path))

	res := req.Do()
	raw, err := res.Raw()
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

// GetConfig returns kubernetes config based on the current environment.
// If fpath is provided, loads configuration from that file. Otherwise,
// GetConfig uses default strategy to load configuration from $KUBECONFIG,
// .kube/config, or just returns in-cluster config.
func GetConfig(fpath, kubeContext string) (*rest.Config, error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	if fpath != "" {
		rules.ExplicitPath = fpath
	}
	overrides := &clientcmd.ConfigOverrides{CurrentContext: kubeContext}
	return clientcmd.
		NewNonInteractiveDeferredLoadingClientConfig(rules, overrides).
		ClientConfig()
}
