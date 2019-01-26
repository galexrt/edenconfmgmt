// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/cronjobs/v1alpha/api.proto

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
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// CronJobsService !!WARNING!! This client must only be used by internal code!
// All "external" clients must use the actual GRPC API Client.
type CronJobsService struct {
	CronJobsServer
	store *github_com_galexrt_edenconfmgmt_pkg_store_object.Store
}

// NewCronJobsService returns a new client to be used to access CronJobs API with the given storage.
func NewCronJobsService(store *github_com_galexrt_edenconfmgmt_pkg_store_object.Store) *CronJobsService {
	return &CronJobsService{
		store: store,
	}
}

// Get
func (this *CronJobsService) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	out, err := this.store.Get(ctx, req.GetOptions())
	if err != nil {
		return nil, err
	}
	target := &CronJob{}
	err = target.Unmarshal(out)
	if err != nil {
		return nil, err
	}
	return &GetResponse{
		CronJob: target,
	}, nil
}

// List
func (this *CronJobsService) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	out, err := this.store.List(ctx, req.GetOptions())
	if err != nil {
		return nil, err
	}
	list := &CronJobList{}
	list.Items = make([]*CronJob, len(out))
	for k := range out {
		list.Items[k] = &CronJob{}
		err = list.Items[k].Unmarshal(out[k])
		if err != nil {
			return nil, err
		}
	}
	return &ListResponse{
		CronJobList: list,
	}, nil
}

// Create
func (this *CronJobsService) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	if err := req.GetCronJob().SetDefaults(); err != nil {
		return nil, err
	}
	out, err := req.GetCronJob().Marshal()
	if err != nil {
		return nil, err
	}
	if err = this.store.Create(ctx, out, req.GetOptions()); err != nil {
		return nil, err
	}
	return &CreateResponse{
		CronJob: req.GetCronJob(),
	}, nil
}

// Update
func (this *CronJobsService) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	if err := req.GetCronJob().SetDefaults(); err != nil {
		return nil, err
	}
	out, err := req.GetCronJob().Marshal()
	if err != nil {
		return nil, err
	}
	if err = this.store.Update(ctx, out, req.GetOptions()); err != nil {
		return nil, err
	}
	return &UpdateResponse{
		CronJob: req.GetCronJob(),
	}, nil
}

// Delete
func (this *CronJobsService) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return &DeleteResponse{}, this.store.Delete(ctx, req.GetOptions())
}

// Watch
func (this *CronJobsService) Watch(req *WatchRequest, stream CronJobs_WatchServer) error {
	watch, err := this.store.Watch(stream.Context(), req.GetOptions())
	if err != nil {
		return nil
	}
	for {
		select {
		case res := <-watch:
			value := res.Value.(*CronJob)
			if err := stream.Send(&WatchResponse{
				CronJob: value,
			}); err != nil {
				return err
			}
		case <-stream.Context().Done():
			return nil
		}
	}
	return nil
}
