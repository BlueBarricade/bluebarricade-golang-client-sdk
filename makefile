build:
	go build -o ./dist/golangapi main.go

run:
	go run main.go

all: build