#!/bin/sh

set -e

# setup
docker-compose stop mysql mysql_test
docker-compose rm -vf mysql mysql_test
docker-compose up --no-start mysql mysql_test
docker-compose start mysql mysql_test
docker-compose exec mysql bash -c "until mysqladmin ping -u root -p12345678; do sleep 3; done; echo 'done'"

# migrate
docker-compose run --rm user_api go run ./hack/database-seeds/main.go

# finished
docker-compose stop mysql mysql_test
