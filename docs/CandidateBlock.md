# CandidateBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | [optional] 
**ExtensionHash** | **string** | Base16-encoded 32 byte digest | 
**Timestamp** | **int64** | Basic timestamp definition | [optional] 
**StateRoot** | **string** | Base16-encoded 33 byte digest - digest with extra byte with tree height | [optional] 
**NBits** | **int64** |  | [optional] 
**AdProofBytes** | **string** | Base16-encoded ad proofs | [optional] 
**ParentId** | **string** | Base16-encoded 32 byte modifier id | 
**TransactionsNumber** | **int32** |  | [optional] 
**Transactions** | [**[]ErgoTransaction**](ErgoTransaction.md) | Ergo transaction objects | [optional] 
**Votes** | **string** | Base16-encoded votes for a soft-fork and parameters | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


