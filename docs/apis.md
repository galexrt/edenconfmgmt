# Protocol Documentation
<a name="top"></a>

## Table of Contents

* [github.com/galexrt/edenrun/pkg/apis/eden.run/runny/playbooks/v1alpha/playbooks.proto](#github.com/galexrt/edenrun/pkg/apis/eden.run/runny/playbooks/v1alpha/playbooks.proto)
    * Description: 
    * Services
        * [PlayBooks](#runny.playbooks.v1alpha.PlayBooks)
    * Messages
        * [Action](#runny.playbooks.v1alpha.Action)
        * [CreateRequest](#runny.playbooks.v1alpha.CreateRequest)
        * [CreateResponse](#runny.playbooks.v1alpha.CreateResponse)
        * [DeleteRequest](#runny.playbooks.v1alpha.DeleteRequest)
        * [DeleteResponse](#runny.playbooks.v1alpha.DeleteResponse)
        * [GetRequest](#runny.playbooks.v1alpha.GetRequest)
        * [GetResponse](#runny.playbooks.v1alpha.GetResponse)
        * [ListRequest](#runny.playbooks.v1alpha.ListRequest)
        * [ListResponse](#runny.playbooks.v1alpha.ListResponse)
        * [PlayBook](#runny.playbooks.v1alpha.PlayBook)
        * [PlayBookList](#runny.playbooks.v1alpha.PlayBookList)
        * [PlayBookSpec](#runny.playbooks.v1alpha.PlayBookSpec)
        * [RunOptions](#runny.playbooks.v1alpha.RunOptions)
        * [Serialize](#runny.playbooks.v1alpha.Serialize)
        * [Step](#runny.playbooks.v1alpha.Step)
        * [Target](#runny.playbooks.v1alpha.Target)
        * [UpdateRequest](#runny.playbooks.v1alpha.UpdateRequest)
        * [UpdateResponse](#runny.playbooks.v1alpha.UpdateResponse)
        * [WatchRequest](#runny.playbooks.v1alpha.WatchRequest)
        * [WatchResponse](#runny.playbooks.v1alpha.WatchResponse)
* [github.com/galexrt/edenrun/pkg/apis/eden.run/core/events/v1alpha/events.proto](#github.com/galexrt/edenrun/pkg/apis/eden.run/core/events/v1alpha/events.proto)
    * Description: 
    * Services
        * [Events](#core.events.v1alpha.Events)
    * Messages
        * [CreateRequest](#core.events.v1alpha.CreateRequest)
        * [CreateResponse](#core.events.v1alpha.CreateResponse)
        * [DeleteRequest](#core.events.v1alpha.DeleteRequest)
        * [DeleteResponse](#core.events.v1alpha.DeleteResponse)
        * [Event](#core.events.v1alpha.Event)
        * [EventList](#core.events.v1alpha.EventList)
        * [EventSpec](#core.events.v1alpha.EventSpec)
        * [GetRequest](#core.events.v1alpha.GetRequest)
        * [GetResponse](#core.events.v1alpha.GetResponse)
        * [ListRequest](#core.events.v1alpha.ListRequest)
        * [ListResponse](#core.events.v1alpha.ListResponse)
        * [UpdateRequest](#core.events.v1alpha.UpdateRequest)
        * [UpdateResponse](#core.events.v1alpha.UpdateResponse)
        * [WatchRequest](#core.events.v1alpha.WatchRequest)
        * [WatchResponse](#core.events.v1alpha.WatchResponse)
* [github.com/galexrt/edenrun/pkg/apis/eden.run/core/meta/v1/meta.proto](#github.com/galexrt/edenrun/pkg/apis/eden.run/core/meta/v1/meta.proto)
    * Description: 
    * Messages
        * [Condition](#core.meta.v1.Condition)
        * [Conditions](#core.meta.v1.Conditions)
        * [CreateOptions](#core.meta.v1.CreateOptions)
        * [DeleteOptions](#core.meta.v1.DeleteOptions)
        * [Error](#core.meta.v1.Error)
        * [GenericObject](#core.meta.v1.GenericObject)
        * [GetOptions](#core.meta.v1.GetOptions)
        * [Include](#core.meta.v1.Include)
        * [ListOptions](#core.meta.v1.ListOptions)
        * [ObjectMetadata](#core.meta.v1.ObjectMetadata)
        * [ObjectMetadata.AnnotationsEntry](#core.meta.v1.ObjectMetadata.AnnotationsEntry)
        * [ObjectMetadata.LabelsEntry](#core.meta.v1.ObjectMetadata.LabelsEntry)
        * [Retry](#core.meta.v1.Retry)
        * [UpdateOptions](#core.meta.v1.UpdateOptions)
        * [VersionRequest](#core.meta.v1.VersionRequest)
        * [VersionResponse](#core.meta.v1.VersionResponse)
        * [WatchOptions](#core.meta.v1.WatchOptions)
* [github.com/galexrt/edenrun/pkg/apis/eden.run/core/nodes/v1alpha/nodes.proto](#github.com/galexrt/edenrun/pkg/apis/eden.run/core/nodes/v1alpha/nodes.proto)
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
* [github.com/galexrt/edenrun/pkg/apis/eden.run/data/variables/v1alpha/variables.proto](#github.com/galexrt/edenrun/pkg/apis/eden.run/data/variables/v1alpha/variables.proto)
    * Description: 
    * Services
        * [Variables](#data.variables.v1alpha.Variables)
    * Messages
        * [CreateRequest](#data.variables.v1alpha.CreateRequest)
        * [CreateResponse](#data.variables.v1alpha.CreateResponse)
        * [DeleteRequest](#data.variables.v1alpha.DeleteRequest)
        * [DeleteResponse](#data.variables.v1alpha.DeleteResponse)
        * [GetRequest](#data.variables.v1alpha.GetRequest)
        * [GetResponse](#data.variables.v1alpha.GetResponse)
        * [ListRequest](#data.variables.v1alpha.ListRequest)
        * [ListResponse](#data.variables.v1alpha.ListResponse)
        * [UpdateRequest](#data.variables.v1alpha.UpdateRequest)
        * [UpdateResponse](#data.variables.v1alpha.UpdateResponse)
        * [Variable](#data.variables.v1alpha.Variable)
        * [VariableKVPair](#data.variables.v1alpha.VariableKVPair)
        * [VariableList](#data.variables.v1alpha.VariableList)
        * [VariableSpec](#data.variables.v1alpha.VariableSpec)
        * [WatchRequest](#data.variables.v1alpha.WatchRequest)
        * [WatchResponse](#data.variables.v1alpha.WatchResponse)
* [github.com/galexrt/edenrun/pkg/apis/eden.run/data/clustervariables/v1alpha/clustervariables.proto](#github.com/galexrt/edenrun/pkg/apis/eden.run/data/clustervariables/v1alpha/clustervariables.proto)
    * Description: 
    * Services
        * [ClusterVariables](#data.clustervariables.v1alpha.ClusterVariables)
    * Messages
        * [ClusterVariable](#data.clustervariables.v1alpha.ClusterVariable)
        * [ClusterVariableList](#data.clustervariables.v1alpha.ClusterVariableList)
        * [ClusterVariableSpec](#data.clustervariables.v1alpha.ClusterVariableSpec)
        * [CreateRequest](#data.clustervariables.v1alpha.CreateRequest)
        * [CreateResponse](#data.clustervariables.v1alpha.CreateResponse)
        * [DeleteRequest](#data.clustervariables.v1alpha.DeleteRequest)
        * [DeleteResponse](#data.clustervariables.v1alpha.DeleteResponse)
        * [GetRequest](#data.clustervariables.v1alpha.GetRequest)
        * [GetResponse](#data.clustervariables.v1alpha.GetResponse)
        * [ListRequest](#data.clustervariables.v1alpha.ListRequest)
        * [ListResponse](#data.clustervariables.v1alpha.ListResponse)
        * [Scope](#data.clustervariables.v1alpha.Scope)
        * [UpdateRequest](#data.clustervariables.v1alpha.UpdateRequest)
        * [UpdateResponse](#data.clustervariables.v1alpha.UpdateResponse)
        * [WatchRequest](#data.clustervariables.v1alpha.WatchRequest)
        * [WatchResponse](#data.clustervariables.v1alpha.WatchResponse)
* [github.com/galexrt/edenrun/pkg/apis/eden.run/data/secrets/v1alpha/secrets.proto](#github.com/galexrt/edenrun/pkg/apis/eden.run/data/secrets/v1alpha/secrets.proto)
    * Description: 
    * Services
        * [Secrets](#data.secrets.v1alpha.Secrets)
    * Messages
        * [CreateRequest](#data.secrets.v1alpha.CreateRequest)
        * [CreateResponse](#data.secrets.v1alpha.CreateResponse)
        * [DeleteRequest](#data.secrets.v1alpha.DeleteRequest)
        * [DeleteResponse](#data.secrets.v1alpha.DeleteResponse)
        * [GetRequest](#data.secrets.v1alpha.GetRequest)
        * [GetResponse](#data.secrets.v1alpha.GetResponse)
        * [ListRequest](#data.secrets.v1alpha.ListRequest)
        * [ListResponse](#data.secrets.v1alpha.ListResponse)
        * [Secret](#data.secrets.v1alpha.Secret)
        * [SecretList](#data.secrets.v1alpha.SecretList)
        * [SecretSpec](#data.secrets.v1alpha.SecretSpec)
        * [SecretSpec.SecretsEntry](#data.secrets.v1alpha.SecretSpec.SecretsEntry)
        * [UpdateRequest](#data.secrets.v1alpha.UpdateRequest)
        * [UpdateResponse](#data.secrets.v1alpha.UpdateResponse)
        * [WatchRequest](#data.secrets.v1alpha.WatchRequest)
        * [WatchResponse](#data.secrets.v1alpha.WatchResponse)
* [github.com/galexrt/edenrun/pkg/apis/eden.run/daemons/configs/v1alpha/configs.proto](#github.com/galexrt/edenrun/pkg/apis/eden.run/daemons/configs/v1alpha/configs.proto)
    * Description: 
    * Services
        * [Configs](#daemons.configs.v1alpha.Configs)
    * Messages
        * [Cache](#daemons.configs.v1alpha.Cache)
        * [Config](#daemons.configs.v1alpha.Config)
        * [ConfigList](#daemons.configs.v1alpha.ConfigList)
        * [ConfigSpec](#daemons.configs.v1alpha.ConfigSpec)
        * [CreateRequest](#daemons.configs.v1alpha.CreateRequest)
        * [CreateResponse](#daemons.configs.v1alpha.CreateResponse)
        * [Data](#daemons.configs.v1alpha.Data)
        * [DeleteRequest](#daemons.configs.v1alpha.DeleteRequest)
        * [DeleteResponse](#daemons.configs.v1alpha.DeleteResponse)
        * [ETCD](#daemons.configs.v1alpha.ETCD)
        * [GRPC](#daemons.configs.v1alpha.GRPC)
        * [GetRequest](#daemons.configs.v1alpha.GetRequest)
        * [GetResponse](#daemons.configs.v1alpha.GetResponse)
        * [HTTP](#daemons.configs.v1alpha.HTTP)
        * [ListRequest](#daemons.configs.v1alpha.ListRequest)
        * [ListResponse](#daemons.configs.v1alpha.ListResponse)
        * [Listener](#daemons.configs.v1alpha.Listener)
        * [Logging](#daemons.configs.v1alpha.Logging)
        * [Master](#daemons.configs.v1alpha.Master)
        * [Node](#daemons.configs.v1alpha.Node)
        * [Store](#daemons.configs.v1alpha.Store)
        * [TLS](#daemons.configs.v1alpha.TLS)
        * [UpdateRequest](#daemons.configs.v1alpha.UpdateRequest)
        * [UpdateResponse](#daemons.configs.v1alpha.UpdateResponse)
        * [WatchRequest](#daemons.configs.v1alpha.WatchRequest)
        * [WatchResponse](#daemons.configs.v1alpha.WatchResponse)
    * Enums
        * [CacheStoreType](#daemons.configs.v1alpha.CacheStoreType)
        * [DataStoreType](#daemons.configs.v1alpha.DataStoreType)
        * [LogLevel](#daemons.configs.v1alpha.LogLevel)
* [github.com/galexrt/edenrun/pkg/apis/eden.run/jobs/cronjobs/v1alpha/cronjobs.proto](#github.com/galexrt/edenrun/pkg/apis/eden.run/jobs/cronjobs/v1alpha/cronjobs.proto)
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
* [github.com/galexrt/edenrun/pkg/apis/eden.run/trigger/beacons/v1alpha/beacons.proto](#github.com/galexrt/edenrun/pkg/apis/eden.run/trigger/beacons/v1alpha/beacons.proto)
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
* [github.com/galexrt/edenrun/pkg/apis/eden.run/trigger/triggers/v1alpha/triggers.proto](#github.com/galexrt/edenrun/pkg/apis/eden.run/trigger/triggers/v1alpha/triggers.proto)
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
        * [WatchResponse](#triggers.v1alpha.WatchResponse)* [Scalar Value Types](#scalar-value-types)

<a name="github.com/galexrt/edenrun/pkg/apis/eden.run/runny/playbooks/v1alpha/playbooks.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/apis/eden.run/runny/playbooks/v1alpha/playbooks.proto



<a name="runny.playbooks.v1alpha.Action"></a>

### Action



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [string](#string) |  | Action name. |
| conditions | [core.meta.v1.Conditions](#core.meta.v1.Conditions) |  | Conditions. |
| runOptions | [RunOptions](#runny.playbooks.v1alpha.RunOptions) |  | RunOptions |
| parameters | [google.protobuf.Any](#google.protobuf.Any) | repeated | Map of string interface{} to be given to the action. |
<a name="runny.playbooks.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.CreateOptions](#core.meta.v1.CreateOptions) |  | core_v1.CreateOptions |
| playBook | [PlayBook](#runny.playbooks.v1alpha.PlayBook) |  | PlayBook object. |
<a name="runny.playbooks.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| playBook | [PlayBook](#runny.playbooks.v1alpha.PlayBook) |  | PlayBook object. |
<a name="runny.playbooks.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.DeleteOptions](#core.meta.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| playBook | [PlayBook](#runny.playbooks.v1alpha.PlayBook) |  | PlayBook object. |
<a name="runny.playbooks.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| playBook | [PlayBook](#runny.playbooks.v1alpha.PlayBook) |  | PlayBook object. |
<a name="runny.playbooks.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.GetOptions](#core.meta.v1.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="runny.playbooks.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| playBook | [PlayBook](#runny.playbooks.v1alpha.PlayBook) |  | PlayBook object. |
<a name="runny.playbooks.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.ListOptions](#core.meta.v1.ListOptions) |  | core_v1.ListOptions |
<a name="runny.playbooks.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| playBookList | [PlayBookList](#runny.playbooks.v1alpha.PlayBookList) |  | PlayBook list. |
<a name="runny.playbooks.v1alpha.PlayBook"></a>

### PlayBook



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for PlayBook object. |
| spec | [PlayBookSpec](#runny.playbooks.v1alpha.PlayBookSpec) |  | Spec for PlayBook. |
<a name="runny.playbooks.v1alpha.PlayBookList"></a>

### PlayBookList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for PlayBookList object. |
| items | [PlayBook](#runny.playbooks.v1alpha.PlayBook) | repeated | List of PlayBook objects. |
<a name="runny.playbooks.v1alpha.PlayBookSpec"></a>

### PlayBookSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| runOptions | [RunOptions](#runny.playbooks.v1alpha.RunOptions) |  | RunOptions. |
| steps | [Step](#runny.playbooks.v1alpha.Step) | repeated | Steps in this PlayBook. |
<a name="runny.playbooks.v1alpha.RunOptions"></a>

### RunOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sync | [bool](#bool) |  |  |
| serialize | [Serialize](#runny.playbooks.v1alpha.Serialize) |  |  |
| target | [Target](#runny.playbooks.v1alpha.Target) |  |  |
<a name="runny.playbooks.v1alpha.Serialize"></a>

### Serialize



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  | Count. |
| failAfterAllTargetsRan | [bool](#bool) |  | FailAfterAllTargetsRan. |
<a name="runny.playbooks.v1alpha.Step"></a>

### Step



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| spec | [Action](#runny.playbooks.v1alpha.Action) |  |  |
<a name="runny.playbooks.v1alpha.Target"></a>

### Target



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hosts | [string](#string) | repeated | List of hosts (either name or label selector). |
| limit | [int64](#int64) |  | Limit of how many hosts of the matches should be used. |
<a name="runny.playbooks.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.UpdateOptions](#core.meta.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| playBook | [PlayBook](#runny.playbooks.v1alpha.PlayBook) |  | PlayBook object. |
<a name="runny.playbooks.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| playBook | [PlayBook](#runny.playbooks.v1alpha.PlayBook) |  | PlayBook object. |
<a name="runny.playbooks.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.WatchOptions](#core.meta.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="runny.playbooks.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| playBook | [PlayBook](#runny.playbooks.v1alpha.PlayBook) |  | PlayBook for watch response. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="runny.playbooks.v1alpha.PlayBooks"></a>

### PlayBooks



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#runny.playbooks.v1alpha.GetRequest) | [GetResponse](#runny.playbooks.v1alpha.GetResponse) | Get a PlayBook. |
| List | [ListRequest](#runny.playbooks.v1alpha.ListRequest) | [ListResponse](#runny.playbooks.v1alpha.ListResponse) | List PlayBooks. |
| Create | [CreateRequest](#runny.playbooks.v1alpha.CreateRequest) | [CreateResponse](#runny.playbooks.v1alpha.CreateResponse) | Create a PlayBook. |
| Update | [UpdateRequest](#runny.playbooks.v1alpha.UpdateRequest) | [UpdateResponse](#runny.playbooks.v1alpha.UpdateResponse) | Update a PlayBook. |
| Delete | [DeleteRequest](#runny.playbooks.v1alpha.DeleteRequest) | [DeleteResponse](#runny.playbooks.v1alpha.DeleteResponse) | Delete a PlayBook. |
| Watch | [WatchRequest](#runny.playbooks.v1alpha.WatchRequest) | [WatchResponse](#runny.playbooks.v1alpha.WatchResponse) stream | Watch PlayBook. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/apis/eden.run/core/events/v1alpha/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/apis/eden.run/core/events/v1alpha/events.proto



<a name="core.events.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.CreateOptions](#core.meta.v1.CreateOptions) |  | core_v1.CreateOptions |
| event | [Event](#core.events.v1alpha.Event) |  | Event object. |
<a name="core.events.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#core.events.v1alpha.Event) |  | Event object. |
<a name="core.events.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.DeleteOptions](#core.meta.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| event | [Event](#core.events.v1alpha.Event) |  | Event object. |
<a name="core.events.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#core.events.v1alpha.Event) |  | Event object. |
<a name="core.events.v1alpha.Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for the Event object. |
| spec | [EventSpec](#core.events.v1alpha.EventSpec) |  | Spec for Event. |
<a name="core.events.v1alpha.EventList"></a>

### EventList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for EventList object. |
| items | [Event](#core.events.v1alpha.Event) | repeated | List of Event objects. |
<a name="core.events.v1alpha.EventSpec"></a>

### EventSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Type of the Event. |
| summary | [string](#string) |  | Summary of Event. |
<a name="core.events.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.GetOptions](#core.meta.v1.GetOptions) |  | core_v1.GetOptions |
<a name="core.events.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#core.events.v1alpha.Event) |  | Event object. |
<a name="core.events.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.ListOptions](#core.meta.v1.ListOptions) |  | core_v1.ListOptions |
<a name="core.events.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| eventList | [EventList](#core.events.v1alpha.EventList) |  | Event list. |
<a name="core.events.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.UpdateOptions](#core.meta.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| event | [Event](#core.events.v1alpha.Event) |  | Event object. |
<a name="core.events.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#core.events.v1alpha.Event) |  | Event object. |
<a name="core.events.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.WatchOptions](#core.meta.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="core.events.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#core.events.v1alpha.Event) |  | Event for watch stream. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="core.events.v1alpha.Events"></a>

### Events



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#core.events.v1alpha.GetRequest) | [GetResponse](#core.events.v1alpha.GetResponse) | Get a Event. |
| List | [ListRequest](#core.events.v1alpha.ListRequest) | [ListResponse](#core.events.v1alpha.ListResponse) | List Events. |
| Create | [CreateRequest](#core.events.v1alpha.CreateRequest) | [CreateResponse](#core.events.v1alpha.CreateResponse) | Create a Event. |
| Update | [UpdateRequest](#core.events.v1alpha.UpdateRequest) | [UpdateResponse](#core.events.v1alpha.UpdateResponse) | Update a Event. |
| Delete | [DeleteRequest](#core.events.v1alpha.DeleteRequest) | [DeleteResponse](#core.events.v1alpha.DeleteResponse) | Delete a Event. |
| Watch | [WatchRequest](#core.events.v1alpha.WatchRequest) | [WatchResponse](#core.events.v1alpha.WatchResponse) stream | Watch for Events. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/apis/eden.run/core/meta/v1/meta.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/apis/eden.run/core/meta/v1/meta.proto



<a name="core.meta.v1.Condition"></a>

### Condition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| condition | [string](#string) |  |  |
| retry | [Retry](#core.meta.v1.Retry) |  |  |
<a name="core.meta.v1.Conditions"></a>

### Conditions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| when | [Condition](#core.meta.v1.Condition) |  | `when` condition. |
| success | [Condition](#core.meta.v1.Condition) |  | `success` condition. |
<a name="core.meta.v1.CreateOptions"></a>

### CreateOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name used for listing. |
| namespace | [string](#string) |  | Namespace used for listing. |
<a name="core.meta.v1.DeleteOptions"></a>

### DeleteOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name used for listing. |
| namespace | [string](#string) |  | Namespace used for listing. |
| labelSelectors | [string](#string) | repeated | Label selector in string form. |
<a name="core.meta.v1.Error"></a>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | Code status code. |
| message | [string](#string) |  | Message error message (can be empty). |
<a name="core.meta.v1.GenericObject"></a>

### GenericObject



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for object. |
<a name="core.meta.v1.GetOptions"></a>

### GetOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name used for listing. |
| namespace | [string](#string) |  | Namespace used for listing. |
| labelSelectors | [string](#string) | repeated | Label selector in string form. |
<a name="core.meta.v1.Include"></a>

### Include



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of include (if no paths are given, will be used as file name/path). |
| paths | [string](#string) | repeated | List of paths. |
<a name="core.meta.v1.ListOptions"></a>

### ListOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name used for listing. |
| namespace | [string](#string) |  | Namespace used for listing. |
| labelSelectors | [string](#string) | repeated | Label selector in string form. |
<a name="core.meta.v1.ObjectMetadata"></a>

### ObjectMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apiVersion | [string](#string) |  | APIVersion of object. |
| kind | [string](#string) |  | Kind of object. |
| uuid | [string](#string) |  | UUID of object. |
| name | [string](#string) |  | Name of object. |
| namespace | [string](#string) |  | Namespace of object. |
| labels | [ObjectMetadata.LabelsEntry](#core.meta.v1.ObjectMetadata.LabelsEntry) | repeated | Labels of object. |
| annotations | [ObjectMetadata.AnnotationsEntry](#core.meta.v1.ObjectMetadata.AnnotationsEntry) | repeated | Annotations of object. |
<a name="core.meta.v1.ObjectMetadata.AnnotationsEntry"></a>

### ObjectMetadata.AnnotationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |
<a name="core.meta.v1.ObjectMetadata.LabelsEntry"></a>

### ObjectMetadata.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |
<a name="core.meta.v1.Retry"></a>

### Retry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| limit | [int64](#int64) |  |  |
<a name="core.meta.v1.UpdateOptions"></a>

### UpdateOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name used for listing. |
| namespace | [string](#string) |  | Namespace used for listing. |
<a name="core.meta.v1.VersionRequest"></a>

### VersionRequest


<a name="core.meta.v1.VersionResponse"></a>

### VersionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  | Semversion compatible version number. |
| major | [int64](#int64) |  | Semversion major number. |
| minor | [int64](#int64) |  | Semversion minor number. |
| patch | [int64](#int64) |  | Semversion patch number. |
| state | [string](#string) |  | API state (e.g., alpha, beta, stable). |
<a name="core.meta.v1.WatchOptions"></a>

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



<a name="github.com/galexrt/edenrun/pkg/apis/eden.run/core/nodes/v1alpha/nodes.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/apis/eden.run/core/nodes/v1alpha/nodes.proto



<a name="nodes.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.CreateOptions](#core.meta.v1.CreateOptions) |  | core_v1.CreateOptions |
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
| options | [core.meta.v1.DeleteOptions](#core.meta.v1.DeleteOptions) |  | core_v1.DeleteOptions |
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
| options | [core.meta.v1.GetOptions](#core.meta.v1.GetOptions) |  | GetOptions options for a GetRequest. |
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
| options | [core.meta.v1.ListOptions](#core.meta.v1.ListOptions) |  | core_v1.ListOptions |
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
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for Node object. |
| spec | [Spec](#nodes.v1alpha.Spec) |  | Spec for Node. |
| status | [Status](#nodes.v1alpha.Status) |  | Status for Node. |
<a name="nodes.v1alpha.NodeList"></a>

### NodeList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for NodeList object. |
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
| options | [core.meta.v1.UpdateOptions](#core.meta.v1.UpdateOptions) |  | core_v1.UpdateOptions |
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
| options | [core.meta.v1.WatchOptions](#core.meta.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="nodes.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
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



<a name="github.com/galexrt/edenrun/pkg/apis/eden.run/data/variables/v1alpha/variables.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/apis/eden.run/data/variables/v1alpha/variables.proto



<a name="data.variables.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.CreateOptions](#core.meta.v1.CreateOptions) |  | core_v1.CreateOptions |
| variable | [Variable](#data.variables.v1alpha.Variable) |  | Variable object. |
<a name="data.variables.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#data.variables.v1alpha.Variable) |  | Variable object. |
<a name="data.variables.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.DeleteOptions](#core.meta.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| variable | [Variable](#data.variables.v1alpha.Variable) |  | Variable object. |
<a name="data.variables.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#data.variables.v1alpha.Variable) |  | Variable object. |
<a name="data.variables.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.GetOptions](#core.meta.v1.GetOptions) |  | core_v1.GetOptions |
<a name="data.variables.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#data.variables.v1alpha.Variable) |  | Variable object. |
<a name="data.variables.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.ListOptions](#core.meta.v1.ListOptions) |  | core_v1.ListOptions |
<a name="data.variables.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variableList | [VariableList](#data.variables.v1alpha.VariableList) |  | Variable list. |
<a name="data.variables.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.UpdateOptions](#core.meta.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| variable | [Variable](#data.variables.v1alpha.Variable) |  | Variable object. |
<a name="data.variables.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#data.variables.v1alpha.Variable) |  | Variable object. |
<a name="data.variables.v1alpha.Variable"></a>

### Variable



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for Variable object. |
| spec | [VariableSpec](#data.variables.v1alpha.VariableSpec) |  | Spec for Variable. |
<a name="data.variables.v1alpha.VariableKVPair"></a>

### VariableKVPair

VariableKVPair

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |
<a name="data.variables.v1alpha.VariableList"></a>

### VariableList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for VariableList object. |
| items | [Variable](#data.variables.v1alpha.Variable) | repeated | List of Variable objects. |
<a name="data.variables.v1alpha.VariableSpec"></a>

### VariableSpec

Variables variables file structure. This covers the "Variables File" and "'From Host' Static Variable File".

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variables | [VariableKVPair](#data.variables.v1alpha.VariableKVPair) | repeated | Variables. |
<a name="data.variables.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.WatchOptions](#core.meta.v1.WatchOptions) |  | WatchOptions options for WatchRequest. |
<a name="data.variables.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| variable | [Variable](#data.variables.v1alpha.Variable) |  | Variable for watch response. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="data.variables.v1alpha.Variables"></a>

### Variables



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#data.variables.v1alpha.GetRequest) | [GetResponse](#data.variables.v1alpha.GetResponse) | Get a Variable. |
| List | [ListRequest](#data.variables.v1alpha.ListRequest) | [ListResponse](#data.variables.v1alpha.ListResponse) | List Variable. |
| Create | [CreateRequest](#data.variables.v1alpha.CreateRequest) | [CreateResponse](#data.variables.v1alpha.CreateResponse) | Create a Variable. |
| Update | [UpdateRequest](#data.variables.v1alpha.UpdateRequest) | [UpdateResponse](#data.variables.v1alpha.UpdateResponse) | Update a Variable. |
| Delete | [DeleteRequest](#data.variables.v1alpha.DeleteRequest) | [DeleteResponse](#data.variables.v1alpha.DeleteResponse) | Delete a Variable. |
| Watch | [WatchRequest](#data.variables.v1alpha.WatchRequest) | [WatchResponse](#data.variables.v1alpha.WatchResponse) stream | Watch Variable. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/apis/eden.run/data/clustervariables/v1alpha/clustervariables.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/apis/eden.run/data/clustervariables/v1alpha/clustervariables.proto



<a name="data.clustervariables.v1alpha.ClusterVariable"></a>

### ClusterVariable



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for ClusterVariable object. |
| spec | [ClusterVariableSpec](#data.clustervariables.v1alpha.ClusterVariableSpec) |  | Spec for ClusterVariable. |
<a name="data.clustervariables.v1alpha.ClusterVariableList"></a>

### ClusterVariableList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for ClusterVariableList object. |
| items | [ClusterVariable](#data.clustervariables.v1alpha.ClusterVariable) | repeated | List of ClusterVariable objects. |
<a name="data.clustervariables.v1alpha.ClusterVariableSpec"></a>

### ClusterVariableSpec

ClusterVariables clustervariables file structure. This covers the "ClusterVariables File" and "'From Host' Static ClusterVariable File".

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| scope | [Scope](#data.clustervariables.v1alpha.Scope) |  | Scope of ClusterVariable. |
| clustervariables | [data.variables.v1alpha.VariableKVPair](#data.variables.v1alpha.VariableKVPair) | repeated | Variables. |
<a name="data.clustervariables.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.CreateOptions](#core.meta.v1.CreateOptions) |  | core_v1.CreateOptions |
| clusterVariable | [ClusterVariable](#data.clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="data.clustervariables.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#data.clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="data.clustervariables.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.DeleteOptions](#core.meta.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| clusterVariable | [ClusterVariable](#data.clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="data.clustervariables.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#data.clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="data.clustervariables.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.GetOptions](#core.meta.v1.GetOptions) |  | core_v1.GetOptions |
<a name="data.clustervariables.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#data.clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="data.clustervariables.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.ListOptions](#core.meta.v1.ListOptions) |  | core_v1.ListOptions |
<a name="data.clustervariables.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariableList | [ClusterVariableList](#data.clustervariables.v1alpha.ClusterVariableList) |  | ClusterVariable list. |
<a name="data.clustervariables.v1alpha.Scope"></a>

### Scope


<a name="data.clustervariables.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.UpdateOptions](#core.meta.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| clusterVariable | [ClusterVariable](#data.clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="data.clustervariables.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#data.clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable object. |
<a name="data.clustervariables.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.WatchOptions](#core.meta.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="data.clustervariables.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clusterVariable | [ClusterVariable](#data.clustervariables.v1alpha.ClusterVariable) |  | ClusterVariable for watch response. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="data.clustervariables.v1alpha.ClusterVariables"></a>

### ClusterVariables



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#data.clustervariables.v1alpha.GetRequest) | [GetResponse](#data.clustervariables.v1alpha.GetResponse) | Get a ClusterVariable. |
| List | [ListRequest](#data.clustervariables.v1alpha.ListRequest) | [ListResponse](#data.clustervariables.v1alpha.ListResponse) | List ClusterVariable. |
| Create | [CreateRequest](#data.clustervariables.v1alpha.CreateRequest) | [CreateResponse](#data.clustervariables.v1alpha.CreateResponse) | Create a ClusterVariable. |
| Update | [UpdateRequest](#data.clustervariables.v1alpha.UpdateRequest) | [UpdateResponse](#data.clustervariables.v1alpha.UpdateResponse) | Update a ClusterVariable. |
| Delete | [DeleteRequest](#data.clustervariables.v1alpha.DeleteRequest) | [DeleteResponse](#data.clustervariables.v1alpha.DeleteResponse) | Delete a ClusterVariable. |
| Watch | [WatchRequest](#data.clustervariables.v1alpha.WatchRequest) | [WatchResponse](#data.clustervariables.v1alpha.WatchResponse) stream | Watch ClusterVariable. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/apis/eden.run/data/secrets/v1alpha/secrets.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/apis/eden.run/data/secrets/v1alpha/secrets.proto



<a name="data.secrets.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.CreateOptions](#core.meta.v1.CreateOptions) |  | core_v1.CreateOptions |
| secret | [Secret](#data.secrets.v1alpha.Secret) |  | Secret object. |
<a name="data.secrets.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#data.secrets.v1alpha.Secret) |  | Secret object. |
<a name="data.secrets.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.DeleteOptions](#core.meta.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| secret | [Secret](#data.secrets.v1alpha.Secret) |  | Secret object. |
<a name="data.secrets.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#data.secrets.v1alpha.Secret) |  | Secret object. |
<a name="data.secrets.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.GetOptions](#core.meta.v1.GetOptions) |  | GetOptions options for a GetRequest. |
<a name="data.secrets.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#data.secrets.v1alpha.Secret) |  | Secret object. |
<a name="data.secrets.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.ListOptions](#core.meta.v1.ListOptions) |  | core_v1.ListOptions |
<a name="data.secrets.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| SecretList | [SecretList](#data.secrets.v1alpha.SecretList) |  | Secret list. |
<a name="data.secrets.v1alpha.Secret"></a>

### Secret



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for Secret object. |
| spec | [SecretSpec](#data.secrets.v1alpha.SecretSpec) |  | Spec for Secret. |
<a name="data.secrets.v1alpha.SecretList"></a>

### SecretList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for SecretList object. |
| items | [Secret](#data.secrets.v1alpha.Secret) | repeated | List of Secret objects. |
<a name="data.secrets.v1alpha.SecretSpec"></a>

### SecretSpec

Secrets secrets file structure. This covers the "Secrets File" and "'From Host' Static Secret File".

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secrets | [SecretSpec.SecretsEntry](#data.secrets.v1alpha.SecretSpec.SecretsEntry) | repeated | Secrets. |
| type | [string](#string) |  | Type of secret. |
<a name="data.secrets.v1alpha.SecretSpec.SecretsEntry"></a>

### SecretSpec.SecretsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |
<a name="data.secrets.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.UpdateOptions](#core.meta.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| secret | [Secret](#data.secrets.v1alpha.Secret) |  | Secret object. |
<a name="data.secrets.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#data.secrets.v1alpha.Secret) |  | Secret object. |
<a name="data.secrets.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.WatchOptions](#core.meta.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="data.secrets.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| secret | [Secret](#data.secrets.v1alpha.Secret) |  | Secret for watch response. |
<!-- end messages -->

<!-- end enums -->

 <!-- end HasExtensions -->


<a name="data.secrets.v1alpha.Secrets"></a>

### Secrets



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#data.secrets.v1alpha.GetRequest) | [GetResponse](#data.secrets.v1alpha.GetResponse) | Get a Secret. |
| List | [ListRequest](#data.secrets.v1alpha.ListRequest) | [ListResponse](#data.secrets.v1alpha.ListResponse) | List Secret. |
| Create | [CreateRequest](#data.secrets.v1alpha.CreateRequest) | [CreateResponse](#data.secrets.v1alpha.CreateResponse) | Create a Secret. |
| Update | [UpdateRequest](#data.secrets.v1alpha.UpdateRequest) | [UpdateResponse](#data.secrets.v1alpha.UpdateResponse) | Update a Secret. |
| Delete | [DeleteRequest](#data.secrets.v1alpha.DeleteRequest) | [DeleteResponse](#data.secrets.v1alpha.DeleteResponse) | Delete a Secret. |
| Watch | [WatchRequest](#data.secrets.v1alpha.WatchRequest) | [WatchResponse](#data.secrets.v1alpha.WatchResponse) stream | Watch Secret. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/apis/eden.run/daemons/configs/v1alpha/configs.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/apis/eden.run/daemons/configs/v1alpha/configs.proto



<a name="daemons.configs.v1alpha.Cache"></a>

### Cache



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [CacheStoreType](#daemons.configs.v1alpha.CacheStoreType) |  | Type. |
<a name="daemons.configs.v1alpha.Config"></a>

### Config

Config object base.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for Config object. |
| spec | [ConfigSpec](#daemons.configs.v1alpha.ConfigSpec) |  | Spec for Config. |
<a name="daemons.configs.v1alpha.ConfigList"></a>

### ConfigList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for ConfigList object. |
| items | [Config](#daemons.configs.v1alpha.Config) | repeated | List of Config objects. |
<a name="daemons.configs.v1alpha.ConfigSpec"></a>

### ConfigSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logging | [Logging](#daemons.configs.v1alpha.Logging) |  | Logging configuration |
| listener | [Listener](#daemons.configs.v1alpha.Listener) |  |  |
| master | [Master](#daemons.configs.v1alpha.Master) |  |  |
| node | [Node](#daemons.configs.v1alpha.Node) |  |  |
<a name="daemons.configs.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.CreateOptions](#core.meta.v1.CreateOptions) |  | core_v1.CreateOptions |
| config | [Config](#daemons.configs.v1alpha.Config) |  | Config object. |
<a name="daemons.configs.v1alpha.CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#daemons.configs.v1alpha.Config) |  | Config object. |
<a name="daemons.configs.v1alpha.Data"></a>

### Data

DataStore config for the data store

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [DataStoreType](#daemons.configs.v1alpha.DataStoreType) |  | Type. |
| etcd | [ETCD](#daemons.configs.v1alpha.ETCD) |  | ETCD config options. |
<a name="daemons.configs.v1alpha.DeleteRequest"></a>

### DeleteRequest

Delete

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.DeleteOptions](#core.meta.v1.DeleteOptions) |  | core_v1.DeleteOptions |
| config | [Config](#daemons.configs.v1alpha.Config) |  | Config object. |
<a name="daemons.configs.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#daemons.configs.v1alpha.Config) |  | Config object. |
<a name="daemons.configs.v1alpha.ETCD"></a>

### ETCD



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| endpoints | [string](#string) |  | List of endpoints if running your own ETCD server (should only be used for testing). |
<a name="daemons.configs.v1alpha.GRPC"></a>

### GRPC



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| tls | [TLS](#daemons.configs.v1alpha.TLS) |  |  |
<a name="daemons.configs.v1alpha.GetRequest"></a>

### GetRequest

Get

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.GetOptions](#core.meta.v1.GetOptions) |  | core_v1.GetOptions |
<a name="daemons.configs.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#daemons.configs.v1alpha.Config) |  | Config object. |
<a name="daemons.configs.v1alpha.HTTP"></a>

### HTTP



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
<a name="daemons.configs.v1alpha.ListRequest"></a>

### ListRequest

List

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.ListOptions](#core.meta.v1.ListOptions) |  | core_v1.ListOptions |
<a name="daemons.configs.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configList | [ConfigList](#daemons.configs.v1alpha.ConfigList) |  | Config list. |
<a name="daemons.configs.v1alpha.Listener"></a>

### Listener



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| http | [HTTP](#daemons.configs.v1alpha.HTTP) |  |  |
| grpc | [GRPC](#daemons.configs.v1alpha.GRPC) |  |  |
<a name="daemons.configs.v1alpha.Logging"></a>

### Logging

Logging configuration

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| logLevel | [LogLevel](#daemons.configs.v1alpha.LogLevel) |  |  |
<a name="daemons.configs.v1alpha.Master"></a>

### Master



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| store | [Store](#daemons.configs.v1alpha.Store) |  |  |
<a name="daemons.configs.v1alpha.Node"></a>

### Node


<a name="daemons.configs.v1alpha.Store"></a>

### Store

Store

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [Data](#daemons.configs.v1alpha.Data) |  |  |
| cache | [Cache](#daemons.configs.v1alpha.Cache) |  |  |
<a name="daemons.configs.v1alpha.TLS"></a>

### TLS



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Certificate | [string](#string) |  |  |
| Key | [string](#string) |  |  |
<a name="daemons.configs.v1alpha.UpdateRequest"></a>

### UpdateRequest

Update

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.UpdateOptions](#core.meta.v1.UpdateOptions) |  | core_v1.UpdateOptions |
| config | [Config](#daemons.configs.v1alpha.Config) |  | Config object. |
<a name="daemons.configs.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#daemons.configs.v1alpha.Config) |  | Config object. |
<a name="daemons.configs.v1alpha.WatchRequest"></a>

### WatchRequest

Watch

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.WatchOptions](#core.meta.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="daemons.configs.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#daemons.configs.v1alpha.Config) |  | Config object. |
<!-- end messages -->


<a name="daemons.configs.v1alpha.CacheStoreType"></a>

### CacheStoreType



| Name | Number | Description |
| ---- | ------ | ----------- |
| MEMORY | 0 |  |



<a name="daemons.configs.v1alpha.DataStoreType"></a>

### DataStoreType



| Name | Number | Description |
| ---- | ------ | ----------- |
| ETCDV3 | 0 |  |



<a name="daemons.configs.v1alpha.LogLevel"></a>

### LogLevel



| Name | Number | Description |
| ---- | ------ | ----------- |
| DEBUG | 0 |  |
| INFO | 1 |  |
| WARN | 2 |  |
| ERROR | 3 |  |


<!-- end enums -->

 <!-- end HasExtensions -->


<a name="daemons.configs.v1alpha.Configs"></a>

### Configs



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#daemons.configs.v1alpha.GetRequest) | [GetResponse](#daemons.configs.v1alpha.GetResponse) | Get a Config. |
| List | [ListRequest](#daemons.configs.v1alpha.ListRequest) | [ListResponse](#daemons.configs.v1alpha.ListResponse) | List Configs. |
| Create | [CreateRequest](#daemons.configs.v1alpha.CreateRequest) | [CreateResponse](#daemons.configs.v1alpha.CreateResponse) | Create a Config. |
| Update | [UpdateRequest](#daemons.configs.v1alpha.UpdateRequest) | [UpdateResponse](#daemons.configs.v1alpha.UpdateResponse) | Update a Config. |
| Delete | [DeleteRequest](#daemons.configs.v1alpha.DeleteRequest) | [DeleteResponse](#daemons.configs.v1alpha.DeleteResponse) | Delete a Config. |
| Watch | [WatchRequest](#daemons.configs.v1alpha.WatchRequest) | [WatchResponse](#daemons.configs.v1alpha.WatchResponse) stream | Watch Config. |

 <!-- end services -->



<a name="github.com/galexrt/edenrun/pkg/apis/eden.run/jobs/cronjobs/v1alpha/cronjobs.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/apis/eden.run/jobs/cronjobs/v1alpha/cronjobs.proto



<a name="cronjobs.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.CreateOptions](#core.meta.v1.CreateOptions) |  | core_v1.CreateOptions |
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
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for CronJob object. |
| spec | [CronJobSpec](#cronjobs.v1alpha.CronJobSpec) |  | Spec for CronJob. |
<a name="cronjobs.v1alpha.CronJobList"></a>

### CronJobList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for CronJobList object. |
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
| options | [core.meta.v1.DeleteOptions](#core.meta.v1.DeleteOptions) |  | core_v1.DeleteOptions |
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
| options | [core.meta.v1.GetOptions](#core.meta.v1.GetOptions) |  | core_v1.GetOptions |
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
| options | [core.meta.v1.ListOptions](#core.meta.v1.ListOptions) |  | core_v1.ListOptions |
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
| options | [core.meta.v1.UpdateOptions](#core.meta.v1.UpdateOptions) |  | core_v1.CreateOptions |
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
| options | [core.meta.v1.WatchOptions](#core.meta.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="cronjobs.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
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



<a name="github.com/galexrt/edenrun/pkg/apis/eden.run/trigger/beacons/v1alpha/beacons.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/apis/eden.run/trigger/beacons/v1alpha/beacons.proto



<a name="beacons.v1alpha.Beacon"></a>

### Beacon



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for Beacon object. |
| spec | [BeaconSpec](#beacons.v1alpha.BeaconSpec) |  | Spec for Beacon. |
<a name="beacons.v1alpha.BeaconList"></a>

### BeaconList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for BeaconList object. |
| items | [Beacon](#beacons.v1alpha.Beacon) | repeated | List of Beacon objects. |
<a name="beacons.v1alpha.BeaconSpec"></a>

### BeaconSpec

TODO
<a name="beacons.v1alpha.CreateRequest"></a>

### CreateRequest

Create Beacon request.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.CreateOptions](#core.meta.v1.CreateOptions) |  | core_v1.CreateOptions |
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
| options | [core.meta.v1.DeleteOptions](#core.meta.v1.DeleteOptions) |  | core_v1.DeleteOptions |
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
| options | [core.meta.v1.GetOptions](#core.meta.v1.GetOptions) |  | core_v1.GetOptions |
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
| options | [core.meta.v1.ListOptions](#core.meta.v1.ListOptions) |  | core_v1.ListOptions |
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
| options | [core.meta.v1.UpdateOptions](#core.meta.v1.UpdateOptions) |  | core_v1.UpdateOptions |
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
| options | [core.meta.v1.WatchOptions](#core.meta.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="beacons.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
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



<a name="github.com/galexrt/edenrun/pkg/apis/eden.run/trigger/triggers/v1alpha/triggers.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenrun/pkg/apis/eden.run/trigger/triggers/v1alpha/triggers.proto



<a name="triggers.v1alpha.Action"></a>

### Action



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [string](#string) |  | Action name. |
| conditions | [core.meta.v1.Conditions](#core.meta.v1.Conditions) |  | Conditions. |
| runOptions | [RunOptions](#triggers.v1alpha.RunOptions) |  | RunOptions |
| parameters | [google.protobuf.Any](#google.protobuf.Any) | repeated | Map of string interface{} to be given to the action. |
<a name="triggers.v1alpha.CreateRequest"></a>

### CreateRequest

Create

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [core.meta.v1.CreateOptions](#core.meta.v1.CreateOptions) |  | core_v1.CreateOptions |
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
| options | [core.meta.v1.DeleteOptions](#core.meta.v1.DeleteOptions) |  | core_v1.DeleteOptions |
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
| options | [core.meta.v1.GetOptions](#core.meta.v1.GetOptions) |  | core_v1.GetOptions |
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
| options | [core.meta.v1.ListOptions](#core.meta.v1.ListOptions) |  | core_v1.ListOptions |
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
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for Trigger object. |
| spec | [TriggerSpec](#triggers.v1alpha.TriggerSpec) |  | Spec for Trigger. |
<a name="triggers.v1alpha.TriggerList"></a>

### TriggerList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.meta.v1.ObjectMetadata](#core.meta.v1.ObjectMetadata) |  | Metadata for TriggerList object. |
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
| options | [core.meta.v1.UpdateOptions](#core.meta.v1.UpdateOptions) |  | core_v1.UpdateOptions |
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
| options | [core.meta.v1.WatchOptions](#core.meta.v1.WatchOptions) |  | core_v1.WatchOptions |
<a name="triggers.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
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
