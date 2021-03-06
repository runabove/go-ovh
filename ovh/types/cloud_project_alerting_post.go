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

// CloudProjectAlertingPost ...
type CloudProjectAlertingPost struct {
	Delay int64 `json:"delay,omitempty"`

	Email string `json:"email,omitempty"`

	MonthlyThreshold int64 `json:"monthlyThreshold,omitempty"`
}
