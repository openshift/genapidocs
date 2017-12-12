all: build
.PHONY: all

build:
	go build github.com/openshift/genapidocs/tools/...
.PHONY: build

clean:
	rm -rf _output
.PHONY: clean
