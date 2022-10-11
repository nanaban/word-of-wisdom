lint:
	golangci-lint run ./...

generate:
	go generate ./...

test:
	go test -v ./... -race

benchmark:
	go test ./... -run=nonthingplease -benchmem -bench=.


