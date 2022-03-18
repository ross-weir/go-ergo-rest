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

// DhtSecret Hex-encoded big-endian 256-bits secret exponent \"w\" along with generators \"g\", \"h\", and group elements \"u\", \"v\", such as g^w = u, h^w = v
type DhtSecret struct {
	// Hex-encoded big-endian 256-bits secret exponent
	Secret string `json:"secret"`
	// Hex-encoded \"g\" generator for the Diffie-Hellman tuple (secp256k1 curve point)
	G string `json:"g"`
	// Hex-encoded \"h\" generator for the Diffie-Hellman tuple (secp256k1 curve point)
	H string `json:"h"`
	// Hex-encoded \"u\" group element of the Diffie-Hellman tuple (secp256k1 curve point)
	U string `json:"u"`
	// Hex-encoded \"v\" group element of the Diffie-Hellman tuple (secp256k1 curve point)
	V string `json:"v"`
}
