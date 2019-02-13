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

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

const (
	flagSwaggerDir = "swagger-dir"
)

var (
	tpl *template.Template
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "swagger-gen",
		Short: "Tool to generate parts for the swagger-yaml.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			swaggerDir := viper.GetString(flagSwaggerDir)
			// Template
			baseFuncMap := template.FuncMap{
				"ToLower": strings.ToLower,
			}
			funcMap := template.FuncMap{
				"IncludeRecursiveFromDir": func(dir string, indent int, data interface{}) (ret template.HTML, err error) {
					var out template.HTML
					filepath.Walk(path.Join(swaggerDir, dir), func(path string, info os.FileInfo, err error) error {
						if err != nil {
							return err
						}
						if info.IsDir() {
							return nil
						}

						file, err := os.Open(path)
						if err != nil {
							log.Fatal(err)
						}
						defer file.Close()

						tmplString := ""
						scanner := bufio.NewScanner(file)
						for scanner.Scan() {
							if indent > 0 {
								for i := 0; i < indent; i++ {
									tmplString += " "
								}
							}
							tmplString += scanner.Text() + "\n"
						}

						if err = scanner.Err(); err != nil {
							log.Fatal(err)
						}

						t, err := template.New("").Funcs(baseFuncMap).Parse(tmplString)
						if err != nil {
							return err
						}
						var buf bytes.Buffer
						if err = t.Execute(&buf, data); err != nil {
							return err
						}
						out += template.HTML(fmt.Sprintf("// File: %s\n", path) + buf.String())
						return nil
					})
					return out, nil
				},
			}

			for k, v := range baseFuncMap {
				funcMap[k] = v
			}

			tpl = template.New("").Funcs(funcMap)
			pathLength := len(swaggerDir) - 1
			err := filepath.Walk(swaggerDir,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if info.IsDir() {
						return nil
					}
					b, err := ioutil.ReadFile(path)
					if err != nil {
						return err
					}
					name := path[pathLength:]
					t := tpl.New(name).Funcs(funcMap)
					_, err = t.Parse(string(b))
					return err
				})
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := loadFile()
			if err != nil {
				return err
			}
			return tpl.ExecuteTemplate(os.Stdout, "swagger.yaml.tmpl", cfg)
		},
	}
)

type config struct {
	Paths map[string][]*pathInfo
}

type pathInfo struct {
	ObjectName string `yaml:"objectName"`
}

func init() {
	rootCmd.PersistentFlags().String("config", "config.yaml", "config.yaml path")
	rootCmd.PersistentFlags().String(flagSwaggerDir, "./swagger", "swagger tmpl file")

	viper.BindPFlags(rootCmd.PersistentFlags())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func loadFile() (*config, error) {
	yamlFile, err := ioutil.ReadFile(viper.GetString("config"))
	if err != nil {
		return nil, err
	}
	cfg := &config{}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
