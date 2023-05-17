#!/bin/bash

DIR_NAME=$(basename "$PWD")
docker load  -i  ${DIR_NAME}.tar.gz