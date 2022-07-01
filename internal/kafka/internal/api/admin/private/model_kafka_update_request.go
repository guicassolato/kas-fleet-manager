/*
 * Kafka Service Fleet Manager Admin APIs
 *
 * The admin APIs for the fleet manager of Kafka service
 *
 * API version: 0.1.0
 * Contact: rhosak-support@redhat.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// KafkaUpdateRequest struct for KafkaUpdateRequest
type KafkaUpdateRequest struct {
	StrimziVersion  string `json:"strimzi_version,omitempty"`
	KafkaVersion    string `json:"kafka_version,omitempty"`
	KafkaIbpVersion string `json:"kafka_ibp_version,omitempty"`
	// Maximum data storage available to this Kafka. This is now deprecated, please use max_data_retention_size instead
	// Deprecated
	DeprecatedKafkaStorageSize string `json:"kafka_storage_size,omitempty"`
	// Maximum data storage available to this Kafka
	MaxDataRetentionSize string `json:"max_data_retention_size,omitempty"`
}
