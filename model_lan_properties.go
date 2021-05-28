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

// LanProperties struct for LanProperties
type LanProperties struct {
	// A name of that resource
	Name *string `json:"name,omitempty"`
	// IP failover configurations for lan
	IpFailover *[]IPFailover `json:"ipFailover,omitempty"`
	// Unique identifier of the private cross connect the given LAN is connected to if any
	Pcc *string `json:"pcc,omitempty"`
	// Does this LAN faces the public Internet or not
	Public *bool `json:"public,omitempty"`
}



// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LanProperties) GetName() *string {
	if o == nil {
		return nil
	}


	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Name, true
}

// SetName sets field value
func (o *LanProperties) SetName(v string) {


	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *LanProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}



// GetIpFailover returns the IpFailover field value
// If the value is explicit nil, the zero value for []IPFailover will be returned
func (o *LanProperties) GetIpFailover() *[]IPFailover {
	if o == nil {
		return nil
	}


	return o.IpFailover

}

// GetIpFailoverOk returns a tuple with the IpFailover field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanProperties) GetIpFailoverOk() (*[]IPFailover, bool) {
	if o == nil {
		return nil, false
	}


	return o.IpFailover, true
}

// SetIpFailover sets field value
func (o *LanProperties) SetIpFailover(v []IPFailover) {


	o.IpFailover = &v

}

// HasIpFailover returns a boolean if a field has been set.
func (o *LanProperties) HasIpFailover() bool {
	if o != nil && o.IpFailover != nil {
		return true
	}

	return false
}



// GetPcc returns the Pcc field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LanProperties) GetPcc() *string {
	if o == nil {
		return nil
	}


	return o.Pcc

}

// GetPccOk returns a tuple with the Pcc field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanProperties) GetPccOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Pcc, true
}

// SetPcc sets field value
func (o *LanProperties) SetPcc(v string) {


	o.Pcc = &v

}

// HasPcc returns a boolean if a field has been set.
func (o *LanProperties) HasPcc() bool {
	if o != nil && o.Pcc != nil {
		return true
	}

	return false
}



// GetPublic returns the Public field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *LanProperties) GetPublic() *bool {
	if o == nil {
		return nil
	}


	return o.Public

}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanProperties) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}


	return o.Public, true
}

// SetPublic sets field value
func (o *LanProperties) SetPublic(v bool) {


	o.Public = &v

}

// HasPublic returns a boolean if a field has been set.
func (o *LanProperties) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}


func (o LanProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	

	if o.IpFailover != nil {
		toSerialize["ipFailover"] = o.IpFailover
	}
	

	if o.Pcc != nil {
		toSerialize["pcc"] = o.Pcc
	}
	

	if o.Public != nil {
		toSerialize["public"] = o.Public
	}
	
	return json.Marshal(toSerialize)
}

type NullableLanProperties struct {
	value *LanProperties
	isSet bool
}

func (v NullableLanProperties) Get() *LanProperties {
	return v.value
}

func (v *NullableLanProperties) Set(val *LanProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableLanProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableLanProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanProperties(val *LanProperties) *NullableLanProperties {
	return &NullableLanProperties{value: val, isSet: true}
}

func (v NullableLanProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


