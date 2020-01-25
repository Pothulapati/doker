RED='\033[0;34m'
NC='\033[0m'


print_red() {
	echo "${RED}$1${NC}" >&2
}

print_red 'Creating a cluster'
gcloud config set compute/zone us-west1-a
gcloud container clusters create images
gcloud container clusters get-credentials images
