# \NipopowApi

All URIs are relative to *http://127.0.0.1:9053*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPopowHeaderByHeight**](NipopowApi.md#GetPopowHeaderByHeight) | **Get** /nipopow/popowHeaderByHeight/{height} | Construct PoPow header for best header at given height
[**GetPopowHeaderById**](NipopowApi.md#GetPopowHeaderById) | **Get** /nipopow/popowHeaderById/{headerId} | Construct PoPow header according to given header id
[**GetPopowProof**](NipopowApi.md#GetPopowProof) | **Get** /nipopow/proof/{minChainLength}/{suffixLength} | Construct PoPoW proof for given min superchain length and suffix length
[**GetPopowProofByHeaderId**](NipopowApi.md#GetPopowProofByHeaderId) | **Get** /nipopow/proof/{minChainLength}/{suffixLength}/{headerId} | Construct PoPoW proof for given min superchain length, suffix length and header ID



## GetPopowHeaderByHeight

> PopowHeader GetPopowHeaderByHeight(ctx, height)

Construct PoPow header for best header at given height

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**height** | **int32**| Height of a wanted header | 

### Return type

[**PopowHeader**](PopowHeader.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPopowHeaderById

> PopowHeader GetPopowHeaderById(ctx, headerId)

Construct PoPow header according to given header id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**headerId** | **string**| ID of wanted header | 

### Return type

[**PopowHeader**](PopowHeader.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPopowProof

> NipopowProof GetPopowProof(ctx, minChainLength, suffixLength)

Construct PoPoW proof for given min superchain length and suffix length

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**minChainLength** | **float32**| Minimal superchain length | 
**suffixLength** | **float32**| Suffix length | 

### Return type

[**NipopowProof**](NipopowProof.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPopowProofByHeaderId

> NipopowProof GetPopowProofByHeaderId(ctx, minChainLength, suffixLength, headerId)

Construct PoPoW proof for given min superchain length, suffix length and header ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**minChainLength** | **float32**| Minimal superchain length | 
**suffixLength** | **float32**| Suffix length | 
**headerId** | **string**| ID of wanted header | 

### Return type

[**NipopowProof**](NipopowProof.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

