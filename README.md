## Intro

Converts xray share links to xray configuration json

## Build

```bash

git submodule update --init --recursive
GOOS=linux GOARCH=amd64 go build -o bin/xray-json-generator.amd64
GOOS=linux GOARCH=arm GOARM=7 GOOS=linux go build -o bin/xray-json-generator.arm7

```
