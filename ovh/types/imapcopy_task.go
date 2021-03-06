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

// ImapcopyTask Task of ImapCopy
type ImapcopyTask struct {

	// FinishDate Finished date of task
	FinishDate *time.Time `json:"finishDate,omitempty"`

	// ID Id of task
	ID int64 `json:"id,omitempty"`

	// LastUpdate Last update of task
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	// TReturn Return message of task
	TReturn string `json:"return,omitempty"`

	// Status Status of task
	Status string `json:"status,omitempty"`

	// TodoDate Creation date of task
	TodoDate *time.Time `json:"todoDate,omitempty"`
}
