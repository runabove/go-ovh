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

// DBaaSLogsOffer DBaaS Logs offer
type DBaaSLogsOffer struct {

	// CurNbAlias Current number of alias booked
	CurNbAlias int64 `json:"curNbAlias,omitempty"`

	// CurNbDashboard Current number of dashboard booked
	CurNbDashboard int64 `json:"curNbDashboard,omitempty"`

	// CurNbIndex Current number of index booked
	CurNbIndex int64 `json:"curNbIndex,omitempty"`

	// CurNbInput Current number of input booked
	CurNbInput int64 `json:"curNbInput,omitempty"`

	// CurNbRole Current number of role booked
	CurNbRole int64 `json:"curNbRole,omitempty"`

	// CurNbStream Current number of stream booked
	CurNbStream int64 `json:"curNbStream,omitempty"`

	// EsStorage Number of GB stored per month included
	EsStorage int64 `json:"esStorage,omitempty"`

	// MaxNbAlias Maximum number of alias allowed
	MaxNbAlias int64 `json:"maxNbAlias,omitempty"`

	// MaxNbDashboard Maximum number of dashboard allowed
	MaxNbDashboard int64 `json:"maxNbDashboard,omitempty"`

	// MaxNbIndex Maximum number of index allowed
	MaxNbIndex int64 `json:"maxNbIndex,omitempty"`

	// MaxNbInput Maximum number of input allowed
	MaxNbInput int64 `json:"maxNbInput,omitempty"`

	// MaxNbRole Maximum number of role allowed
	MaxNbRole int64 `json:"maxNbRole,omitempty"`

	// MaxNbStream Maximum number of stream allowed
	MaxNbStream int64 `json:"maxNbStream,omitempty"`

	// Reference Option unique reference
	Reference string `json:"reference,omitempty"`

	// Retention Data retention in hours
	Retention int64 `json:"retention,omitempty"`
}