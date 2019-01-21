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
package setdefaults

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/gogo/protobuf/vanity"
)

type plugin struct {
	*generator.Generator
	generator.PluginImports
	validatorPkg  generator.Single
	useGogoImport bool
}

func NewPlugin(useGogoImport bool) generator.Plugin {
	return &plugin{useGogoImport: useGogoImport}
}

func (p *plugin) Name() string {
	return "setdefaults"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	if !p.useGogoImport {
		vanity.TurnOffGogoImport(file.FileDescriptorProto)
	}
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.NewImport("fmt")
	p.validatorPkg = p.NewImport("github.com/galexrt/edenconfmgmt/pkg/grpc/plugins/setdefaults")

	for _, msg := range file.Messages() {
		ccTypeName := generator.CamelCaseSlice(msg.TypeName())

		p.P(`func (this *`, ccTypeName, `) SetDefaults() error {`)
		p.In()
		p.P(`i, ok := `, ccTypeName, `.(setDefaults)`)
		p.P(`fmt.Printf("TEST NOPE: %+v, %+v", i, ok)`)
		p.Out()
		p.P(`}`)
	}
}
