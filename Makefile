all: build

fmt:
	gofmt -l -w -s .

dep: fmt
	go mod download

build:
	go build -v

clean:
	go clean
