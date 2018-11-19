# Protocol Documentation
<a name="top"></a>

## Table of Contents

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

