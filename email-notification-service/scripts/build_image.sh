#!/bin/bash
IMAGE_VERSION=1.1.0

env=''
registry=''

print_usage() {
  printf "Usage: '-r' - prod, '-p' - pre_prod, '-d' - dev"
}

while getopts 'dpr' flag; do
  case "${flag}" in
    d) env='dev' && registry='10.0.8.1:5000' ;;
    p) env='pre_prod' && registry='10.0.8.1:5000' ;;
    r) env='prod' && registry='g99vzm03.gra7.container-registry.ovh.net' ;;
    *) print_usage
       exit 1 ;;
  esac
done

# Copy configs etc
cp config/config-local-email.yaml deployments/docker/email-notification-service.yaml

# Copy binary
cp cmd/email_service/email_service deployments/docker/email_service

# Build docker image
cd deployments/docker && docker build --network=host -t $registry/c4e/email-notification-service:$env-$IMAGE_VERSION . && cd -

# Remove temporary copied files
rm deployments/docker/email-notification-service.yaml && \
 rm deployments/docker/email_service
