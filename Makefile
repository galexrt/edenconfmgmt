DIR = $(strip $(dir $(realpath $(lastword $(MAKEFILE_LIST)))))

GO = go

PROTOFILES = $(shell find -type f -name '*.proto' | sed 's~\.\/~~g')
GOPROTOFILES = $(patsubst %.proto,%.pb.go,$(PROTOFILES))

.DEFAULT: all

all: prepare build

prepare: protoc protoc-gen-doc

build: proto-gen

protoc:
	$(GO) get -v -u github.com/gogo/protobuf/protoc-gen-gofast
	$(GO) get -v -u github.com/golang/protobuf/proto
	$(GO) get -v -u github.com/golang/protobuf/protoc-gen-go

protoc-gen-doc:
	$(GO) get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

proto-gen: $(GOPROTOFILES) docs/apis.md

proto-doc-gen: proto-clean
	$(MAKE) proto-gen

%.pb.go: %.proto
	protoc \
		-I $$GOPATH/src/ \
		--go_out=plugins=grpc:$$GOPATH/src/ \
		github.com/galexrt/edenconfmgmt/$^

docs/apis.md: $(GOPROTOFILES)
	protoc \
		-I $$GOPATH/src/ \
		--doc_out=$(DIR)docs/ \
		--doc_opt=markdown,apis.md \
		$(addprefix github.com/galexrt/edenconfmgmt/,$(PROTOFILES))

proto-clean:
	rm -f $(GOPROTOFILES)

test-env: build
	docker run \
		-d \
		-p 2379:2379 \
		-p 2380:2380 \
		-v /usr/share/ca-certificates/:/etc/ssl/certs \
		--name etcd \
		quay.io/coreos/etcd:v3.3.10 \
		/usr/local/bin/etcd \
		--data-dir=/etcd-data --name node1 \
		--initial-advertise-peer-urls http://0.0.0.0:2380 \
		--listen-peer-urls http://0.0.0.0:2380 \
		--advertise-client-urls http://0.0.0.0:2379 \
		--listen-client-urls http://0.0.0.0:2379 \
		--initial-cluster node1=http://0.0.0.0:2380

.PHONY: all protoc protoc-gen-doc proto-clean proto-doc-gen proto-gen
