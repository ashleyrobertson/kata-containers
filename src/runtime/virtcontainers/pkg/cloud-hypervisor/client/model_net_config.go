/*
 * Cloud Hypervisor API
 *
 * Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// NetConfig struct for NetConfig
type NetConfig struct {
	Tap               string            `json:"tap,omitempty"`
	Ip                string            `json:"ip,omitempty"`
	Mask              string            `json:"mask,omitempty"`
	Mac               string            `json:"mac,omitempty"`
	Iommu             bool              `json:"iommu,omitempty"`
	NumQueues         int32             `json:"num_queues,omitempty"`
	QueueSize         int32             `json:"queue_size,omitempty"`
	VhostUser         bool              `json:"vhost_user,omitempty"`
	VhostSocket       string            `json:"vhost_socket,omitempty"`
	VhostMode         string            `json:"vhost_mode,omitempty"`
	Id                string            `json:"id,omitempty"`
	Fd                []int32           `json:"fd,omitempty"`
	RateLimiterConfig RateLimiterConfig `json:"rate_limiter_config,omitempty"`
}
