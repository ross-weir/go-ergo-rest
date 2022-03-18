# \WalletApi

All URIs are relative to *http://127.0.0.1:9053*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBox**](WalletApi.md#AddBox) | **Post** /scan/addBox | Adds a box to scans, writes box to database if it is not there. You can use scan number 10 to add a box to the wallet.
[**CheckSeed**](WalletApi.md#CheckSeed) | **Post** /wallet/check | Check whether mnemonic phrase is corresponding to the wallet seed
[**ExtractHints**](WalletApi.md#ExtractHints) | **Post** /wallet/extractHints | Extract hints from a transaction
[**GenerateCommitments**](WalletApi.md#GenerateCommitments) | **Post** /wallet/generateCommitments | Generate signature commitments for inputs of an unsigned transaction
[**GetWalletStatus**](WalletApi.md#GetWalletStatus) | **Get** /wallet/status | Get wallet status
[**WalletAddresses**](WalletApi.md#WalletAddresses) | **Get** /wallet/addresses | Get wallet addresses
[**WalletBalances**](WalletApi.md#WalletBalances) | **Get** /wallet/balances | Get total amount of confirmed Ergo tokens and assets
[**WalletBalancesUnconfirmed**](WalletApi.md#WalletBalancesUnconfirmed) | **Get** /wallet/balances/withUnconfirmed | Get summary amount of confirmed plus unconfirmed Ergo tokens and assets
[**WalletBoxes**](WalletApi.md#WalletBoxes) | **Get** /wallet/boxes | Get a list of all wallet-related boxes, both spent and unspent. Set minConfirmations to -1 to get mempool boxes included.
[**WalletBoxesCollect**](WalletApi.md#WalletBoxesCollect) | **Post** /wallet/boxes/collect | Get a list of collected boxes.
[**WalletDeriveKey**](WalletApi.md#WalletDeriveKey) | **Post** /wallet/deriveKey | Derive new key according to a provided path
[**WalletDeriveNextKey**](WalletApi.md#WalletDeriveNextKey) | **Get** /wallet/deriveNextKey | Derive next key
[**WalletGetTransaction**](WalletApi.md#WalletGetTransaction) | **Get** /wallet/transactionById | Get wallet-related transaction by id
[**WalletInit**](WalletApi.md#WalletInit) | **Post** /wallet/init | Initialize new wallet with randomly generated seed
[**WalletLock**](WalletApi.md#WalletLock) | **Get** /wallet/lock | Lock wallet
[**WalletPaymentTransactionGenerateAndSend**](WalletApi.md#WalletPaymentTransactionGenerateAndSend) | **Post** /wallet/payment/send | Generate and send payment transaction (default fee of 0.001 Erg is used)
[**WalletRescan**](WalletApi.md#WalletRescan) | **Get** /wallet/rescan | Rescan wallet (all the available full blocks)
[**WalletRestore**](WalletApi.md#WalletRestore) | **Post** /wallet/restore | Create new wallet from existing mnemonic seed
[**WalletTransactionGenerate**](WalletApi.md#WalletTransactionGenerate) | **Post** /wallet/transaction/generate | Generate arbitrary transaction from array of requests.
[**WalletTransactionGenerateAndSend**](WalletApi.md#WalletTransactionGenerateAndSend) | **Post** /wallet/transaction/send | Generate and send arbitrary transaction
[**WalletTransactionSign**](WalletApi.md#WalletTransactionSign) | **Post** /wallet/transaction/sign | Sign arbitrary unsigned transaction with wallet secrets and also secrets provided.
[**WalletTransactions**](WalletApi.md#WalletTransactions) | **Get** /wallet/transactions | Get a list of all wallet-related transactions
[**WalletTransactionsByScanId**](WalletApi.md#WalletTransactionsByScanId) | **Get** /wallet/transactionsByScanId/{scanId} | Get scan-related transactions by scan id
[**WalletUnlock**](WalletApi.md#WalletUnlock) | **Post** /wallet/unlock | Unlock wallet
[**WalletUnsignedTransactionGenerate**](WalletApi.md#WalletUnsignedTransactionGenerate) | **Post** /wallet/transaction/generateUnsigned | Generate unsigned transaction from array of requests.
[**WalletUnspentBoxes**](WalletApi.md#WalletUnspentBoxes) | **Get** /wallet/boxes/unspent | Get a list of unspent boxes. Set minConfirmations to -1 to have mempool boxes considered.
[**WalletUpdateChangeAddress**](WalletApi.md#WalletUpdateChangeAddress) | **Post** /wallet/updateChangeAddress | Update address to be used to send change to



## AddBox

> string AddBox(ctx, scanIdsBox)

Adds a box to scans, writes box to database if it is not there. You can use scan number 10 to add a box to the wallet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanIdsBox** | [**ScanIdsBox**](ScanIdsBox.md)|  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckSeed

> PassphraseMatch CheckSeed(ctx, checkWallet)

Check whether mnemonic phrase is corresponding to the wallet seed

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkWallet** | [**CheckWallet**](CheckWallet.md)|  | 

### Return type

[**PassphraseMatch**](PassphraseMatch.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtractHints

> TransactionHintsBag ExtractHints(ctx, hintExtractionRequest)

Extract hints from a transaction

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hintExtractionRequest** | [**HintExtractionRequest**](HintExtractionRequest.md)|  | 

### Return type

[**TransactionHintsBag**](TransactionHintsBag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateCommitments

> TransactionHintsBag GenerateCommitments(ctx, generateCommitmentsRequest)

Generate signature commitments for inputs of an unsigned transaction

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**generateCommitmentsRequest** | [**GenerateCommitmentsRequest**](GenerateCommitmentsRequest.md)|  | 

### Return type

[**TransactionHintsBag**](TransactionHintsBag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletStatus

> WalletStatus GetWalletStatus(ctx, )

Get wallet status

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**WalletStatus**](WalletStatus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletAddresses

> []string WalletAddresses(ctx, )

Get wallet addresses

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletBalances

> BalancesSnapshot WalletBalances(ctx, )

Get total amount of confirmed Ergo tokens and assets

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**BalancesSnapshot**](BalancesSnapshot.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletBalancesUnconfirmed

> BalancesSnapshot WalletBalancesUnconfirmed(ctx, )

Get summary amount of confirmed plus unconfirmed Ergo tokens and assets

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**BalancesSnapshot**](BalancesSnapshot.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletBoxes

> []WalletBox WalletBoxes(ctx, optional)

Get a list of all wallet-related boxes, both spent and unspent. Set minConfirmations to -1 to get mempool boxes included.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletBoxesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WalletBoxesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minConfirmations** | **optional.Int32**| Minimal number of confirmations | [default to 0]
 **minInclusionHeight** | **optional.Int32**| Minimal box inclusion height | [default to 0]

### Return type

[**[]WalletBox**](WalletBox.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletBoxesCollect

> []WalletBox WalletBoxesCollect(ctx, boxesRequestHolder)

Get a list of collected boxes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boxesRequestHolder** | [**BoxesRequestHolder**](BoxesRequestHolder.md)| This API method recieves balance and assets, according to which, it&#39;s collecting result | 

### Return type

[**[]WalletBox**](WalletBox.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletDeriveKey

> DeriveKeyResult WalletDeriveKey(ctx, deriveKey)

Derive new key according to a provided path

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deriveKey** | [**DeriveKey**](DeriveKey.md)|  | 

### Return type

[**DeriveKeyResult**](DeriveKeyResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletDeriveNextKey

> DeriveNextKeyResult WalletDeriveNextKey(ctx, )

Derive next key

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DeriveNextKeyResult**](DeriveNextKeyResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGetTransaction

> []WalletTransaction WalletGetTransaction(ctx, id)

Get wallet-related transaction by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| Transaction id | 

### Return type

[**[]WalletTransaction**](WalletTransaction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletInit

> InitWalletResult WalletInit(ctx, initWallet)

Initialize new wallet with randomly generated seed

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**initWallet** | [**InitWallet**](InitWallet.md)|  | 

### Return type

[**InitWalletResult**](InitWalletResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletLock

> WalletLock(ctx, )

Lock wallet

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletPaymentTransactionGenerateAndSend

> string WalletPaymentTransactionGenerateAndSend(ctx, paymentRequest)

Generate and send payment transaction (default fee of 0.001 Erg is used)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentRequest** | [**[]PaymentRequest**](PaymentRequest.md)|  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletRescan

> WalletRescan(ctx, )

Rescan wallet (all the available full blocks)

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletRestore

> WalletRestore(ctx, restoreWallet)

Create new wallet from existing mnemonic seed

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restoreWallet** | [**RestoreWallet**](RestoreWallet.md)|  | 

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


## WalletTransactionGenerate

> ErgoTransaction WalletTransactionGenerate(ctx, requestsHolder)

Generate arbitrary transaction from array of requests.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestsHolder** | [**RequestsHolder**](RequestsHolder.md)| This API method receives a sequence of requests as an input. Each request will produce an output of the resulting transaction (with fee output created automatically). Currently supported types of requests are payment and asset issuance requests. An example for a transaction with requests of both kinds is provided below. Please note that for the payment request \&quot;assets\&quot; and \&quot;registers\&quot; fields are not needed. For asset issuance request, \&quot;registers\&quot; field is not needed. You may specify boxes to spend by providing them in \&quot;inputsRaw\&quot;. Please note you need to have strict equality between input and output total amounts of Ergs in this case. If you want wallet to pick up the boxes, leave \&quot;inputsRaw\&quot; empty. | 

### Return type

[**ErgoTransaction**](ErgoTransaction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletTransactionGenerateAndSend

> string WalletTransactionGenerateAndSend(ctx, requestsHolder)

Generate and send arbitrary transaction

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestsHolder** | [**RequestsHolder**](RequestsHolder.md)| See description of /wallet/transaction/generate | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletTransactionSign

> ErgoTransaction WalletTransactionSign(ctx, transactionSigningRequest)

Sign arbitrary unsigned transaction with wallet secrets and also secrets provided.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionSigningRequest** | [**TransactionSigningRequest**](TransactionSigningRequest.md)| With this API method an arbitrary unsigned transaction can be signed with secrets provided or stored in the wallet. Both DLOG and Diffie-Hellman tuple secrets are supported. Please note that the unsigned transaction contains only identifiers of inputs and data inputs. If the node holds UTXO set, it is able to extract boxes needed. Otherwise, input (and data-input) boxes can be provided in \&quot;inputsRaw\&quot; and \&quot;dataInputsRaw\&quot; fields. | 

### Return type

[**ErgoTransaction**](ErgoTransaction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletTransactions

> []WalletTransaction WalletTransactions(ctx, optional)

Get a list of all wallet-related transactions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WalletTransactionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minInclusionHeight** | **optional.Int32**| Minimal tx inclusion height | 
 **maxInclusionHeight** | **optional.Int32**| Maximal tx inclusion height | 
 **minConfirmations** | **optional.Int32**| Minimal confirmations number | 
 **maxConfirmations** | **optional.Int32**| Maximal confirmations number | 

### Return type

[**[]WalletTransaction**](WalletTransaction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletTransactionsByScanId

> []WalletTransaction WalletTransactionsByScanId(ctx, scanId, optional)

Get scan-related transactions by scan id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scanId** | **int32**| identifier of a scan | 
 **optional** | ***WalletTransactionsByScanIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WalletTransactionsByScanIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **minInclusionHeight** | **optional.Int32**| Minimal tx inclusion height | 
 **maxInclusionHeight** | **optional.Int32**| Maximal tx inclusion height | 
 **minConfirmations** | **optional.Int32**| Minimal confirmations number | 
 **maxConfirmations** | **optional.Int32**| Maximal confirmations number | 

### Return type

[**[]WalletTransaction**](WalletTransaction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletUnlock

> WalletUnlock(ctx, unlockWallet)

Unlock wallet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**unlockWallet** | [**UnlockWallet**](UnlockWallet.md)|  | 

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


## WalletUnsignedTransactionGenerate

> UnsignedErgoTransaction WalletUnsignedTransactionGenerate(ctx, requestsHolder)

Generate unsigned transaction from array of requests.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestsHolder** | [**RequestsHolder**](RequestsHolder.md)| The same as /wallet/transaction/generate but generates unsigned transaction. | 

### Return type

[**UnsignedErgoTransaction**](UnsignedErgoTransaction.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletUnspentBoxes

> []WalletBox WalletUnspentBoxes(ctx, optional)

Get a list of unspent boxes. Set minConfirmations to -1 to have mempool boxes considered.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WalletUnspentBoxesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a WalletUnspentBoxesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minConfirmations** | **optional.Int32**| Minimal number of confirmations | [default to 0]
 **minInclusionHeight** | **optional.Int32**| Minimal box inclusion height | [default to 0]

### Return type

[**[]WalletBox**](WalletBox.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletUpdateChangeAddress

> WalletUpdateChangeAddress(ctx, body)

Update address to be used to send change to

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

