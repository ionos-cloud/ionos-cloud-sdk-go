/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ionossdk
// SnapshotProperties struct for SnapshotProperties
type SnapshotProperties struct {
	// A name of that resource
	Name *string `json:"name,omitempty"` /* nullable: false */
	// Human readable description
	Description *string `json:"description,omitempty"` /* nullable: false */
	// Location of that image/snapshot. 
	Location *string `json:"location,omitempty"` /* nullable: false */
	// The size of the image in GB
	Size *float32 `json:"size,omitempty"` /* nullable: false */
	// Boolean value representing if the snapshot requires extra protection e.g. two factor protection
	SecAuthProtection *bool `json:"secAuthProtection,omitempty"` /* nullable: false */
	// Is capable of CPU hot plug (no reboot required)
	CpuHotPlug *bool `json:"cpuHotPlug,omitempty"` /* nullable: false */
	// Is capable of CPU hot unplug (no reboot required)
	CpuHotUnplug *bool `json:"cpuHotUnplug,omitempty"` /* nullable: false */
	// Is capable of memory hot plug (no reboot required)
	RamHotPlug *bool `json:"ramHotPlug,omitempty"` /* nullable: false */
	// Is capable of memory hot unplug (no reboot required)
	RamHotUnplug *bool `json:"ramHotUnplug,omitempty"` /* nullable: false */
	// Is capable of nic hot plug (no reboot required)
	NicHotPlug *bool `json:"nicHotPlug,omitempty"` /* nullable: false */
	// Is capable of nic hot unplug (no reboot required)
	NicHotUnplug *bool `json:"nicHotUnplug,omitempty"` /* nullable: false */
	// Is capable of Virt-IO drive hot plug (no reboot required)
	DiscVirtioHotPlug *bool `json:"discVirtioHotPlug,omitempty"` /* nullable: false */
	// Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.
	DiscVirtioHotUnplug *bool `json:"discVirtioHotUnplug,omitempty"` /* nullable: false */
	// Is capable of SCSI drive hot plug (no reboot required)
	DiscScsiHotPlug *bool `json:"discScsiHotPlug,omitempty"` /* nullable: false */
	// Is capable of SCSI drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.
	DiscScsiHotUnplug *bool `json:"discScsiHotUnplug,omitempty"` /* nullable: false */
	// OS type of this Snapshot
	LicenceType *string `json:"licenceType,omitempty"` /* nullable: false */
}
