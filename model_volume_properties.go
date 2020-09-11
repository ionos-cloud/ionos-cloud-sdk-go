/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ionossdk
// VolumeProperties struct for VolumeProperties
type VolumeProperties struct {
	// A name of that resource
	Name *string `json:"name,omitempty"` /* nullable: false */
	// Hardware type of the volume.
	Type *string `json:"type,omitempty"` /* nullable: false */
	// The size of the volume in GB
	Size *float32 `json:"size,omitempty"` /* nullable: false */
	// The availability zone in which the volume should exist. The storage volume will be provisioned on as less physical storages as possible but cannot guarantee upfront
	AvailabilityZone *string `json:"availabilityZone,omitempty"` /* nullable: false */
	// Image or snapshot ID to be used as template for this volume
	Image *string `json:"image,omitempty"` /* nullable: false */
	ImageAlias *string `json:"imageAlias,omitempty"` /* nullable: false */
	// Initial password to be set for installed OS. Works with public images only. Not modifiable, forbidden in update requests. Password rules allows all characters from a-z, A-Z, 0-9
	ImagePassword *string `json:"imagePassword,omitempty"` /* nullable: false */
	// Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key. This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation.
	SshKeys *[]string `json:"sshKeys,omitempty"` /* nullable: false */
	// The bus type of the volume. Default is VIRTIO
	Bus *string `json:"bus,omitempty"` /* nullable: false */
	// OS type of this volume
	LicenceType *string `json:"licenceType,omitempty"` /* nullable: false */
	// Is capable of CPU hot plug (no reboot required)
	CpuHotPlug *bool `json:"cpuHotPlug,omitempty"` /* nullable: false */
	// Is capable of memory hot plug (no reboot required)
	RamHotPlug *bool `json:"ramHotPlug,omitempty"` /* nullable: false */
	// Is capable of nic hot plug (no reboot required)
	NicHotPlug *bool `json:"nicHotPlug,omitempty"` /* nullable: false */
	// Is capable of nic hot unplug (no reboot required)
	NicHotUnplug *bool `json:"nicHotUnplug,omitempty"` /* nullable: false */
	// Is capable of Virt-IO drive hot plug (no reboot required)
	DiscVirtioHotPlug *bool `json:"discVirtioHotPlug,omitempty"` /* nullable: false */
	// Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.
	DiscVirtioHotUnplug *bool `json:"discVirtioHotUnplug,omitempty"` /* nullable: false */
	// The LUN ID of the storage volume. Null for volumes not mounted to any VM
	DeviceNumber *int64 `json:"deviceNumber,omitempty"` /* nullable: false */
	// The uuid of the Backup Unit that user has access to. The property is immutable and is only allowed to be set on a new volume creation. It is mandatory to provied either public image or imageAlias in conjunction with this property.
	BackupunitId *string `json:"backupunitId,omitempty"` /* nullable: false */
}
