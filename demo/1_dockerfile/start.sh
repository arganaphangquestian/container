#!/bin/sh

docker stop web
docker rm web
docker build -t container/web .
docker run -d --name web -p 8000:80 container/web