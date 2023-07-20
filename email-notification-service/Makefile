SHELL := /bin/bash
.SILENT:

run:

# App binary file build
app_build:
	echo "Building app..." && \
	./scripts/build.sh && \
	echo "App binary file built!"

# Build dev docker image
docker_build_dev:
	echo "Building dev docker image..." && \
	./scripts/build_image.sh -d && \
	echo "Docker image built!"

# Push dev docker image to registry
docker_push_dev:
	echo "Pushing dev docker image to registry..." && \
	./scripts/push_image.sh -d && \
	echo "Docker image pushed!"

# Build pre_prod docker image
docker_build_pre_prod:
	echo "Building pre_prod docker image..." && \
	./scripts/build_image.sh -p && \
	echo "Docker image built!"

# Push pre_prod docker image to registry
docker_push_pre_prod:
	echo "Pushing pre_prod docker image to registry..." && \
	./scripts/push_image.sh -p && \
	echo "Docker image pushed!"

# Build prod docker image
docker_build_prod:
	echo "Building prod docker image..." && \
	./scripts/build_image.sh -r && \
	echo "Docker image built!"

# Push prod docker image to registry
docker_push_prod:
	echo "Pushing prod docker image to registry..." && \
	./scripts/push_image.sh -r && \
	echo "Docker image pushed!"

# Deploy to dev k8s cluster
k8s_deploy_dev:
	echo "Deploying resources to dev k8s cluster..." && \
	./scripts/k8s_deploy.sh -d && \
	echo "Resources deployed to dev k8s cluster!"

# Remove objects from dev k8s cluster
k8s_clear_dev:
	echo "Removing resources from dev k8s cluster..." && \
	./scripts/k8s_clear.sh -d && \
	echo "Resources removed from dev k8s cluster!"

k8s_deploy_pre_prod:
	echo "Deploying resources to pre_prod k8s cluster..." && \
	./scripts/k8s_deploy.sh -p && \
	echo "Resources deployed to pre_prod k8s cluster!"

# Remove objects from pre_prod k8s cluster
k8s_clear_pre_prod:
	echo "Removing resources from pre_prod k8s cluster..." && \
	./scripts/k8s_clear.sh -p && \
	echo "Resources removed from pre_prod k8s cluster!"

k8s_deploy_prod:
	echo "Deploying resources to prod k8s cluster..." && \
	./scripts/k8s_deploy.sh -r && \
	echo "Resources deployed to prod k8s cluster!"

# Remove objects from prod k8s cluster
k8s_clear_prod:
	echo "Removing resources from prod k8s cluster..." && \
	./scripts/k8s_clear.sh -r && \
	echo "Resources removed from prod k8s cluster!"

# ------------- dev ------------------
build: app_build docker_build_dev docker_push_dev

install: build k8s_deploy_dev

clean: k8s_clear_dev

# ------------- pre_prod ------------------
build_pre_prod: app_build docker_build_pre_prod docker_push_pre_prod

install_pre_prod: build_pre_prod k8s_deploy_pre_prod

clean_pre_prod: k8s_clear_pre_prod

# ------------- prod ------------------
build_prod: app_build docker_build_prod docker_push_prod

install_prod: build_prod k8s_deploy_prod

clean_prod: k8s_clear_prod