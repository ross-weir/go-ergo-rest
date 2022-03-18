# \UtilsApi

All URIs are relative to *http://127.0.0.1:9053*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressToRaw**](UtilsApi.md#AddressToRaw) | **Get** /utils/addressToRaw/{address} | Convert Pay-To-Public-Key Address to raw representation (hex-encoded serialized curve point)
[**CheckAddressValidity**](UtilsApi.md#CheckAddressValidity) | **Get** /utils/address/{address} | Check address validity
[**ErgoTreeToAddress**](UtilsApi.md#ErgoTreeToAddress) | **Get** /utils/ergoTreeToAddress/{ergoTreeHex} | Generate Ergo address from hex-encoded ErgoTree
[**GetRandomSeed**](UtilsApi.md#GetRandomSeed) | **Get** /utils/seed | Get random seed of 32 bytes
[**GetRandomSeedWithLength**](UtilsApi.md#GetRandomSeedWithLength) | **Get** /utils/seed/{length} | Generate random seed of specified length in bytes
[**HashBlake2b**](UtilsApi.md#HashBlake2b) | **Post** /utils/hash/blake2b | Return Blake2b hash of specified message
[**RawToAddress**](UtilsApi.md#RawToAddress) | **Get** /utils/rawToAddress/{pubkeyHex} | Generate Pay-To-Public-Key address from hex-encoded raw pubkey (secp256k1 serialized point)



## AddressToRaw

> string AddressToRaw(ctx, address)

Convert Pay-To-Public-Key Address to raw representation (hex-encoded serialized curve point)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string**| address to extract public key from | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckAddressValidity

> AddressValidity CheckAddressValidity(ctx, address)

Check address validity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string**| address to check | 

### Return type

[**AddressValidity**](AddressValidity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ErgoTreeToAddress

> string ErgoTreeToAddress(ctx, ergoTreeHex)

Generate Ergo address from hex-encoded ErgoTree

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ergoTreeHex** | **string**| ErgoTree to derive an address from | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRandomSeed

> string GetRandomSeed(ctx, )

Get random seed of 32 bytes

### Required Parameters

This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRandomSeedWithLength

> string GetRandomSeedWithLength(ctx, length)

Generate random seed of specified length in bytes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**length** | **string**| seed length in bytes | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HashBlake2b

> string HashBlake2b(ctx, body)

Return Blake2b hash of specified message

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RawToAddress

> string RawToAddress(ctx, pubkeyHex)

Generate Pay-To-Public-Key address from hex-encoded raw pubkey (secp256k1 serialized point)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pubkeyHex** | **string**| public key to get address from | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

