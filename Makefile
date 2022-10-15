ARCH := $(shell uname -m)

all: test build

build:
	go build -o dist/upnt .

build_linux_64:
	GOOS=linux GOARCH=amd64 go build -o dist/linux/upnt.64 .

build_linux_32:
	GOOS=linux GOARCH=386 go build -o dist/linux/upnt.32 .

build_win_64:
	GOOS=windows GOARCH=amd64 go build -o dist/win/upnt.64.exe  .

build_win_32:
	GOOS=windows GOARCH=386 go build -o dist/win/upnt.32.exe .

build_all_32: build_linux_32 build_win_32

build_all_64: build_win_64 build_linux_64

build_all: build_all_64 build_all_32

test:
	go test -v

bench:
	go test -bench=.

clean:
	rm -rf dist/
