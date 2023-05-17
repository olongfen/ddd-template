#!/bin/bash

DIR_NAME=$(basename "$PWD")

docker build -f Dockerfile ./ -t  ${DIR_NAME}/server