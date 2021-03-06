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

// VpsIP Information about an IP address for a VPS Virtual Machine
type VpsIP struct {
	Gateway string `json:"gateway,omitempty"`

	Geolocation string `json:"geolocation,omitempty"`

	// IPAddress The effective ip address of the Ip object
	IPAddress string `json:"ipAddress,omitempty"`

	MacAddress string `json:"macAddress,omitempty"`

	Reverse string `json:"reverse,omitempty"`

	TType string `json:"type,omitempty"`

	Version string `json:"version,omitempty"`
}
