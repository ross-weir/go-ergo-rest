# \TransactionsApi

All URIs are relative to *http://127.0.0.1:9053*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckTransaction**](TransactionsApi.md#CheckTransaction) | **Post** /transactions/check | Checks an Ergo transaction without sending it over the network. Checks that transaction is valid and its inputs are in the UTXO set. Returns transaction identifier if the transaction is passing the checks.
[**GetExpectedWaitTime**](TransactionsApi.md#GetExpectedWaitTime) | **Get** /transactions/waitTime | Get expected wait time for the transaction with specified fee and size
[**GetFeeHistogram**](TransactionsApi.md#GetFeeHistogram) | **Get** /transactions/poolHistogram | Get histogram (waittime, (n_trans, sum(fee)) for transactions in mempool. It contains \&quot;bins\&quot;+1 bins, where i-th elements corresponds to transaction with wait time [i*maxtime/bins, (i+1)*maxtime/bins), and last bin corresponds to the transactions with wait time &gt;&#x3D; maxtime.
[**GetRecommendedFee**](TransactionsApi.md#GetRecommendedFee) | **Get** /transactions/getFee | Get recommended fee (in nanoErgs) for a transaction with specified size (in bytes) to be proceeded in specified time (in minutes)
[**GetUnconfirmedTransactions**](TransactionsApi.md#GetUnconfirmedTransactions) | **Get** /transactions/unconfirmed | Get current pool of the unconfirmed transactions pool
[**SendTransaction**](TransactionsApi.md#SendTransaction) | **Post** /transactions | Submit an Ergo transaction to unconfirmed pool to send it over the network



## CheckTransaction

> string CheckTransaction(ctx, ergoTransaction)

Checks an Ergo transaction without sending it over the network. Checks that transaction is valid and its inputs are in the UTXO set. Returns transaction identifier if the transaction is passing the checks.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ergoTransaction** | [**ErgoTransaction**](ErgoTransaction.md)|  | 

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


## GetExpectedWaitTime

> int32 GetExpectedWaitTime(ctx, fee, txSize)

Get expected wait time for the transaction with specified fee and size

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fee** | **int32**| Transaction fee (in nanoErgs) | [default to 1]
**txSize** | **int32**| Transaction size | [default to 100]

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeeHistogram

> []FeeHistogramBin GetFeeHistogram(ctx, optional)

Get histogram (waittime, (n_trans, sum(fee)) for transactions in mempool. It contains \"bins\"+1 bins, where i-th elements corresponds to transaction with wait time [i*maxtime/bins, (i+1)*maxtime/bins), and last bin corresponds to the transactions with wait time >= maxtime.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetFeeHistogramOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFeeHistogramOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bins** | **optional.Int32**| The number of bins in histogram | [default to 10]
 **maxtime** | **optional.Int64**| Maximal wait time in milliseconds | [default to 60000]

### Return type

[**[]FeeHistogramBin**](FeeHistogramBin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecommendedFee

> int32 GetRecommendedFee(ctx, waitTime, txSize)

Get recommended fee (in nanoErgs) for a transaction with specified size (in bytes) to be proceeded in specified time (in minutes)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waitTime** | **int32**| Maximum transaction wait time in minutes | [default to 1]
**txSize** | **int32**| Transaction size | [default to 100]

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnconfirmedTransactions

> []ErgoTransaction GetUnconfirmedTransactions(ctx, optional)

Get current pool of the unconfirmed transactions pool

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUnconfirmedTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUnconfirmedTransactionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of items in list to return | [default to 50]
 **offset** | **optional.Int32**| The number of items in list to skip | [default to 0]

### Return type

[**[]ErgoTransaction**](ErgoTransaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTransaction

> string SendTransaction(ctx, ergoTransaction)

Submit an Ergo transaction to unconfirmed pool to send it over the network

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ergoTransaction** | [**ErgoTransaction**](ErgoTransaction.md)|  | 

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

