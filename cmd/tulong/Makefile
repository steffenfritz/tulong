ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BINARY=tulong
VERSION=0.1.0-DEV

BUILD=`git rev-parse --short HEAD`
PLATFORMS=darwin linux windows
ARCHITECTURES=amd64

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD} -w -s"
BLDFLAGS=-buildmode=pie

all: clean build_all

build:
	go build ${BLDFLAGS} ${LDFLAGS} -o ${BINARY}

build_all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build $(BLDFLAGS) $(LDFLAGS) -v -o $(BINARY)-$(GOOS))))
	mv tulong-darwin tulong && tar cvfz tulong_macos_$(VERSION).tar.gz tulong
	rm tulong
	mv tulong-linux tulong && tar cvfz tulong_linux_$(VERSION).tar.gz tulong
	rm tulong
	mv tulong-windows tulong && tar cvfz tulong_win_$(VERSION).tar.gz tulong
	rm tulong

clean:
	rm -f '${BINARY}-linux'
	rm -f '${BINARY}-darwin'
	rm -f '${BINARY}-windows'

.PHONY: clean build build_all all
