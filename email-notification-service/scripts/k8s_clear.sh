#!/bin/bash

env=''

print_usage() {
  printf "Usage: '-r' - prod, '-p' - pre_prod, '-d' - dev"
}

while getopts 'dpr' flag; do
  case "${flag}" in
    d) env='dev' ;;
    p) env='pre_prod' ;;
    r) env='prod' ;;
    *) print_usage
       exit 1 ;;
  esac
done

# Create all resources
kubectl delete -k deployments/k8s/overlays/$env





