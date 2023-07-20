#!/bin/bash
IMAGE_VERSION=IMAGE_TAG

env=''
registry=''

print_usage() {
  printf "./build_image.sh {ENVIRONMENT}\nAccepted param values: dev, pre_prod, prod\n"
}

if [ -z "$1" ]
then
  printf "Not enough arguments passed"
  print_usage
  exit 1
fi

case "$1" in
  dev)
    env='dev' && registry='10.0.8.1:5000' ;;
  pre_prod)
    env='pre_prod' && registry='10.0.8.1:5000' ;;
  prod)
    env='prod' && registry='g99vzm03.gra7.container-registry.ovh.net' ;;
  *)
    print_usage
    exit 1 ;;
esac

# Copy configs etc
cp config/config-local-email.yaml deployments/docker/email-notification-service.yaml

# Copy binary
cp cmd/email_service/email_service deployments/docker/email_service

# Build docker image
cd deployments/docker >> /dev/null && docker build --network=host -t $registry/c4e/email-notification-service:$env-$IMAGE_VERSION . >> /dev/null && cd - >> /dev/null

#echo $(docker images | awk '{print $3}' | awk 'NR==2')
echo $registry/c4e/email-notification-service:$env-$IMAGE_VERSION