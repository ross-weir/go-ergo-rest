/*
 * Ergo Node API
 *
 * API docs for Ergo Node. Models are shared between all Ergo products
 *
 * API version: 4.0.23
 * Contact: ergoplatform@protonmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ergo_rest

// ErgoTransactionInput struct for ErgoTransactionInput
type ErgoTransactionInput struct {
	// Base16-encoded transaction box id bytes. Should be 32 bytes long
	BoxID         string        `json:"boxID"`
	SpendingProof SpendingProof `json:"spendingProof"`
}