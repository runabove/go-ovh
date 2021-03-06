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

// DedicatedServerBackupFtp Backup Ftp assigned to this server
type DedicatedServerBackupFtp struct {

	// FtpBackupName The backup FTP server name
	FtpBackupName string `json:"ftpBackupName,omitempty"`

	Quota *DedicatedServerBackupFtpQuota `json:"quota,omitempty"`

	// ReadOnlyDate If not-null, gives the date since when your account was set in read-only mode because the quota was exceeded
	ReadOnlyDate *time.Time `json:"readOnlyDate,omitempty"`

	// TType The backup FTP type
	TType string `json:"type,omitempty"`

	Usage *DedicatedServerBackupFtpUsage `json:"usage,omitempty"`
}
