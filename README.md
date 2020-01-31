# Do*k*er

A docker-cli like experience for managing your docker daemons on a kubernetes cluster.

*do**k**er* is a CLI tool that enables the use of common docker-cli commands but on the context of a kubernetes cluster. It allows users to manage and operate both containers and images mainly by connecting to the docker daemons running on each node in the kubernetes cluster. This is done by running a daemonset which acts like an agent that talks to the docker daemon on the node, and performs tasks.

## Features

- Allows users to load images from the local docker daemon to all the nodes in your remote Kubernetes cluster.
- Allows users to list images present across all the nodes in the cluster.
, etc


## Architecture


## Problem

Images not being a first class concept in Kubernetes, dosen't give users a nice way of managing them that are present on your nodes across the cluster. This repo contains a set of tools, to help you manage your images across your nodes, with the user experience of the docker-cli.

Think