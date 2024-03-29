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

// HintExtractionRequest request to extract prover hints from a transaction
type HintExtractionRequest struct {
	Tx ErgoTransaction `json:"tx"`
	// Real signers of the transaction
	Real []SigmaBoolean `json:"real"`
	// Simulated signers of the transaction
	Simulated []SigmaBoolean `json:"simulated"`
	// Optional list of inputs to be used in serialized form
	InputsRaw []string `json:"inputsRaw,omitempty"`
	// Optional list of inputs to be used in serialized form
	DataInputsRaw []string `json:"dataInputsRaw,omitempty"`
}
