// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/triggers/v1alpha/api.proto

package v1alpha

import (
	fmt "fmt"
	_ "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1"
	_ "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	_ "github.com/galexrt/edenconfmgmt/pkg/grpc/plugins/apiserver"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Trigger) Validate() error {
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	if this.Spec != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Spec); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Spec", err)
		}
	}
	return nil
}
func (this *TriggerList) Validate() error {
	if nil == this.Metadata {
		return github_com_mwitkow_go_proto_validators.FieldError("Metadata", fmt.Errorf("message must exist"))
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}
func (this *TriggerSpec) Validate() error {
	if this.RunOptions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RunOptions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RunOptions", err)
		}
	}
	for _, item := range this.Steps {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Steps", err)
			}
		}
	}
	return nil
}
func (this *Step) Validate() error {
	if this.Spec != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Spec); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Spec", err)
		}
	}
	return nil
}
func (this *RunOptions) Validate() error {
	if this.Serialize != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Serialize); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Serialize", err)
		}
	}
	if this.Target != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Target); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Target", err)
		}
	}
	return nil
}
func (this *Serialize) Validate() error {
	return nil
}
func (this *Target) Validate() error {
	return nil
}
func (this *Action) Validate() error {
	if this.Conditions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Conditions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Conditions", err)
		}
	}
	if this.RunOptions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RunOptions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RunOptions", err)
		}
	}
	for _, item := range this.Parameters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Parameters", err)
			}
		}
	}
	return nil
}
func (this *GetRequest) Validate() error {
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	return nil
}
func (this *GetResponse) Validate() error {
	if this.Trigger != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Trigger); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Trigger", err)
		}
	}
	return nil
}
func (this *ListRequest) Validate() error {
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	return nil
}
func (this *ListResponse) Validate() error {
	if this.TriggerList != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TriggerList); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TriggerList", err)
		}
	}
	return nil
}
func (this *CreateRequest) Validate() error {
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	if this.Trigger != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Trigger); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Trigger", err)
		}
	}
	return nil
}
func (this *CreateResponse) Validate() error {
	if this.Trigger != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Trigger); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Trigger", err)
		}
	}
	return nil
}
func (this *UpdateRequest) Validate() error {
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	if this.Trigger != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Trigger); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Trigger", err)
		}
	}
	return nil
}
func (this *UpdateResponse) Validate() error {
	if this.Trigger != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Trigger); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Trigger", err)
		}
	}
	return nil
}
func (this *DeleteRequest) Validate() error {
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	if this.Trigger != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Trigger); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Trigger", err)
		}
	}
	return nil
}
func (this *DeleteResponse) Validate() error {
	if this.Trigger != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Trigger); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Trigger", err)
		}
	}
	return nil
}
func (this *WatchRequest) Validate() error {
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	return nil
}
func (this *WatchResponse) Validate() error {
	if this.Event != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Event); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Event", err)
		}
	}
	if this.Trigger != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Trigger); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Trigger", err)
		}
	}
	return nil
}
