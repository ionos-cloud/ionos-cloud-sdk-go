/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ionossdk
// Peer struct for Peer
type Peer struct {
	Id *string `json:"id,omitempty"` /* nullable: false */
	Name *string `json:"name,omitempty"` /* nullable: false */
	DatacenterId *string `json:"datacenterId,omitempty"` /* nullable: false */
	DatacenterName *string `json:"datacenterName,omitempty"` /* nullable: false */
	Location *string `json:"location,omitempty"` /* nullable: false */
}
