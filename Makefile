run:
	go run ./cmd/api

format:
	go fmt ./...

test: format
	go test -v ./...

build: test
	go build -o go-webapp ./cmd/api

clean:
	rm -rf ./go-webapp