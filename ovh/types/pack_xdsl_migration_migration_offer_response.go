/*
 * OVH API - EU
 *
 * Build your own OVH world.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-subscribe@ml.ovh.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package types

// PackXdslMigrationOfferResponse List of available Migration offer
type PackXdslMigrationOfferResponse struct {

	// Offers Array of offers
	Offers []*PackXdslMigrationOffer `json:"offers,omitempty"`
}
