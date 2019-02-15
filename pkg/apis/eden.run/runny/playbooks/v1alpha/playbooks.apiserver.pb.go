// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/galexrt/edenrun/pkg/apis/eden.run/runny/playbooks/v1alpha/playbooks.proto

package v1alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/galexrt/edenrun/pkg/apis/eden.run/core/meta/v1"
	_ "github.com/galexrt/edenrun/pkg/grpc/plugins/apiserver"
	github_com_galexrt_edenrun_pkg_grpc_plugins_apiserver "github.com/galexrt/edenrun/pkg/grpc/plugins/apiserver"
	github_com_galexrt_edenrun_pkg_store_object "github.com/galexrt/edenrun/pkg/store/object"
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

// PlayBooksService !!WARNING!! This client must only be used by internal code!
// All "external" clients must use the actual GRPC API Client.
type PlayBooksService struct {
	PlayBooksServer
	store *github_com_galexrt_edenrun_pkg_store_object.Store
}

// NewPlayBooksService returns a new client to be used to access PlayBooks API with the given storage.
func NewPlayBooksService(store *github_com_galexrt_edenrun_pkg_store_object.Store) *PlayBooksService {
	return &PlayBooksService{
		store: store,
	}
}

// Get
func (this *PlayBooksService) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	out, err := this.store.Get(ctx, req.GetOptions())
	if err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	target := &PlayBook{}
	err = target.Unmarshal(out)
	if err != nil {
		return nil, err
	}
	return &GetResponse{
		PlayBook: target,
	}, nil
}

// List
func (this *PlayBooksService) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	out, err := this.store.List(ctx, req.GetOptions())
	if err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	list := &PlayBookList{}
	list.Items = make([]*PlayBook, len(out))
	for k := range out {
		list.Items[k] = &PlayBook{}
		err = list.Items[k].Unmarshal(out[k])
		if err != nil {
			return nil, err
		}
	}
	return &ListResponse{
		PlayBookList: list,
	}, nil
}

// Create
func (this *PlayBooksService) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	if err := req.GetPlayBook().SetDefaults(github_com_galexrt_edenrun_pkg_grpc_plugins_apiserver.MethodCreate); err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	req.GetPlayBook().GetMetadata().Namespace = ""
	if err := this.store.Create(ctx, req.GetPlayBook(), req.GetOptions()); err != nil {
		return nil, err
	}
	return &CreateResponse{
		PlayBook: req.GetPlayBook(),
	}, nil
}

// Update
func (this *PlayBooksService) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	if err := req.GetPlayBook().SetDefaults(github_com_galexrt_edenrun_pkg_grpc_plugins_apiserver.MethodUpdate); err != nil {
		return nil, err
	}
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	req.GetPlayBook().GetMetadata().Namespace = ""
	if err := this.store.Update(ctx, req.GetPlayBook(), req.GetOptions()); err != nil {
		return nil, err
	}
	return &UpdateResponse{
		PlayBook: req.GetPlayBook(),
	}, nil
}

// Delete
func (this *PlayBooksService) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	if req.Options != nil {
		req.Options.Namespace = ""
	}
	return &DeleteResponse{}, this.store.Delete(ctx, req.GetPlayBook(), req.GetOptions())
}

// Watch
func (this *PlayBooksService) Watch(req *WatchRequest, stream PlayBooks_WatchServer) error {
	req.GetOptions().Namespace = ""
	watch, err := this.store.Watch(stream.Context(), req.GetOptions())
	if err != nil {
		return nil
	}
	for {
		select {
		case out := <-watch:
			target := &PlayBook{}
			if err = target.Unmarshal(out.Value); err != nil {
				return err
			}
			if err = stream.Send(&WatchResponse{
				PlayBook: target,
			}); err != nil {
				return err
			}
		case <-stream.Context().Done():
			return stream.Context().Err()
		}
	}
}
