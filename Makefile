BINARIES := euler
SOURCES := $(shell find . -name "*.go")

all: build

.PHONY: build
build: $(BINARIES)

$(BINARIES): $(SOURCES)
	go build ./cmd/$@
