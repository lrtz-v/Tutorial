FROM golang:latest AS build

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
ENV GO111MODULE=on
ENV GOOS=linux

COPY . .
RUN go mod tidy
RUN go build -o app .


FROM alpine:latest
WORKDIR /root/
COPY --from=build /app .
COPY start.sh .
CMD ["sh", "start.sh"]
