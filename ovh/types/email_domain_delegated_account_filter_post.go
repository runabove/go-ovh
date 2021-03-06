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

// EmailDomainDelegatedAccountFilterPost ...
type EmailDomainDelegatedAccountFilterPost struct {
	Action string `json:"action,omitempty"`

	ActionParam string `json:"actionParam,omitempty"`

	Active bool `json:"active,omitempty"`

	Header string `json:"header,omitempty"`

	Name string `json:"name,omitempty"`

	Operand string `json:"operand,omitempty"`

	Priority int64 `json:"priority,omitempty"`

	Value string `json:"value,omitempty"`
}
