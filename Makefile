.PHONY: test
test:
	go test ./... -cover

run:
	go run main.go

lint:
	golangci-lint run
