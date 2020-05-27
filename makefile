build:
	GO_ENV=CGO_ENABLED=0 GOOS=linux go build -o ./dist/golangapi main.go

run:
	go run main.go

all: build