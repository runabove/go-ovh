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

// MsServicesAccount Active Directory Account
type MsServicesAccount struct {

	// SAMAccountName SAM account name
	SAMAccountName string `json:"SAMAccountName,omitempty"`

	// CreationDate Creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// DisplayName Account display name
	DisplayName string `json:"displayName,omitempty"`

	// FirstName Account first name
	FirstName string `json:"firstName,omitempty"`

	// GUID GUID for user in active directory
	GUID string `json:"guid,omitempty"`

	// ID Account id
	ID int64 `json:"id,omitempty"`

	// Initials Account initials
	Initials string `json:"initials,omitempty"`

	// LastLogoffDate Last logoff
	LastLogoffDate *time.Time `json:"lastLogoffDate,omitempty"`

	// LastLogonDate Last logon
	LastLogonDate *time.Time `json:"lastLogonDate,omitempty"`

	// LastName Account last name
	LastName string `json:"lastName,omitempty"`

	// LastUpdateDate Last update
	LastUpdateDate *time.Time `json:"lastUpdateDate,omitempty"`

	// PasswordLastUpdate Time of account's password last update
	PasswordLastUpdate *time.Time `json:"passwordLastUpdate,omitempty"`

	// State Account state
	State string `json:"state,omitempty"`

	// TaskPendingID Pending task for this account
	TaskPendingID int64 `json:"taskPendingId,omitempty"`

	// UserPrincipalName User Principal Name
	UserPrincipalName string `json:"userPrincipalName,omitempty"`
}
