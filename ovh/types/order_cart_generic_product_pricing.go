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

// OrderCartGenericProductPricing Representation of a product pricing
type OrderCartGenericProductPricing struct {

	// Capacities Capacities of the pricing (type of pricing)
	Capacities []string `json:"capacities,omitempty"`

	// Description Description of the pricing
	Description string `json:"description,omitempty"`

	// Duration Duration for ordering the product
	Duration string `json:"duration,omitempty"`

	// Interval Interval of renewal
	Interval int64 `json:"interval,omitempty"`

	// MaximumQuantity Maximum quantity that can be ordered
	MaximumQuantity int64 `json:"maximumQuantity,omitempty"`

	// MaximumRepeat Maximum repeat for renewal
	MaximumRepeat int64 `json:"maximumRepeat,omitempty"`

	// MinimumQuantity Minimum quantity that can be ordered
	MinimumQuantity int64 `json:"minimumQuantity,omitempty"`

	// MinimumRepeat Minimum repeat for renewal
	MinimumRepeat int64 `json:"minimumRepeat,omitempty"`

	Price *OrderPrice `json:"price,omitempty"`

	// PriceInUcents Price of the product in micro-centims
	PriceInUcents int64 `json:"priceInUcents,omitempty"`

	// PricingMode Pricing model identifier
	PricingMode string `json:"pricingMode,omitempty"`

	// PricingType Pricing type
	PricingType string `json:"pricingType,omitempty"`
}