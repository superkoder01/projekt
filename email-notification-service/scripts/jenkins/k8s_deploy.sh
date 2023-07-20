#!/bin/bash

env=''
ns=''

print_usage() {
  printf "./k8s_deploy.sh {ENVIRONMENT}\nAccepted param values: dev, pre_prod, prod\n"
}

if [ -z "$1" ]
then
  printf "Not enough arguments passed"
  print_usage
  exit 1
fi

case "$1" in
  dev)
    env='dev' && ns='billing-ovoo' ;;
  pre_prod)
    env='pre_prod' && ns='billing' ;;
  prod)
    env='prod' && ns='billing' ;;
  *)
    print_usage
    exit 1 ;;
esac

# Create all resources
kubectl apply -k deployments/k8s/overlays/$env

# Delete old pod
kubectl get pod -n $ns | grep email-notification-service | awk '{print $1}' | xargs kubectl delete pod -n $ns