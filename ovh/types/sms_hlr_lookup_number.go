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

import (
	"time"
)

// SmsHlrLookupNumber Home Location Register informations. Give informations about a given cellular phone.
type SmsHlrLookupNumber struct {

	// Datetime HLR creation datetime
	Datetime *time.Time `json:"datetime,omitempty"`

	// ID HLR id
	ID int64 `json:"id,omitempty"`

	// Msisdn MSISDN
	Msisdn string `json:"msisdn,omitempty"`

	// OperatorCode The {Mobile Country Code, Mobile Network Code} unique identifier
	OperatorCode string `json:"operatorCode,omitempty"`

	// Ported Has the MSISDN been ported from its original network
	Ported bool `json:"ported,omitempty"`

	// Reachable Is the MSISDN currently reachable
	Reachable bool `json:"reachable,omitempty"`

	// Roaming Is the MSISDN currently roaming outside its natinal network
	Roaming bool `json:"roaming,omitempty"`

	// Status Status of the HLR request
	Status string `json:"status,omitempty"`

	// Valid Is the MSISDN valid
	Valid bool `json:"valid,omitempty"`
}