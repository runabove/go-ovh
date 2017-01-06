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

// HostingPrivateDatabaseDatabaseImportPost ...
type HostingPrivateDatabaseDatabaseImportPost struct {
	DocumentID string `json:"documentId,omitempty"`

	FlushDatabase bool `json:"flushDatabase,omitempty"`

	SendEmail bool `json:"sendEmail,omitempty"`
}