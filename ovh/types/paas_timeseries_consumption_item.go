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

// PaasTimeseriesConsumptionItem ConsumptionItem
type PaasTimeseriesConsumptionItem struct {
	MetricName string `json:"metricName,omitempty"`

	Price *OrderPrice `json:"price,omitempty"`

	Quantity *PaasTimeseriesConsumptionItemQuantity `json:"quantity,omitempty"`

	UnitPrice *OrderPrice `json:"unitPrice,omitempty"`
}
