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

// CloudBillingViewHourlyStorage HourlyStorage
type CloudBillingViewHourlyStorage struct {
	IncomingBandwidth *CloudBillingViewBandwidthStorage `json:"incomingBandwidth,omitempty"`

	OutgoingBandwidth *CloudBillingViewBandwidthStorage `json:"outgoingBandwidth,omitempty"`

	// Region Region
	Region string `json:"region,omitempty"`

	Stored *CloudBillingViewStoredStorage `json:"stored,omitempty"`

	// TotalPrice Total price
	TotalPrice float64 `json:"totalPrice,omitempty"`

	// TType Storage type
	TType string `json:"type,omitempty"`
}