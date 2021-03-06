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

// InlineResponseDefault ...
type InlineResponseDefault struct {
	ErrorCode int32 `json:"errorCode,omitempty"`

	HTTPCode int32 `json:"httpCode,omitempty"`

	Message string `json:"message,omitempty"`
}
