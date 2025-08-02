build:
	go build -o bin/omniforge ./cmd/omniforge/main.go

run:
	go run ./cmd/omniforge/main.go

run-api:
	go run ./cmd/omniforge/api_main.go

test:
	go test ./...

lint:
	golangci-lint run

docker-build:
	docker build -t omniforge .

clean:
	rm -rf bin/