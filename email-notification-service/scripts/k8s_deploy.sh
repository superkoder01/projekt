#!/bin/bash

env=''
ns=''

print_usage() {
  printf "Usage: '-r' - prod, '-p' - pre_prod, '-d' - dev"
}

while getopts 'dpr' flag; do
  case "${flag}" in
    d) env='dev' && ns='billing-ovoo' ;;
    p) env='pre_prod' && ns='billing' ;;
    r) env='prod' && ns='billing' ;;
    *) print_usage
       exit 1 ;;
  esac
done

# Create all resources
kubectl apply -k deployments/k8s/overlays/$env

# Delete old pod
kubectl get pod -n $ns | grep email-notification-service | awk '{print $1}' | xargs kubectl delete pod -n $ns