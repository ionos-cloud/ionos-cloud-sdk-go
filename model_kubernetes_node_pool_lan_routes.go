/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 6.0-SDK.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// KubernetesNodePoolLanRoutes struct for KubernetesNodePoolLanRoutes
type KubernetesNodePoolLanRoutes struct {
	// IPv4 or IPv6 CIDR to be routed via the interface.
	Network *string `json:"network,omitempty"`
	// IPv4 or IPv6 Gateway IP for the route.
	GatewayIp *string `json:"gatewayIp,omitempty"`
}



// GetNetwork returns the Network field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolLanRoutes) GetNetwork() *string {
	if o == nil {
		return nil
	}


	return o.Network

}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolLanRoutes) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Network, true
}

// SetNetwork sets field value
func (o *KubernetesNodePoolLanRoutes) SetNetwork(v string) {


	o.Network = &v

}

// HasNetwork returns a boolean if a field has been set.
func (o *KubernetesNodePoolLanRoutes) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}



// GetGatewayIp returns the GatewayIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesNodePoolLanRoutes) GetGatewayIp() *string {
	if o == nil {
		return nil
	}


	return o.GatewayIp

}

// GetGatewayIpOk returns a tuple with the GatewayIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodePoolLanRoutes) GetGatewayIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.GatewayIp, true
}

// SetGatewayIp sets field value
func (o *KubernetesNodePoolLanRoutes) SetGatewayIp(v string) {


	o.GatewayIp = &v

}

// HasGatewayIp returns a boolean if a field has been set.
func (o *KubernetesNodePoolLanRoutes) HasGatewayIp() bool {
	if o != nil && o.GatewayIp != nil {
		return true
	}

	return false
}


func (o KubernetesNodePoolLanRoutes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	

	if o.GatewayIp != nil {
		toSerialize["gatewayIp"] = o.GatewayIp
	}
	
	return json.Marshal(toSerialize)
}

type NullableKubernetesNodePoolLanRoutes struct {
	value *KubernetesNodePoolLanRoutes
	isSet bool
}

func (v NullableKubernetesNodePoolLanRoutes) Get() *KubernetesNodePoolLanRoutes {
	return v.value
}

func (v *NullableKubernetesNodePoolLanRoutes) Set(val *KubernetesNodePoolLanRoutes) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodePoolLanRoutes) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodePoolLanRoutes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodePoolLanRoutes(val *KubernetesNodePoolLanRoutes) *NullableKubernetesNodePoolLanRoutes {
	return &NullableKubernetesNodePoolLanRoutes{value: val, isSet: true}
}

func (v NullableKubernetesNodePoolLanRoutes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodePoolLanRoutes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


