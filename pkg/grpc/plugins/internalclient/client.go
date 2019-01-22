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

package internalclient

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
	useGogoImport bool
	validatorPkg  generator.Single
	storeCachePkg generator.Single
	utilsAPIPkg   generator.Single
	packagesUsed  map[string]generator.Single
}

// NewPlugin return a GRPC generator plugin, see `plugin` struct.
func NewPlugin(useGogoImport bool) generator.Plugin {
	return &plugin{useGogoImport: useGogoImport}
}

func (p *plugin) Name() string {
	return "internalclient"
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
	p.NewImport("context")
	p.storeCachePkg = p.NewImport("github.com/galexrt/edenconfmgmt/pkg/store/cache")
	p.utilsAPIPkg = p.NewImport("github.com/galexrt/edenconfmgmt/pkg/utils/api")
	p.validatorPkg = p.NewImport("github.com/galexrt/edenconfmgmt/pkg/grpc/plugins/internalclient")

	generatedClient := false

	for _, msg := range file.Messages() {
		ccTypeName := generator.CamelCaseSlice(msg.TypeName())
		varName := strings.ToLower(ccTypeName)

		options := msg.GetOptions()
		if options == nil {
			continue
		}

		if !strings.Contains(options.String(), "internalclient.options") {
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

		if !generatedClient {
			p.P(`// Client !!WARNING!! This client must only be used by internal code!`)
			p.P(`// All "external" clients must use the actual GRPC API Client.`)
			p.P(`type Client struct {`)
			p.In()
			p.P(`store *`, p.storeCachePkg.Use(), `.Store`)
			p.Out()
			p.P(`}`)
			p.P(``)
			p.P(`// NewClient returns a new Client to be used with the given storage.`)
			p.P(`func NewClient(store *`, p.storeCachePkg.Use(), `.Store) *Client {`)
			p.In()
			p.P(`return &Client{`)
			p.In()
			p.P(`store: store,`)
			p.Out()
			p.P(`}`)
			p.Out()
			p.P(`}`)
			p.P(``)
			generatedClient = true
		}

		coreAPIPkg := "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
		if extOptsCasted.CoreAPIPkg != nil && *extOptsCasted.CoreAPIPkg != "" {
			coreAPIPkg = *extOptsCasted.CoreAPIPkg
		}
		if _, ok := p.packagesUsed[coreAPIPkg]; !ok {
			p.packagesUsed[coreAPIPkg] = p.NewImport(coreAPIPkg)
		}

		// Get
		p.P(`// Get`)
		if extOptsCasted.Namespaced != nil && *extOptsCasted.Namespaced {
			p.P(`func (this *Client) Get(ctx context.Context, opts *`, p.packagesUsed[coreAPIPkg].Use(), `.GetOptions) (*`, ccTypeName, `, error) {`)
		} else {
			p.P(`func (this *Client) Get(ctx context.Context, opts *`, p.packagesUsed[coreAPIPkg].Use(), `.GetOptions) (*`, ccTypeName, `, error) {`)
		}
		p.In()
		if extOptsCasted.Namespaced != nil && *extOptsCasted.Namespaced {
			p.P(`out, err := this.store.Get(context.Background(), `, p.utilsAPIPkg.Use(), `.NamespacedObjectPath(APIPath, opts.Namespace, opts.Name))`)
		} else {
			p.P(`out, err := this.store.Get(context.Background(), `, p.utilsAPIPkg.Use(), `.ObjectPath(APIPath, opts.Name))`)
		}
		p.P(`if err != nil {`)
		p.In()
		p.P(`return nil, nil`)
		p.Out()
		p.P(`}`)
		p.P(varName, ` := &`, ccTypeName, `{}`)
		p.P(`err = proto.Unmarshal(out, `, varName, `)`)
		p.P(`return `, varName, `, err`)
		p.Out()
		p.P(`}`)
		p.P(``)

		// List
		p.P(`// List`)
		p.P(`func (this *Client) List(name string) (error) {`)
		p.In()
		p.P(`return nil`)
		p.Out()
		p.P(`}`)
		p.P(``)

		// Create
		p.P(`// Create`)
		p.P(`func (this *Client) Create(name string) (error) {`)
		p.In()
		p.P(`return nil`)
		p.Out()
		p.P(`}`)
		p.P(``)

		// Update
		p.P(`// Update`)
		p.P(`func (this *Client) Update(name string) (error) {`)
		p.In()
		p.P(`return nil`)
		p.Out()
		p.P(`}`)
		p.P(``)

		// Delete
		p.P(`// Delete`)
		p.P(`func (this *Client) Delete(name string) (error) {`)
		p.In()
		p.P(`return nil`)
		p.Out()
		p.P(`}`)
	}

	if !generatedClient {
		p.P("// No Client generated, because no `option (internalclient.options).generate = true;` is given for a `Message`")
	}
}
