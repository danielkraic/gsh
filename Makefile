APP_NAME := gsh

default: build

vet:
	go list ./... | grep -v "./vendor*" | xargs go vet

fmt:
	find . -type f -name "*.go" | grep -v "./vendor*" | xargs gofmt -s -w

build: vet fmt
	CGO_ENABLED=0 GOOS=linux go build -o $(APP_NAME) .

test:
	go test ./... -v

install:
	go install -v

clean:
	rm -f $(APP_NAME)

.PHONY: vet fmt build test install clean
