# Stage 1: compile mbtileserver
FROM golang:1.19 as builder

WORKDIR /app
COPY . .

RUN GOOS=linux GOPROXY="https://goproxy.io,direct" go build  -o server ./


# Stage 2: start from a smaller image
FROM  debian:stable-slim
WORKDIR /app
# copy the executable to the empty container
COPY --from=builder /app/server /app/server

