#!/bin/sh
CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ./build/9rece ./main.go