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

// FeeHistogramBin Fee histogram bin
type FeeHistogramBin struct {
	NTxns    int32 `json:"nTxns,omitempty"`
	TotalFee int64 `json:"totalFee,omitempty"`
}
