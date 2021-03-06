/*
 * Finnhub API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package finnhub
// NewsSentiment struct for NewsSentiment
type NewsSentiment struct {
	Buzz CompanyNewsStatistics `json:"buzz,omitempty"`
	// News score.
	CompanyNewsScore float32 `json:"companyNewsScore,omitempty"`
	// Sector average bullish percent.
	SectorAverageBullishPercent float32 `json:"sectorAverageBullishPercent,omitempty"`
	// Sectore average score.
	SectorAverageNewsScore float32 `json:"sectorAverageNewsScore,omitempty"`
	Sentiment Sentiment `json:"sentiment,omitempty"`
	// Requested symbol.
	Symbol string `json:"symbol,omitempty"`
}
