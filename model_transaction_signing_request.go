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

// TransactionSigningRequest A request to sign a transaction with secrets provided
type TransactionSigningRequest struct {
	Tx UnsignedErgoTransaction `json:"tx"`
	// Optional list of inputs to be used in serialized form
	InputsRaw []string `json:"inputsRaw,omitempty"`
	// Optional list of inputs to be used in serialized form
	DataInputsRaw []string                         `json:"dataInputsRaw,omitempty"`
	Hints         TransactionHintsBag              `json:"hints,omitempty"`
	Secrets       TransactionSigningRequestSecrets `json:"secrets"`
}