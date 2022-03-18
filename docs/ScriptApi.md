# \ScriptApi

All URIs are relative to *http://127.0.0.1:9053*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressToBytes**](ScriptApi.md#AddressToBytes) | **Get** /script/addressToBytes/{address} | Convert an address to hex-encoded Sigma byte array constant which contains script bytes
[**AddressToTree**](ScriptApi.md#AddressToTree) | **Get** /script/addressToTree/{address} | Convert an address to hex-encoded serialized ErgoTree (script)
[**ExecuteWithContext**](ScriptApi.md#ExecuteWithContext) | **Post** /script/executeWithContext | Execute script with context
[**ScriptP2SAddress**](ScriptApi.md#ScriptP2SAddress) | **Post** /script/p2sAddress | Create P2SAddress from Sigma source
[**ScriptP2SHAddress**](ScriptApi.md#ScriptP2SHAddress) | **Post** /script/p2shAddress | Create P2SHAddress from Sigma source



## AddressToBytes

> ScriptBytes AddressToBytes(ctx, address)

Convert an address to hex-encoded Sigma byte array constant which contains script bytes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string**| address to get a script from | 

### Return type

[**ScriptBytes**](ScriptBytes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressToTree

> ErgoTreeObject AddressToTree(ctx, address)

Convert an address to hex-encoded serialized ErgoTree (script)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string**| address to get a script from | 

### Return type

[**ErgoTreeObject**](ErgoTreeObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteWithContext

> CryptoResult ExecuteWithContext(ctx, executeScript)

Execute script with context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executeScript** | [**ExecuteScript**](ExecuteScript.md)|  | 

### Return type

[**CryptoResult**](CryptoResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScriptP2SAddress

> AddressHolder ScriptP2SAddress(ctx, sourceHolder)

Create P2SAddress from Sigma source

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceHolder** | [**SourceHolder**](SourceHolder.md)|  | 

### Return type

[**AddressHolder**](AddressHolder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScriptP2SHAddress

> AddressHolder ScriptP2SHAddress(ctx, sourceHolder)

Create P2SHAddress from Sigma source

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceHolder** | [**SourceHolder**](SourceHolder.md)|  | 

### Return type

[**AddressHolder**](AddressHolder.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

