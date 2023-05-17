#!/bin/bash
DEPLOYBUNDLE=deployment

echo 'build all ...'
# build
make swag && make
DIR_NAME=$(basename "$PWD")


# bundles
mkdir -p ${DEPLOYBUNDLE}/${DIR_NAME}

#
cp -r bin ${DEPLOYBUNDLE}/${DIR_NAME}
cp -r config ${DEPLOYBUNDLE}/${DIR_NAME}
cp docker-build.sh ${DEPLOYBUNDLE}/${DIR_NAME}
cp docker-save.sh ${DEPLOYBUNDLE}/${DIR_NAME}
cp docker-load.sh ${DEPLOYBUNDLE}/${DIR_NAME}
cp start.sh ${DEPLOYBUNDLE}/${DIR_NAME}
cp stop.sh ${DEPLOYBUNDLE}/${DIR_NAME}
cp Dockerfile ${DEPLOYBUNDLE}/${DIR_NAME}
cp docker-compose.* ${DEPLOYBUNDLE}/${DIR_NAME}
cp entrypoint.sh ${DEPLOYBUNDLE}/${DIR_NAME}
cp .env ${DEPLOYBUNDLE}/${DIR_NAME}

# docker build image
sh docker-build.sh
# save image
docker save ${DIR_NAME}/server -o ${DEPLOYBUNDLE}/${DIR_NAME}/${DIR_NAME}.tar.gz