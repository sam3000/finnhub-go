/*
 * Finnhub API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finnhub
// IndexHistoricalConstituent struct for IndexHistoricalConstituent
type IndexHistoricalConstituent struct {
	// Symbol
	Symbol string `json:"symbol,omitempty"`
	// <code>add</code> or <code>remove</code>.
	Action string `json:"action,omitempty"`
	// Date of joining or leaving the index.
	Date string `json:"date,omitempty"`
}
