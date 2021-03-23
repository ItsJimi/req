#!/usr/bin/env bash

# MacOS
GOOS=darwin GOARCH=amd64 go build -o req-macos main.go
# Linux 32 bit
GOOS=linux GOARCH=386 go build -o req-linux32 main.go
# Linux 64 bit
GOOS=linux GOARCH=amd64 go build -o req-linux64 main.go