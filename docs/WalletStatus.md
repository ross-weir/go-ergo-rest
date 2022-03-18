# WalletStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsInitialized** | **bool** | true if wallet is initialized, false otherwise | 
**IsUnlocked** | **bool** | true if wallet is unlocked, false otherwise | 
**ChangeAddress** | **string** | address to send change to. Empty when wallet is not initialized or locked. By default change address correponds to root key address, could be set via /wallet/updateChangeAddress method. | 
**WalletHeight** | **int32** | last scanned height for the wallet (and external scans) | 
**Error** | **string** | last wallet error caught | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


