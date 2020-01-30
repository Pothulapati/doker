
# Deploy: tell Tilt what YAML to deploy
k8s_yaml('./deploy/manifests.yaml')

# Build: tell Tilt what images to build from which directories
docker_build('tarunpothulapati/image-loader', './')
