# SigmaHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Base16-encoded 32 byte modifier id | [optional] 
**Timestamp** | **int64** | Basic timestamp definition | 
**Version** | **int32** | Ergo blockchain protocol version | 
**AdProofsRoot** | **string** | Base16-encoded 32 byte digest | 
**AdProofsId** | **string** | Base16-encoded 32 byte modifier id | [optional] 
**StateRoot** | [**AvlTreeData**](AvlTreeData.md) |  | 
**TransactionsRoot** | **string** | Base16-encoded 32 byte digest | 
**TransactionsId** | **string** | Base16-encoded 32 byte modifier id | [optional] 
**NBits** | **int64** |  | 
**ExtensionHash** | **string** | Base16-encoded 32 byte digest | 
**ExtensionRoot** | **string** | Base16-encoded 32 byte digest | [optional] 
**ExtensionId** | **string** | Base16-encoded 32 byte modifier id | [optional] 
**Height** | **int32** |  | 
**Size** | **int32** |  | [optional] 
**ParentId** | **string** | Base16-encoded 32 byte modifier id | 
**PowSolutions** | [**PowSolutions**](PowSolutions.md) |  | [optional] 
**Votes** | **string** | Base16-encoded votes for a soft-fork and parameters | 
**MinerPk** | **string** |  | [optional] 
**PowOnetimePk** | **string** |  | [optional] 
**PowNonce** | **string** | Base16-encoded 32 byte digest | [optional] 
**PowDistance** | **float32** | sigma.BigInt | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


