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

// DeriveKey struct for DeriveKey
type DeriveKey struct {
	// Derivation path for a new secret to derive
	DerivationPath string `json:"derivationPath"`
}
