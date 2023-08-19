.PHONY: all build

all:
	go build
	./wait

build:
	go mod tidy
	go build

