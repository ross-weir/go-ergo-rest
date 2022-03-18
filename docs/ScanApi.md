# \ScanApi

All URIs are relative to *http://127.0.0.1:9053*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBox**](ScanApi.md#AddBox) | **Post** /scan/addBox | Adds a box to scans, writes box to database if it is not there. You can use scan number 10 to add a box to the wallet.
[**DeregisterScan**](ScanApi.md#DeregisterScan) | **Post** /scan/deregister | Stop tracking and deregister scan
[**ListAllScans**](ScanApi.md#ListAllScans) | **Get** /scan/listAll | List all the registered scans
[**ListUnspentScans**](ScanApi.md#ListUnspentScans) | **Get** /scan/unspentBoxes/{scanId} | List boxes which are not spent.
[**RegisterScan**](ScanApi.md#RegisterScan) | **Post** /scan/register | Register a scan
[**ScanStopTracking**](ScanApi.md#ScanStopTracking) | **Post** /scan/stopTracking | Stop scan-related box tracking



## AddBox

> string AddBox(ctx, scanIdsBox)

Adds a box to scans, writes box to database if it is not there. You can use scan number 10 to add a box to the wallet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanIdsBox** | [**ScanIdsBox**](ScanIdsBox.md)|  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeregisterScan

> ScanId DeregisterScan(ctx, scanId)

Stop tracking and deregister scan

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanId** | [**ScanId**](ScanId.md)|  | 

### Return type

[**ScanId**](ScanId.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllScans

> []Scan ListAllScans(ctx, )

List all the registered scans

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Scan**](Scan.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUnspentScans

> []WalletBox ListUnspentScans(ctx, scanId, optional)

List boxes which are not spent.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanId** | **int32**| identifier of a scan | 
 **optional** | ***ListUnspentScansOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUnspentScansOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **minConfirmations** | **optional.Int32**| Minimal number of confirmations | [default to 0]
 **minInclusionHeight** | **optional.Int32**| Minimal box inclusion height | [default to 0]

### Return type

[**[]WalletBox**](WalletBox.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterScan

> ScanId RegisterScan(ctx, scanRequest)

Register a scan

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanRequest** | [**ScanRequest**](ScanRequest.md)|  | 

### Return type

[**ScanId**](ScanId.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScanStopTracking

> ScanIdBoxId ScanStopTracking(ctx, scanIdBoxId)

Stop scan-related box tracking

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanIdBoxId** | [**ScanIdBoxId**](ScanIdBoxId.md)|  | 

### Return type

[**ScanIdBoxId**](ScanIdBoxId.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

