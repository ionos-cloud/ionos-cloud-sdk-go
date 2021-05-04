/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0-SDK.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// LanEntities struct for LanEntities
type LanEntities struct {
	Nics *LanNics `json:"nics,omitempty"`
}



// GetNics returns the Nics field value
// If the value is explicit nil, the zero value for LanNics will be returned
func (o *LanEntities) GetNics() *LanNics {
	if o == nil {
		return nil
	}


	return o.Nics

}

// GetNicsOk returns a tuple with the Nics field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanEntities) GetNicsOk() (*LanNics, bool) {
	if o == nil {
		return nil, false
	}


	return o.Nics, true
}

// SetNics sets field value
func (o *LanEntities) SetNics(v LanNics) {


	o.Nics = &v

}

// HasNics returns a boolean if a field has been set.
func (o *LanEntities) HasNics() bool {
	if o != nil && o.Nics != nil {
		return true
	}

	return false
}


func (o LanEntities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Nics != nil {
		toSerialize["nics"] = o.Nics
	}
	
	return json.Marshal(toSerialize)
}

type NullableLanEntities struct {
	value *LanEntities
	isSet bool
}

func (v NullableLanEntities) Get() *LanEntities {
	return v.value
}

func (v *NullableLanEntities) Set(val *LanEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableLanEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableLanEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanEntities(val *LanEntities) *NullableLanEntities {
	return &NullableLanEntities{value: val, isSet: true}
}

func (v NullableLanEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


