# Kubernetes Images

A docker-cli like experience for managing your images on your kubernetes cluster.

Kubernetes Images is a Kubectl plug-in that allows users to manage docker images that are present across their nodes in a Kubernetes Cluster. This requires a daemonSet to be deployed which acts like an agent that talks to the docker daemon on the node, and performs tasks.

The Kubectl plugin gives you a docker-cli like experience but in the context of a Kubernetes cluster i.e multiple docker daemons. 

## Features

- Allows users to load images from the local docker daemon to all the nodes in your remote Kubernetes cluster.
- Allows users to list images present across all the nodes in the cluster.
- Allows users to prune images in the kubernetes cluster.
- Allows to check if an image present in the cluster.


## Architecture


## Problem

Images not being a first class concept in Kubernetes, dosen't give users a nice way of managing them that are present on your nodes across the cluster. This repo contains a set of tools, to help you manage your images across your nodes, with the user experience of the docker-cli.

Think