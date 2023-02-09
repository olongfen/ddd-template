# Stage 1: compile mbtileserver
FROM golang:1.19
ENV CONFIG="./configs/config.yaml"
ENV TZ="Asia/Shanghai"
# 设置时区
RUN ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && dpkg-reconfigure --frontend noninteractive tzdata \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY start.sh /app/start.sh
RUN go env -w GOPROXY="https://goproxy.io,direct"
ENTRYPOINT ["./start.sh"]
#RUN GOOS=linux GOPROXY="https://goproxy.io,direct" go build  -o server ./

