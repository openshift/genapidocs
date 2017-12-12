all: build
.PHONY: all

build:
	go build -o _output/local/bin/genapidocs github.com/openshift/genapidocs/tools/...
.PHONY: build

clean:
	rm -rf _output
.PHONY: clean

update-deps:
	hack/update-deps.sh
.PHONY: update-deps
