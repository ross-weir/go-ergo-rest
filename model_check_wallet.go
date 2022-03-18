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

// CheckWallet struct for CheckWallet
type CheckWallet struct {
	// Mnemonic seed (optional)
	Mnemonic string `json:"mnemonic"`
	// Optional pass to password-protect mnemonic seed
	MnemonicPass string `json:"mnemonicPass,omitempty"`
}
