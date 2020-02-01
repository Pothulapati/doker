# Do*k*er

*do**k**er* is a CLI tool that enables the use of common docker-cli commands but on the context of a kubernetes cluster. It allows users to manage and operate both containers and images mainly by connecting to the docker daemons running on each node in the kubernetes cluster. This is done by running a daemonset which acts like an agent that talks to the docker daemon on the node, and performs tasks.

This would require the dockerd to be present on the node, whose sock is mounted to the daemonset allowing it to interact.

The inital focus of this tool is to help you manage images in your cluster as there is no concrete kubernetes primitive for images. But all the commands from the `docker-cli` will be replicated here.

## Installation

First, The Daemonset can be installed on the cluster by running

```bash
kubectl apply -f https://raw.githubusercontent.com/Pothulapati/doker/master/deploy/manifests.yaml
```

Now, *do**k**er* (this project's cli tool) can be installed by running

```bash
 curl -sL https://raw.githubusercontent.com/Pothulapati/doker/master/bin/install.sh | sh
 export PATH=$PATH:/home/tarun/.doker/bin
```

## Demo
This requires the `kubeConfig` file to be present and set to the current cluster.

Now that you have installed, you can list the images in your cluster by runnning
```bash
tarun@tarun-Inspiron-5559:work/doker ‹master*›$ doker images

gke-images-default-pool-aafc8b9e-lqkd

REPOSITORY                                            TAG                 IMAGE ID            CREATED             SIZE
tarunpothulapati/dokerd                               latest              5f96e6ae49f2        4 hours ago         1.65GB
k8s.gcr.io/kube-proxy                                 v1.13.11-gke.23     f2cfe256a767        2 weeks ago         83.7MB
k8s.gcr.io/prometheus-to-sd                           v0.8.2              3df88c7a6ea8        3 months ago        37.9MB
k8s.gcr.io/k8s-dns-sidecar-amd64                      1.15.4              928a271628ea        8 months ago        79.3MB
k8s.gcr.io/k8s-dns-kube-dns-amd64                     1.15.4              305ac63fd465        8 months ago        87MB
k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64                1.15.4              d16f387fc6c5        8 months ago        77.8MB
gcr.io/stackdriver-agents/stackdriver-logging-agent   1.6.8               4270b0a6eeb1        9 months ago        265MB
k8s.gcr.io/prometheus-to-sd                           v0.5.0              42e4387da83f        11 months ago       41.9MB
k8s.gcr.io/prometheus-to-sd                           v0.4.2              626698890281        11 months ago       41.9MB
k8s.gcr.io/addon-resizer                              1.8.4               5ec630648120        15 months ago       38.3MB
k8s.gcr.io/metrics-server-amd64                       v0.3.1              61a0c90da56e        16 months ago       40.8MB
k8s.gcr.io/pause                                      3.1                 da86e6ba6ca1        2 years ago         742kB

gke-images-default-pool-aafc8b9e-tj3m

REPOSITORY                                            TAG                 IMAGE ID            CREATED             SIZE
tarunpothulapati/dokerd                               latest              5f96e6ae49f2        4 hours ago         1.65GB
k8s.gcr.io/kube-proxy                                 v1.13.11-gke.23     f2cfe256a767        2 weeks ago         83.7MB
k8s.gcr.io/prometheus-to-sd                           v0.8.2              3df88c7a6ea8        3 months ago        37.9MB
k8s.gcr.io/k8s-dns-sidecar-amd64                      1.15.4              928a271628ea        8 months ago        79.3MB
k8s.gcr.io/k8s-dns-kube-dns-amd64                     1.15.4              305ac63fd465        8 months ago        87MB
k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64                1.15.4              d16f387fc6c5        8 months ago        77.8MB
gcr.io/stackdriver-agents/stackdriver-logging-agent   1.6.8               4270b0a6eeb1        9 months ago        265MB
k8s.gcr.io/prometheus-to-sd                           v0.5.0              42e4387da83f        11 months ago       41.9MB
k8s.gcr.io/event-exporter                             v0.2.4              e7c317b95d73        11 months ago       47.3MB
k8s.gcr.io/prometheus-to-sd                           v0.4.2              626698890281        11 months ago       41.9MB
k8s.gcr.io/defaultbackend-amd64                       1.5                 b5af743e5984        16 months ago       5.13MB
k8s.gcr.io/pause                                      3.1                 da86e6ba6ca1        2 years ago         742kB

gke-images-default-pool-aafc8b9e-3xhn

REPOSITORY                                            TAG                 IMAGE ID            CREATED             SIZE
tarunpothulapati/dokerd                               latest              5f96e6ae49f2        4 hours ago         1.65GB
k8s.gcr.io/kube-proxy                                 v1.13.11-gke.23     f2cfe256a767        2 weeks ago         83.7MB
k8s.gcr.io/prometheus-to-sd                           v0.8.2              3df88c7a6ea8        3 months ago        37.9MB
gke.gcr.io/heapster                                   v1.7.0              c153bcbe94af        5 months ago        84MB
gcr.io/stackdriver-agents/stackdriver-logging-agent   1.6.8               4270b0a6eeb1        9 months ago        265MB
k8s.gcr.io/fluentd-gcp-scaler                         0.5.2               3dfc22ad2d25        9 months ago        90.5MB
k8s.gcr.io/prometheus-to-sd                           v0.5.0              42e4387da83f        11 months ago       41.9MB
k8s.gcr.io/cluster-proportional-autoscaler-amd64      1.3.0               33813c948942        16 months ago       45.8MB
k8s.gcr.io/addon-resizer                              1.8.3               b57c00a12f6c        18 months ago       33.1MB
k8s.gcr.io/pause                                      3.1                 da86e6ba6ca1        2 years ago         742kB
```
As you can see all the images categorized by the node name are printed out.

Now, Let's try loading hello-world docker image into your kubernetes nodes.
```bash
tarun@tarun-Inspiron-5559:work/doker ‹master*›$ doker load hello-world
Loading images to  gke-images-default-pool-aafc8b9e-lqkd
Loading images to  gke-images-default-pool-aafc8b9e-tj3m
Loading images to  gke-images-default-pool-aafc8b9e-3xhn
```
Now,
You should be able to see the `hello-world` image in your kubernetes nodes.
```bash
tarun@tarun-Inspiron-5559:work/doker ‹master*›$ doker images | grep hello-world
hello-world                                           latest              fce289e99eb9        13 months ago       1.84kB
hello-world                                           latest              fce289e99eb9        13 months ago       1.84kB
hello-world                                           latest              fce289e99eb9        13 months ago       1.84kB
```

## Components
- **Dokerd**: This is the daemon that is deployed as a web-server and listens for HTTP requests from the cli.
- **doker**: This is the docker like cli, which talks to the dokerd agents through the kubenetes proxy sub-resources and passes commands.

## TODO
- Add more container management commands.
- Support Labeling of Nodes.
- Write tests for dokerd client.
- Support for more larger files.
