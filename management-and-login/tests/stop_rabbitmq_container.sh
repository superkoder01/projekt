#!/bin/bash
docker stop $(docker ps --all | grep rabbit | grep -v CONTAINER | awk '{print $1}')
docker container rm $(docker ps --all | grep rabbit | grep -v CONTAINER | awk '{print $1}')
docker volume prune --force