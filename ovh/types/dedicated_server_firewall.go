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

// DedicatedServerFirewall Firewall attached to this server
type DedicatedServerFirewall struct {
	Enabled bool `json:"enabled,omitempty"`

	Firewall string `json:"firewall,omitempty"`

	// IP Firewall management IP
	IP string `json:"ip,omitempty"`

	// Mode transparent mode : device is invisible on the network; routed mode : the security appliance is considered to be a router hop in the network
	Mode string `json:"mode,omitempty"`

	// Model Firewall model
	Model string `json:"model,omitempty"`
}