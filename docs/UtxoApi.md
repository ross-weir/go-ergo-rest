# \UtxoApi

All URIs are relative to *http://127.0.0.1:9053*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenesisBoxes**](UtxoApi.md#GenesisBoxes) | **Get** /utxo/genesis | Get genesis boxes (boxes existed before the very first block)
[**GetBoxById**](UtxoApi.md#GetBoxById) | **Get** /utxo/byId/{boxId} | Get box contents for a box by a unique identifier.
[**GetBoxByIdBinary**](UtxoApi.md#GetBoxByIdBinary) | **Get** /utxo/byIdBinary/{boxId} | Get serialized box from UTXO pool in Base16 encoding by an identifier.
[**GetBoxWithPoolById**](UtxoApi.md#GetBoxWithPoolById) | **Get** /utxo/withPool/byId/{boxId} | Get box contents for a box by a unique identifier, from UTXO set and also the mempool.
[**GetBoxWithPoolByIdBinary**](UtxoApi.md#GetBoxWithPoolByIdBinary) | **Get** /utxo/withPool/byIdBinary/{boxId} | Get serialized box in Base16 encoding by an identifier, considering also the mempool.
[**GetBoxesBinaryProof**](UtxoApi.md#GetBoxesBinaryProof) | **Post** /utxo/getBoxesBinaryProof | Get serialized batch proof for given set of boxes



## GenesisBoxes

> []ErgoTransactionOutput GenesisBoxes(ctx, )

Get genesis boxes (boxes existed before the very first block)

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ErgoTransactionOutput**](ErgoTransactionOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoxById

> ErgoTransactionOutput GetBoxById(ctx, boxId)

Get box contents for a box by a unique identifier.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boxId** | **string**| ID of a wanted box | 

### Return type

[**ErgoTransactionOutput**](ErgoTransactionOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoxByIdBinary

> SerializedBox GetBoxByIdBinary(ctx, boxId)

Get serialized box from UTXO pool in Base16 encoding by an identifier.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boxId** | **string**| ID of a wanted box | 

### Return type

[**SerializedBox**](SerializedBox.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoxWithPoolById

> ErgoTransactionOutput GetBoxWithPoolById(ctx, boxId)

Get box contents for a box by a unique identifier, from UTXO set and also the mempool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boxId** | **string**| ID of a box to obtain | 

### Return type

[**ErgoTransactionOutput**](ErgoTransactionOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoxWithPoolByIdBinary

> SerializedBox GetBoxWithPoolByIdBinary(ctx, boxId)

Get serialized box in Base16 encoding by an identifier, considering also the mempool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boxId** | **string**| ID of a wanted box | 

### Return type

[**SerializedBox**](SerializedBox.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoxesBinaryProof

> string GetBoxesBinaryProof(ctx, requestBody)

Get serialized batch proof for given set of boxes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestBody** | [**[]string**](string.md)|  | 

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

