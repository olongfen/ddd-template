#!/bin/bash

#生成私钥
openssl genrsa -out server.key 2048

#证书文件
openssl req -nodes -new -x509 -sha256 -days 1825 -config server.conf -extensions 'req_ext' -key server.key -out server.crt

openssl x509 -in server.crt -text