BINARIES := euler
SOURCES := $(shell find . -name "*.go")
PACKAGES := ./challenges ./utils

all: build

.PHONY: build
build: $(BINARIES)

$(BINARIES): $(SOURCES)
	go get ./...
	go build ./cmd/$@

.PHONY: test
test:
	go get -t ./...
	go test -v $(PACKAGES)

.PHONY: cover
cover: profile.out

profile.out: $(SOURCES)
	find . -name profile.out -delete
	for package in $(PACKAGES); do \
	  go test -covermode=count -coverpkg=./... -coverprofile=$$package/profile.out $$package; \
	done
	echo "mode: count" > profile.out.tmp
	cat `find . -name profile.out` | grep -v  mode: | sort -r | awk '{if($$1 != last) {print $$0;last=$$1}}' >> profile.out.tmp
	mv profile.out.tmp profile.out

.PHONY: convey
convey:
	go get github.com/smartystreets/goconvey
	goconvey -cover -port=9042 -workDir="$(realpath .)" -depth=-1
