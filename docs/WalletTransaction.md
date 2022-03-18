# WalletTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Base16-encoded transaction id bytes | [optional] 
**Inputs** | [**[]ErgoTransactionInput**](ErgoTransactionInput.md) | Transaction inputs | 
**DataInputs** | [**[]ErgoTransactionDataInput**](ErgoTransactionDataInput.md) | Transaction data inputs | 
**Outputs** | [**[]ErgoTransactionOutput**](ErgoTransactionOutput.md) | Transaction outputs | 
**InclusionHeight** | **int32** | Height of a block the transaction was included in | 
**NumConfirmations** | **int32** | Number of transaction confirmations | 
**Scans** | **[]int32** | Scan identifiers the transaction relates to | 
**Size** | **int32** | Size in bytes | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


