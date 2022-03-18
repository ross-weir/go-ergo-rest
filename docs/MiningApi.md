# \MiningApi

All URIs are relative to *http://127.0.0.1:9053*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MiningReadMinerRewardAddress**](MiningApi.md#MiningReadMinerRewardAddress) | **Get** /mining/rewardAddress | Read miner reward address
[**MiningReadMinerRewardPubkey**](MiningApi.md#MiningReadMinerRewardPubkey) | **Get** /mining/rewardPublicKey | Read public key associated with miner rewards
[**MiningRequestBlockCandidate**](MiningApi.md#MiningRequestBlockCandidate) | **Get** /mining/candidate | Request block candidate
[**MiningRequestBlockCandidateWithMandatoryTransactions**](MiningApi.md#MiningRequestBlockCandidateWithMandatoryTransactions) | **Post** /mining/candidateWithTxs | Request block candidate
[**MiningSubmitSolution**](MiningApi.md#MiningSubmitSolution) | **Post** /mining/solution | Submit solution for current candidate



## MiningReadMinerRewardAddress

> RewardAddress MiningReadMinerRewardAddress(ctx, )

Read miner reward address

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**RewardAddress**](RewardAddress.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MiningReadMinerRewardPubkey

> RewardPubKey MiningReadMinerRewardPubkey(ctx, )

Read public key associated with miner rewards

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**RewardPubKey**](RewardPubKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MiningRequestBlockCandidate

> WorkMessage MiningRequestBlockCandidate(ctx, )

Request block candidate

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WorkMessage**](WorkMessage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MiningRequestBlockCandidateWithMandatoryTransactions

> WorkMessage MiningRequestBlockCandidateWithMandatoryTransactions(ctx, ergoTransaction)

Request block candidate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ergoTransaction** | [**[]ErgoTransaction**](ErgoTransaction.md)|  | 

### Return type

[**WorkMessage**](WorkMessage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MiningSubmitSolution

> MiningSubmitSolution(ctx, powSolutions)

Submit solution for current candidate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**powSolutions** | [**PowSolutions**](PowSolutions.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

