#!/bin/sh


RED='\033[0;34m'
NC='\033[0m'

print_red() {
	echo "${RED}$1${NC}" >&2
}

print_red 'Deleting the cluster'
gcloud container clusters delete images --quiet --async