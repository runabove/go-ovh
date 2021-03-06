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

// MePaymentMeanBankAccountPost ...
type MePaymentMeanBankAccountPost struct {
	Bic string `json:"bic,omitempty"`

	Description string `json:"description,omitempty"`

	Iban string `json:"iban,omitempty"`

	OwnerAddress string `json:"ownerAddress,omitempty"`

	OwnerName string `json:"ownerName,omitempty"`
}
