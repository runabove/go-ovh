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

// DomainZoneDynHostLoginPost ...
type DomainZoneDynHostLoginPost struct {
	LoginSuffix string `json:"loginSuffix,omitempty"`

	Password string `json:"password,omitempty"`

	SubDomain string `json:"subDomain,omitempty"`
}
