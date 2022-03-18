# NodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AppVersion** | **string** |  | 
**FullHeight** | Pointer to **int32** | Can be &#39;null&#39; if state is empty (no full block is applied since node launch) | 
**HeadersHeight** | Pointer to **int32** | Can be &#39;null&#39; if state is empty (no header applied since node launch) | 
**BestFullHeaderId** | Pointer to **string** | Can be &#39;null&#39; if no full block is applied since node launch | 
**PreviousFullHeaderId** | Pointer to **string** | Can be &#39;null&#39; if no full block is applied since node launch | 
**BestHeaderId** | Pointer to **string** | Can be &#39;null&#39; if no header applied since node launch | 
**StateRoot** | Pointer to **string** | Can be &#39;null&#39; if state is empty (no full block is applied since node launch) | 
**StateType** | **string** |  | 
**StateVersion** | Pointer to **string** | Can be &#39;null&#39; if no full block is applied since node launch | 
**IsMining** | **bool** |  | 
**PeersCount** | **int32** | Number of connected peers | 
**UnconfirmedCount** | **int32** | Current unconfirmed transactions count | 
**Difficulty** | Pointer to **int32** | Difficulty on current bestFullHeaderId. Can be &#39;null&#39; if no full block is applied since node launch. Difficulty is a BigInt integer.  | 
**CurrentTime** | **int64** | Current internal node time | 
**LaunchTime** | **int64** | Time when the node was started | 
**HeadersScore** | Pointer to **int32** | Can be &#39;null&#39; if no headers is applied since node launch. headersScore is a BigInt integer. | 
**FullBlocksScore** | Pointer to **int32** | Can be &#39;null&#39; if no full block is applied since node launch. fullBlocksScore is a BigInt integer. | 
**GenesisBlockId** | Pointer to **string** | Can be &#39;null&#39; if genesis blocks is not produced yet | 
**Parameters** | [**Parameters**](.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


