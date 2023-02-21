#!/bin/sh
#
export DB_NAME=todo
export DB_USER=root
export DB_PASSWORD=root
export DB_HOST=localhost:3306

docker run \
    --name mysql-todo \
    --rm \
    -e MYSQL_DATABASE=$DB_NAME\
    -e MYSQL_ROOT_PASSWORD=$DB_PASSWORD\
    -p 3306:3306 \
    -d mysql:latest

go run .
