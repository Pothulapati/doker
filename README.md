# Do*k*er

*do**k**er* is a CLI tool that enables the use of common docker-cli commands but on the context of a kubernetes cluster. It allows users to manage and operate both containers and images mainly by connecting to the docker daemons running on each node in the kubernetes cluster. This is done by running a daemonset which acts like an agent that talks to the docker daemon on the node, and performs tasks.

This would require the dockerd to be present on the node, whose sock is mounted to the daemonset allowing it to interact.

The inital focus of this tool is to help you manage images in your cluster as there is no concrete kubernetes primitive for images. But all the commands from the `docker-cli` will be replicated here.

## Demo

First, The Daemonset can be installed on the cluster by running

```bash
 curl https://raw.githubusercontent.com/Pothulapati/doker/master/deploy/manifests.yaml | kubectl apply -f -
```

Now, *do**k**er* (this project's cli tool) can be installed by running

```bash

```


## Architecture



## TODO

