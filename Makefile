
build: 
		docker build . -t tarunpothulapati/image-loader

push: 
		docker push tarunpothulapati/image-loader

dep:
		kubectl apply -f ./deploy/manifests.yaml