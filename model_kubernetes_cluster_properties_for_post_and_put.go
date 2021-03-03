/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// KubernetesClusterPropertiesForPostAndPut struct for KubernetesClusterPropertiesForPostAndPut
type KubernetesClusterPropertiesForPostAndPut struct {
	// A Kubernetes Cluster Name. Valid Kubernetes Cluster name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Name *string `json:"name"`
	// The kubernetes version in which a cluster is running. This imposes restrictions on what kubernetes versions can be run in a cluster's nodepools. Additionally, not all kubernetes versions are viable upgrade targets for all prior versions.
	K8sVersion *string `json:"k8sVersion,omitempty"`
	MaintenanceWindow *KubernetesMaintenanceWindow `json:"maintenanceWindow,omitempty"`
}



// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesClusterPropertiesForPostAndPut) GetName() *string {
	if o == nil {
		return nil
	}


	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterPropertiesForPostAndPut) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Name, true
}

// SetName sets field value
func (o *KubernetesClusterPropertiesForPostAndPut) SetName(v string) {


	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *KubernetesClusterPropertiesForPostAndPut) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}



// GetK8sVersion returns the K8sVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesClusterPropertiesForPostAndPut) GetK8sVersion() *string {
	if o == nil {
		return nil
	}


	return o.K8sVersion

}

// GetK8sVersionOk returns a tuple with the K8sVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterPropertiesForPostAndPut) GetK8sVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.K8sVersion, true
}

// SetK8sVersion sets field value
func (o *KubernetesClusterPropertiesForPostAndPut) SetK8sVersion(v string) {


	o.K8sVersion = &v

}

// HasK8sVersion returns a boolean if a field has been set.
func (o *KubernetesClusterPropertiesForPostAndPut) HasK8sVersion() bool {
	if o != nil && o.K8sVersion != nil {
		return true
	}

	return false
}



// GetMaintenanceWindow returns the MaintenanceWindow field value
// If the value is explicit nil, the zero value for KubernetesMaintenanceWindow will be returned
func (o *KubernetesClusterPropertiesForPostAndPut) GetMaintenanceWindow() *KubernetesMaintenanceWindow {
	if o == nil {
		return nil
	}


	return o.MaintenanceWindow

}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterPropertiesForPostAndPut) GetMaintenanceWindowOk() (*KubernetesMaintenanceWindow, bool) {
	if o == nil {
		return nil, false
	}


	return o.MaintenanceWindow, true
}

// SetMaintenanceWindow sets field value
func (o *KubernetesClusterPropertiesForPostAndPut) SetMaintenanceWindow(v KubernetesMaintenanceWindow) {


	o.MaintenanceWindow = &v

}

// HasMaintenanceWindow returns a boolean if a field has been set.
func (o *KubernetesClusterPropertiesForPostAndPut) HasMaintenanceWindow() bool {
	if o != nil && o.MaintenanceWindow != nil {
		return true
	}

	return false
}


func (o KubernetesClusterPropertiesForPostAndPut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	

	if o.K8sVersion != nil {
		toSerialize["k8sVersion"] = o.K8sVersion
	}
	

	if o.MaintenanceWindow != nil {
		toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	}
	
	return json.Marshal(toSerialize)
}

type NullableKubernetesClusterPropertiesForPostAndPut struct {
	value *KubernetesClusterPropertiesForPostAndPut
	isSet bool
}

func (v NullableKubernetesClusterPropertiesForPostAndPut) Get() *KubernetesClusterPropertiesForPostAndPut {
	return v.value
}

func (v *NullableKubernetesClusterPropertiesForPostAndPut) Set(val *KubernetesClusterPropertiesForPostAndPut) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterPropertiesForPostAndPut) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterPropertiesForPostAndPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterPropertiesForPostAndPut(val *KubernetesClusterPropertiesForPostAndPut) *NullableKubernetesClusterPropertiesForPostAndPut {
	return &NullableKubernetesClusterPropertiesForPostAndPut{value: val, isSet: true}
}

func (v NullableKubernetesClusterPropertiesForPostAndPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterPropertiesForPostAndPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


