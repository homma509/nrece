FROM golang:1.15.6-alpine3.12 AS build-env
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN apk add --no-cache --virtual .go-builddeps git gcc make build-base alpine-sdk \
    && go mod download \
    && go get golang.org/x/lint/golint \
    && make go-lint \
    && make go-build

FROM alpine:3.12.3
RUN mkdir /app \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
WORKDIR /app
COPY --from=build-env /app/build/9rece /app/
EXPOSE 80
ENTRYPOINT ["/app/9rece"]