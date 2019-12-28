FROM golang:latest AS build
WORKDIR /app
COPY ./release/app-latest-linux-amd64 .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/app-latest-linux-amd64 .
CMD ["./app-latest-linux-amd64"]