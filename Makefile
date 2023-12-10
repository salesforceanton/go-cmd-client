start: build run

build:
	go mod download
	go build -o ./bin/app ./cmd/main.go

run:
	./bin/app