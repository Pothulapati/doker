package k8s

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/julienschmidt/httprouter"

	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	KubekerAppLabel      = "app"
	KubekerAppLabelValue = "kubekerd"
	KubekerNamespace     = "default"
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
	pods, err := k.CoreV1().Pods(KubekerNamespace).List(v12.ListOptions{LabelSelector: fmt.Sprintf("%s=%s", KubekerAppLabel, KubekerAppLabelValue)})
	if err != nil {
		return nil, err
	}

	return pods, nil
}

// SendMultiPartHttpRequest Creates a post kubernetes request to the pod, with the given reader
// as a multipart form data
func (k *KubernetesAPI) SendMultiPartHttpRequest(name, namespace, path string, r io.Reader) ([]byte, error) {
	req := k.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(namespace).
		Name(name).
		SubResource("proxy").Suffix(fmt.Sprintf("/%s", path))

	// Push the reader into the http client as a file
	var data bytes.Buffer
	_, err := data.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "docker.tar")
	if err != nil {
		return nil, err
	}
	_, err = part.Write(data.Bytes())
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req.Body(body)
	req.SetHeader("Content-Type", writer.FormDataContentType())

	return req.DoRaw()
}

func (k *KubernetesAPI) SendPodGetRequestWithParams(name, namespace, path string, params httprouter.Params) (string, error) {
	req := k.CoreV1().RESTClient().Get().
		Resource("pods").
		Namespace(namespace).
		Name(name).
		SubResource("proxy").Suffix(fmt.Sprintf("/%s", path))

	// Add all http parameters
	for _, param := range params {
		req = req.Param(param.Key, param.Value)
	}

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
