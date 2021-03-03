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

// IpConsumer struct for IpConsumer
type IpConsumer struct {
	Ip *string `json:"ip,omitempty"`
	Mac *string `json:"mac,omitempty"`
	NicId *string `json:"nicId,omitempty"`
	ServerId *string `json:"serverId,omitempty"`
	ServerName *string `json:"serverName,omitempty"`
	DatacenterId *string `json:"datacenterId,omitempty"`
	DatacenterName *string `json:"datacenterName,omitempty"`
}



// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IpConsumer) GetIp() *string {
	if o == nil {
		return nil
	}


	return o.Ip

}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpConsumer) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Ip, true
}

// SetIp sets field value
func (o *IpConsumer) SetIp(v string) {


	o.Ip = &v

}

// HasIp returns a boolean if a field has been set.
func (o *IpConsumer) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}



// GetMac returns the Mac field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IpConsumer) GetMac() *string {
	if o == nil {
		return nil
	}


	return o.Mac

}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpConsumer) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Mac, true
}

// SetMac sets field value
func (o *IpConsumer) SetMac(v string) {


	o.Mac = &v

}

// HasMac returns a boolean if a field has been set.
func (o *IpConsumer) HasMac() bool {
	if o != nil && o.Mac != nil {
		return true
	}

	return false
}



// GetNicId returns the NicId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IpConsumer) GetNicId() *string {
	if o == nil {
		return nil
	}


	return o.NicId

}

// GetNicIdOk returns a tuple with the NicId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpConsumer) GetNicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.NicId, true
}

// SetNicId sets field value
func (o *IpConsumer) SetNicId(v string) {


	o.NicId = &v

}

// HasNicId returns a boolean if a field has been set.
func (o *IpConsumer) HasNicId() bool {
	if o != nil && o.NicId != nil {
		return true
	}

	return false
}



// GetServerId returns the ServerId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IpConsumer) GetServerId() *string {
	if o == nil {
		return nil
	}


	return o.ServerId

}

// GetServerIdOk returns a tuple with the ServerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpConsumer) GetServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.ServerId, true
}

// SetServerId sets field value
func (o *IpConsumer) SetServerId(v string) {


	o.ServerId = &v

}

// HasServerId returns a boolean if a field has been set.
func (o *IpConsumer) HasServerId() bool {
	if o != nil && o.ServerId != nil {
		return true
	}

	return false
}



// GetServerName returns the ServerName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IpConsumer) GetServerName() *string {
	if o == nil {
		return nil
	}


	return o.ServerName

}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpConsumer) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.ServerName, true
}

// SetServerName sets field value
func (o *IpConsumer) SetServerName(v string) {


	o.ServerName = &v

}

// HasServerName returns a boolean if a field has been set.
func (o *IpConsumer) HasServerName() bool {
	if o != nil && o.ServerName != nil {
		return true
	}

	return false
}



// GetDatacenterId returns the DatacenterId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IpConsumer) GetDatacenterId() *string {
	if o == nil {
		return nil
	}


	return o.DatacenterId

}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpConsumer) GetDatacenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.DatacenterId, true
}

// SetDatacenterId sets field value
func (o *IpConsumer) SetDatacenterId(v string) {


	o.DatacenterId = &v

}

// HasDatacenterId returns a boolean if a field has been set.
func (o *IpConsumer) HasDatacenterId() bool {
	if o != nil && o.DatacenterId != nil {
		return true
	}

	return false
}



// GetDatacenterName returns the DatacenterName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IpConsumer) GetDatacenterName() *string {
	if o == nil {
		return nil
	}


	return o.DatacenterName

}

// GetDatacenterNameOk returns a tuple with the DatacenterName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpConsumer) GetDatacenterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.DatacenterName, true
}

// SetDatacenterName sets field value
func (o *IpConsumer) SetDatacenterName(v string) {


	o.DatacenterName = &v

}

// HasDatacenterName returns a boolean if a field has been set.
func (o *IpConsumer) HasDatacenterName() bool {
	if o != nil && o.DatacenterName != nil {
		return true
	}

	return false
}


func (o IpConsumer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	

	if o.Mac != nil {
		toSerialize["mac"] = o.Mac
	}
	

	if o.NicId != nil {
		toSerialize["nicId"] = o.NicId
	}
	

	if o.ServerId != nil {
		toSerialize["serverId"] = o.ServerId
	}
	

	if o.ServerName != nil {
		toSerialize["serverName"] = o.ServerName
	}
	

	if o.DatacenterId != nil {
		toSerialize["datacenterId"] = o.DatacenterId
	}
	

	if o.DatacenterName != nil {
		toSerialize["datacenterName"] = o.DatacenterName
	}
	
	return json.Marshal(toSerialize)
}

type NullableIpConsumer struct {
	value *IpConsumer
	isSet bool
}

func (v NullableIpConsumer) Get() *IpConsumer {
	return v.value
}

func (v *NullableIpConsumer) Set(val *IpConsumer) {
	v.value = val
	v.isSet = true
}

func (v NullableIpConsumer) IsSet() bool {
	return v.isSet
}

func (v *NullableIpConsumer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpConsumer(val *IpConsumer) *NullableIpConsumer {
	return &NullableIpConsumer{value: val, isSet: true}
}

func (v NullableIpConsumer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpConsumer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


