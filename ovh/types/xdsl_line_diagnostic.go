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

// XdslLineDiagnostic Detailed line tests
type XdslLineDiagnostic struct {
	LineTest string `json:"lineTest,omitempty"`

	LineTestTime *time.Time `json:"lineTestTime,omitempty"`

	Number string `json:"number,omitempty"`

	ProposedProfileID int64 `json:"proposedProfileId,omitempty"`

	Sync bool `json:"sync,omitempty"`
}