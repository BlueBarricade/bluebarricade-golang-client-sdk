SHELL := /bin/bash

OUT := golangapi

ifeq ($(OS),Windows_NT)
	GO_ENV=CGO_ENABLED=0 GOOS=windows
else
	GO_ENV=CGO_ENABLED=0 GOOS=linux
endif

ifeq ($(OS),Windows_NT)
	$(GO_ENV) go build -o ./dist/${OUT}.exe
else
	$(GO_ENV) go build -o ./dist/${OUT}
endif
cp -r ./.env ./dist
