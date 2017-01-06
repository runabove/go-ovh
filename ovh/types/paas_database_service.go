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

// PaasDatabaseService PAAS Project
type PaasDatabaseService struct {

	// CreationDate Project creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// LastUpdate Project last update date
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`

	// Name Project custom name
	Name string `json:"name,omitempty"`

	// ProjectID Project id
	ProjectID string `json:"projectId,omitempty"`

	// Status Project status
	Status string `json:"status,omitempty"`

	// TType Project type
	TType string `json:"type,omitempty"`
}