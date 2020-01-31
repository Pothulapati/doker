
build: 
		docker build . -t tarunpothulapati/dokerd

push: 
		docker push tarunpothulapati/dokerd

dep:
		kubectl apply -f ./deploy/manifests.yaml

plugin:
		GO111MODULE=on  go build -o doker ./cmd/doker

release:
		# Perform a tag here
		# and find the diff from the last tag, and make it as release notes
		# add the kubectl-images plugin as a artifact
		# also a do a docker push to docker hub & github packages, with the same version
		