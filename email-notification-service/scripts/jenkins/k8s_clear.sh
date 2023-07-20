#!/bin/bash

env=''

print_usage() {
  printf "./k8s_clear.sh {ENVIRONMENT}\nAccepted param values: dev, pre_prod, prod\n"
}

if [ -z "$1" ]
then
  printf "Not enough arguments passed"
  print_usage
  exit 1
fi

case "$1" in
  dev)
    env='dev' ;;
  pre_prod)
    env='pre_prod' ;;
  prod)
    env='prod' ;;
  *)
    print_usage
    exit 1 ;;
esac

# Create all resources
kubectl delete -k deployments/k8s/overlays/$env