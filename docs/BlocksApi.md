# \BlocksApi

All URIs are relative to *http://127.0.0.1:9053*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlockHeaderById**](BlocksApi.md#GetBlockHeaderById) | **Get** /blocks/{headerId}/header | Get the block header info by a given signature
[**GetBlockTransactionsById**](BlocksApi.md#GetBlockTransactionsById) | **Get** /blocks/{headerId}/transactions | Get the block transactions info by a given signature
[**GetChainSlice**](BlocksApi.md#GetChainSlice) | **Get** /blocks/chainSlice | Get headers in a specified range
[**GetFullBlockAt**](BlocksApi.md#GetFullBlockAt) | **Get** /blocks/at/{blockHeight} | Get the header ids at a given height
[**GetFullBlockById**](BlocksApi.md#GetFullBlockById) | **Get** /blocks/{headerId} | Get the full block info by a given signature
[**GetHeaderIds**](BlocksApi.md#GetHeaderIds) | **Get** /blocks | Get the Array of header ids
[**GetLastHeaders**](BlocksApi.md#GetLastHeaders) | **Get** /blocks/lastHeaders/{count} | Get the last headers objects
[**GetModifierById**](BlocksApi.md#GetModifierById) | **Get** /blocks/modifier/{modifierId} | Get the persistent modifier by its id
[**GetProofForTx**](BlocksApi.md#GetProofForTx) | **Get** /blocks/{headerId}/proofFor/{txId} | Get Merkle proof for transaction
[**SendMinedBlock**](BlocksApi.md#SendMinedBlock) | **Post** /blocks | Send a mined block



## GetBlockHeaderById

> BlockHeader GetBlockHeaderById(ctx, headerId)

Get the block header info by a given signature

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**headerId** | **string**| ID of a wanted block header | 

### Return type

[**BlockHeader**](BlockHeader.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockTransactionsById

> BlockTransactions GetBlockTransactionsById(ctx, headerId)

Get the block transactions info by a given signature

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**headerId** | **string**| ID of a wanted block transactions | 

### Return type

[**BlockTransactions**](BlockTransactions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChainSlice

> []BlockHeader GetChainSlice(ctx, optional)

Get headers in a specified range

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetChainSliceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetChainSliceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromHeight** | **optional.Int32**| Min header height | [default to 0]
 **toHeight** | **optional.Int32**| Max header height (best header height by default) | [default to -1]

### Return type

[**[]BlockHeader**](BlockHeader.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFullBlockAt

> []string GetFullBlockAt(ctx, blockHeight)

Get the header ids at a given height

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockHeight** | **int32**| Height of a wanted block | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFullBlockById

> FullBlock GetFullBlockById(ctx, headerId)

Get the full block info by a given signature

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**headerId** | **string**| ID of a wanted block | 

### Return type

[**FullBlock**](FullBlock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHeaderIds

> []string GetHeaderIds(ctx, optional)

Get the Array of header ids

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetHeaderIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetHeaderIdsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items in list to return | [default to 50]
 **offset** | **optional.Int32**| The number of items in list to skip | [default to 0]

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastHeaders

> []BlockHeader GetLastHeaders(ctx, count)

Get the last headers objects

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**count** | **float32**| count of a wanted block headers | 

### Return type

[**[]BlockHeader**](BlockHeader.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModifierById

> GetModifierById(ctx, modifierId)

Get the persistent modifier by its id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modifierId** | **string**| ID of a wanted modifier | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProofForTx

> MerkleProof GetProofForTx(ctx, headerId, txId)

Get Merkle proof for transaction

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**headerId** | **string**| ID of a wanted block transactions | 
**txId** | **string**| ID of a wanted transaction | 

### Return type

[**MerkleProof**](MerkleProof.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMinedBlock

> SendMinedBlock(ctx, fullBlock)

Send a mined block

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fullBlock** | [**FullBlock**](FullBlock.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

