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

// Scan struct for Scan
type Scan struct {
	ScanName          string            `json:"scanName,omitempty"`
	ScanID            int32             `json:"scanID,omitempty"`
	WalletInteraction string            `json:"walletInteraction,omitempty"`
	RemoveOffchain    bool              `json:"removeOffchain,omitempty"`
	TrackingRule      ScanningPredicate `json:"trackingRule,omitempty"`
}
