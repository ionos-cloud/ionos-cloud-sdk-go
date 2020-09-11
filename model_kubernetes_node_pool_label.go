/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ionossdk
// KubernetesNodePoolLabel map of labels attached to node pool
type KubernetesNodePoolLabel struct {
	// Key of the label. String part must consist of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character.
	Key *string `json:"key,omitempty"` /* nullable: false */
	// Value of the label. String part must consist of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character.
	Value *string `json:"value,omitempty"` /* nullable: false */
}
