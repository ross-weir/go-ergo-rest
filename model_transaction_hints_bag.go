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

// TransactionHintsBag prover hints extracted from a transaction
type TransactionHintsBag struct {
	// Hints which contain secrets, do not share them!
	SecretHints []map[string][]interface{} `json:"secretHints,omitempty"`
	// Hints which contain public data only, share them freely!
	PublicHints []map[string][]interface{} `json:"publicHints,omitempty"`
}
