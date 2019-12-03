#!/bin/bash


cd ./go.yyue.dev

# GOOS=linux go build -a -o ../dict/go-apps ./...


cd ../dict

scp -r -P 28013 . root@162.219.127.105:~/workspace