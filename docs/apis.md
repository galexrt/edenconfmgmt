# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [github.com/galexrt/edenconfmgmt/pkg/apis/eventreactors/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/eventreactors/v1alpha/api.proto)
    - [AddRequest](#eventreactors.v1alpha.AddRequest)
    - [AddResponse](#eventreactors.v1alpha.AddResponse)
    - [DeleteRequest](#eventreactors.v1alpha.DeleteRequest)
    - [DeleteResponse](#eventreactors.v1alpha.DeleteResponse)
    - [EventReactor](#eventreactors.v1alpha.EventReactor)
    - [EventReactorSpec](#eventreactors.v1alpha.EventReactorSpec)
    - [GetRequest](#eventreactors.v1alpha.GetRequest)
    - [GetResponse](#eventreactors.v1alpha.GetResponse)
    - [ListRequest](#eventreactors.v1alpha.ListRequest)
    - [ListResponse](#eventreactors.v1alpha.ListResponse)
    - [Reaction](#eventreactors.v1alpha.Reaction)
    - [TriggerLimiter](#eventreactors.v1alpha.TriggerLimiter)
    - [UpdateRequest](#eventreactors.v1alpha.UpdateRequest)
    - [UpdateResponse](#eventreactors.v1alpha.UpdateResponse)
    - [WaitForMore](#eventreactors.v1alpha.WaitForMore)
    - [WatchRequest](#eventreactors.v1alpha.WatchRequest)
    - [WatchResponse](#eventreactors.v1alpha.WatchResponse)
  
  
  
    - [EventReactors](#eventreactors.v1alpha.EventReactors)
  

- [github.com/galexrt/edenconfmgmt/pkg/apis/jobs/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/jobs/v1alpha/api.proto)
    - [AddRequest](#jobs.v1alpha.AddRequest)
    - [AddResponse](#jobs.v1alpha.AddResponse)
    - [DeleteRequest](#jobs.v1alpha.DeleteRequest)
    - [DeleteResponse](#jobs.v1alpha.DeleteResponse)
    - [GetRequest](#jobs.v1alpha.GetRequest)
    - [GetResponse](#jobs.v1alpha.GetResponse)
    - [Job](#jobs.v1alpha.Job)
    - [ListRequest](#jobs.v1alpha.ListRequest)
    - [ListResponse](#jobs.v1alpha.ListResponse)
    - [UpdateRequest](#jobs.v1alpha.UpdateRequest)
    - [UpdateResponse](#jobs.v1alpha.UpdateResponse)
    - [WatchRequest](#jobs.v1alpha.WatchRequest)
    - [WatchResponse](#jobs.v1alpha.WatchResponse)
  
  
  
    - [Jobs](#jobs.v1alpha.Jobs)
  

- [github.com/galexrt/edenconfmgmt/pkg/apis/configs/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/configs/v1alpha/api.proto)
    - [AddRequest](#configs.v1alpha.AddRequest)
    - [AddResponse](#configs.v1alpha.AddResponse)
    - [Config](#configs.v1alpha.Config)
    - [ConfigSpec](#configs.v1alpha.ConfigSpec)
    - [DataStore](#configs.v1alpha.DataStore)
    - [DeleteRequest](#configs.v1alpha.DeleteRequest)
    - [DeleteResponse](#configs.v1alpha.DeleteResponse)
    - [ETCD](#configs.v1alpha.ETCD)
    - [GetRequest](#configs.v1alpha.GetRequest)
    - [GetResponse](#configs.v1alpha.GetResponse)
    - [ListRequest](#configs.v1alpha.ListRequest)
    - [ListResponse](#configs.v1alpha.ListResponse)
    - [UpdateRequest](#configs.v1alpha.UpdateRequest)
    - [UpdateResponse](#configs.v1alpha.UpdateResponse)
    - [WatchRequest](#configs.v1alpha.WatchRequest)
    - [WatchResponse](#configs.v1alpha.WatchResponse)
  
  
  
    - [Configs](#configs.v1alpha.Configs)
  

- [github.com/galexrt/edenconfmgmt/pkg/apis/variables/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/variables/v1alpha/api.proto)
    - [AddRequest](#variables.v1alpha.AddRequest)
    - [AddResponse](#variables.v1alpha.AddResponse)
    - [DeleteRequest](#variables.v1alpha.DeleteRequest)
    - [DeleteResponse](#variables.v1alpha.DeleteResponse)
    - [GetRequest](#variables.v1alpha.GetRequest)
    - [GetResponse](#variables.v1alpha.GetResponse)
    - [ListRequest](#variables.v1alpha.ListRequest)
    - [ListResponse](#variables.v1alpha.ListResponse)
    - [UpdateRequest](#variables.v1alpha.UpdateRequest)
    - [UpdateResponse](#variables.v1alpha.UpdateResponse)
    - [Variable](#variables.v1alpha.Variable)
    - [VariableSpec](#variables.v1alpha.VariableSpec)
    - [VariableSpec.VariablesEntry](#variables.v1alpha.VariableSpec.VariablesEntry)
    - [WatchRequest](#variables.v1alpha.WatchRequest)
    - [WatchResponse](#variables.v1alpha.WatchResponse)
  
  
  
    - [Variables](#variables.v1alpha.Variables)
  

- [github.com/galexrt/edenconfmgmt/pkg/apis/tasks/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/tasks/v1alpha/api.proto)
    - [Action](#tasks.v1alpha.Action)
    - [AddRequest](#tasks.v1alpha.AddRequest)
    - [AddResponse](#tasks.v1alpha.AddResponse)
    - [DeleteRequest](#tasks.v1alpha.DeleteRequest)
    - [DeleteResponse](#tasks.v1alpha.DeleteResponse)
    - [GetRequest](#tasks.v1alpha.GetRequest)
    - [GetResponse](#tasks.v1alpha.GetResponse)
    - [ListRequest](#tasks.v1alpha.ListRequest)
    - [ListResponse](#tasks.v1alpha.ListResponse)
    - [RunOptions](#tasks.v1alpha.RunOptions)
    - [Serialize](#tasks.v1alpha.Serialize)
    - [Step](#tasks.v1alpha.Step)
    - [Target](#tasks.v1alpha.Target)
    - [Task](#tasks.v1alpha.Task)
    - [TaskSpec](#tasks.v1alpha.TaskSpec)
    - [UpdateRequest](#tasks.v1alpha.UpdateRequest)
    - [UpdateResponse](#tasks.v1alpha.UpdateResponse)
    - [WatchRequest](#tasks.v1alpha.WatchRequest)
    - [WatchResponse](#tasks.v1alpha.WatchResponse)
  
  
  
    - [Tasks](#tasks.v1alpha.Tasks)
  

- [github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha/api.proto)
    - [Condition](#core.v1alpha.Condition)
    - [Conditions](#core.v1alpha.Conditions)
    - [Error](#core.v1alpha.Error)
    - [GetOptions](#core.v1alpha.GetOptions)
    - [Include](#core.v1alpha.Include)
    - [ListOptions](#core.v1alpha.ListOptions)
    - [ObjectMetadata](#core.v1alpha.ObjectMetadata)
    - [ObjectMetadata.AnnotationsEntry](#core.v1alpha.ObjectMetadata.AnnotationsEntry)
    - [ObjectMetadata.LabelsEntry](#core.v1alpha.ObjectMetadata.LabelsEntry)
    - [Retry](#core.v1alpha.Retry)
    - [VersionRequest](#core.v1alpha.VersionRequest)
    - [VersionResponse](#core.v1alpha.VersionResponse)
    - [WatchOptions](#core.v1alpha.WatchOptions)
  
  
  
  

- [github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha/api.proto)
    - [AddRequest](#events.v1alpha.AddRequest)
    - [AddResponse](#events.v1alpha.AddResponse)
    - [DeleteRequest](#events.v1alpha.DeleteRequest)
    - [DeleteResponse](#events.v1alpha.DeleteResponse)
    - [Event](#events.v1alpha.Event)
    - [EventSpec](#events.v1alpha.EventSpec)
    - [GetRequest](#events.v1alpha.GetRequest)
    - [GetResponse](#events.v1alpha.GetResponse)
    - [ListRequest](#events.v1alpha.ListRequest)
    - [ListResponse](#events.v1alpha.ListResponse)
    - [UpdateRequest](#events.v1alpha.UpdateRequest)
    - [UpdateResponse](#events.v1alpha.UpdateResponse)
    - [WatchRequest](#events.v1alpha.WatchRequest)
    - [WatchResponse](#events.v1alpha.WatchResponse)
  
  
  
    - [Events](#events.v1alpha.Events)
  

- [github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/nodes/v1alpha/api.proto)
    - [AddRequest](#nodes.v1alpha.AddRequest)
    - [AddResponse](#nodes.v1alpha.AddResponse)
    - [DeleteRequest](#nodes.v1alpha.DeleteRequest)
    - [DeleteResponse](#nodes.v1alpha.DeleteResponse)
    - [GetRequest](#nodes.v1alpha.GetRequest)
    - [GetResponse](#nodes.v1alpha.GetResponse)
    - [ListRequest](#nodes.v1alpha.ListRequest)
    - [ListResponse](#nodes.v1alpha.ListResponse)
    - [Network](#nodes.v1alpha.Network)
    - [Node](#nodes.v1alpha.Node)
    - [Port](#nodes.v1alpha.Port)
    - [Spec](#nodes.v1alpha.Spec)
    - [Status](#nodes.v1alpha.Status)
    - [UpdateRequest](#nodes.v1alpha.UpdateRequest)
    - [UpdateResponse](#nodes.v1alpha.UpdateResponse)
    - [WatchRequest](#nodes.v1alpha.WatchRequest)
    - [WatchResponse](#nodes.v1alpha.WatchResponse)
  
  
  
    - [Nodes](#nodes.v1alpha.Nodes)
  

- [github.com/galexrt/edenconfmgmt/pkg/apis/templatemacros/v1alpha/api.proto](#github.com/galexrt/edenconfmgmt/pkg/apis/templatemacros/v1alpha/api.proto)
    - [AddRequest](#templatemacros.v1alpha.AddRequest)
    - [AddResponse](#templatemacros.v1alpha.AddResponse)
    - [DeleteRequest](#templatemacros.v1alpha.DeleteRequest)
    - [DeleteResponse](#templatemacros.v1alpha.DeleteResponse)
    - [GetRequest](#templatemacros.v1alpha.GetRequest)
    - [GetResponse](#templatemacros.v1alpha.GetResponse)
    - [ListRequest](#templatemacros.v1alpha.ListRequest)
    - [ListResponse](#templatemacros.v1alpha.ListResponse)
    - [TemplateMacro](#templatemacros.v1alpha.TemplateMacro)
    - [TemplateMacroItem](#templatemacros.v1alpha.TemplateMacroItem)
    - [TemplateMacroSpec](#templatemacros.v1alpha.TemplateMacroSpec)
    - [UpdateRequest](#templatemacros.v1alpha.UpdateRequest)
    - [UpdateResponse](#templatemacros.v1alpha.UpdateResponse)
    - [WatchRequest](#templatemacros.v1alpha.WatchRequest)
    - [WatchResponse](#templatemacros.v1alpha.WatchResponse)
  
  
  
    - [TemplateMacros](#templatemacros.v1alpha.TemplateMacros)
  

- [Scalar Value Types](#scalar-value-types)



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/eventreactors/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/eventreactors/v1alpha/api.proto



<a name="eventreactors.v1alpha.AddRequest"></a>

### AddRequest
Add


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| eventReactor | [EventReactor](#eventreactors.v1alpha.EventReactor) |  | EventReactor object. |






<a name="eventreactors.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| eventReactor | [EventReactor](#eventreactors.v1alpha.EventReactor) |  | EventReactor object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="eventreactors.v1alpha.DeleteRequest"></a>

### DeleteRequest
Delete


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| eventReactor | [EventReactor](#eventreactors.v1alpha.EventReactor) |  | EventReactor object. |






<a name="eventreactors.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| eventReactor | [EventReactor](#eventreactors.v1alpha.EventReactor) |  | EventReactor object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="eventreactors.v1alpha.EventReactor"></a>

### EventReactor
EventReactor Reaction to certain events.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for EventReactor object. |
| spec | [EventReactorSpec](#eventreactors.v1alpha.EventReactorSpec) |  | Spec for EventReactor. |






<a name="eventreactors.v1alpha.EventReactorSpec"></a>

### EventReactorSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reaction | [Reaction](#eventreactors.v1alpha.Reaction) |  | Reaction. |
| task | [core.v1alpha.Conditions](#core.v1alpha.Conditions) |  | Conditions. |
| triggerLimiter | [TriggerLimiter](#eventreactors.v1alpha.TriggerLimiter) |  | TriggerLimiter. |






<a name="eventreactors.v1alpha.GetRequest"></a>

### GetRequest
Get


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |






<a name="eventreactors.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| eventReactor | [EventReactor](#eventreactors.v1alpha.EventReactor) |  | EventReactor object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="eventreactors.v1alpha.ListRequest"></a>

### ListRequest
List


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |






<a name="eventreactors.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| eventReactors | [EventReactor](#eventreactors.v1alpha.EventReactor) | repeated | EventReactor list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="eventreactors.v1alpha.Reaction"></a>

### Reaction
Reaction currently only jobs can be triggered when an event occurs.

Jobs jobs.Jobs `yaml:&#34;jobs&#34;`
TODO Allow other events to be triggered here.
Events events.Events `yaml:&#34;events&#34;`






<a name="eventreactors.v1alpha.TriggerLimiter"></a>

### TriggerLimiter
TriggerLimiter trigger limiter options.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| MinOccurences | [int64](#int64) |  | MinOccurences |
| reReactDelay | [google.protobuf.Duration](#google.protobuf.Duration) |  | ReReactDelay. |
| waitForMore | [WaitForMore](#eventreactors.v1alpha.WaitForMore) |  | WaitForMore. |






<a name="eventreactors.v1alpha.UpdateRequest"></a>

### UpdateRequest
Update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| eventReactor | [EventReactor](#eventreactors.v1alpha.EventReactor) |  | EventReactor object. |






<a name="eventreactors.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| eventReactor | [EventReactor](#eventreactors.v1alpha.EventReactor) |  | EventReactor object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="eventreactors.v1alpha.WaitForMore"></a>

### WaitForMore
WaitForMore trigger limiter WaitForMore options.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  | Timeout. |
| resetTimeoutOnMore | [bool](#bool) |  | ResetTimeoutOnMore |






<a name="eventreactors.v1alpha.WatchRequest"></a>

### WatchRequest
Watch


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |






<a name="eventreactors.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Event info for watch response. |
| EventReactor | [EventReactor](#eventreactors.v1alpha.EventReactor) |  | EventReactor for watch response. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |





 

 

 


<a name="eventreactors.v1alpha.EventReactors"></a>

### EventReactors


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#eventreactors.v1alpha.GetRequest) | [GetResponse](#eventreactors.v1alpha.GetResponse) | Get a EventReactor. |
| List | [ListRequest](#eventreactors.v1alpha.ListRequest) | [ListResponse](#eventreactors.v1alpha.ListResponse) | List EventReactors. |
| Add | [AddRequest](#eventreactors.v1alpha.AddRequest) | [AddResponse](#eventreactors.v1alpha.AddResponse) | Add a EventReactor. |
| Update | [UpdateRequest](#eventreactors.v1alpha.UpdateRequest) | [UpdateResponse](#eventreactors.v1alpha.UpdateResponse) | Update a EventReactor. |
| Delete | [DeleteRequest](#eventreactors.v1alpha.DeleteRequest) | [DeleteResponse](#eventreactors.v1alpha.DeleteResponse) | Delete a EventReactor. |
| Watch | [WatchRequest](#eventreactors.v1alpha.WatchRequest) | [WatchResponse](#eventreactors.v1alpha.WatchResponse) stream | Watch EventReactor. |

 



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/jobs/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/jobs/v1alpha/api.proto



<a name="jobs.v1alpha.AddRequest"></a>

### AddRequest
Add


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| job | [Job](#jobs.v1alpha.Job) |  | Job object. |






<a name="jobs.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| job | [Job](#jobs.v1alpha.Job) |  | Job object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="jobs.v1alpha.DeleteRequest"></a>

### DeleteRequest
Delete


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| job | [Job](#jobs.v1alpha.Job) |  | Job object. |






<a name="jobs.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| job | [Job](#jobs.v1alpha.Job) |  | Job object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="jobs.v1alpha.GetRequest"></a>

### GetRequest
Get


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |






<a name="jobs.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| job | [Job](#jobs.v1alpha.Job) |  | Job object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="jobs.v1alpha.Job"></a>

### Job



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for Job object. |
| spec | [core.v1alpha.Include](#core.v1alpha.Include) |  | Spec for Job. |






<a name="jobs.v1alpha.ListRequest"></a>

### ListRequest
List


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |






<a name="jobs.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| jobs | [Job](#jobs.v1alpha.Job) | repeated | Config list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="jobs.v1alpha.UpdateRequest"></a>

### UpdateRequest
Update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| job | [Job](#jobs.v1alpha.Job) |  | Job object. |






<a name="jobs.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| job | [Job](#jobs.v1alpha.Job) |  | Job object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="jobs.v1alpha.WatchRequest"></a>

### WatchRequest
Watch


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |






<a name="jobs.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Event info for watch response. |
| job | [Job](#jobs.v1alpha.Job) |  | Job for watch response. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |





 

 

 


<a name="jobs.v1alpha.Jobs"></a>

### Jobs


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#jobs.v1alpha.GetRequest) | [GetResponse](#jobs.v1alpha.GetResponse) | Get a Job. |
| List | [ListRequest](#jobs.v1alpha.ListRequest) | [ListResponse](#jobs.v1alpha.ListResponse) | List Jobs. |
| Add | [AddRequest](#jobs.v1alpha.AddRequest) | [AddResponse](#jobs.v1alpha.AddResponse) | Add a Job. |
| Update | [UpdateRequest](#jobs.v1alpha.UpdateRequest) | [UpdateResponse](#jobs.v1alpha.UpdateResponse) | Update a Job. |
| Delete | [DeleteRequest](#jobs.v1alpha.DeleteRequest) | [DeleteResponse](#jobs.v1alpha.DeleteResponse) | Delete a Job. |
| Watch | [WatchRequest](#jobs.v1alpha.WatchRequest) | [WatchResponse](#jobs.v1alpha.WatchResponse) stream | Watch Job. |

 



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
Variables variables file structure. This covers the &#34;Variables File&#34; and &#34;&#39;From Host&#39; Static Variable File&#34;.


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

 



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/tasks/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/tasks/v1alpha/api.proto



<a name="tasks.v1alpha.Action"></a>

### Action



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [string](#string) |  | Action name. |
| conditions | [core.v1alpha.Conditions](#core.v1alpha.Conditions) |  | Conditions. |
| runOptions | [RunOptions](#tasks.v1alpha.RunOptions) |  | RunOptions |
| parameters | [google.protobuf.Any](#google.protobuf.Any) | repeated | Map of string interface{} to be given to the action. |






<a name="tasks.v1alpha.AddRequest"></a>

### AddRequest
Add


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Task](#tasks.v1alpha.Task) |  | Task object. |






<a name="tasks.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Task](#tasks.v1alpha.Task) |  | Task object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="tasks.v1alpha.DeleteRequest"></a>

### DeleteRequest
Delete


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Task](#tasks.v1alpha.Task) |  | Task object. |






<a name="tasks.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Task](#tasks.v1alpha.Task) |  | Task object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="tasks.v1alpha.GetRequest"></a>

### GetRequest
Get


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |






<a name="tasks.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Task](#tasks.v1alpha.Task) |  | Task object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="tasks.v1alpha.ListRequest"></a>

### ListRequest
List


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |






<a name="tasks.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tasks | [Task](#tasks.v1alpha.Task) | repeated | Task list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="tasks.v1alpha.RunOptions"></a>

### RunOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sync | [bool](#bool) |  |  |
| serialize | [Serialize](#tasks.v1alpha.Serialize) |  |  |
| target | [Target](#tasks.v1alpha.Target) |  |  |






<a name="tasks.v1alpha.Serialize"></a>

### Serialize



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  | Count. |
| failAfterAllTargetsRan | [bool](#bool) |  | FailAfterAllTargetsRan. |






<a name="tasks.v1alpha.Step"></a>

### Step



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| spec | [Action](#tasks.v1alpha.Action) |  |  |






<a name="tasks.v1alpha.Target"></a>

### Target



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hosts | [string](#string) | repeated | List of hosts (either name or label selector). |
| limit | [int64](#int64) |  | Limit of how many hosts of the matches should be used. |






<a name="tasks.v1alpha.Task"></a>

### Task



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for Task object. |
| spec | [TaskSpec](#tasks.v1alpha.TaskSpec) |  | Spec for Task. |






<a name="tasks.v1alpha.TaskSpec"></a>

### TaskSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| runOptions | [RunOptions](#tasks.v1alpha.RunOptions) |  | RunOptions. |
| steps | [Step](#tasks.v1alpha.Step) | repeated | Steps in this Task. |






<a name="tasks.v1alpha.UpdateRequest"></a>

### UpdateRequest
Update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Task](#tasks.v1alpha.Task) |  | Task object. |






<a name="tasks.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task | [Task](#tasks.v1alpha.Task) |  | Task object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="tasks.v1alpha.WatchRequest"></a>

### WatchRequest
Watch


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |






<a name="tasks.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | Task info for watch response. |
| task | [Task](#tasks.v1alpha.Task) |  | Task for watch response. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |





 

 

 


<a name="tasks.v1alpha.Tasks"></a>

### Tasks


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#tasks.v1alpha.GetRequest) | [GetResponse](#tasks.v1alpha.GetResponse) | Get a Task. |
| List | [ListRequest](#tasks.v1alpha.ListRequest) | [ListResponse](#tasks.v1alpha.ListResponse) | List Tasks. |
| Add | [AddRequest](#tasks.v1alpha.AddRequest) | [AddResponse](#tasks.v1alpha.AddResponse) | Add a Task. |
| Update | [UpdateRequest](#tasks.v1alpha.UpdateRequest) | [UpdateResponse](#tasks.v1alpha.UpdateResponse) | Update a Task. |
| Delete | [DeleteRequest](#tasks.v1alpha.DeleteRequest) | [DeleteResponse](#tasks.v1alpha.DeleteResponse) | Delete a Task. |
| Watch | [WatchRequest](#tasks.v1alpha.WatchRequest) | [WatchResponse](#tasks.v1alpha.WatchResponse) stream | Watch Task. |

 



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

 



<a name="github.com/galexrt/edenconfmgmt/pkg/apis/templatemacros/v1alpha/api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## github.com/galexrt/edenconfmgmt/pkg/apis/templatemacros/v1alpha/api.proto



<a name="templatemacros.v1alpha.AddRequest"></a>

### AddRequest
Add


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TemplateMacro | [TemplateMacro](#templatemacros.v1alpha.TemplateMacro) |  | TemplateMacro object. |






<a name="templatemacros.v1alpha.AddResponse"></a>

### AddResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TemplateMacro | [TemplateMacro](#templatemacros.v1alpha.TemplateMacro) |  | TemplateMacro object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="templatemacros.v1alpha.DeleteRequest"></a>

### DeleteRequest
Delete


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TemplateMacro | [TemplateMacro](#templatemacros.v1alpha.TemplateMacro) |  | TemplateMacro object. |






<a name="templatemacros.v1alpha.DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TemplateMacro | [TemplateMacro](#templatemacros.v1alpha.TemplateMacro) |  | TemplateMacro object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="templatemacros.v1alpha.GetRequest"></a>

### GetRequest
Get


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| getOptions | [core.v1alpha.GetOptions](#core.v1alpha.GetOptions) |  | GetOptions options for a GetRequest. |






<a name="templatemacros.v1alpha.GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TemplateMacro | [TemplateMacro](#templatemacros.v1alpha.TemplateMacro) |  | TemplateMacros object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="templatemacros.v1alpha.ListRequest"></a>

### ListRequest
List


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| listOptions | [core.v1alpha.ListOptions](#core.v1alpha.ListOptions) |  | ListOptions options for a ListRequest. |






<a name="templatemacros.v1alpha.ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TemplateMacros | [TemplateMacro](#templatemacros.v1alpha.TemplateMacro) | repeated | TemplateMacro list. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="templatemacros.v1alpha.TemplateMacro"></a>

### TemplateMacro
DataStore templatemacros for the data store.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metadata | [core.v1alpha.ObjectMetadata](#core.v1alpha.ObjectMetadata) |  | Metadata for TemplateMacro object. |
| spec | [TemplateMacroSpec](#templatemacros.v1alpha.TemplateMacroSpec) |  | Spec for TemplateMacro. |






<a name="templatemacros.v1alpha.TemplateMacroItem"></a>

### TemplateMacroItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [string](#string) |  | Action. |
| parameters | [google.protobuf.Any](#google.protobuf.Any) | repeated | Parameters. |






<a name="templatemacros.v1alpha.TemplateMacroSpec"></a>

### TemplateMacroSpec



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| templateMacros | [TemplateMacroItem](#templatemacros.v1alpha.TemplateMacroItem) | repeated |  |






<a name="templatemacros.v1alpha.UpdateRequest"></a>

### UpdateRequest
Update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TemplateMacro | [TemplateMacro](#templatemacros.v1alpha.TemplateMacro) |  | TemplateMacro object. |






<a name="templatemacros.v1alpha.UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| TemplateMacro | [TemplateMacro](#templatemacros.v1alpha.TemplateMacro) |  | TemplateMacro object. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |






<a name="templatemacros.v1alpha.WatchRequest"></a>

### WatchRequest
Watch


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| watchOptions | [core.v1alpha.WatchOptions](#core.v1alpha.WatchOptions) |  | WatchOptions options for WatchRequest. |






<a name="templatemacros.v1alpha.WatchResponse"></a>

### WatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [events.v1alpha.Event](#events.v1alpha.Event) |  | TemplateMacro info for watch response. |
| TemplateMacro | [TemplateMacro](#templatemacros.v1alpha.TemplateMacro) |  | TemplateMacro for watch response. |
| error | [core.v1alpha.Error](#core.v1alpha.Error) |  | Error object. |





 

 

 


<a name="templatemacros.v1alpha.TemplateMacros"></a>

### TemplateMacros


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#templatemacros.v1alpha.GetRequest) | [GetResponse](#templatemacros.v1alpha.GetResponse) | Get a TemplateMacro. |
| List | [ListRequest](#templatemacros.v1alpha.ListRequest) | [ListResponse](#templatemacros.v1alpha.ListResponse) | List TemplateMacros. |
| Add | [AddRequest](#templatemacros.v1alpha.AddRequest) | [AddResponse](#templatemacros.v1alpha.AddResponse) | Add a TemplateMacro. |
| Update | [UpdateRequest](#templatemacros.v1alpha.UpdateRequest) | [UpdateResponse](#templatemacros.v1alpha.UpdateResponse) | Update a TemplateMacro. |
| Delete | [DeleteRequest](#templatemacros.v1alpha.DeleteRequest) | [DeleteResponse](#templatemacros.v1alpha.DeleteResponse) | Delete a TemplateMacro. |
| Watch | [WatchRequest](#templatemacros.v1alpha.WatchRequest) | [WatchResponse](#templatemacros.v1alpha.WatchResponse) stream | Watch TemplateMacros. |

 



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

