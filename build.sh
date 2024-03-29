#!/usr/bin/env bash
mkdir -p bin
GOOS=windows GOARCH=amd64 go build  -ldflags="-w -s" -o bin/eureka-cli-Windows-x86_64.exe
GOOS=linux GOARCH=amd64 go build  -ldflags="-w -s" -o bin/eureka-cli-Linux-x86_64
GOOS=darwin GOARCH=amd64 go build  -ldflags="-w -s" -o bin/eureka-cli-Darwin-x86_64
