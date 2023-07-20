#!/bin/bash
docker run -d -e MARIADB_ROOT_PASSWORD=root mariadb:10.7

# get container IP
IP=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps | grep mariadb:10.7 | awk '{print $1}'))
# Wait for container start
while ! nc -z $IP 3306; do
  sleep 0.5 # wait for 1/2 of the second before check again
done

mysql -uroot -proot -h $IP < ./tests/data/sql/schema.sql
mysql -uroot -proot -h $IP < ./tests/data/sql/data.sql