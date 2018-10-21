DIR = $(strip $(dir $(realpath $(lastword $(MAKEFILE_LIST)))))
PROTOFILES = $(shell find -type f -name '*.proto')
GOPROTOFILES = $(patsubst %.proto,%.pb.go,$(PROTOFILES))

.DEFAULT: build

build: proto-gen

proto-gen: $(GOPROTOFILES)

%.pb.go: %.proto
	protoc \
		-I $$GOPATH/src/ \
		--go_out=plugins=grpc:$$GOPATH/src/ \
		github.com/galexrt/edenconfmgmt/$^

proto-clean:
	rm -f $(GOPROTOFILES)

.PHONY: all proto-gen proto-clean
