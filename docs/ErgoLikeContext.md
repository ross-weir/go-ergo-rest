# ErgoLikeContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastBlockUtxoRoot** | [**AvlTreeData**](.md) |  | 
**Headers** | [**[]SigmaHeader**](SigmaHeader.md) | fixed number of last block headers in descending order (first header is the newest one) | 
**PreHeader** | [**PreHeader**](.md) |  | 
**DataBoxes** | [**[]ErgoTransactionOutput**](ErgoTransactionOutput.md) | boxes, that corresponds to id&#39;s of &#x60;spendingTransaction.dataInputs&#x60; | 
**BoxesToSpend** | [**[]ErgoTransactionOutput**](ErgoTransactionOutput.md) | boxes, that corresponds to id&#39;s of &#x60;spendingTransaction.inputs&#x60; | 
**SpendingTransaction** | [**ErgoLikeTransaction**](.md) |  | 
**SelfIndex** | **int64** | index of the box in &#x60;boxesToSpend&#x60; that contains the script we&#39;re evaluating | 
**Extension** | [**map[string]interface{}**](.md) | prover-defined key-value pairs, that may be used inside a script | 
**ValidationSettings** | **string** | validation parameters passed to Interpreter.verify to detect soft-fork conditions | 
**CostLimit** | **int64** | hard limit on accumulated execution cost, if exceeded lead to CostLimitException to be thrown | 
**InitCost** | **int64** | initial value of execution cost already accumulated before Interpreter.verify is called | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


