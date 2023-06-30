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

profile: test
	cd regress ; go run regress.go -cpuprofile cpu.prof
	echo "top10" | go tool pprof regress/cpu.prof

# Targets that do not represent actual files
.PHONY: fmt vet test run regress profile
