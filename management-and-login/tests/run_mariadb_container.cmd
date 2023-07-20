#!/bin/bash
docker run -d -p 33306:3306 -e MARIADB_ROOT_PASSWORD=root mariadb:10.7

while ! nc -z localhost 33306; do
  sleep 0.5 # wait for 1/2 of the second before check again
done

sleep 10

mysql -uroot -proot -h localhost --port=33306 --protocol=TCP < ./tests/data/sql/schema.sql
mysql -uroot -proot -h localhost --port=33306 --protocol=TCP < ./tests/data/sql/data.sql