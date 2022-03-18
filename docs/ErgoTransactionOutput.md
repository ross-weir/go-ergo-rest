# ErgoTransactionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoxId** | **string** | Base16-encoded transaction box id bytes. Should be 32 bytes long | [optional] 
**Value** | **int64** | Amount of Ergo token | 
**ErgoTree** | **string** | Base16-encoded ergo tree bytes | 
**CreationHeight** | **int32** | Height the output was created at | 
**Assets** | [**[]Asset**](Asset.md) | Assets list in the transaction | [optional] 
**AdditionalRegisters** | **map[string]string** | Ergo box registers | 
**TransactionId** | **string** | Base16-encoded transaction id bytes | [optional] 
**Index** | **int32** | Index in the transaction outputs | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


