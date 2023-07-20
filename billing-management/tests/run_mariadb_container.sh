#!/bin/bash
docker run -d -e MARIADB_ROOT_PASSWORD=root mariadb:10.7
# Wait for container start
while ! nc -z 172.17.0.2 3306; do
  sleep 0.5 # wait for 1/2 of the second before check again
done

mysql -uroot -proot -h 172.17.0.2 < ./data/sql/schema.sql
mysql -uroot -proot -h 172.17.0.2 < ./data/sql/data.sql