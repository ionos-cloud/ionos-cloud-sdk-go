/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ionossdk
// GroupProperties struct for GroupProperties
type GroupProperties struct {
	// A name of that resource
	Name *string `json:"name,omitempty"` /* nullable: false */
	// create data center privilege
	CreateDataCenter *bool `json:"createDataCenter,omitempty"` /* nullable: false */
	// create snapshot privilege
	CreateSnapshot *bool `json:"createSnapshot,omitempty"` /* nullable: false */
	// reserve ip block privilege
	ReserveIp *bool `json:"reserveIp,omitempty"` /* nullable: false */
	// activity log access privilege
	AccessActivityLog *bool `json:"accessActivityLog,omitempty"` /* nullable: false */
	// create pcc privilege
	CreatePcc *bool `json:"createPcc,omitempty"` /* nullable: false */
	// S3 privilege
	S3Privilege *bool `json:"s3Privilege,omitempty"` /* nullable: false */
	// create backup unit privilege
	CreateBackupUnit *bool `json:"createBackupUnit,omitempty"` /* nullable: false */
	// create internet access privilege
	CreateInternetAccess *bool `json:"createInternetAccess,omitempty"` /* nullable: false */
	// create kubernetes cluster privilege
	CreateK8sCluster *bool `json:"createK8sCluster,omitempty"` /* nullable: false */
}
