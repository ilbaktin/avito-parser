#!/usr/bin/env bash
set -ex
GOOS=windows GOARCH=amd64 go build -o main.exe main.go
GOOS=linux GOARCH=amd64 go build -o main_linux main.go
GOOS=darwin GOARCH=amd64 go build -o main_osx main.go

