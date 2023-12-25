/*
 * Nexus-MP Sidecar Agent API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type HealthCheckPayload struct {

	ServerId string `json:"serverId,omitempty"`

	HealthStatus int32 `json:"healthStatus,omitempty"`
}

// AssertHealthCheckPayloadRequired checks if the required fields are not zero-ed
func AssertHealthCheckPayloadRequired(obj HealthCheckPayload) error {
	return nil
}

// AssertHealthCheckPayloadConstraints checks if the values respects the defined constraints
func AssertHealthCheckPayloadConstraints(obj HealthCheckPayload) error {
	return nil
}