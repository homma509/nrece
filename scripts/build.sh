#!/bin/sh

# for macos
# CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ./build/nrece ./main.go

# for linux
CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ./build/nrece ./main.go