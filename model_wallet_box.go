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

// WalletBox struct for WalletBox
type WalletBox struct {
	Box ErgoTransactionOutput `json:"box"`
	// Number of confirmations, if the box is included into the blockchain
	ConfirmationsNum *int32 `json:"confirmationsNum"`
	// Encoded Ergo Address
	Address string `json:"address"`
	// Base16-encoded 32 byte modifier id
	CreationTransaction string `json:"creationTransaction"`
	// Base16-encoded 32 byte modifier id
	SpendingTransaction string `json:"spendingTransaction"`
	// The height the box was spent at
	SpendingHeight *int32 `json:"spendingHeight"`
	// The height the transaction containing the box was included in a block at
	InclusionHeight int32 `json:"inclusionHeight"`
	// A flag signalling whether the box is created on main chain
	Onchain bool `json:"onchain"`
	// A flag signalling whether the box was spent
	Spent bool `json:"spent"`
	// An index of a box in the creating transaction
	CreationOutIndex int32 `json:"creationOutIndex"`
	// Scan identifiers the box relates to
	Scans []int32 `json:"scans"`
}
