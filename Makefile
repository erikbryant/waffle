fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test ./...

run: test
	cd main ; go run ./...

regress: test
	cd test ; go run test.go -cpuprofile cpu.prof
	echo "top10" | go tool pprof test/cpu.prof

# Targets that do not represent actual files
.PHONY: fmt vet test regress run
