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

import (
	"time"
)

// PaasTimeseriesConsumption Consumption
type PaasTimeseriesConsumption struct {

	// From Consumption start date
	From *time.Time `json:"from,omitempty"`

	// Generated Timestamp of consumption generation
	Generated *time.Time `json:"generated,omitempty"`

	// Items List of consumption items
	Items []*PaasTimeseriesConsumptionItem `json:"items,omitempty"`

	// To Consumption end date
	To *time.Time `json:"to,omitempty"`

	Total *OrderPrice `json:"total,omitempty"`
}
