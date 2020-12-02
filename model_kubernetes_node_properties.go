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

// KubernetesNodeProperties struct for KubernetesNodeProperties
type KubernetesNodeProperties struct {
	// A Kubernetes Node Name.
	Name *string `json:"name"`
	// A valid public IP.
	PublicIP *string `json:"publicIP"`
	// The kubernetes version in which a nodepool is running. This imposes restrictions on what kubernetes versions can be run in a cluster's nodepools. Additionally, not all kubernetes versions are viable upgrade targets for all prior versions.
	K8sVersion *string `json:"k8sVersion"`
}



// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodeProperties) GetName() *string {
	if o == nil {
		return nil
	}


	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Name, true
}

// SetName sets field value
func (o *KubernetesNodeProperties) SetName(v string) {


	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *KubernetesNodeProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}



// GetPublicIP returns the PublicIP field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodeProperties) GetPublicIP() *string {
	if o == nil {
		return nil
	}


	return o.PublicIP

}

// GetPublicIPOk returns a tuple with the PublicIP field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeProperties) GetPublicIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.PublicIP, true
}

// SetPublicIP sets field value
func (o *KubernetesNodeProperties) SetPublicIP(v string) {


	o.PublicIP = &v

}

// HasPublicIP returns a boolean if a field has been set.
func (o *KubernetesNodeProperties) HasPublicIP() bool {
	if o != nil && o.PublicIP != nil {
		return true
	}

	return false
}



// GetK8sVersion returns the K8sVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodeProperties) GetK8sVersion() *string {
	if o == nil {
		return nil
	}


	return o.K8sVersion

}

// GetK8sVersionOk returns a tuple with the K8sVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeProperties) GetK8sVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.K8sVersion, true
}

// SetK8sVersion sets field value
func (o *KubernetesNodeProperties) SetK8sVersion(v string) {


	o.K8sVersion = &v

}

// HasK8sVersion returns a boolean if a field has been set.
func (o *KubernetesNodeProperties) HasK8sVersion() bool {
	if o != nil && o.K8sVersion != nil {
		return true
	}

	return false
}


func (o KubernetesNodeProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	

	if o.PublicIP != nil {
		toSerialize["publicIP"] = o.PublicIP
	}
	

	if o.K8sVersion != nil {
		toSerialize["k8sVersion"] = o.K8sVersion
	}
	
	return json.Marshal(toSerialize)
}

type NullableKubernetesNodeProperties struct {
	value *KubernetesNodeProperties
	isSet bool
}

func (v NullableKubernetesNodeProperties) Get() *KubernetesNodeProperties {
	return v.value
}

func (v *NullableKubernetesNodeProperties) Set(val *KubernetesNodeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeProperties(val *KubernetesNodeProperties) *NullableKubernetesNodeProperties {
	return &NullableKubernetesNodeProperties{value: val, isSet: true}
}

func (v NullableKubernetesNodeProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


