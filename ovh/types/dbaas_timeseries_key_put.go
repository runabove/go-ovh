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

// DBaasTimeseriesKeyPut ...
type DBaasTimeseriesKeyPut struct {
	Description string `json:"description,omitempty"`

	Permissions []string `json:"permissions,omitempty"`

	Tags []*PaasTimeseriesTag `json:"tags,omitempty"`
}
