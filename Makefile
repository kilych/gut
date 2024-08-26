tidy:
	go mod tidy

fmt:
	go fmt

check-style:
	golangci-lint run

test:
	go test -cover
