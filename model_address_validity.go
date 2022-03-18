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

// AddressValidity Validity status of Ergo address
type AddressValidity struct {
	// Encoded Ergo Address
	Address string `json:"address"`
	IsValid bool   `json:"isValid"`
	Error   string `json:"error,omitempty"`
}
