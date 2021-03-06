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

// XdslDslamLineProfile Profile on the DSLAM
type XdslDslamLineProfile struct {
	ID int64 `json:"id,omitempty"`

	IsCurrent bool `json:"isCurrent,omitempty"`

	Name string `json:"name,omitempty"`
}
