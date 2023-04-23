## Stage 1: compile mbtileserver
#ARG GO_VERSION=1.19
#FROM golang:${GO_VERSION}
#ENV TZ="Asia/Shanghai"
## 设置时区
#RUN ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
#    && echo ${TZ} > /etc/timezone \
#    && dpkg-reconfigure --frontend noninteractive tzdata \
#    && rm -rf /var/lib/apt/lists/*
#WORKDIR /app
#COPY start.sh /app/start.sh
#RUN go env -w GOPROXY="https://goproxy.io,direct"
#RUN chmod +x ./start.sh
#ENTRYPOINT ["./start.sh"]
##RUN GOOS=linux GOPROXY="https://goproxy.io,direct" go build  -o server ./

FROM ubuntu:22.04

ENV TZ="Asia/Shanghai" \
    LANG="en_US.utf8" \
    LC_ALL="en_US.utf8"
RUN echo 'LANG="en_US.utf8"' > /etc/locale.conf
WORKDIR /app
COPY bin/server bin/server
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]