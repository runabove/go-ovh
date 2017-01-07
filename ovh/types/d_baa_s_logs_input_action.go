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

// DBaasLogsInputAction Action on input
type DBaasLogsInputAction struct {

	// IsAllowed Indicates if action is allowed
	IsAllowed bool `json:"isAllowed,omitempty"`

	// TType Action type
	TType string `json:"type,omitempty"`
}
