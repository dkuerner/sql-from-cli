#!/usr/bin/env bash

export GOARCH=amd64
export GOOS=windows
go build -o bin/windows-amd64/sql-cli.exe