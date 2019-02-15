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

package etcd

import (
	storecommon "github.com/galexrt/edenrun/pkg/store/common"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

// ConvertETCDtoResultState convert given etcd `clientv3.Event` type to a `storecommon.ResultState`.
func ConvertETCDtoResultState(event *clientv3.Event) storecommon.ResultState {
	if event.Type == mvccpb.DELETE {
		return storecommon.ResultStateDeleted
	}
	if event.Type == mvccpb.PUT {
		if event.Kv.CreateRevision == event.Kv.Version {
			return storecommon.ResultStateUpdated
		}
		return storecommon.ResultStateCreated
	}
	return storecommon.ResultStateUnknown
}
