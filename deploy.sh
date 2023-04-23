#!/bin/bash

echo 'build all ...'
# build
make swag && make gqlgen && make

# docker build image
sh docker-build.sh

# docker deploy server
docker-compose up -d