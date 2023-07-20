#!/bin/bash
docker run -d -e PASSWORD=redis 10.0.8.1:5000/redis-env
# get container IP
IP=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps | grep redis | awk '{print $1}'))
# Wait for container start
while ! nc -z $IP 6379; do
  sleep 0.5 # wait for 1/2 of the second before check again
done