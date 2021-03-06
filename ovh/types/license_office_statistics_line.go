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

// LicenseOfficeStatisticsLine License usage statistics line.
type LicenseOfficeStatisticsLine struct {

	// EndOfDayCount Count of activated licenses at the end of the day.
	EndOfDayCount int64 `json:"endOfDayCount,omitempty"`

	// LicenceType Type of the Office license.
	LicenceType string `json:"licenceType,omitempty"`

	// PeakCount Maximum count of simultaneous activated licences.
	PeakCount int64 `json:"peakCount,omitempty"`
}
