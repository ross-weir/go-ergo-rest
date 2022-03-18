# \PeersApi

All URIs are relative to *http://127.0.0.1:9053*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectToPeer**](PeersApi.md#ConnectToPeer) | **Post** /peers/connect | Add address to peers list
[**GetAllPeers**](PeersApi.md#GetAllPeers) | **Get** /peers/all | Get all known peers
[**GetBlacklistedPeers**](PeersApi.md#GetBlacklistedPeers) | **Get** /peers/blacklisted | Get blacklisted peers
[**GetConnectedPeers**](PeersApi.md#GetConnectedPeers) | **Get** /peers/connected | Get current connected peers
[**GetPeersStatus**](PeersApi.md#GetPeersStatus) | **Get** /peers/status | Get last incoming message timestamp and current network time
[**GetPeersSyncInfo**](PeersApi.md#GetPeersSyncInfo) | **Get** /peers/syncInfo | Get sync info reported by peers, including versions, current status and height (if available)
[**GetPeersTrackInfo**](PeersApi.md#GetPeersTrackInfo) | **Get** /peers/trackInfo | Get track info reported by peers, including count of invalid modifiers and details of requested and received modifiers



## ConnectToPeer

> ConnectToPeer(ctx, body)

Add address to peers list

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **string**|  | 

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


## GetAllPeers

> []Peer GetAllPeers(ctx, )

Get all known peers

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Peer**](Peer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlacklistedPeers

> BlacklistedPeers GetBlacklistedPeers(ctx, )

Get blacklisted peers

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**BlacklistedPeers**](BlacklistedPeers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectedPeers

> []Peer GetConnectedPeers(ctx, )

Get current connected peers

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]Peer**](Peer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeersStatus

> []PeersStatus GetPeersStatus(ctx, )

Get last incoming message timestamp and current network time

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PeersStatus**](PeersStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeersSyncInfo

> []SyncInfo GetPeersSyncInfo(ctx, )

Get sync info reported by peers, including versions, current status and height (if available)

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SyncInfo**](SyncInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPeersTrackInfo

> []TrackInfo GetPeersTrackInfo(ctx, )

Get track info reported by peers, including count of invalid modifiers and details of requested and received modifiers

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]TrackInfo**](TrackInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

