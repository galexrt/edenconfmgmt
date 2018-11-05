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

.PHONY: all proto-gen proto-clean
