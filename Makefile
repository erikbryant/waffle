fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test ./...

run: test
	cd main ; go run ./...

# Targets that do not represent actual files
.PHONY: fmt vet test run
