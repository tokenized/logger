
all: clean deps tools test

clean:
	go clean -testcache

prepare:
	mkdir -p tmp

test: prepare
	go test -coverprofile=tmp/coverage.out ./...

test-race:
	go test -race ./...

bench:
	go test -bench . ./...
