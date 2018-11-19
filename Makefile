DIR = $(strip $(dir $(realpath $(lastword $(MAKEFILE_LIST)))))

GO = go

PROTOFILES = $(shell find -type f -name '*.proto')
GOPROTOFILES = $(patsubst %.proto,%.pb.go,$(PROTOFILES))

.DEFAULT: build

prepare: protoc protoc-gen-doc

build: proto-gen

doc-gen: proto-clean
	$(MAKE) proto-gen

protoc:
	$(GO) get -v -u github.com/gogo/protobuf/protoc-gen-gofast
	$(GO) get -v -u github.com/golang/protobuf/proto
	$(GO) get -v -u github.com/golang/protobuf/protoc-gen-go

protoc-gen-doc:
	$(GO) get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

proto-gen: $(GOPROTOFILES)

%.pb.go: %.proto
	protoc \
		-I $$GOPATH/src/ \
		--go_out=plugins=grpc:$$GOPATH/src/ \
		--doc_out=$(DIR)docs/ \
		--doc_opt=markdown,apis.md \
		github.com/galexrt/edenconfmgmt/$^

proto-clean:
	rm -f $(GOPROTOFILES)

test-env: build
	docker run \
		-d \
		-v /usr/share/ca-certificates/:/etc/ssl/certs \
		-p 4001:4001 \
		-p 2380:2380 \
		-p 2379:2379 \
		--name etcd quay.io/coreos/etcd:v2.3.8 \
		-name etcd0 \
		-advertise-client-urls http://0.0.0.0:2379 \
		-listen-client-urls http://0.0.0.0:2379 \
		-initial-advertise-peer-urls http://0.0.0.0:2380 \
		-listen-peer-urls http://0.0.0.0:2380 \
		-initial-cluster-token etcd-cluster-1 \
		-initial-cluster etcd0=http://0.0.0.0:2380 \
		-initial-cluster-state new

.PHONY: all protoc proto-gen protoc-gen-doc proto-clean
