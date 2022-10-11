deps:
	go mod download

lint:
	golangci-lint run ./...

generate:
	go generate ./...

test:
	go test -v ./... -race

benchmark:
	go test ./... -run=nonthingplease -benchmem -bench=.

.PHONY: build
build:
	mkdir -p ./bin && rm -rf ./bin/*
	go build -o ./bin/server ./cmd/server/...
	go build -o ./bin/client ./cmd/client/...

docker-build:
	docker build --no-cache -f ./build/server/Dockerfile -t wow/server .
	docker build --no-cache -f ./build/client/Dockerfile -t wow/client .
