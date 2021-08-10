SERVICE = insights-cache-projects
BUILD_TIME=`date -u '+%Y-%m-%d_%I:%M:%S%p'`
COMMIT=`git rev-parse HEAD`
LDFLAGS=-ldflags "-s -w -extldflags '-static' -X main.BuildStamp=$(BUILD_TIME) -X main.GitHash=$(COMMIT)"
GO_LINT=golint -set_exit_status
PWD=$(shell pwd)


clean:
	rm -rf ./bin

deps: clean
	mkdir bin
	go mod tidy

build-mac: deps
	env GOOS=darwin GOARCH=amd64 go build -o bin/$(SERVICE) -a $(LDFLAGS) .
	chmod +x bin/$(SERVICE)

build-win: deps
	env GOOS=windows GOARCH=amd64 go build -o bin/$(SERVICE) -a $(LDFLAGS) .
	chmod +x bin/$(SERVICE)

build: clean deps
	GOOS=linux GOARCH=amd64 go build -o bin/$(SERVICE) -a $(LDFLAGS) .

docker-build:
	docker build -t insights-cache .

