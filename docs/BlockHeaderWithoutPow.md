# BlockHeaderWithoutPow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Base16-encoded 32 byte modifier id | 
**Timestamp** | **int64** | Basic timestamp definition | 
**Version** | **int32** | Ergo blockchain protocol version | 
**AdProofsRoot** | **string** | Base16-encoded 32 byte digest | 
**StateRoot** | **string** | Base16-encoded 33 byte digest - digest with extra byte with tree height | 
**TransactionsRoot** | **string** | Base16-encoded 32 byte digest | 
**NBits** | **int64** |  | 
**ExtensionHash** | **string** | Base16-encoded 32 byte digest | 
**Height** | **int32** |  | 
**Difficulty** | **int32** |  | 
**ParentId** | **string** | Base16-encoded 32 byte modifier id | 
**Votes** | **string** | Base16-encoded votes for a soft-fork and parameters | 
**Size** | **int32** | Size in bytes | [optional] 
**ExtensionId** | **string** | Base16-encoded 32 byte modifier id | [optional] 
**TransactionsId** | **string** | Base16-encoded 32 byte modifier id | [optional] 
**AdProofsId** | **string** | Base16-encoded 32 byte modifier id | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


