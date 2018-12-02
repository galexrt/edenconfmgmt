# Protocol Documentation
<a name="top"></a>

## Table of Contents

* [github.com/galexrt/edenconfmgmt/pkg/apis/beacons/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/beacons/v1alpha/api.proto)
    * Description: 
    * Services
        * [Beacons](#beacons.v1alpha.Beacons)
    * Messages
        * [AddRequest](#beacons.v1alpha.AddRequest)
        * [AddResponse](#beacons.v1alpha.AddResponse)
        * [Beacon](#beacons.v1alpha.Beacon)
        * [BeaconSpec](#beacons.v1alpha.BeaconSpec)
        * [DeleteRequest](#beacons.v1alpha.DeleteRequest)
        * [DeleteResponse](#beacons.v1alpha.DeleteResponse)
        * [GetRequest](#beacons.v1alpha.GetRequest)
        * [GetResponse](#beacons.v1alpha.GetResponse)
        * [ListRequest](#beacons.v1alpha.ListRequest)
        * [ListResponse](#beacons.v1alpha.ListResponse)
        * [UpdateRequest](#beacons.v1alpha.UpdateRequest)
        * [UpdateResponse](#beacons.v1alpha.UpdateResponse)
        * [WatchRequest](#beacons.v1alpha.WatchRequest)
        * [WatchResponse](#beacons.v1alpha.WatchResponse)
* [github.com/galexrt/edenconfmgmt/pkg/apis/configs/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/configs/v1alpha/api.proto)
    * Description: 
    * Services
        * [Configs](#configs.v1alpha.Configs)
    * Messages
        * [AddRequest](#configs.v1alpha.AddRequest)
        * [AddResponse](#configs.v1alpha.AddResponse)
        * [Config](#configs.v1alpha.Config)
        * [ConfigSpec](#configs.v1alpha.ConfigSpec)
        * [DataStore](#configs.v1alpha.DataStore)
        * [DeleteRequest](#configs.v1alpha.DeleteRequest)
        * [DeleteResponse](#configs.v1alpha.DeleteResponse)
        * [ETCD](#configs.v1alpha.ETCD)
        * [GetRequest](#configs.v1alpha.GetRequest)
        * [GetResponse](#configs.v1alpha.GetResponse)
        * [ListRequest](#configs.v1alpha.ListRequest)
        * [ListResponse](#configs.v1alpha.ListResponse)
        * [UpdateRequest](#configs.v1alpha.UpdateRequest)
        * [UpdateResponse](#configs.v1alpha.UpdateResponse)
        * [WatchRequest](#configs.v1alpha.WatchRequest)
        * [WatchResponse](#configs.v1alpha.WatchResponse)
* [github.com/galexrt/edenconfmgmt/pkg/apis/secrets/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/secrets/v1alpha/api.proto)
    * Description: 
    * Services
        * [Secrets](#secrets.v1alpha.Secrets)
    * Messages
        * [AddRequest](#secrets.v1alpha.AddRequest)
        * [AddResponse](#secrets.v1alpha.AddResponse)
        * [DeleteRequest](#secrets.v1alpha.DeleteRequest)
        * [DeleteResponse](#secrets.v1alpha.DeleteResponse)
        * [GetRequest](#secrets.v1alpha.GetRequest)
        * [GetResponse](#secrets.v1alpha.GetResponse)
        * [ListRequest](#secrets.v1alpha.ListRequest)
        * [ListResponse](#secrets.v1alpha.ListResponse)
        * [Secret](#secrets.v1alpha.Secret)
        * [SecretSpec](#secrets.v1alpha.SecretSpec)
        * [SecretSpec.SecretsEntry](#secrets.v1alpha.SecretSpec.SecretsEntry)
        * [UpdateRequest](#secrets.v1alpha.UpdateRequest)
        * [UpdateResponse](#secrets.v1alpha.UpdateResponse)
        * [WatchRequest](#secrets.v1alpha.WatchRequest)
        * [WatchResponse](#secrets.v1alpha.WatchResponse)
* [github.com/galexrt/edenconfmgmt/pkg/apis/clustervariables/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/clustervariables/v1alpha/api.proto)
    * Description: 
    * Services
        * [ClusterVariables](#clustervariables.v1alpha.ClusterVariables)
    * Messages
        * [AddRequest](#clustervariables.v1alpha.AddRequest)
        * [AddResponse](#clustervariables.v1alpha.AddResponse)
        * [ClusterVariable](#clustervariables.v1alpha.ClusterVariable)
        * [DeleteRequest](#clustervariables.v1alpha.DeleteRequest)
        * [DeleteResponse](#clustervariables.v1alpha.DeleteResponse)
        * [GetRequest](#clustervariables.v1alpha.GetRequest)
        * [GetResponse](#clustervariables.v1alpha.GetResponse)
        * [ListRequest](#clustervariables.v1alpha.ListRequest)
        * [ListResponse](#clustervariables.v1alpha.ListResponse)
        * [UpdateRequest](#clustervariables.v1alpha.UpdateRequest)
        * [UpdateResponse](#clustervariables.v1alpha.UpdateResponse)
        * [VariableSpec](#clustervariables.v1alpha.VariableSpec)
        * [VariableSpec.ClustervariablesEntry](#clustervariables.v1alpha.VariableSpec.ClustervariablesEntry)
        * [WatchRequest](#clustervariables.v1alpha.WatchRequest)
        * [WatchResponse](#clustervariables.v1alpha.WatchResponse)
* [github.com/galexrt/edenconfmgmt/pkg/apis/variables/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/variables/v1alpha/api.proto)
    * Description: 
    * Services
        * [Variables](#variables.v1alpha.Variables)
    * Messages
        * [AddRequest](#variables.v1alpha.AddRequest)
        * [AddResponse](#variables.v1alpha.AddResponse)
        * [DeleteRequest](#variables.v1alpha.DeleteRequest)
        * [DeleteResponse](#variables.v1alpha.DeleteResponse)
        * [GetRequest](#variables.v1alpha.GetRequest)
        * [GetResponse](#variables.v1alpha.GetResponse)
        * [ListRequest](#variables.v1alpha.ListRequest)
        * [ListResponse](#variables.v1alpha.ListResponse)
        * [UpdateRequest](#variables.v1alpha.UpdateRequest)
        * [UpdateResponse](#variables.v1alpha.UpdateResponse)
        * [Variable](#variables.v1alpha.Variable)
        * [VariableSpec](#variables.v1alpha.VariableSpec)
        * [VariableSpec.VariablesEntry](#variables.v1alpha.VariableSpec.VariablesEntry)
        * [WatchRequest](#variables.v1alpha.WatchRequest)
        * [WatchResponse](#variables.v1alpha.WatchResponse)
* [github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha/api.proto)
    * Description: 
    * Messages
        * [Condition](#core.v1alpha.Condition)
        * [Conditions](#core.v1alpha.Conditions)
        * [Error](#core.v1alpha.Error)
        * [GetOptions](#core.v1alpha.GetOptions)
        * [Include](#core.v1alpha.Include)
        * [ListOptions](#core.v1alpha.ListOptions)
        * [ObjectMetadata](#core.v1alpha.ObjectMetadata)
        * [ObjectMetadata.AnnotationsEntry](#core.v1alpha.ObjectMetadata.AnnotationsEntry)
        * [ObjectMetadata.LabelsEntry](#core.v1alpha.ObjectMetadata.LabelsEntry)
        * [Retry](#core.v1alpha.Retry)
        * [VersionRequest](#core.v1alpha.VersionRequest)
        * [VersionResponse](#core.v1alpha.VersionResponse)
        * [WatchOptions](#core.v1alpha.WatchOptions)
* [github.com/galexrt/edenconfmgmt/pkg/apis/triggers/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/triggers/v1alpha/api.proto)
    * Description: 
    * Services
        * [Triggers](#triggers.v1alpha.Triggers)
    * Messages
        * [Action](#triggers.v1alpha.Action)
        * [AddRequest](#triggers.v1alpha.AddRequest)
        * [AddResponse](#triggers.v1alpha.AddResponse)
        * [DeleteRequest](#triggers.v1alpha.DeleteRequest)
        * [DeleteResponse](#triggers.v1alpha.DeleteResponse)
        * [GetRequest](#triggers.v1alpha.GetRequest)
        * [GetResponse](#triggers.v1alpha.GetResponse)
        * [ListRequest](#triggers.v1alpha.ListRequest)
        * [ListResponse](#triggers.v1alpha.ListResponse)
        * [RunOptions](#triggers.v1alpha.RunOptions)
        * [Serialize](#triggers.v1alpha.Serialize)
        * [Step](#triggers.v1alpha.Step)
        * [Target](#triggers.v1alpha.Target)
        * [Trigger](#triggers.v1alpha.Trigger)
        * [TriggerSpec](#triggers.v1alpha.TriggerSpec)
        * [UpdateRequest](#triggers.v1alpha.UpdateRequest)
        * [UpdateResponse](#triggers.v1alpha.UpdateResponse)
        * [WatchRequest](#triggers.v1alpha.WatchRequest)
        * [WatchResponse](#triggers.v1alpha.WatchResponse)
* [github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha/api.proto)
    * Description: 
    * Services
        * [Events](#events.v1alpha.Events)
    * Messages
        * [AddRequest](#events.v1alpha.AddRequest)
        * [AddResponse](#events.v1alpha.AddResponse)
        * [DeleteRequest](#events.v1alpha.DeleteRequest)
        * [DeleteResponse](#events.v1alpha.DeleteResponse)
        * [Event](#events.v1alpha.Event)
        * [EventSpec](#events.v1alpha.EventSpec)
        * [GetRequest](#events.v1alpha.GetRequest)
        * [GetResponse](#events.v1alpha.GetResponse)
        * [ListRequest](#events.v1alpha.ListRequest)
        * [ListResponse](#events.v1alpha.ListResponse)
        * [UpdateRequest](#events.v1alpha.UpdateRequest)
        * [UpdateResponse](#events.v1alpha.UpdateResponse)
        * [WatchRequest](#events.v1alpha.WatchRequest)
        * [WatchResponse](#events.v1alpha.WatchResponse)
* [github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto)
    * Description: 
    * Services
        * [Nodes](#nodes.v1alpha.Nodes)
    * Messages
        * [AddRequest](#nodes.v1alpha.AddRequest)
        * [AddResponse](#nodes.v1alpha.AddResponse)
        * [DeleteRequest](#nodes.v1alpha.DeleteRequest)
        * [DeleteResponse](#nodes.v1alpha.DeleteResponse)
        * [GetRequest](#nodes.v1alpha.GetRequest)
        * [GetResponse](#nodes.v1alpha.GetResponse)
        * [ListRequest](#nodes.v1alpha.ListRequest)
        * [ListResponse](#nodes.v1alpha.ListResponse)
        * [Network](#nodes.v1alpha.Network)
        * [Node](#nodes.v1alpha.Node)
        * [Port](#nodes.v1alpha.Port)
        * [Spec](#nodes.v1alpha.Spec)
        * [Status](#nodes.v1alpha.Status)
        * [UpdateRequest](#nodes.v1alpha.UpdateRequest)
        * [UpdateResponse](#nodes.v1alpha.UpdateResponse)
        * [WatchRequest](#nodes.v1alpha.WatchRequest)
        * [WatchResponse](#nodes.v1alpha.WatchResponse)
* [github.com/galexrt/edenconfmgmt/pkg/apis/cronjobs/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/cronjobs/v1alpha/api.proto)
    * Description: 
    * Services
        * [CronJobs](#cronjobs.v1alpha.CronJobs)
    * Messages
        * [AddRequest](#cronjobs.v1alpha.AddRequest)
        * [AddResponse](#cronjobs.v1alpha.AddResponse)
        * [CronJob](#cronjobs.v1alpha.CronJob)
        * [CronJobSpec](#cronjobs.v1alpha.CronJobSpec)
        * [DeleteRequest](#cronjobs.v1alpha.DeleteRequest)
        * [DeleteResponse](#cronjobs.v1alpha.DeleteResponse)
        * [GetRequest](#cronjobs.v1alpha.GetRequest)
        * [GetResponse](#cronjobs.v1alpha.GetResponse)
        * [ListRequest](#cronjobs.v1alpha.ListRequest)
        * [ListResponse](#cronjobs.v1alpha.ListResponse)
        * [Schedule](#cronjobs.v1alpha.Schedule)
        * [UpdateRequest](#cronjobs.v1alpha.UpdateRequest)
        * [UpdateResponse](#cronjobs.v1alpha.UpdateResponse)
        * [WatchRequest](#cronjobs.v1alpha.WatchRequest)
        * [WatchResponse](#cronjobs.v1alpha.WatchResponse)
* [github.com/galexrt/edenconfmgmt/pkg/apis/taskbooks/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/taskbooks/v1alpha/api.proto)
    * Description: 
    * Services
        * [TaskBookBooks](#taskbooks.v1alpha.TaskBookBooks)
    * Messages
        * [Action](#taskbooks.v1alpha.Action)
        * [AddRequest](#taskbooks.v1alpha.AddRequest)
        * [AddResponse](#taskbooks.v1alpha.AddResponse)
        * [DeleteRequest](#taskbooks.v1alpha.DeleteRequest)
        * [DeleteResponse](#taskbooks.v1alpha.DeleteResponse)
        * [GetRequest](#taskbooks.v1alpha.GetRequest)
        * [GetResponse](#taskbooks.v1alpha.GetResponse)
        * [ListRequest](#taskbooks.v1alpha.ListRequest)
        * [ListResponse](#taskbooks.v1alpha.ListResponse)
        * [RunOptions](#taskbooks.v1alpha.RunOptions)
        * [Serialize](#taskbooks.v1alpha.Serialize)
        * [Step](#taskbooks.v1alpha.Step)
        * [Target](#taskbooks.v1alpha.Target)
        * [TaskBook](#taskbooks.v1alpha.TaskBook)
        * [TaskBookSpec](#taskbooks.v1alpha.TaskBookSpec)
        * [UpdateRequest](#taskbooks.v1alpha.UpdateRequest)
        * [UpdateResponse](#taskbooks.v1alpha.UpdateResponse)
        * [WatchRequest](#taskbooks.v1alpha.WatchRequest)
        * [WatchResponse](#taskbooks.v1alpha.WatchResponse)* [Scalar Value Types](#scalar-value-types)

<a name="github.com/galexrt/edenconfmgmt/pkg/apis/beacons/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/beacons/v1alpha/api.proto



<a name="beacons.v1alpha.AddRequest"></a>

### AddRequest

Add Beacon request.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
<a name="beacons.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="beacons.v1alpha.Beacon"></a>

### Beacon



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for Beacon object. |
| spec | [BeaconSpec](#beacons.v1alpha.BeaconSpec) |  | Spec for Beacon. |
<a name="beacons.v1alpha.BeaconSpec"></a>

### BeaconSpec

TODO
<a name="beacons.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
<a name="beacons.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="beacons.v1alpha.GetRequest"></a>

### GetRequest

Get request content.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="beacons.v1alpha.GetResponse"></a>

### GetResponse

Get response content.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="beacons.v1alpha.ListRequest"></a>

### ListRequest

List request content.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |
<a name="beacons.v1alpha.ListResponse"></a>

### ListResponse

List response content.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beacons | [Beacon](#beacons.v1alpha.Beacon) | repeated | Beacon list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="beacons.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
<a name="beacons.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="beacons.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |
<a name="beacons.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Beacon info for watch response. |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon for watch response. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="beacons.v1alpha.Beacons"></a>

### Beacons



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#beacons.v1alpha.GetRequest) | [GetResponse](#beacons.v1alpha.GetResponse) | Get a Beacon. |
| List | [ListRequest](#beacons.v1alpha.ListRequest) | [ListResponse](#beacons.v1alpha.ListResponse) | List Beacons. |
| Add | [AddRequest](#beacons.v1alpha.AddRequest) | [AddResponse](#beacons.v1alpha.AddResponse) | Add a Beacon. |
| Update | [UpdateRequest](#beacons.v1alpha.UpdateRequest) | [UpdateResponse](#beacons.v1alpha.UpdateResponse) | Update a Beacon. |
| Delete | [DeleteRequest](#beacons.v1alpha.DeleteRequest) | [DeleteResponse](#beacons.v1alpha.DeleteResponse) | Delete a Beacon. |
| Watch | [WatchRequest](#beacons.v1alpha.WatchRequest) | [WatchResponse](#beacons.v1alpha.WatchResponse) stream | Watch Beacon. |

 <!-- end services -->



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/configs/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/configs/v1alpha/api.proto



<a name="configs.v1alpha.AddRequest"></a>

### AddRequest

Add

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
<a name="configs.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="configs.v1alpha.Config"></a>

### Config

Config object base.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for Config object. |
| spec | [ConfigSpec](#configs.v1alpha.ConfigSpec) |  | Spec for Config. |
<a name="configs.v1alpha.ConfigSpec"></a>

### ConfigSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dataStore | [DataStore](#configs.v1alpha.DataStore) |  | DataStore. |
<a name="configs.v1alpha.DataStore"></a>

### DataStore

DataStore config for the data store

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Type. |
| count | [int64](#int64) |  | Count. |
| etcd | [ETCD](#configs.v1alpha.ETCD) |  | ETCD config options. |
<a name="configs.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
<a name="configs.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="configs.v1alpha.ETCD"></a>

### ETCD



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoints | [string](#string) |  | List of endpoints if running your own ETCD server (should only be used for testing). |
<a name="configs.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="configs.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="configs.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |
<a name="configs.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configs | [Config](#configs.v1alpha.Config) | repeated | Config list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="configs.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
<a name="configs.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="configs.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |
<a name="configs.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Event info for watch response. |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="configs.v1alpha.Configs"></a>

### Configs



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#configs.v1alpha.GetRequest) | [GetResponse](#configs.v1alpha.GetResponse) | Get a Config. |
| List | [ListRequest](#configs.v1alpha.ListRequest) | [ListResponse](#configs.v1alpha.ListResponse) | List Configs. |
| Add | [AddRequest](#configs.v1alpha.AddRequest) | [AddResponse](#configs.v1alpha.AddResponse) | Add a Config. |
| Update | [UpdateRequest](#configs.v1alpha.UpdateRequest) | [UpdateResponse](#configs.v1alpha.UpdateResponse) | Update a Config. |
| Delete | [DeleteRequest](#configs.v1alpha.DeleteRequest) | [DeleteResponse](#configs.v1alpha.DeleteResponse) | Delete a Config. |
| Watch | [WatchRequest](#configs.v1alpha.WatchRequest) | [WatchResponse](#configs.v1alpha.WatchResponse) stream | Watch Config. |

 <!-- end services -->



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/secrets/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/secrets/v1alpha/api.proto



<a name="secrets.v1alpha.AddRequest"></a>

### AddRequest

Add

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
<a name="secrets.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="secrets.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
<a name="secrets.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="secrets.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="secrets.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="secrets.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |
<a name="secrets.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secrets | [Secret](#secrets.v1alpha.Secret) | repeated | Secret list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="secrets.v1alpha.Secret"></a>

### Secret



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for Secret object. |
| spec | [SecretSpec](#secrets.v1alpha.SecretSpec) |  | Spec for Secret. |
<a name="secrets.v1alpha.SecretSpec"></a>

### SecretSpec

Secrets secrets file structure. This covers the "Secrets File" and "'From Host' Static Secret File".

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secrets | [SecretSpec.SecretsEntry](#secrets.v1alpha.SecretSpec.SecretsEntry) | repeated | Secrets. |
| type | [string](#string) |  | Type of secret. |
<a name="secrets.v1alpha.SecretSpec.SecretsEntry"></a>

### SecretSpec.SecretsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Any](#google.protobuf.Any) |  |  |
<a name="secrets.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
<a name="secrets.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="secrets.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |
<a name="secrets.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Secret info for watch response. |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret for watch response. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="secrets.v1alpha.Secrets"></a>

### Secrets



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#secrets.v1alpha.GetRequest) | [GetResponse](#secrets.v1alpha.GetResponse) | Get a Secret. |
| List | [ListRequest](#secrets.v1alpha.ListRequest) | [ListResponse](#secrets.v1alpha.ListResponse) | List Secret. |
| Add | [AddRequest](#secrets.v1alpha.AddRequest) | [AddResponse](#secrets.v1alpha.AddResponse) | Add a Secret. |
| Update | [UpdateRequest](#secrets.v1alpha.UpdateRequest) | [UpdateResponse](#secrets.v1alpha.UpdateResponse) | Update a Secret. |
| Delete | [DeleteRequest](#secrets.v1alpha.DeleteRequest) | [DeleteResponse](#secrets.v1alpha.DeleteResponse) | Delete a Secret. |
| Watch | [WatchRequest](#secrets.v1alpha.WatchRequest) | [WatchResponse](#secrets.v1alpha.WatchResponse) stream | Watch Secret. |

 <!-- end services -->



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/clustervariables/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/clustervariables/v1alpha/api.proto



<a name="clustervariables.v1alpha.AddRequest"></a>

### AddRequest

Add

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="clustervariables.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="clustervariables.v1alpha.ClusterVariable"></a>

### ClusterVariable



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for ClusterVariable object. |
| spec | [VariableSpec](#clustervariables.v1alpha.VariableSpec) |  | Spec for ClusterVariable. |
<a name="clustervariables.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="clustervariables.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="clustervariables.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="clustervariables.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="clustervariables.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |
<a name="clustervariables.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clustervariables | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) | repeated | ClusterVariable list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="clustervariables.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="clustervariables.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="clustervariables.v1alpha.VariableSpec"></a>

### VariableSpec

ClusterVariables clustervariables file structure. This covers the "ClusterVariables File" and "'From Host' Static ClusterVariable File".

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clustervariables | [VariableSpec.ClustervariablesEntry](#clustervariables.v1alpha.VariableSpec.ClustervariablesEntry) | repeated | ClusterVariables. |
| type | [string](#string) |  | Type of ClusterVariable. |
<a name="clustervariables.v1alpha.VariableSpec.ClustervariablesEntry"></a>

### VariableSpec.ClustervariablesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Any](#google.protobuf.Any) |  |  |
<a name="clustervariables.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |
<a name="clustervariables.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | ClusterVariable info for watch response. |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable for watch response. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="clustervariables.v1alpha.ClusterVariables"></a>

### ClusterVariables



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#clustervariables.v1alpha.GetRequest) | [GetResponse](#clustervariables.v1alpha.GetResponse) | Get a ClusterVariable. |
| List | [ListRequest](#clustervariables.v1alpha.ListRequest) | [ListResponse](#clustervariables.v1alpha.ListResponse) | List ClusterVariable. |
| Add | [AddRequest](#clustervariables.v1alpha.AddRequest) | [AddResponse](#clustervariables.v1alpha.AddResponse) | Add a ClusterVariable. |
| Update | [UpdateRequest](#clustervariables.v1alpha.UpdateRequest) | [UpdateResponse](#clustervariables.v1alpha.UpdateResponse) | Update a ClusterVariable. |
| Delete | [DeleteRequest](#clustervariables.v1alpha.DeleteRequest) | [DeleteResponse](#clustervariables.v1alpha.DeleteResponse) | Delete a ClusterVariable. |
| Watch | [WatchRequest](#clustervariables.v1alpha.WatchRequest) | [WatchResponse](#clustervariables.v1alpha.WatchResponse) stream | Watch ClusterVariable. |

 <!-- end services -->



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/variables/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/variables/v1alpha/api.proto



<a name="variables.v1alpha.AddRequest"></a>

### AddRequest

Add

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
<a name="variables.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="variables.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
<a name="variables.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="variables.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="variables.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="variables.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |
<a name="variables.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variables | [Variable](#variables.v1alpha.Variable) | repeated | Variable list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="variables.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
<a name="variables.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="variables.v1alpha.Variable"></a>

### Variable



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for Variable object. |
| spec | [VariableSpec](#variables.v1alpha.VariableSpec) |  | Spec for Variable. |
<a name="variables.v1alpha.VariableSpec"></a>

### VariableSpec

Variables variables file structure. This covers the "Variables File" and "'From Host' Static Variable File".

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variables | [VariableSpec.VariablesEntry](#variables.v1alpha.VariableSpec.VariablesEntry) | repeated | Variables. |
| type | [string](#string) |  | Type of variable. |
<a name="variables.v1alpha.VariableSpec.VariablesEntry"></a>

### VariableSpec.VariablesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.Any](#google.protobuf.Any) |  |  |
<a name="variables.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |
<a name="variables.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Variable info for watch response. |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable for watch response. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="variables.v1alpha.Variables"></a>

### Variables



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#variables.v1alpha.GetRequest) | [GetResponse](#variables.v1alpha.GetResponse) | Get a Variable. |
| List | [ListRequest](#variables.v1alpha.ListRequest) | [ListResponse](#variables.v1alpha.ListResponse) | List Variable. |
| Add | [AddRequest](#variables.v1alpha.AddRequest) | [AddResponse](#variables.v1alpha.AddResponse) | Add a Variable. |
| Update | [UpdateRequest](#variables.v1alpha.UpdateRequest) | [UpdateResponse](#variables.v1alpha.UpdateResponse) | Update a Variable. |
| Delete | [DeleteRequest](#variables.v1alpha.DeleteRequest) | [DeleteResponse](#variables.v1alpha.DeleteResponse) | Delete a Variable. |
| Watch | [WatchRequest](#variables.v1alpha.WatchRequest) | [WatchResponse](#variables.v1alpha.WatchResponse) stream | Watch Variable. |

 <!-- end services -->



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha/api.proto



<a name="core.v1alpha.Condition"></a>

### Condition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [string](#string) |  |  |
| retry | [Retry](#core.v1alpha.Retry) |  |  |
<a name="core.v1alpha.Conditions"></a>

### Conditions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| when | [Condition](#core.v1alpha.Condition) |  | `when` condition. |
| success | [Condition](#core.v1alpha.Condition) |  | `success` condition. |
<a name="core.v1alpha.Error"></a>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | Code status code. |
| message | [string](#string) |  | Message error message (can be empty). |
<a name="core.v1alpha.GetOptions"></a>

### GetOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name used for listing. |
| labelSelectors | [string](#string) | repeated | Label selector in string form. |
<a name="core.v1alpha.Include"></a>

### Include



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of include (if no paths are given, will be used as file name/path). |
| paths | [string](#string) | repeated | List of paths. |
<a name="core.v1alpha.ListOptions"></a>

### ListOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| labelSelectors | [string](#string) | repeated | Label selector in string form. |
<a name="core.v1alpha.ObjectMetadata"></a>

### ObjectMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apiVersion | [string](#string) |  | APIVersion of object. |
| kind | [string](#string) |  | Kind of object. |
| uuid | [string](#string) |  | UUID of object. |
| name | [string](#string) |  | Name of object. |
| labels | [ObjectMetadata.LabelsEntry](#core.v1alpha.ObjectMetadata.LabelsEntry) | repeated | Labels of object. |
| annotations | [ObjectMetadata.AnnotationsEntry](#core.v1alpha.ObjectMetadata.AnnotationsEntry) | repeated | Annotations of object. |
<a name="core.v1alpha.ObjectMetadata.AnnotationsEntry"></a>

### ObjectMetadata.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |
<a name="core.v1alpha.ObjectMetadata.LabelsEntry"></a>

### ObjectMetadata.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |
<a name="core.v1alpha.Retry"></a>

### Retry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| limit | [int64](#int64) |  |  |
<a name="core.v1alpha.VersionRequest"></a>

### VersionRequest


<a name="core.v1alpha.VersionResponse"></a>

### VersionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  | Semversion compatible version number. |
| major | [int64](#int64) |  | Semversion major number. |
| minor | [int64](#int64) |  | Semversion minor number. |
| patch | [int64](#int64) |  | Semversion patch number. |
| state | [string](#string) |  | API state (e.g., alpha, beta, stable). |
<a name="core.v1alpha.WatchOptions"></a>

### WatchOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name used for listing. |
| labelSelectors | [string](#string) | repeated | Label selector in string form. |
| conitions | [string](#string) | repeated | Conditions to look for when watching Events on an object. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/triggers/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/triggers/v1alpha/api.proto



<a name="triggers.v1alpha.Action"></a>

### Action



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [string](#string) |  | Action name. |
| conditions | [core.v1alpha.Conditions](#core.v1alpha.Conditions) |  | Conditions. |
| runOptions | [RunOptions](#triggers.v1alpha.RunOptions) |  | RunOptions |
| parameters | [google.protobuf.Any](#google.protobuf.Any) | repeated | Map of string interface{} to be given to the action. |
<a name="triggers.v1alpha.AddRequest"></a>

### AddRequest

Add

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
<a name="triggers.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="triggers.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
<a name="triggers.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="triggers.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="triggers.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="triggers.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |
<a name="triggers.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| triggers | [Trigger](#triggers.v1alpha.Trigger) | repeated | Trigger list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="triggers.v1alpha.RunOptions"></a>

### RunOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sync | [bool](#bool) |  |  |
| serialize | [Serialize](#triggers.v1alpha.Serialize) |  |  |
| target | [Target](#triggers.v1alpha.Target) |  |  |
<a name="triggers.v1alpha.Serialize"></a>

### Serialize



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  | Count. |
| failAfterAllTargetsRan | [bool](#bool) |  | FailAfterAllTargetsRan. |
<a name="triggers.v1alpha.Step"></a>

### Step



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| spec | [Action](#triggers.v1alpha.Action) |  |  |
<a name="triggers.v1alpha.Target"></a>

### Target



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hosts | [string](#string) | repeated | List of hosts (either name or label selector). |
| limit | [int64](#int64) |  | Limit of how many hosts of the matches should be used. |
<a name="triggers.v1alpha.Trigger"></a>

### Trigger



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for Trigger object. |
| spec | [TriggerSpec](#triggers.v1alpha.TriggerSpec) |  | Spec for Trigger. |
<a name="triggers.v1alpha.TriggerSpec"></a>

### TriggerSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| runOptions | [RunOptions](#triggers.v1alpha.RunOptions) |  | RunOptions. |
| steps | [Step](#triggers.v1alpha.Step) | repeated | Steps in this Trigger. |
<a name="triggers.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
<a name="triggers.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="triggers.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |
<a name="triggers.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Trigger info for watch response. |
| task | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger for watch response. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="triggers.v1alpha.Triggers"></a>

### Triggers



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#triggers.v1alpha.GetRequest) | [GetResponse](#triggers.v1alpha.GetResponse) | Get a Trigger. |
| List | [ListRequest](#triggers.v1alpha.ListRequest) | [ListResponse](#triggers.v1alpha.ListResponse) | List Triggers. |
| Add | [AddRequest](#triggers.v1alpha.AddRequest) | [AddResponse](#triggers.v1alpha.AddResponse) | Add a Trigger. |
| Update | [UpdateRequest](#triggers.v1alpha.UpdateRequest) | [UpdateResponse](#triggers.v1alpha.UpdateResponse) | Update a Trigger. |
| Delete | [DeleteRequest](#triggers.v1alpha.DeleteRequest) | [DeleteResponse](#triggers.v1alpha.DeleteResponse) | Delete a Trigger. |
| Watch | [WatchRequest](#triggers.v1alpha.WatchRequest) | [WatchResponse](#triggers.v1alpha.WatchResponse) stream | Watch Trigger. |

 <!-- end services -->



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha/api.proto



<a name="events.v1alpha.AddRequest"></a>

### AddRequest

Add

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
<a name="events.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="events.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
<a name="events.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="events.v1alpha.Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for the Event object. |
| spec | [EventSpec](#events.v1alpha.EventSpec) |  | Spec for Event. |
<a name="events.v1alpha.EventSpec"></a>

### EventSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Type of the Event. |
| summary | [string](#string) |  | Summary of Event. |
| data | [google.protobuf.Any](#google.protobuf.Any) |  | Data of Event. |
<a name="events.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="events.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="events.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |
<a name="events.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| events | [Event](#events.v1alpha.Event) | repeated | Event list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="events.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
<a name="events.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="events.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |
<a name="events.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#events.v1alpha.Event) |  | Event for watch stream. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="events.v1alpha.Events"></a>

### Events



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#events.v1alpha.GetRequest) | [GetResponse](#events.v1alpha.GetResponse) | Get a Event. |
| List | [ListRequest](#events.v1alpha.ListRequest) | [ListResponse](#events.v1alpha.ListResponse) | List Events. |
| Add | [AddRequest](#events.v1alpha.AddRequest) | [AddResponse](#events.v1alpha.AddResponse) | Add a Event. |
| Update | [UpdateRequest](#events.v1alpha.UpdateRequest) | [UpdateResponse](#events.v1alpha.UpdateResponse) | Update a Event. |
| Delete | [DeleteRequest](#events.v1alpha.DeleteRequest) | [DeleteResponse](#events.v1alpha.DeleteResponse) | Delete a Event. |
| Watch | [WatchRequest](#events.v1alpha.WatchRequest) | [WatchResponse](#events.v1alpha.WatchResponse) stream | Watch for Events. |

 <!-- end services -->



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto



<a name="nodes.v1alpha.AddRequest"></a>

### AddRequest

Add

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
<a name="nodes.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="nodes.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
<a name="nodes.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="nodes.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="nodes.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="nodes.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |
<a name="nodes.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodes | [Node](#nodes.v1alpha.Node) | repeated | Node list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="nodes.v1alpha.Network"></a>

### Network



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| daemonPorts | [Port](#nodes.v1alpha.Port) | repeated | Ports of edenconfmgmt daemon (should only be two right now, 1. GRPC, 2. HTTP). |
| fqdn | [string](#string) |  | FQDN of Node. |
| addresses | [string](#string) | repeated | IP addresses of Node. |
<a name="nodes.v1alpha.Node"></a>

### Node



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for Node object. |
| spec | [Spec](#nodes.v1alpha.Spec) |  | Spec for Node. |
| status | [Status](#nodes.v1alpha.Status) |  | Status for Node. |
<a name="nodes.v1alpha.Port"></a>

### Port



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of port. |
| port | [int32](#int32) |  | Port number. |
| protocol | [string](#string) |  | Protocol name. |
<a name="nodes.v1alpha.Spec"></a>

### Spec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| network | [Network](#nodes.v1alpha.Network) |  | Network specifications of Node. |
<a name="nodes.v1alpha.Status"></a>

### Status



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [string](#string) |  | State of Node. |
<a name="nodes.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
<a name="nodes.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="nodes.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |
<a name="nodes.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Event info for watch response. |
| node | [Node](#nodes.v1alpha.Node) |  | Node for watch response. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="nodes.v1alpha.Nodes"></a>

### Nodes



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#nodes.v1alpha.GetRequest) | [GetResponse](#nodes.v1alpha.GetResponse) | Get a Node. |
| List | [ListRequest](#nodes.v1alpha.ListRequest) | [ListResponse](#nodes.v1alpha.ListResponse) | List Nodes. |
| Add | [AddRequest](#nodes.v1alpha.AddRequest) | [AddResponse](#nodes.v1alpha.AddResponse) | Add a Node. |
| Update | [UpdateRequest](#nodes.v1alpha.UpdateRequest) | [UpdateResponse](#nodes.v1alpha.UpdateResponse) | Update a Node. |
| Delete | [DeleteRequest](#nodes.v1alpha.DeleteRequest) | [DeleteResponse](#nodes.v1alpha.DeleteResponse) | Delete a Node. |
| Watch | [WatchRequest](#nodes.v1alpha.WatchRequest) | [WatchResponse](#nodes.v1alpha.WatchResponse) stream | Watch a Node(s). |

 <!-- end services -->



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/cronjobs/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/cronjobs/v1alpha/api.proto



<a name="cronjobs.v1alpha.AddRequest"></a>

### AddRequest

Add

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob object. |
<a name="cronjobs.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="cronjobs.v1alpha.CronJob"></a>

### CronJob



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for CronJob object. |
| spec | [CronJobSpec](#cronjobs.v1alpha.CronJobSpec) |  | Spec for CronJob. |
<a name="cronjobs.v1alpha.CronJobSpec"></a>

### CronJobSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedule | [Schedule](#cronjobs.v1alpha.Schedule) |  | Schedule of CronJob. |
<a name="cronjobs.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob object. |
<a name="cronjobs.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="cronjobs.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="cronjobs.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="cronjobs.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |
<a name="cronjobs.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cronjobs | [CronJob](#cronjobs.v1alpha.CronJob) | repeated | Config list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="cronjobs.v1alpha.Schedule"></a>

### Schedule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| crontTab | [string](#string) |  | Crontab style schedule time. |
<a name="cronjobs.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob object. |
<a name="cronjobs.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="cronjobs.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |
<a name="cronjobs.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Event info for watch response. |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob for watch response. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="cronjobs.v1alpha.CronJobs"></a>

### CronJobs



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#cronjobs.v1alpha.GetRequest) | [GetResponse](#cronjobs.v1alpha.GetResponse) | Get a Job. |
| List | [ListRequest](#cronjobs.v1alpha.ListRequest) | [ListResponse](#cronjobs.v1alpha.ListResponse) | List CronJobs. |
| Add | [AddRequest](#cronjobs.v1alpha.AddRequest) | [AddResponse](#cronjobs.v1alpha.AddResponse) | Add a Job. |
| Update | [UpdateRequest](#cronjobs.v1alpha.UpdateRequest) | [UpdateResponse](#cronjobs.v1alpha.UpdateResponse) | Update a Job. |
| Delete | [DeleteRequest](#cronjobs.v1alpha.DeleteRequest) | [DeleteResponse](#cronjobs.v1alpha.DeleteResponse) | Delete a Job. |
| Watch | [WatchRequest](#cronjobs.v1alpha.WatchRequest) | [WatchResponse](#cronjobs.v1alpha.WatchResponse) stream | Watch Job. |

 <!-- end services -->



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/taskbooks/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/taskbooks/v1alpha/api.proto



<a name="taskbooks.v1alpha.Action"></a>

### Action



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [string](#string) |  | Action name. |
| conditions | [core.v1alpha.Conditions](#core.v1alpha.Conditions) |  | Conditions. |
| runOptions | [RunOptions](#taskbooks.v1alpha.RunOptions) |  | RunOptions |
| parameters | [google.protobuf.Any](#google.protobuf.Any) | repeated | Map of string interface{} to be given to the action. |
<a name="taskbooks.v1alpha.AddRequest"></a>

### AddRequest

Add

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
<a name="taskbooks.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="taskbooks.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
<a name="taskbooks.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="taskbooks.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="taskbooks.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="taskbooks.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |
<a name="taskbooks.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| taskbooks | [TaskBook](#taskbooks.v1alpha.TaskBook) | repeated | TaskBook list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="taskbooks.v1alpha.RunOptions"></a>

### RunOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sync | [bool](#bool) |  |  |
| serialize | [Serialize](#taskbooks.v1alpha.Serialize) |  |  |
| target | [Target](#taskbooks.v1alpha.Target) |  |  |
<a name="taskbooks.v1alpha.Serialize"></a>

### Serialize



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  | Count. |
| failAfterAllTargetsRan | [bool](#bool) |  | FailAfterAllTargetsRan. |
<a name="taskbooks.v1alpha.Step"></a>

### Step



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| spec | [Action](#taskbooks.v1alpha.Action) |  |  |
<a name="taskbooks.v1alpha.Target"></a>

### Target



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hosts | [string](#string) | repeated | List of hosts (either name or label selector). |
| limit | [int64](#int64) |  | Limit of how many hosts of the matches should be used. |
<a name="taskbooks.v1alpha.TaskBook"></a>

### TaskBook



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for TaskBook object. |
| spec | [TaskBookSpec](#taskbooks.v1alpha.TaskBookSpec) |  | Spec for TaskBook. |
<a name="taskbooks.v1alpha.TaskBookSpec"></a>

### TaskBookSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| runOptions | [RunOptions](#taskbooks.v1alpha.RunOptions) |  | RunOptions. |
| steps | [Step](#taskbooks.v1alpha.Step) | repeated | Steps in this TaskBook. |
<a name="taskbooks.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
<a name="taskbooks.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<a name="taskbooks.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |
<a name="taskbooks.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | TaskBook info for watch response. |
| task | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook for watch response. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="taskbooks.v1alpha.TaskBookBooks"></a>

### TaskBookBooks



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#taskbooks.v1alpha.GetRequest) | [GetResponse](#taskbooks.v1alpha.GetResponse) | Get a TaskBook. |
| List | [ListRequest](#taskbooks.v1alpha.ListRequest) | [ListResponse](#taskbooks.v1alpha.ListResponse) | List TaskBookBooks. |
| Add | [AddRequest](#taskbooks.v1alpha.AddRequest) | [AddResponse](#taskbooks.v1alpha.AddResponse) | Add a TaskBook. |
| Update | [UpdateRequest](#taskbooks.v1alpha.UpdateRequest) | [UpdateResponse](#taskbooks.v1alpha.UpdateResponse) | Update a TaskBook. |
| Delete | [DeleteRequest](#taskbooks.v1alpha.DeleteRequest) | [DeleteResponse](#taskbooks.v1alpha.DeleteResponse) | Delete a TaskBook. |
| Watch | [WatchRequest](#taskbooks.v1alpha.WatchRequest) | [WatchResponse](#taskbooks.v1alpha.WatchResponse) stream | Watch TaskBook. |

 <!-- end services -->

## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |
