NAME	= wegotickets
BIN 	= bin/$(NAME)

all: build

test: deps
	go test $(glide novendor)

deps:
	glide install

build: deps test
	go build -o $(GOPATH)/$(BIN)

run:
	$(GOPATH)/$(BIN)

fmt:
	go fmt ./...

lint:
	golint ./...

clean:
	rm $(GOPATH)/$(BIN)