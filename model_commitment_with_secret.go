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

// CommitmentWithSecret commitment to secret along with secret (!) randomness
type CommitmentWithSecret struct {
	Hint     string       `json:"hint"`
	Pubkey   SigmaBoolean `json:"pubkey"`
	Position string       `json:"position"`
	Type     string       `json:"type,omitempty"`
	// a group element of the commitment
	A string `json:"a"`
	// b group element of the commitment (needed for DHT protocol only)
	B string `json:"b,omitempty"`
}
