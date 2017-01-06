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

// TelephonyEasyMiniPabxHuntingAgent Easy/Mini PABX agent
type TelephonyEasyMiniPabxHuntingAgent struct {

	// AgentNumber The phone number of the agent
	AgentNumber string `json:"agentNumber,omitempty"`

	// Logged True if the agent is logged
	Logged bool `json:"logged,omitempty"`

	// NoReplyTimer The maxium ringing time
	NoReplyTimer int64 `json:"noReplyTimer,omitempty"`

	// Position The position in the hunting
	Position int64 `json:"position,omitempty"`
}