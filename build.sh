#!/bin/sh

GOOS=linux GOARCH=amd64 go build -o bin/wkhtmltopdf-server-linux-amd64

GOOS=darwin GOARCH=amd64 go build -o bin/wkhtmltopdf-server-darwin-amd64

GOOS=windows GOARCH=amd64 go build -o bin/wkhtmltopdf-server-windows-amd64.exe
