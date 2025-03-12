.Phony: run fmt lint build clean

APP_NAME=topmusicstreaming

fmt:
	gofmt -s -l .

lint: fmt
	golangci-lint run

run:
	go run main.go

build:
	go build -o bin/$(APP_NAME)

clean:
	rm -rf ./bin
	rm -rf json