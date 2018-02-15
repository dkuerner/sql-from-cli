#!/usr/bin/env bash

export GOARCH=amd64
export GOOS=darwin
go build -o bin/darwin-amd64/sql-cli