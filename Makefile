.PHONY: build clean

build:
	go build -o bin/golox cli/cli.go

clean:
	rm -rf bin/