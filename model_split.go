/*
 * Finnhub API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finnhub
// Split struct for Split
type Split struct {
	// Symbol.
	Symbol string `json:"symbol,omitempty"`
	// Split date.
	Date string `json:"Date,omitempty"`
	// From factor.
	FromFactor float32 `json:"fromFactor,omitempty"`
	// To factor.
	ToFactor float32 `json:"toFactor,omitempty"`
}