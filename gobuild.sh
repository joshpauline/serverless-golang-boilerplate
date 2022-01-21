#!/bin/bash
for r in handlers/*; do
    if [ -d "$r" ]; then
        r=$(basename "$r")
        env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/$r handlers/$r/main.go
    fi
done


