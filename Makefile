
build: 
		docker build . -t tarunpothulapati/image-loader

push: 
		docker push tarunpothulapati/image-loader

dep:
		kubectl apply -f ./deploy/manifests.yaml

plugin:
		GO111MODULE=on  go build -o images ./cmd/plugin