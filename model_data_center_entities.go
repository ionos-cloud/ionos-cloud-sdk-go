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

// DataCenterEntities struct for DataCenterEntities
type DataCenterEntities struct {
	Servers *Servers `json:"servers,omitempty"`
	Volumes *Volumes `json:"volumes,omitempty"`
	Loadbalancers *Loadbalancers `json:"loadbalancers,omitempty"`
	Lans *Lans `json:"lans,omitempty"`
	Networkloadbalancers *NetworkLoadBalancers `json:"networkloadbalancers,omitempty"`
	Natgateways *NatGateways `json:"natgateways,omitempty"`
}



// GetServers returns the Servers field value
// If the value is explicit nil, the zero value for Servers will be returned
func (o *DataCenterEntities) GetServers() *Servers {
	if o == nil {
		return nil
	}


	return o.Servers

}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataCenterEntities) GetServersOk() (*Servers, bool) {
	if o == nil {
		return nil, false
	}


	return o.Servers, true
}

// SetServers sets field value
func (o *DataCenterEntities) SetServers(v Servers) {


	o.Servers = &v

}

// HasServers returns a boolean if a field has been set.
func (o *DataCenterEntities) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}



// GetVolumes returns the Volumes field value
// If the value is explicit nil, the zero value for Volumes will be returned
func (o *DataCenterEntities) GetVolumes() *Volumes {
	if o == nil {
		return nil
	}


	return o.Volumes

}

// GetVolumesOk returns a tuple with the Volumes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataCenterEntities) GetVolumesOk() (*Volumes, bool) {
	if o == nil {
		return nil, false
	}


	return o.Volumes, true
}

// SetVolumes sets field value
func (o *DataCenterEntities) SetVolumes(v Volumes) {


	o.Volumes = &v

}

// HasVolumes returns a boolean if a field has been set.
func (o *DataCenterEntities) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}



// GetLoadbalancers returns the Loadbalancers field value
// If the value is explicit nil, the zero value for Loadbalancers will be returned
func (o *DataCenterEntities) GetLoadbalancers() *Loadbalancers {
	if o == nil {
		return nil
	}


	return o.Loadbalancers

}

// GetLoadbalancersOk returns a tuple with the Loadbalancers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataCenterEntities) GetLoadbalancersOk() (*Loadbalancers, bool) {
	if o == nil {
		return nil, false
	}


	return o.Loadbalancers, true
}

// SetLoadbalancers sets field value
func (o *DataCenterEntities) SetLoadbalancers(v Loadbalancers) {


	o.Loadbalancers = &v

}

// HasLoadbalancers returns a boolean if a field has been set.
func (o *DataCenterEntities) HasLoadbalancers() bool {
	if o != nil && o.Loadbalancers != nil {
		return true
	}

	return false
}



// GetLans returns the Lans field value
// If the value is explicit nil, the zero value for Lans will be returned
func (o *DataCenterEntities) GetLans() *Lans {
	if o == nil {
		return nil
	}


	return o.Lans

}

// GetLansOk returns a tuple with the Lans field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataCenterEntities) GetLansOk() (*Lans, bool) {
	if o == nil {
		return nil, false
	}


	return o.Lans, true
}

// SetLans sets field value
func (o *DataCenterEntities) SetLans(v Lans) {


	o.Lans = &v

}

// HasLans returns a boolean if a field has been set.
func (o *DataCenterEntities) HasLans() bool {
	if o != nil && o.Lans != nil {
		return true
	}

	return false
}



// GetNetworkloadbalancers returns the Networkloadbalancers field value
// If the value is explicit nil, the zero value for NetworkLoadBalancers will be returned
func (o *DataCenterEntities) GetNetworkloadbalancers() *NetworkLoadBalancers {
	if o == nil {
		return nil
	}


	return o.Networkloadbalancers

}

// GetNetworkloadbalancersOk returns a tuple with the Networkloadbalancers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataCenterEntities) GetNetworkloadbalancersOk() (*NetworkLoadBalancers, bool) {
	if o == nil {
		return nil, false
	}


	return o.Networkloadbalancers, true
}

// SetNetworkloadbalancers sets field value
func (o *DataCenterEntities) SetNetworkloadbalancers(v NetworkLoadBalancers) {


	o.Networkloadbalancers = &v

}

// HasNetworkloadbalancers returns a boolean if a field has been set.
func (o *DataCenterEntities) HasNetworkloadbalancers() bool {
	if o != nil && o.Networkloadbalancers != nil {
		return true
	}

	return false
}



// GetNatgateways returns the Natgateways field value
// If the value is explicit nil, the zero value for NatGateways will be returned
func (o *DataCenterEntities) GetNatgateways() *NatGateways {
	if o == nil {
		return nil
	}


	return o.Natgateways

}

// GetNatgatewaysOk returns a tuple with the Natgateways field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataCenterEntities) GetNatgatewaysOk() (*NatGateways, bool) {
	if o == nil {
		return nil, false
	}


	return o.Natgateways, true
}

// SetNatgateways sets field value
func (o *DataCenterEntities) SetNatgateways(v NatGateways) {


	o.Natgateways = &v

}

// HasNatgateways returns a boolean if a field has been set.
func (o *DataCenterEntities) HasNatgateways() bool {
	if o != nil && o.Natgateways != nil {
		return true
	}

	return false
}


func (o DataCenterEntities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	

	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}
	

	if o.Loadbalancers != nil {
		toSerialize["loadbalancers"] = o.Loadbalancers
	}
	

	if o.Lans != nil {
		toSerialize["lans"] = o.Lans
	}
	

	if o.Networkloadbalancers != nil {
		toSerialize["networkloadbalancers"] = o.Networkloadbalancers
	}
	

	if o.Natgateways != nil {
		toSerialize["natgateways"] = o.Natgateways
	}
	
	return json.Marshal(toSerialize)
}

type NullableDataCenterEntities struct {
	value *DataCenterEntities
	isSet bool
}

func (v NullableDataCenterEntities) Get() *DataCenterEntities {
	return v.value
}

func (v *NullableDataCenterEntities) Set(val *DataCenterEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableDataCenterEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableDataCenterEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataCenterEntities(val *DataCenterEntities) *NullableDataCenterEntities {
	return &NullableDataCenterEntities{value: val, isSet: true}
}

func (v NullableDataCenterEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataCenterEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


