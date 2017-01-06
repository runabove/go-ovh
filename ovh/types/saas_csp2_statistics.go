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

// SaasCsp2Statistics License usage statistics.
type SaasCsp2Statistics struct {

	// Date Date of the statistics.
	Date *time.Time `json:"date,omitempty"`

	// Lines List of lines associated to this statistics entity.
	Lines []*SaasCsp2StatisticsLine `json:"lines,omitempty"`
}