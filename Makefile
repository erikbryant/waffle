fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test ./...

run: test
	cd main ; go run ./...

pprof: test
	cd main ; go run waffle.go -cpuprofile cpu.prof
	echo "top10" | go tool pprof main/cpu.prof

# Targets that do not represent actual files
.PHONY: fmt vet test run pprof
