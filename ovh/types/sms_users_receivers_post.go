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

// SmsUsersReceiversPost ...
type SmsUsersReceiversPost struct {
	AutoUpdate bool `json:"autoUpdate,omitempty"`

	CsvURL string `json:"csvUrl,omitempty"`

	Description string `json:"description,omitempty"`

	SlotID int64 `json:"slotId,omitempty"`
}