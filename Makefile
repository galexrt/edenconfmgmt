DIR = $(strip $(dir $(realpath $(lastword $(MAKEFILE_LIST)))))

GO = go
GO_GET_FLAGS = -v -u

PROTOFILES = $(shell find -type f -name '*.proto' -not -path "./pkg/grpc/plugins/*" | sed 's~\.\/~~g')
GOPROTOFILES = $(patsubst %.proto,%.pb.go,$(PROTOFILES))
GOFILESDIRS = $(shell find -type d)
GOFILES = $(shell find -type f -name '*.go')

.DEFAULT: all

FORCE:

all: prepare build

prepare: protoc protoc-gen-doc go-proto-validators

build: proto-gen

protoc:
	$(GO) get $(GO_GET_FLAGS) github.com/gogo/protobuf/protoc-gen-gofast

go-proto-validators:
	$(GO) get $(GO_GET_FLAGS) github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

protoc-gen-doc:
	$(GO) get $(GO_GET_FLAGS) github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

license-bill-of-materials:
	$(GO) get $(GO_GET_FLAGS) github.com/coreos/license-bill-of-materials

bill-of-materials.json: license-bill-of-materials FORCE
	license-bill-of-materials \
		github.com/galexrt/edenconfmgmt/cmd/edenconfmgmt \
		github.com/galexrt/edenconfmgmt/cmd/edenctl \
		> bill-of-materials.json

license-header-check:
	@noHeaderError=false; \
	for f in $(GOFILES); do \
		if [[ "$$f" =~ .*(apipb_test.go|api.pb.go) ]]; then echo "-> File $$f skipped."; continue; fi; \
		if ! awk -v strings="$$(cat .license-header)" -f ./build/license-header-check.awk "$$f"; then \
			echo "-> File $$f has no license header."; \
			noHeaderError=true; \
		fi; \
	done; \
	if $$noHeaderError; then echo "=> Files with no license header found."; exit 1; fi

proto-gen: $(GOPROTOFILES) docs/apis.md

%.pb.go: %.proto
	protoc \
		-I $$GOPATH/src/ \
		--gofast_out=\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
plugins=grpc:$$GOPATH/src/ \
		--govalidators_out=gogoimport=true:$$GOPATH/src/ \
		github.com/galexrt/edenconfmgmt/$^
	go fmt github.com/galexrt/edenconfmgmt/$(dir $^)...

proto-gen-plugins: $(patsubst %.proto,%.pb.go,$(shell find -type f -name '*.proto' -path "./pkg/grpc/plugins/*" | sed 's~\.\/~~g'))

proto-gen-doc: proto-clean
	$(MAKE) proto-gen

proto-gen-all: proto-gen

docs/apis.md: $(GOPROTOFILES)
	protoc \
		-I $$GOPATH/src/ \
		--doc_out=$(DIR)docs/ \
		--doc_opt=$(DIR)build/docs/proto-doc-template.tmpl,apis.md \
		$(addprefix github.com/galexrt/edenconfmgmt/,$(PROTOFILES))

proto-clean:
	rm -f $(shell find -type f -name '*.pb.go' | sed 's~\.\/~~g')

test-env:
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

.PHONY: all protoc protoc-gen-doc proto-clean proto-gen-doc proto-gen test-env
