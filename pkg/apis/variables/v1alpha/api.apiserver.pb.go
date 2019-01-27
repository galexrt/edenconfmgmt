// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/variables/v1alpha/api.proto

package v1alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1"
	_ "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	_ "github.com/galexrt/edenconfmgmt/pkg/grpc/plugins/apiserver"
	github_com_galexrt_edenconfmgmt_pkg_store_object "github.com/galexrt/edenconfmgmt/pkg/store/object"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// VariablesService !!WARNING!! This client must only be used by internal code!
// All "external" clients must use the actual GRPC API Client.
type VariablesService struct {
	VariablesServer
	store *github_com_galexrt_edenconfmgmt_pkg_store_object.Store
}

// NewVariablesService returns a new client to be used to access Variables API with the given storage.
func NewVariablesService(store *github_com_galexrt_edenconfmgmt_pkg_store_object.Store) *VariablesService {
	return &VariablesService{
		store: store,
	}
}

// Get
func (this *VariablesService) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	out, err := this.store.Get(ctx, req.GetOptions())
	if err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	target := &Variable{}
	err = target.Unmarshal(out)
	if err != nil {
		return nil, err
	}
	return &GetResponse{
		Variable: target,
	}, nil
}

// List
func (this *VariablesService) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	out, err := this.store.List(ctx, req.GetOptions())
	if err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	list := &VariableList{}
	list.Items = make([]*Variable, len(out))
	for k := range out {
		list.Items[k] = &Variable{}
		err = list.Items[k].Unmarshal(out[k])
		if err != nil {
			return nil, err
		}
	}
	return &ListResponse{
		VariableList: list,
	}, nil
}

// Create
func (this *VariablesService) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	if err := req.GetVariable().SetDefaults(); err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	req.GetVariable().GetMetadata().Namespace = ""
	if err := this.store.Create(ctx, req.GetVariable(), req.GetOptions()); err != nil {
		return nil, err
	}
	return &CreateResponse{
		Variable: req.GetVariable(),
	}, nil
}

// Update
func (this *VariablesService) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	if err := req.GetVariable().SetDefaults(); err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	req.GetVariable().GetMetadata().Namespace = ""
	if err := this.store.Update(ctx, req.GetVariable(), req.GetOptions()); err != nil {
		return nil, err
	}
	return &UpdateResponse{
		Variable: req.GetVariable(),
	}, nil
}

// Delete
func (this *VariablesService) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	return &DeleteResponse{}, this.store.Delete(ctx, req.GetVariable(), req.GetOptions())
}

// Watch
func (this *VariablesService) Watch(req *WatchRequest, stream Variables_WatchServer) error {
	req.GetOptions().Namespace = ""
	watch, err := this.store.Watch(stream.Context(), req.GetOptions())
	if err != nil {
		return nil
	}
	for {
		select {
		case out := <-watch:
			target := &Variable{}
			if err = target.Unmarshal(out.Value); err != nil {
				return err
			}
			if err := stream.Send(&WatchResponse{
				Variable: target,
			}); err != nil {
				return err
			}
		case <-stream.Context().Done():
			return nil
		}
	}
}
