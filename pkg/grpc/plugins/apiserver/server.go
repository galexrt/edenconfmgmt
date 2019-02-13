/*
Copyright 2019 Alexander Trost. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apiserver

import (
	"log"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/gogo/protobuf/vanity"
)

const protoClientGenerateOptionID = 1337

type plugin struct {
	*generator.Generator
	generator.PluginImports
	useGogoImport  bool
	validatorPkg   generator.Single
	cacheStorePkg  generator.Single
	objectStorePkg generator.Single
	utilsAPIPkg    generator.Single
	packagesUsed   map[string]generator.Single
}

// NewPlugin return a GRPC generator plugin, see `plugin` struct.
func NewPlugin(useGogoImport bool) generator.Plugin {
	return &plugin{useGogoImport: useGogoImport}
}

func (p *plugin) Name() string {
	return "apiserver"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	p.packagesUsed = map[string]generator.Single{}
	if !p.useGogoImport {
		vanity.TurnOffGogoImport(file.FileDescriptorProto)
	}
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.NewImport("fmt")
	p.NewImport("math")
	p.cacheStorePkg = p.NewImport("github.com/galexrt/edenrun/pkg/store/cache")
	p.objectStorePkg = p.NewImport("github.com/galexrt/edenrun/pkg/store/object")
	p.utilsAPIPkg = p.NewImport("github.com/galexrt/edenrun/pkg/utils/api")
	p.validatorPkg = p.NewImport("github.com/galexrt/edenrun/pkg/grpc/plugins/apiserver")

	for _, msg := range file.Messages() {
		ccTypeName := generator.CamelCaseSlice(msg.TypeName())

		options := msg.GetOptions()
		if options == nil {
			continue
		}

		if !strings.Contains(options.String(), "apiserver.options") {
			continue
		}

		extOpts, err := proto.GetExtension(options, E_Options)
		if err != nil && err != proto.ErrMissingExtension {
			log.Fatal(err)
		}

		extOptsCasted := extOpts.(*MessageOptions)
		if extOptsCasted.Generate == nil || !*extOptsCasted.Generate {
			continue
		}

		objectIsNamespaced := false
		if extOptsCasted.Namespaced != nil && *extOptsCasted.Namespaced {
			objectIsNamespaced = true
		}

		p.NewImport("context").Use()
		p.P(`// `, ccTypeName, `sService !!WARNING!! This client must only be used by internal code!`)
		p.P(`// All "external" clients must use the actual GRPC API Client.`)
		p.P(`type `, ccTypeName, `sService struct {`)
		p.In()
		p.P(ccTypeName, `sServer`)
		p.P(`store *`, p.objectStorePkg.Use(), `.Store`)
		p.Out()
		p.P(`}`)
		p.P(``)
		p.P(`// New`, ccTypeName, `sService returns a new client to be used to access `, ccTypeName, `s API with the given storage.`)
		p.P(`func New`, ccTypeName, `sService(store *`, p.objectStorePkg.Use(), `.Store) *`, ccTypeName, `sService {`)
		p.In()
		p.P(`return &`, ccTypeName, `sService{`)
		p.In()
		p.P(`store: store,`)
		p.Out()
		p.P(`}`)
		p.Out()
		p.P(`}`)
		p.P(``)

		coreAPIPkg := "github.com/galexrt/edenrun/pkg/grpcapi/core/v1"
		if extOptsCasted.CoreAPIPkg != nil && *extOptsCasted.CoreAPIPkg != "" {
			coreAPIPkg = *extOptsCasted.CoreAPIPkg
		}
		if _, ok := p.packagesUsed[coreAPIPkg]; !ok {
			p.packagesUsed[coreAPIPkg] = p.NewImport(coreAPIPkg)
		}

		// Get
		p.P(`// Get`)
		p.P(`func (this *`, ccTypeName, `sService) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {`)
		p.In()
		p.P(`out, err := this.store.Get(ctx, req.GetOptions())`)
		p.P(`if err != nil {`)
		p.In()
		p.P(`return nil, err`)
		p.Out()
		p.P(`}`)
		if !objectIsNamespaced {
			p.P(`if req.Options != nil {`)
			p.In()
			p.P(`req.Options.Namespace = ""`)
			p.Out()
			p.P(`}`)
		}
		p.P(`target := &`, ccTypeName, `{}`)
		p.P(`err = target.Unmarshal(out)`)
		p.P(`if err != nil {`)
		p.In()
		p.P(`return nil, err`)
		p.Out()
		p.P(`}`)
		p.P(`return &GetResponse{`)
		p.In()
		p.P(ccTypeName, `: target,`)
		p.Out()
		p.P(`}, nil`)
		p.Out()
		p.P(`}`)
		p.P(``)

		// List
		p.P(`// List`)
		p.P(`func (this *`, ccTypeName, `sService) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {`)
		p.In()
		p.P(`out, err := this.store.List(ctx, req.GetOptions())`)
		p.P(`if err != nil {`)
		p.In()
		p.P(`return nil, err`)
		p.Out()
		p.P(`}`)
		if !objectIsNamespaced {
			p.P(`if req.Options != nil {`)
			p.In()
			p.P(`req.Options.Namespace = ""`)
			p.Out()
			p.P(`}`)
		}
		p.P(`list := &`, ccTypeName, `List{}`)
		p.P(`list.Items = make([]*`, ccTypeName, `, len(out))`)
		p.P(`for k := range out {`)
		p.In()
		p.P(`list.Items[k] = &`, ccTypeName, `{}`)
		p.P(`err = list.Items[k].Unmarshal(out[k])`)
		p.P(`if err != nil {`)
		p.In()
		p.P(`return nil, err`)
		p.Out()
		p.P(`}`)
		p.Out()
		p.P(`}`)
		p.P(`return &ListResponse{`)
		p.In()
		p.P(ccTypeName, `List: list,`)
		p.Out()
		p.P(`}, nil`)
		p.Out()
		p.P(`}`)
		p.P(``)

		// Create
		p.P(`// Create`)
		p.P(`func (this *`, ccTypeName, `sService) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {`)
		p.In()
		p.P(`if err := req.Get`, ccTypeName, `().SetDefaults(); err != nil {`)
		p.In()
		p.P(`return nil, err`)
		p.Out()
		p.P(`}`)
		if !objectIsNamespaced {
			p.P(`if req.Options != nil {`)
			p.In()
			p.P(`req.Options.Namespace = ""`)
			p.Out()
			p.P(`}`)
			p.P(`req.Get`, ccTypeName, `().GetMetadata().Namespace = ""`)
		}
		// Store
		p.P(`if err := this.store.Create(ctx, req.Get`, ccTypeName, `(), req.GetOptions()); err != nil {`)
		p.In()
		p.P(`return nil, err`)
		p.Out()
		p.P(`}`)
		p.P(`return &CreateResponse{`)
		p.In()
		p.P(ccTypeName, `: req.Get`, ccTypeName, `(),`)
		p.Out()
		p.P(`}, nil`)
		p.Out()
		p.P(`}`)
		p.P(``)

		// Update
		p.P(`// Update`)
		p.P(`func (this *`, ccTypeName, `sService) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {`)
		p.In()
		p.P(`if err := req.Get`, ccTypeName, `().SetDefaults(); err != nil {`)
		p.In()
		p.P(`return nil, err`)
		p.Out()
		p.P(`}`)
		if !objectIsNamespaced {
			p.P(`if req.Options != nil {`)
			p.In()
			p.P(`req.Options.Namespace = ""`)
			p.Out()
			p.P(`}`)
			p.P(`req.Get`, ccTypeName, `().GetMetadata().Namespace = ""`)
		}
		// Store
		p.P(`if err := this.store.Update(ctx, req.Get`, ccTypeName, `(), req.GetOptions()); err != nil {`)
		p.In()
		p.P(`return nil, err`)
		p.Out()
		p.P(`}`)
		p.P(`return &UpdateResponse{`)
		p.In()
		p.P(ccTypeName, `: req.Get`, ccTypeName, `(),`)
		p.Out()
		p.P(`}, nil`)
		p.Out()
		p.P(`}`)
		p.P(``)

		// Delete
		p.P(`// Delete`)
		p.P(`func (this *`, ccTypeName, `sService) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {`)
		p.In()
		if !objectIsNamespaced {
			p.P(`if req.Options != nil {`)
			p.In()
			p.P(`req.Options.Namespace = ""`)
			p.Out()
			p.P(`}`)
		}
		p.P(`return &DeleteResponse{}, this.store.Delete(ctx, req.Get`, ccTypeName, `(), req.GetOptions())`)
		p.Out()
		p.P(`}`)

		// Watch
		p.P(`// Watch`)
		p.P(`func (this *`, ccTypeName, `sService) Watch(req *WatchRequest, stream `, ccTypeName, `s_WatchServer) error {`)
		p.In()
		if !objectIsNamespaced {
			p.P(`req.GetOptions().Namespace = ""`)
		}
		p.P(`watch, err := this.store.Watch(stream.Context(), req.GetOptions())`)
		p.P(`if err != nil {`)
		p.In()
		p.P(`return nil`)
		p.Out()
		p.P(`}`)
		p.P(`for {`)
		p.In()
		p.P(`select {`)
		p.P(`case out := <-watch:`)
		p.In()
		// TODO Make use of other values from the cache.InformerResult besides `Value`
		p.P(`target := &`, ccTypeName, `{}`)
		p.P(`if err = target.Unmarshal(out.Value); err != nil {`)
		p.In()
		p.P(`return err`)
		p.Out()
		p.P(`}`)
		p.P(`if err := stream.Send(&WatchResponse{`)
		p.In()
		p.P(ccTypeName, `: target,`)
		p.Out()
		p.P(`}); err != nil {`)
		p.In()
		p.P(`return err`)
		p.Out()
		p.P(`}`)
		p.Out()
		p.P(`case <-stream.Context().Done():`)
		p.In()
		p.P(`return nil`)
		p.Out()
		p.P(`}`)
		p.Out()
		p.P(`}`)
		p.Out()
		p.P(`}`)
	}
}
