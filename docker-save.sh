#!/bin/bash

DIR_NAME=$(basename "$PWD")

docker save ${DIR_NAME}/server -o ${DIR_NAME}.tar.gz