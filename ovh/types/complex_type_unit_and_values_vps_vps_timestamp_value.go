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

// ComplexTypeUnitAndValuesVpsTimestampValue A value set tagged with its unit
type ComplexTypeUnitAndValuesVpsTimestampValue struct {
	Unit string `json:"unit,omitempty"`

	Values []*VpsTimestampValue `json:"values,omitempty"`
}
