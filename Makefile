GOARCH ?= amd64
GOOS ?= linux
GO ?= env GOOS=$(GOOS) GOARCH=$(GOARCH) go
UNAME := $(shell uname)
BUILD ?= $(abspath build)
BUILDARCH ?= $(BUILD)/$(GOOS)-$(GOARCH)

all: $(BUILD)/fkc

$(BUILD)/fkc: cmd/*.go *.go
	go build -o $(BUILD)/fkc cmd/*.go

clean:
	rm -rf $(BUILD)
