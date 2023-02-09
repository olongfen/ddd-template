#!/bin/bash
echo $CONFIG
go run . --config ${CONFIG:-./configs/config.yaml}