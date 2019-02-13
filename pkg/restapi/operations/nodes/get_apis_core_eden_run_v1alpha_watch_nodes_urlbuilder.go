// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GetApisCoreEdenRunV1alphaWatchNodesURL generates an URL for the get apis core eden run v1alpha watch nodes operation
type GetApisCoreEdenRunV1alphaWatchNodesURL struct {
	LabelSelector *string
	Name          *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetApisCoreEdenRunV1alphaWatchNodesURL) WithBasePath(bp string) *GetApisCoreEdenRunV1alphaWatchNodesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetApisCoreEdenRunV1alphaWatchNodesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetApisCoreEdenRunV1alphaWatchNodesURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/apis/core.eden.run/v1alpha/watch/nodes"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var labelSelector string
	if o.LabelSelector != nil {
		labelSelector = *o.LabelSelector
	}
	if labelSelector != "" {
		qs.Set("labelSelector", labelSelector)
	}

	var name string
	if o.Name != nil {
		name = *o.Name
	}
	if name != "" {
		qs.Set("name", name)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetApisCoreEdenRunV1alphaWatchNodesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetApisCoreEdenRunV1alphaWatchNodesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetApisCoreEdenRunV1alphaWatchNodesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetApisCoreEdenRunV1alphaWatchNodesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetApisCoreEdenRunV1alphaWatchNodesURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetApisCoreEdenRunV1alphaWatchNodesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
