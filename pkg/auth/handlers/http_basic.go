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

package handlers

import (
	"context"

	"google.golang.org/grpc"
)

// HTTPBasicOptions options for HTTPBasic.
type HTTPBasicOptions struct {
	Username  string
	Passwords []string
	Options
}

// HTTPBasic struct
type HTTPBasic struct {
}

// NewHTTPBasic return a http basic auth handler
func NewHTTPBasic(opts Options) (*HTTPBasic, error) {
	// TODO use opts
	return &HTTPBasic{}, nil
}

// StreamAuth stream auth function
func (htb *HTTPBasic) StreamAuth(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (bool, error) {
	return false, nil
}

// UnaryAuth unary auth function
func (htb *HTTPBasic) UnaryAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (bool, interface{}, error) {
	return false, nil, nil
}
