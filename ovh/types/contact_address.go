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

// ContactAddress Representation of an Address
type ContactAddress struct {

	// City City
	City string `json:"city,omitempty"`

	// Country Country
	Country string `json:"country,omitempty"`

	// Line1 First line of the address
	Line1 string `json:"line1,omitempty"`

	// Line2 Second line of the address
	Line2 string `json:"line2,omitempty"`

	// Line3 Third line of the address
	Line3 string `json:"line3,omitempty"`

	// OtherDetails Others details
	OtherDetails string `json:"otherDetails,omitempty"`

	// Province Province name
	Province string `json:"province,omitempty"`

	// Zip Zipcode
	Zip string `json:"zip,omitempty"`
}
