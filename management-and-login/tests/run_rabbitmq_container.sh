#!/bin/bash
docker run -d -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=pass rabbitmq:3.8
# get container IP
IP=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps | grep rabbitmq:3.8 | awk '{print $1}'))
# Wait for container start
while ! nc -z $IP 5672; do
  sleep 0.5 # wait for 1/2 of the second before check again
done