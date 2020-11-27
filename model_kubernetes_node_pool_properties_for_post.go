/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionossdk

import (
	"encoding/json"
)

// KubernetesNodePoolPropertiesForPost struct for KubernetesNodePoolPropertiesForPost
type KubernetesNodePoolPropertiesForPost struct {
	// A Kubernetes Node Pool Name. Valid Kubernetes Node Pool name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Name *string `json:"name"`
	// A valid uuid of the datacenter on which user has access
	DatacenterId *string `json:"datacenterId"`
	// Number of nodes part of the Node Pool
	NodeCount *int32 `json:"nodeCount"`
	// A valid cpu family name
	CpuFamily *string `json:"cpuFamily"`
	// Number of cores for node
	CoresCount *int32 `json:"coresCount"`
	// RAM size for node, minimum size 2048MB is recommended. Ram size must be set to multiple of 1024MB.
	RamSize *int32 `json:"ramSize"`
	// The availability zone in which the server should exist
	AvailabilityZone *string `json:"availabilityZone"`
	// Hardware type of the volume
	StorageType *string `json:"storageType"`
	// The size of the volume in GB. The size should be greater than 10GB.
	StorageSize *int32 `json:"storageSize"`
	// The kubernetes version in which a nodepool is running. This imposes restrictions on what kubernetes versions can be run in a cluster's nodepools. Additionally, not all kubernetes versions are viable upgrade targets for all prior versions.
	K8sVersion *string `json:"k8sVersion,omitempty"`
	MaintenanceWindow *KubernetesMaintenanceWindow `json:"maintenanceWindow,omitempty"`
	AutoScaling *KubernetesAutoScaling `json:"autoScaling,omitempty"`
	// array of additional LANs attached to worker nodes
	Lans *[]KubernetesNodePoolLan `json:"lans,omitempty"`
	Labels *KubernetesNodePoolLabel `json:"labels,omitempty"`
	Annotations *KubernetesNodePoolAnnotation `json:"annotations,omitempty"`
	// Optional array of reserved public IP addresses to be used by the nodes. IPs must be from same location as the data center used for the node pool. The array must contain one extra IP than maximum number of nodes could be. (nodeCount+1 if fixed node amount or maxNodeCount+1 if auto scaling is used) The extra provided IP Will be used during rebuilding of nodes.
	PublicIps *[]string `json:"publicIps,omitempty"`
}



// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetName(v string) {
	o.Name = &v
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}



// GetDatacenterId returns the DatacenterId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetDatacenterId() *string {
	if o == nil {
		return nil
	}

	return o.DatacenterId
}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetDatacenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatacenterId, true
}

// SetDatacenterId sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetDatacenterId(v string) {
	o.DatacenterId = &v
}

// HasDatacenterId returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasDatacenterId() bool {
	if o != nil && o.DatacenterId != nil {
		return true
	}

	return false
}



// GetNodeCount returns the NodeCount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetNodeCount() *int32 {
	if o == nil {
		return nil
	}

	return o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetNodeCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeCount, true
}

// SetNodeCount sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetNodeCount(v int32) {
	o.NodeCount = &v
}

// HasNodeCount returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasNodeCount() bool {
	if o != nil && o.NodeCount != nil {
		return true
	}

	return false
}



// GetCpuFamily returns the CpuFamily field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetCpuFamily() *string {
	if o == nil {
		return nil
	}

	return o.CpuFamily
}

// GetCpuFamilyOk returns a tuple with the CpuFamily field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetCpuFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuFamily, true
}

// SetCpuFamily sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetCpuFamily(v string) {
	o.CpuFamily = &v
}

// HasCpuFamily returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasCpuFamily() bool {
	if o != nil && o.CpuFamily != nil {
		return true
	}

	return false
}



// GetCoresCount returns the CoresCount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetCoresCount() *int32 {
	if o == nil {
		return nil
	}

	return o.CoresCount
}

// GetCoresCountOk returns a tuple with the CoresCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetCoresCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoresCount, true
}

// SetCoresCount sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetCoresCount(v int32) {
	o.CoresCount = &v
}

// HasCoresCount returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasCoresCount() bool {
	if o != nil && o.CoresCount != nil {
		return true
	}

	return false
}



// GetRamSize returns the RamSize field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetRamSize() *int32 {
	if o == nil {
		return nil
	}

	return o.RamSize
}

// GetRamSizeOk returns a tuple with the RamSize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetRamSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RamSize, true
}

// SetRamSize sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetRamSize(v int32) {
	o.RamSize = &v
}

// HasRamSize returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasRamSize() bool {
	if o != nil && o.RamSize != nil {
		return true
	}

	return false
}



// GetAvailabilityZone returns the AvailabilityZone field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetAvailabilityZone() *string {
	if o == nil {
		return nil
	}

	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetAvailabilityZone(v string) {
	o.AvailabilityZone = &v
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasAvailabilityZone() bool {
	if o != nil && o.AvailabilityZone != nil {
		return true
	}

	return false
}



// GetStorageType returns the StorageType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetStorageType() *string {
	if o == nil {
		return nil
	}

	return o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetStorageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageType, true
}

// SetStorageType sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetStorageType(v string) {
	o.StorageType = &v
}

// HasStorageType returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasStorageType() bool {
	if o != nil && o.StorageType != nil {
		return true
	}

	return false
}



// GetStorageSize returns the StorageSize field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetStorageSize() *int32 {
	if o == nil {
		return nil
	}

	return o.StorageSize
}

// GetStorageSizeOk returns a tuple with the StorageSize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetStorageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageSize, true
}

// SetStorageSize sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetStorageSize(v int32) {
	o.StorageSize = &v
}

// HasStorageSize returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasStorageSize() bool {
	if o != nil && o.StorageSize != nil {
		return true
	}

	return false
}



// GetK8sVersion returns the K8sVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetK8sVersion() *string {
	if o == nil {
		return nil
	}

	return o.K8sVersion
}

// GetK8sVersionOk returns a tuple with the K8sVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetK8sVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.K8sVersion, true
}

// SetK8sVersion sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetK8sVersion(v string) {
	o.K8sVersion = &v
}

// HasK8sVersion returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasK8sVersion() bool {
	if o != nil && o.K8sVersion != nil {
		return true
	}

	return false
}



// GetMaintenanceWindow returns the MaintenanceWindow field value
// If the value is explicit nil, the zero value for KubernetesMaintenanceWindow will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetMaintenanceWindow() *KubernetesMaintenanceWindow {
	if o == nil {
		return nil
	}

	return o.MaintenanceWindow
}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetMaintenanceWindowOk() (*KubernetesMaintenanceWindow, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaintenanceWindow, true
}

// SetMaintenanceWindow sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetMaintenanceWindow(v KubernetesMaintenanceWindow) {
	o.MaintenanceWindow = &v
}

// HasMaintenanceWindow returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasMaintenanceWindow() bool {
	if o != nil && o.MaintenanceWindow != nil {
		return true
	}

	return false
}



// GetAutoScaling returns the AutoScaling field value
// If the value is explicit nil, the zero value for KubernetesAutoScaling will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetAutoScaling() *KubernetesAutoScaling {
	if o == nil {
		return nil
	}

	return o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetAutoScalingOk() (*KubernetesAutoScaling, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoScaling, true
}

// SetAutoScaling sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetAutoScaling(v KubernetesAutoScaling) {
	o.AutoScaling = &v
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasAutoScaling() bool {
	if o != nil && o.AutoScaling != nil {
		return true
	}

	return false
}



// GetLans returns the Lans field value
// If the value is explicit nil, the zero value for []KubernetesNodePoolLan will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetLans() *[]KubernetesNodePoolLan {
	if o == nil {
		return nil
	}

	return o.Lans
}

// GetLansOk returns a tuple with the Lans field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetLansOk() (*[]KubernetesNodePoolLan, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lans, true
}

// SetLans sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetLans(v []KubernetesNodePoolLan) {
	o.Lans = &v
}

// HasLans returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasLans() bool {
	if o != nil && o.Lans != nil {
		return true
	}

	return false
}



// GetLabels returns the Labels field value
// If the value is explicit nil, the zero value for KubernetesNodePoolLabel will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetLabels() *KubernetesNodePoolLabel {
	if o == nil {
		return nil
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetLabelsOk() (*KubernetesNodePoolLabel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetLabels(v KubernetesNodePoolLabel) {
	o.Labels = &v
}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}



// GetAnnotations returns the Annotations field value
// If the value is explicit nil, the zero value for KubernetesNodePoolAnnotation will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetAnnotations() *KubernetesNodePoolAnnotation {
	if o == nil {
		return nil
	}

	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetAnnotationsOk() (*KubernetesNodePoolAnnotation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Annotations, true
}

// SetAnnotations sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetAnnotations(v KubernetesNodePoolAnnotation) {
	o.Annotations = &v
}

// HasAnnotations returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}



// GetPublicIps returns the PublicIps field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetPublicIps() *[]string {
	if o == nil {
		return nil
	}

	return o.PublicIps
}

// GetPublicIpsOk returns a tuple with the PublicIps field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolPropertiesForPost) GetPublicIpsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicIps, true
}

// SetPublicIps sets field value
func (o *KubernetesNodePoolPropertiesForPost) SetPublicIps(v []string) {
	o.PublicIps = &v
}

// HasPublicIps returns a boolean if a field has been set.
func (o *KubernetesNodePoolPropertiesForPost) HasPublicIps() bool {
	if o != nil && o.PublicIps != nil {
		return true
	}

	return false
}


func (o KubernetesNodePoolPropertiesForPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	

	if o.DatacenterId != nil {
		toSerialize["datacenterId"] = o.DatacenterId
	}
	

	if o.NodeCount != nil {
		toSerialize["nodeCount"] = o.NodeCount
	}
	

	if o.CpuFamily != nil {
		toSerialize["cpuFamily"] = o.CpuFamily
	}
	

	if o.CoresCount != nil {
		toSerialize["coresCount"] = o.CoresCount
	}
	

	if o.RamSize != nil {
		toSerialize["ramSize"] = o.RamSize
	}
	

	if o.AvailabilityZone != nil {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	

	if o.StorageType != nil {
		toSerialize["storageType"] = o.StorageType
	}
	

	if o.StorageSize != nil {
		toSerialize["storageSize"] = o.StorageSize
	}
	

	if o.K8sVersion != nil {
		toSerialize["k8sVersion"] = o.K8sVersion
	}
	

	if o.MaintenanceWindow != nil {
		toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	}
	

	if o.AutoScaling != nil {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	

	if o.Lans != nil {
		toSerialize["lans"] = o.Lans
	}
	

	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	

	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	

	if o.PublicIps != nil {
		toSerialize["publicIps"] = o.PublicIps
	}
	
	return json.Marshal(toSerialize)
}

type NullableKubernetesNodePoolPropertiesForPost struct {
	value *KubernetesNodePoolPropertiesForPost
	isSet bool
}

func (v NullableKubernetesNodePoolPropertiesForPost) Get() *KubernetesNodePoolPropertiesForPost {
	return v.value
}

func (v *NullableKubernetesNodePoolPropertiesForPost) Set(val *KubernetesNodePoolPropertiesForPost) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodePoolPropertiesForPost) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodePoolPropertiesForPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodePoolPropertiesForPost(val *KubernetesNodePoolPropertiesForPost) *NullableKubernetesNodePoolPropertiesForPost {
	return &NullableKubernetesNodePoolPropertiesForPost{value: val, isSet: true}
}

func (v NullableKubernetesNodePoolPropertiesForPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodePoolPropertiesForPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

