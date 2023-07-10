fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test ./...

run: test
	cd waffle ; go run waffle.go

regress: test
	cd regress ; time go run regress.go

# Targets that do not represent actual files
.PHONY: fmt vet test run regress
