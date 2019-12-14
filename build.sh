#!/usr/bin/env bash
# Stops the process if something fails
set -xe

# All of the dependencies needed/fetched for your project.
go get "github.com/jc3m/ridge"
go get "github.com/gorilla/handlers"
go get "github.com/joho/godotenv"

# create the application binary that eb uses
GOOS=linux GOARCH=amd64 go build -o bin/application -ldflags="-s -w"
