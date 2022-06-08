#!/bin/sh

docker stop web
docker rm web
docker rmi container/web
docker system prune