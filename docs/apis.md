# Protocol Documentation
<a name="top"></a>

## Table of Contents

* [github.com/galexrt/edenrun/pkg/grpcapi/clustervariables/v1alpha/api.proto](#github.com/galexrt/edenrun/pkg/grpcapi/clustervariables/v1alpha/api.proto)
    * Description: 
    * Services
        * [ClusterVariables](#clustervariables.v1alpha.ClusterVariables)
    * Messages
        * [ClusterVariable](#clustervariables.v1alpha.ClusterVariable)
        * [ClusterVariableList](#clustervariables.v1alpha.ClusterVariableList)
        * [CreateRequest](#clustervariables.v1alpha.CreateRequest)
        * [CreateResponse](#clustervariables.v1alpha.CreateResponse)
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
* [github.com/galexrt/edenrun/pkg/grpcapi/configs/v1alpha/api.proto](#github.com/galexrt/edenrun/pkg/grpcapi/configs/v1alpha/api.proto)
    * Description: 
    * Services
        * [Configs](#configs.v1alpha.Configs)
    * Messages
        * [Config](#configs.v1alpha.Config)
        * [ConfigList](#configs.v1alpha.ConfigList)
        * [ConfigSpec](#configs.v1alpha.ConfigSpec)
        * [CreateRequest](#configs.v1alpha.CreateRequest)
        * [CreateResponse](#configs.v1alpha.CreateResponse)
        * [DataStore](#configs.v1alpha.DataStore)
        * [DeleteRequest](#configs.v1alpha.DeleteRequest)
        * [DeleteResponse](#configs.v1alpha.DeleteResponse)
        * [ETCD](#configs.v1alpha.ETCD)
        * [GetRequest](#configs.v1alpha.GetRequest)
        * [GetResponse](#configs.v1alpha.GetResponse)
        * [ListRequest](#configs.v1alpha.ListRequest)
        * [ListResponse](#configs.v1alpha.ListResponse)
        * [Listen](#configs.v1alpha.Listen)
        * [Logging](#configs.v1alpha.Logging)
        * [UpdateRequest](#configs.v1alpha.UpdateRequest)
        * [UpdateResponse](#configs.v1alpha.UpdateResponse)
        * [WatchRequest](#configs.v1alpha.WatchRequest)
        * [WatchResponse](#configs.v1alpha.WatchResponse)
    * Enums
        * [Logging.LogLevel](#configs.v1alpha.Logging.LogLevel)
* [github.com/galexrt/edenrun/pkg/grpcapi/secrets/v1alpha/api.proto](#github.com/galexrt/edenrun/pkg/grpcapi/secrets/v1alpha/api.proto)
    * Description: 
    * Services
        * [Secrets](#secrets.v1alpha.Secrets)
    * Messages
        * [CreateRequest](#secrets.v1alpha.CreateRequest)
        * [CreateResponse](#secrets.v1alpha.CreateResponse)
        * [DeleteRequest](#secrets.v1alpha.DeleteRequest)
        * [DeleteResponse](#secrets.v1alpha.DeleteResponse)
        * [GetRequest](#secrets.v1alpha.GetRequest)
        * [GetResponse](#secrets.v1alpha.GetResponse)
        * [ListRequest](#secrets.v1alpha.ListRequest)
        * [ListResponse](#secrets.v1alpha.ListResponse)
        * [Secret](#secrets.v1alpha.Secret)
        * [SecretList](#secrets.v1alpha.SecretList)
        * [SecretSpec](#secrets.v1alpha.SecretSpec)
        * [SecretSpec.SecretsEntry](#secrets.v1alpha.SecretSpec.SecretsEntry)
        * [UpdateRequest](#secrets.v1alpha.UpdateRequest)
        * [UpdateResponse](#secrets.v1alpha.UpdateResponse)
        * [WatchRequest](#secrets.v1alpha.WatchRequest)
        * [WatchResponse](#secrets.v1alpha.WatchResponse)
* [github.com/galexrt/edenrun/pkg/grpcapi/beacons/v1alpha/api.proto](#github.com/galexrt/edenrun/pkg/grpcapi/beacons/v1alpha/api.proto)
    * Description: 
    * Services
        * [Beacons](#beacons.v1alpha.Beacons)
    * Messages
        * [Beacon](#beacons.v1alpha.Beacon)
        * [BeaconList](#beacons.v1alpha.BeaconList)
        * [BeaconSpec](#beacons.v1alpha.BeaconSpec)
        * [CreateRequest](#beacons.v1alpha.CreateRequest)
        * [CreateResponse](#beacons.v1alpha.CreateResponse)
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
* [github.com/galexrt/edenrun/pkg/grpcapi/cronjobs/v1alpha/api.proto](#github.com/galexrt/edenrun/pkg/grpcapi/cronjobs/v1alpha/api.proto)
    * Description: 
    * Services
        * [CronJobs](#cronjobs.v1alpha.CronJobs)
    * Messages
        * [CreateRequest](#cronjobs.v1alpha.CreateRequest)
        * [CreateResponse](#cronjobs.v1alpha.CreateResponse)
        * [CronJob](#cronjobs.v1alpha.CronJob)
        * [CronJobList](#cronjobs.v1alpha.CronJobList)
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
* [github.com/galexrt/edenrun/pkg/grpcapi/nodes/v1alpha/api.proto](#github.com/galexrt/edenrun/pkg/grpcapi/nodes/v1alpha/api.proto)
    * Description: 
    * Services
        * [Nodes](#nodes.v1alpha.Nodes)
    * Messages
        * [CreateRequest](#nodes.v1alpha.CreateRequest)
        * [CreateResponse](#nodes.v1alpha.CreateResponse)
        * [DeleteRequest](#nodes.v1alpha.DeleteRequest)
        * [DeleteResponse](#nodes.v1alpha.DeleteResponse)
        * [GetRequest](#nodes.v1alpha.GetRequest)
        * [GetResponse](#nodes.v1alpha.GetResponse)
        * [ListRequest](#nodes.v1alpha.ListRequest)
        * [ListResponse](#nodes.v1alpha.ListResponse)
        * [Network](#nodes.v1alpha.Network)
        * [Node](#nodes.v1alpha.Node)
        * [NodeList](#nodes.v1alpha.NodeList)
        * [Port](#nodes.v1alpha.Port)
        * [Spec](#nodes.v1alpha.Spec)
        * [Status](#nodes.v1alpha.Status)
        * [UpdateRequest](#nodes.v1alpha.UpdateRequest)
        * [UpdateResponse](#nodes.v1alpha.UpdateResponse)
        * [WatchRequest](#nodes.v1alpha.WatchRequest)
        * [WatchResponse](#nodes.v1alpha.WatchResponse)
* [github.com/galexrt/edenrun/pkg/grpcapi/events/v1alpha/api.proto](#github.com/galexrt/edenrun/pkg/grpcapi/events/v1alpha/api.proto)
    * Description: 
    * Services
        * [Events](#events.v1alpha.Events)
    * Messages
        * [CreateRequest](#events.v1alpha.CreateRequest)
        * [CreateResponse](#events.v1alpha.CreateResponse)
        * [DeleteRequest](#events.v1alpha.DeleteRequest)
        * [DeleteResponse](#events.v1alpha.DeleteResponse)
        * [Event](#events.v1alpha.Event)
        * [EventList](#events.v1alpha.EventList)
        * [EventSpec](#events.v1alpha.EventSpec)
        * [GetRequest](#events.v1alpha.GetRequest)
        * [GetResponse](#events.v1alpha.GetResponse)
        * [ListRequest](#events.v1alpha.ListRequest)
        * [ListResponse](#events.v1alpha.ListResponse)
        * [UpdateRequest](#events.v1alpha.UpdateRequest)
        * [UpdateResponse](#events.v1alpha.UpdateResponse)
        * [WatchRequest](#events.v1alpha.WatchRequest)
        * [WatchResponse](#events.v1alpha.WatchResponse)
* [github.com/galexrt/edenrun/pkg/grpcapi/variables/v1alpha/api.proto](#github.com/galexrt/edenrun/pkg/grpcapi/variables/v1alpha/api.proto)
    * Description: 
    * Services
        * [Variables](#variables.v1alpha.Variables)
    * Messages
        * [CreateRequest](#variables.v1alpha.CreateRequest)
        * [CreateResponse](#variables.v1alpha.CreateResponse)
        * [DeleteRequest](#variables.v1alpha.DeleteRequest)
        * [DeleteResponse](#variables.v1alpha.DeleteResponse)
        * [GetRequest](#variables.v1alpha.GetRequest)
        * [GetResponse](#variables.v1alpha.GetResponse)
        * [ListRequest](#variables.v1alpha.ListRequest)
        * [ListResponse](#variables.v1alpha.ListResponse)
        * [UpdateRequest](#variables.v1alpha.UpdateRequest)
        * [UpdateResponse](#variables.v1alpha.UpdateResponse)
        * [Variable](#variables.v1alpha.Variable)
        * [VariableList](#variables.v1alpha.VariableList)
        * [VariableSpec](#variables.v1alpha.VariableSpec)
        * [VariableSpec.VariablesEntry](#variables.v1alpha.VariableSpec.VariablesEntry)
        * [WatchRequest](#variables.v1alpha.WatchRequest)
        * [WatchResponse](#variables.v1alpha.WatchResponse)
* [github.com/galexrt/edenrun/pkg/grpcapi/core/v1/api.proto](#github.com/galexrt/edenrun/pkg/grpcapi/core/v1/api.proto)
    * Description: 
    * Messages
        * [Condition](#core.v1.Condition)
        * [Conditions](#core.v1.Conditions)
        * [CreateOptions](#core.v1.CreateOptions)
        * [DeleteOptions](#core.v1.DeleteOptions)
        * [Error](#core.v1.Error)
        * [GetOptions](#core.v1.GetOptions)
        * [Include](#core.v1.Include)
        * [ListOptions](#core.v1.ListOptions)
        * [ObjectMetadata](#core.v1.ObjectMetadata)
        * [ObjectMetadata.AnnotationsEntry](#core.v1.ObjectMetadata.AnnotationsEntry)
        * [ObjectMetadata.LabelsEntry](#core.v1.ObjectMetadata.LabelsEntry)
        * [Retry](#core.v1.Retry)
        * [UpdateOptions](#core.v1.UpdateOptions)
        * [VersionRequest](#core.v1.VersionRequest)
        * [VersionResponse](#core.v1.VersionResponse)
        * [WatchOptions](#core.v1.WatchOptions)
* [github.com/galexrt/edenrun/pkg/grpcapi/triggers/v1alpha/api.proto](#github.com/galexrt/edenrun/pkg/grpcapi/triggers/v1alpha/api.proto)
    * Description: 
    * Services
        * [Triggers](#triggers.v1alpha.Triggers)
    * Messages
        * [Action](#triggers.v1alpha.Action)
        * [CreateRequest](#triggers.v1alpha.CreateRequest)
        * [CreateResponse](#triggers.v1alpha.CreateResponse)
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
        * [TriggerList](#triggers.v1alpha.TriggerList)
        * [TriggerSpec](#triggers.v1alpha.TriggerSpec)
        * [UpdateRequest](#triggers.v1alpha.UpdateRequest)
        * [UpdateResponse](#triggers.v1alpha.UpdateResponse)
        * [WatchRequest](#triggers.v1alpha.WatchRequest)
        * [WatchResponse](#triggers.v1alpha.WatchResponse)
* [github.com/galexrt/edenrun/pkg/grpcapi/taskbooks/v1alpha/api.proto](#github.com/galexrt/edenrun/pkg/grpcapi/taskbooks/v1alpha/api.proto)
    * Description: 
    * Services
        * [TaskBooks](#taskbooks.v1alpha.TaskBooks)
    * Messages
        * [Action](#taskbooks.v1alpha.Action)
        * [CreateRequest](#taskbooks.v1alpha.CreateRequest)
        * [CreateResponse](#taskbooks.v1alpha.CreateResponse)
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
        * [TaskBookList](#taskbooks.v1alpha.TaskBookList)
        * [TaskBookSpec](#taskbooks.v1alpha.TaskBookSpec)
        * [UpdateRequest](#taskbooks.v1alpha.UpdateRequest)
        * [UpdateResponse](#taskbooks.v1alpha.UpdateResponse)
        * [WatchRequest](#taskbooks.v1alpha.WatchRequest)
        * [WatchResponse](#taskbooks.v1alpha.WatchResponse)* [Scalar Value Types](#scalar-value-types)

<a name="github.com/galexrt/edenrun/pkg/grpcapi/clustervariables/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/grpcapi/clustervariables/v1alpha/api.proto



<a name="clustervariables.v1alpha.ClusterVariable"></a>

### ClusterVariable



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for ClusterVariable object. |
| spec | [VariableSpec](#clustervariables.v1alpha.VariableSpec) |  | Spec for ClusterVariable. |
<a name="clustervariables.v1alpha.ClusterVariableList"></a>

### ClusterVariableList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for ClusterVariableList object. |
| items | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) | repeated | List of ClusterVariable objects. |
<a name="clustervariables.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.CreateOptions](#core.v1.CreateOptions) |  | core_v1.CreateOptions |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="clustervariables.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="clustervariables.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.DeleteOptions](#core.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="clustervariables.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="clustervariables.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.GetOptions](#core.v1.GetOptions) |  | core_v1.GetOptions |
<a name="clustervariables.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="clustervariables.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.ListOptions](#core.v1.ListOptions) |  | core_v1.ListOptions |
<a name="clustervariables.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariableList | [ClusterVariableList](#clustervariables.v1alpha.ClusterVariableList) |  | ClusterVariable list. |
<a name="clustervariables.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.UpdateOptions](#core.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="clustervariables.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
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
| options | [core.v1.WatchOptions](#core.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="clustervariables.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | ClusterVariable info for watch response. |
| clusterVariable | [ClusterVariable](#clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable for watch response. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="clustervariables.v1alpha.ClusterVariables"></a>

### ClusterVariables



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#clustervariables.v1alpha.GetRequest) | [GetResponse](#clustervariables.v1alpha.GetResponse) | Get a ClusterVariable. |
| List | [ListRequest](#clustervariables.v1alpha.ListRequest) | [ListResponse](#clustervariables.v1alpha.ListResponse) | List ClusterVariable. |
| Create | [CreateRequest](#clustervariables.v1alpha.CreateRequest) | [CreateResponse](#clustervariables.v1alpha.CreateResponse) | Create a ClusterVariable. |
| Update | [UpdateRequest](#clustervariables.v1alpha.UpdateRequest) | [UpdateResponse](#clustervariables.v1alpha.UpdateResponse) | Update a ClusterVariable. |
| Delete | [DeleteRequest](#clustervariables.v1alpha.DeleteRequest) | [DeleteResponse](#clustervariables.v1alpha.DeleteResponse) | Delete a ClusterVariable. |
| Watch | [WatchRequest](#clustervariables.v1alpha.WatchRequest) | [WatchResponse](#clustervariables.v1alpha.WatchResponse) stream | Watch ClusterVariable. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/grpcapi/configs/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/grpcapi/configs/v1alpha/api.proto



<a name="configs.v1alpha.Config"></a>

### Config

Config object base.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for Config object. |
| spec | [ConfigSpec](#configs.v1alpha.ConfigSpec) |  | Spec for Config. |
<a name="configs.v1alpha.ConfigList"></a>

### ConfigList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for ConfigList object. |
| items | [Config](#configs.v1alpha.Config) | repeated | List of Config objects. |
<a name="configs.v1alpha.ConfigSpec"></a>

### ConfigSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logging | [Logging](#configs.v1alpha.Logging) |  | Logging |
| dataStore | [DataStore](#configs.v1alpha.DataStore) |  | DataStore. |
| listen | [Listen](#configs.v1alpha.Listen) |  | Listen. |
<a name="configs.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.CreateOptions](#core.v1.CreateOptions) |  | core_v1.CreateOptions |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
<a name="configs.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
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
| options | [core.v1.DeleteOptions](#core.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
<a name="configs.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
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
| options | [core.v1.GetOptions](#core.v1.GetOptions) |  | core_v1.GetOptions |
<a name="configs.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
<a name="configs.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.ListOptions](#core.v1.ListOptions) |  | core_v1.ListOptions |
<a name="configs.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configList | [ConfigList](#configs.v1alpha.ConfigList) |  | Config list. |
<a name="configs.v1alpha.Listen"></a>

### Listen

Listen

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| httpAddress | [string](#string) |  | HTTP Port |
| grpcAddress | [string](#string) |  | gRPC Port |
<a name="configs.v1alpha.Logging"></a>

### Logging



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mode | [string](#string) |  | Mode |
| level | [Logging.LogLevel](#configs.v1alpha.Logging.LogLevel) |  |  |
<a name="configs.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.UpdateOptions](#core.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
<a name="configs.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
<a name="configs.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.WatchOptions](#core.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="configs.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Event info for watch response. |
| config | [Config](#configs.v1alpha.Config) |  | Config object. |
<!-- end messages -->


<a name="configs.v1alpha.Logging.LogLevel"></a>

### Logging.LogLevel

LogLevel

| Name | Number | Description |
| ---- | ------ | ----------- |
| DEBUG | 0 |  |
| INFO | 1 |  |
| WARN | 2 |  |
| ERROR | 3 |  |
| FATAL | 4 |  |
| PANIC | 5 |  |


<!-- end enums -->

 <!-- end HasExtensions -->


<a name="configs.v1alpha.Configs"></a>

### Configs



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#configs.v1alpha.GetRequest) | [GetResponse](#configs.v1alpha.GetResponse) | Get a Config. |
| List | [ListRequest](#configs.v1alpha.ListRequest) | [ListResponse](#configs.v1alpha.ListResponse) | List Configs. |
| Create | [CreateRequest](#configs.v1alpha.CreateRequest) | [CreateResponse](#configs.v1alpha.CreateResponse) | Create a Config. |
| Update | [UpdateRequest](#configs.v1alpha.UpdateRequest) | [UpdateResponse](#configs.v1alpha.UpdateResponse) | Update a Config. |
| Delete | [DeleteRequest](#configs.v1alpha.DeleteRequest) | [DeleteResponse](#configs.v1alpha.DeleteResponse) | Delete a Config. |
| Watch | [WatchRequest](#configs.v1alpha.WatchRequest) | [WatchResponse](#configs.v1alpha.WatchResponse) stream | Watch Config. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/grpcapi/secrets/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/grpcapi/secrets/v1alpha/api.proto



<a name="secrets.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.CreateOptions](#core.v1.CreateOptions) |  | core_v1.CreateOptions |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
<a name="secrets.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
<a name="secrets.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.DeleteOptions](#core.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
<a name="secrets.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
<a name="secrets.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.GetOptions](#core.v1.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="secrets.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
<a name="secrets.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.ListOptions](#core.v1.ListOptions) |  | core_v1.ListOptions |
<a name="secrets.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| SecretList | [SecretList](#secrets.v1alpha.SecretList) |  | Secret list. |
<a name="secrets.v1alpha.Secret"></a>

### Secret



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for Secret object. |
| spec | [SecretSpec](#secrets.v1alpha.SecretSpec) |  | Spec for Secret. |
<a name="secrets.v1alpha.SecretList"></a>

### SecretList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for SecretList object. |
| items | [Secret](#secrets.v1alpha.Secret) | repeated | List of Secret objects. |
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
| options | [core.v1.UpdateOptions](#core.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
<a name="secrets.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret object. |
<a name="secrets.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.WatchOptions](#core.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="secrets.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Secret info for watch response. |
| secret | [Secret](#secrets.v1alpha.Secret) |  | Secret for watch response. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="secrets.v1alpha.Secrets"></a>

### Secrets



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#secrets.v1alpha.GetRequest) | [GetResponse](#secrets.v1alpha.GetResponse) | Get a Secret. |
| List | [ListRequest](#secrets.v1alpha.ListRequest) | [ListResponse](#secrets.v1alpha.ListResponse) | List Secret. |
| Create | [CreateRequest](#secrets.v1alpha.CreateRequest) | [CreateResponse](#secrets.v1alpha.CreateResponse) | Create a Secret. |
| Update | [UpdateRequest](#secrets.v1alpha.UpdateRequest) | [UpdateResponse](#secrets.v1alpha.UpdateResponse) | Update a Secret. |
| Delete | [DeleteRequest](#secrets.v1alpha.DeleteRequest) | [DeleteResponse](#secrets.v1alpha.DeleteResponse) | Delete a Secret. |
| Watch | [WatchRequest](#secrets.v1alpha.WatchRequest) | [WatchResponse](#secrets.v1alpha.WatchResponse) stream | Watch Secret. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/grpcapi/beacons/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/grpcapi/beacons/v1alpha/api.proto



<a name="beacons.v1alpha.Beacon"></a>

### Beacon



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for Beacon object. |
| spec | [BeaconSpec](#beacons.v1alpha.BeaconSpec) |  | Spec for Beacon. |
<a name="beacons.v1alpha.BeaconList"></a>

### BeaconList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for BeaconList object. |
| items | [Beacon](#beacons.v1alpha.Beacon) | repeated | List of Beacon objects. |
<a name="beacons.v1alpha.BeaconSpec"></a>

### BeaconSpec

TODO
<a name="beacons.v1alpha.CreateRequest"></a>

### CreateRequest

Create Beacon request.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.CreateOptions](#core.v1.CreateOptions) |  | core_v1.CreateOptions |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
<a name="beacons.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
<a name="beacons.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.DeleteOptions](#core.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
<a name="beacons.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
<a name="beacons.v1alpha.GetRequest"></a>

### GetRequest

Get request content.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.GetOptions](#core.v1.GetOptions) |  | core_v1.GetOptions |
<a name="beacons.v1alpha.GetResponse"></a>

### GetResponse

Get response content.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
<a name="beacons.v1alpha.ListRequest"></a>

### ListRequest

List request content.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.ListOptions](#core.v1.ListOptions) |  | core_v1.ListOptions |
<a name="beacons.v1alpha.ListResponse"></a>

### ListResponse

List response content.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beaconList | [BeaconList](#beacons.v1alpha.BeaconList) |  | Beacon list. |
<a name="beacons.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.UpdateOptions](#core.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
<a name="beacons.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon object. |
<a name="beacons.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.WatchOptions](#core.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="beacons.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Beacon info for watch response. |
| beacon | [Beacon](#beacons.v1alpha.Beacon) |  | Beacon for watch response. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="beacons.v1alpha.Beacons"></a>

### Beacons



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#beacons.v1alpha.GetRequest) | [GetResponse](#beacons.v1alpha.GetResponse) | Get a Beacon. |
| List | [ListRequest](#beacons.v1alpha.ListRequest) | [ListResponse](#beacons.v1alpha.ListResponse) | List Beacons. |
| Create | [CreateRequest](#beacons.v1alpha.CreateRequest) | [CreateResponse](#beacons.v1alpha.CreateResponse) | Create a Beacon. |
| Update | [UpdateRequest](#beacons.v1alpha.UpdateRequest) | [UpdateResponse](#beacons.v1alpha.UpdateResponse) | Update a Beacon. |
| Delete | [DeleteRequest](#beacons.v1alpha.DeleteRequest) | [DeleteResponse](#beacons.v1alpha.DeleteResponse) | Delete a Beacon. |
| Watch | [WatchRequest](#beacons.v1alpha.WatchRequest) | [WatchResponse](#beacons.v1alpha.WatchResponse) stream | Watch Beacon. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/grpcapi/cronjobs/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/grpcapi/cronjobs/v1alpha/api.proto



<a name="cronjobs.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.CreateOptions](#core.v1.CreateOptions) |  | core_v1.CreateOptions |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob |
<a name="cronjobs.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob object. |
<a name="cronjobs.v1alpha.CronJob"></a>

### CronJob



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for CronJob object. |
| spec | [CronJobSpec](#cronjobs.v1alpha.CronJobSpec) |  | Spec for CronJob. |
<a name="cronjobs.v1alpha.CronJobList"></a>

### CronJobList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for CronJobList object. |
| items | [CronJob](#cronjobs.v1alpha.CronJob) | repeated | List of CronJob objects. |
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
| options | [core.v1.DeleteOptions](#core.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob object. |
<a name="cronjobs.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob object. |
<a name="cronjobs.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.GetOptions](#core.v1.GetOptions) |  | core_v1.GetOptions |
<a name="cronjobs.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob object. |
<a name="cronjobs.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.ListOptions](#core.v1.ListOptions) |  | core_v1.ListOptions |
<a name="cronjobs.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cronJobList | [CronJobList](#cronjobs.v1alpha.CronJobList) |  | Config list. |
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
| options | [core.v1.UpdateOptions](#core.v1.UpdateOptions) |  | core_v1.CreateOptions |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob object. |
<a name="cronjobs.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob object. |
<a name="cronjobs.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.WatchOptions](#core.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="cronjobs.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Event info for watch response. |
| cronJob | [CronJob](#cronjobs.v1alpha.CronJob) |  | CronJob for watch response. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="cronjobs.v1alpha.CronJobs"></a>

### CronJobs



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#cronjobs.v1alpha.GetRequest) | [GetResponse](#cronjobs.v1alpha.GetResponse) | Get a Job. |
| List | [ListRequest](#cronjobs.v1alpha.ListRequest) | [ListResponse](#cronjobs.v1alpha.ListResponse) | List CronJobs. |
| Create | [CreateRequest](#cronjobs.v1alpha.CreateRequest) | [CreateResponse](#cronjobs.v1alpha.CreateResponse) | Create a Job. |
| Update | [UpdateRequest](#cronjobs.v1alpha.UpdateRequest) | [UpdateResponse](#cronjobs.v1alpha.UpdateResponse) | Update a Job. |
| Delete | [DeleteRequest](#cronjobs.v1alpha.DeleteRequest) | [DeleteResponse](#cronjobs.v1alpha.DeleteResponse) | Delete a Job. |
| Watch | [WatchRequest](#cronjobs.v1alpha.WatchRequest) | [WatchResponse](#cronjobs.v1alpha.WatchResponse) stream | Watch Job. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/grpcapi/nodes/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/grpcapi/nodes/v1alpha/api.proto



<a name="nodes.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.CreateOptions](#core.v1.CreateOptions) |  | core_v1.CreateOptions |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
<a name="nodes.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
<a name="nodes.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.DeleteOptions](#core.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
<a name="nodes.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
<a name="nodes.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.GetOptions](#core.v1.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="nodes.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
<a name="nodes.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.ListOptions](#core.v1.ListOptions) |  | core_v1.ListOptions |
<a name="nodes.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nodeList | [NodeList](#nodes.v1alpha.NodeList) |  | Node list. |
<a name="nodes.v1alpha.Network"></a>

### Network



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| daemonPorts | [Port](#nodes.v1alpha.Port) | repeated | Ports of edenrun daemon (should only be two right now, 1. GRPC, 2. HTTP). |
| fqdn | [string](#string) |  | FQDN of Node. |
| addresses | [string](#string) | repeated | IP addresses of Node. |
<a name="nodes.v1alpha.Node"></a>

### Node



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for Node object. |
| spec | [Spec](#nodes.v1alpha.Spec) |  | Spec for Node. |
| status | [Status](#nodes.v1alpha.Status) |  | Status for Node. |
<a name="nodes.v1alpha.NodeList"></a>

### NodeList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for NodeList object. |
| items | [Node](#nodes.v1alpha.Node) | repeated | List of Node objects. |
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
| options | [core.v1.UpdateOptions](#core.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
<a name="nodes.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| node | [Node](#nodes.v1alpha.Node) |  | Node object. |
<a name="nodes.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.WatchOptions](#core.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="nodes.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Event info for watch response. |
| node | [Node](#nodes.v1alpha.Node) |  | Node for watch response. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="nodes.v1alpha.Nodes"></a>

### Nodes



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#nodes.v1alpha.GetRequest) | [GetResponse](#nodes.v1alpha.GetResponse) | Get a Node. |
| List | [ListRequest](#nodes.v1alpha.ListRequest) | [ListResponse](#nodes.v1alpha.ListResponse) | List Nodes. |
| Create | [CreateRequest](#nodes.v1alpha.CreateRequest) | [CreateResponse](#nodes.v1alpha.CreateResponse) | Create a Node. |
| Update | [UpdateRequest](#nodes.v1alpha.UpdateRequest) | [UpdateResponse](#nodes.v1alpha.UpdateResponse) | Update a Node. |
| Delete | [DeleteRequest](#nodes.v1alpha.DeleteRequest) | [DeleteResponse](#nodes.v1alpha.DeleteResponse) | Delete a Node. |
| Watch | [WatchRequest](#nodes.v1alpha.WatchRequest) | [WatchResponse](#nodes.v1alpha.WatchResponse) stream | Watch a Node(s). |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/grpcapi/events/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/grpcapi/events/v1alpha/api.proto



<a name="events.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.CreateOptions](#core.v1.CreateOptions) |  | core_v1.CreateOptions |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
<a name="events.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
<a name="events.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.DeleteOptions](#core.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
<a name="events.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
<a name="events.v1alpha.Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for the Event object. |
| spec | [EventSpec](#events.v1alpha.EventSpec) |  | Spec for Event. |
<a name="events.v1alpha.EventList"></a>

### EventList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for EventList object. |
| items | [Event](#events.v1alpha.Event) | repeated | List of Event objects. |
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
| options | [core.v1.GetOptions](#core.v1.GetOptions) |  | core_v1.GetOptions |
<a name="events.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
<a name="events.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.ListOptions](#core.v1.ListOptions) |  | core_v1.ListOptions |
<a name="events.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| eventList | [EventList](#events.v1alpha.EventList) |  | Event list. |
<a name="events.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.UpdateOptions](#core.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
<a name="events.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#events.v1alpha.Event) |  | Event object. |
<a name="events.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.WatchOptions](#core.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="events.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#events.v1alpha.Event) |  | Event for watch stream. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="events.v1alpha.Events"></a>

### Events



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#events.v1alpha.GetRequest) | [GetResponse](#events.v1alpha.GetResponse) | Get a Event. |
| List | [ListRequest](#events.v1alpha.ListRequest) | [ListResponse](#events.v1alpha.ListResponse) | List Events. |
| Create | [CreateRequest](#events.v1alpha.CreateRequest) | [CreateResponse](#events.v1alpha.CreateResponse) | Create a Event. |
| Update | [UpdateRequest](#events.v1alpha.UpdateRequest) | [UpdateResponse](#events.v1alpha.UpdateResponse) | Update a Event. |
| Delete | [DeleteRequest](#events.v1alpha.DeleteRequest) | [DeleteResponse](#events.v1alpha.DeleteResponse) | Delete a Event. |
| Watch | [WatchRequest](#events.v1alpha.WatchRequest) | [WatchResponse](#events.v1alpha.WatchResponse) stream | Watch for Events. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/grpcapi/variables/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/grpcapi/variables/v1alpha/api.proto



<a name="variables.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.CreateOptions](#core.v1.CreateOptions) |  | core_v1.CreateOptions |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
<a name="variables.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
<a name="variables.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.DeleteOptions](#core.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
<a name="variables.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
<a name="variables.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.GetOptions](#core.v1.GetOptions) |  | core_v1.GetOptions |
<a name="variables.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
<a name="variables.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.ListOptions](#core.v1.ListOptions) |  | core_v1.ListOptions |
<a name="variables.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variableList | [VariableList](#variables.v1alpha.VariableList) |  | Variable list. |
<a name="variables.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.UpdateOptions](#core.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
<a name="variables.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable object. |
<a name="variables.v1alpha.Variable"></a>

### Variable



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for Variable object. |
| spec | [VariableSpec](#variables.v1alpha.VariableSpec) |  | Spec for Variable. |
<a name="variables.v1alpha.VariableList"></a>

### VariableList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for VariableList object. |
| items | [Variable](#variables.v1alpha.Variable) | repeated | List of Variable objects. |
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
| options | [core.v1.WatchOptions](#core.v1.WatchOptions) |  | WatchOptions options for WatchRequest. |
<a name="variables.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Variable info for watch response. |
| variable | [Variable](#variables.v1alpha.Variable) |  | Variable for watch response. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="variables.v1alpha.Variables"></a>

### Variables



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#variables.v1alpha.GetRequest) | [GetResponse](#variables.v1alpha.GetResponse) | Get a Variable. |
| List | [ListRequest](#variables.v1alpha.ListRequest) | [ListResponse](#variables.v1alpha.ListResponse) | List Variable. |
| Create | [CreateRequest](#variables.v1alpha.CreateRequest) | [CreateResponse](#variables.v1alpha.CreateResponse) | Create a Variable. |
| Update | [UpdateRequest](#variables.v1alpha.UpdateRequest) | [UpdateResponse](#variables.v1alpha.UpdateResponse) | Update a Variable. |
| Delete | [DeleteRequest](#variables.v1alpha.DeleteRequest) | [DeleteResponse](#variables.v1alpha.DeleteResponse) | Delete a Variable. |
| Watch | [WatchRequest](#variables.v1alpha.WatchRequest) | [WatchResponse](#variables.v1alpha.WatchResponse) stream | Watch Variable. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/grpcapi/core/v1/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/grpcapi/core/v1/api.proto



<a name="core.v1.Condition"></a>

### Condition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [string](#string) |  |  |
| retry | [Retry](#core.v1.Retry) |  |  |
<a name="core.v1.Conditions"></a>

### Conditions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| when | [Condition](#core.v1.Condition) |  | `when` condition. |
| success | [Condition](#core.v1.Condition) |  | `success` condition. |
<a name="core.v1.CreateOptions"></a>

### CreateOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  | Namespace used for listing. |
<a name="core.v1.DeleteOptions"></a>

### DeleteOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  | Namespace used for listing. |
| labelSelectors | [string](#string) | repeated | Label selector in string form. |
<a name="core.v1.Error"></a>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | Code status code. |
| message | [string](#string) |  | Message error message (can be empty). |
<a name="core.v1.GetOptions"></a>

### GetOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name used for listing. |
| namespace | [string](#string) |  | Namespace used for listing. |
| labelSelectors | [string](#string) | repeated | Label selector in string form. |
<a name="core.v1.Include"></a>

### Include



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of include (if no paths are given, will be used as file name/path). |
| paths | [string](#string) | repeated | List of paths. |
<a name="core.v1.ListOptions"></a>

### ListOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name used for listing. |
| namespace | [string](#string) |  | Namespace used for listing. |
| labelSelectors | [string](#string) | repeated | Label selector in string form. |
<a name="core.v1.ObjectMetadata"></a>

### ObjectMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apiVersion | [string](#string) |  | APIVersion of object. |
| kind | [string](#string) |  | Kind of object. |
| uuid | [string](#string) |  | UUID of object. |
| name | [string](#string) |  | Name of object. |
| namespace | [string](#string) |  | Namespace of object. |
| labels | [ObjectMetadata.LabelsEntry](#core.v1.ObjectMetadata.LabelsEntry) | repeated | Labels of object. |
| annotations | [ObjectMetadata.AnnotationsEntry](#core.v1.ObjectMetadata.AnnotationsEntry) | repeated | Annotations of object. |
<a name="core.v1.ObjectMetadata.AnnotationsEntry"></a>

### ObjectMetadata.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |
<a name="core.v1.ObjectMetadata.LabelsEntry"></a>

### ObjectMetadata.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |
<a name="core.v1.Retry"></a>

### Retry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| limit | [int64](#int64) |  |  |
<a name="core.v1.UpdateOptions"></a>

### UpdateOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  | Namespace used for listing. |
<a name="core.v1.VersionRequest"></a>

### VersionRequest


<a name="core.v1.VersionResponse"></a>

### VersionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  | Semversion compatible version number. |
| major | [int64](#int64) |  | Semversion major number. |
| minor | [int64](#int64) |  | Semversion minor number. |
| patch | [int64](#int64) |  | Semversion patch number. |
| state | [string](#string) |  | API state (e.g., alpha, beta, stable). |
<a name="core.v1.WatchOptions"></a>

### WatchOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name used for listing. |
| namespace | [string](#string) |  | Namespace used for listing. |
| labelSelectors | [string](#string) | repeated | Label selector in string form. |
| conitions | [string](#string) | repeated | Conditions to look for when watching Events on an object. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/grpcapi/triggers/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/grpcapi/triggers/v1alpha/api.proto



<a name="triggers.v1alpha.Action"></a>

### Action



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [string](#string) |  | Action name. |
| conditions | [core.v1.Conditions](#core.v1.Conditions) |  | Conditions. |
| runOptions | [RunOptions](#triggers.v1alpha.RunOptions) |  | RunOptions |
| parameters | [google.protobuf.Any](#google.protobuf.Any) | repeated | Map of string interface{} to be given to the action. |
<a name="triggers.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.CreateOptions](#core.v1.CreateOptions) |  | core_v1.CreateOptions |
| trigger | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
<a name="triggers.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trigger | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
<a name="triggers.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.DeleteOptions](#core.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| trigger | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
<a name="triggers.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trigger | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
<a name="triggers.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.GetOptions](#core.v1.GetOptions) |  | core_v1.GetOptions |
<a name="triggers.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trigger | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
<a name="triggers.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.ListOptions](#core.v1.ListOptions) |  | core_v1.ListOptions |
<a name="triggers.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| triggerList | [TriggerList](#triggers.v1alpha.TriggerList) |  | Trigger list. |
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
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for Trigger object. |
| spec | [TriggerSpec](#triggers.v1alpha.TriggerSpec) |  | Spec for Trigger. |
<a name="triggers.v1alpha.TriggerList"></a>

### TriggerList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for TriggerList object. |
| items | [Trigger](#triggers.v1alpha.Trigger) | repeated | List of Trigger objects. |
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
| options | [core.v1.UpdateOptions](#core.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| trigger | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
<a name="triggers.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trigger | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger object. |
<a name="triggers.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.WatchOptions](#core.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="triggers.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Trigger info for watch response. |
| trigger | [Trigger](#triggers.v1alpha.Trigger) |  | Trigger for watch response. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="triggers.v1alpha.Triggers"></a>

### Triggers



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#triggers.v1alpha.GetRequest) | [GetResponse](#triggers.v1alpha.GetResponse) | Get a Trigger. |
| List | [ListRequest](#triggers.v1alpha.ListRequest) | [ListResponse](#triggers.v1alpha.ListResponse) | List Triggers. |
| Create | [CreateRequest](#triggers.v1alpha.CreateRequest) | [CreateResponse](#triggers.v1alpha.CreateResponse) | Create a Trigger. |
| Update | [UpdateRequest](#triggers.v1alpha.UpdateRequest) | [UpdateResponse](#triggers.v1alpha.UpdateResponse) | Update a Trigger. |
| Delete | [DeleteRequest](#triggers.v1alpha.DeleteRequest) | [DeleteResponse](#triggers.v1alpha.DeleteResponse) | Delete a Trigger. |
| Watch | [WatchRequest](#triggers.v1alpha.WatchRequest) | [WatchResponse](#triggers.v1alpha.WatchResponse) stream | Watch Trigger. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/grpcapi/taskbooks/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/grpcapi/taskbooks/v1alpha/api.proto



<a name="taskbooks.v1alpha.Action"></a>

### Action



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [string](#string) |  | Action name. |
| conditions | [core.v1.Conditions](#core.v1.Conditions) |  | Conditions. |
| runOptions | [RunOptions](#taskbooks.v1alpha.RunOptions) |  | RunOptions |
| parameters | [google.protobuf.Any](#google.protobuf.Any) | repeated | Map of string interface{} to be given to the action. |
<a name="taskbooks.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.CreateOptions](#core.v1.CreateOptions) |  | core_v1.CreateOptions |
| taskBook | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
<a name="taskbooks.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| taskBook | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
<a name="taskbooks.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.DeleteOptions](#core.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| taskBook | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
<a name="taskbooks.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| taskBook | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
<a name="taskbooks.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.GetOptions](#core.v1.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="taskbooks.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| taskBook | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
<a name="taskbooks.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.ListOptions](#core.v1.ListOptions) |  | core_v1.ListOptions |
<a name="taskbooks.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| taskBookList | [TaskBookList](#taskbooks.v1alpha.TaskBookList) |  | TaskBook list. |
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
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for TaskBook object. |
| spec | [TaskBookSpec](#taskbooks.v1alpha.TaskBookSpec) |  | Spec for TaskBook. |
<a name="taskbooks.v1alpha.TaskBookList"></a>

### TaskBookList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1.ObjectMetadata](#core.v1.ObjectMetadata) |  | Metadata for TaskBookList object. |
| items | [TaskBook](#taskbooks.v1alpha.TaskBook) | repeated | List of TaskBook objects. |
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
| options | [core.v1.UpdateOptions](#core.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| taskBook | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
<a name="taskbooks.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| taskBook | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook object. |
<a name="taskbooks.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.v1.WatchOptions](#core.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="taskbooks.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | TaskBook info for watch response. |
| taskBook | [TaskBook](#taskbooks.v1alpha.TaskBook) |  | TaskBook for watch response. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="taskbooks.v1alpha.TaskBooks"></a>

### TaskBooks



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#taskbooks.v1alpha.GetRequest) | [GetResponse](#taskbooks.v1alpha.GetResponse) | Get a TaskBook. |
| List | [ListRequest](#taskbooks.v1alpha.ListRequest) | [ListResponse](#taskbooks.v1alpha.ListResponse) | List TaskBooks. |
| Create | [CreateRequest](#taskbooks.v1alpha.CreateRequest) | [CreateResponse](#taskbooks.v1alpha.CreateResponse) | Create a TaskBook. |
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
