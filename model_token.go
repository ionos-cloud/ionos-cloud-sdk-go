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

// Token struct for Token
type Token struct {
	// The jwToken for the server
	Token *string `json:"token,omitempty"`
}



// GetToken returns the Token field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Token) GetToken() *string {
	if o == nil {
		return nil
	}


	return o.Token

}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Token) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Token, true
}

// SetToken sets field value
func (o *Token) SetToken(v string) {


	o.Token = &v

}

// HasToken returns a boolean if a field has been set.
func (o *Token) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}


func (o Token) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	
	return json.Marshal(toSerialize)
}

type NullableToken struct {
	value *Token
	isSet bool
}

func (v NullableToken) Get() *Token {
	return v.value
}

func (v *NullableToken) Set(val *Token) {
	v.value = val
	v.isSet = true
}

func (v NullableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToken(val *Token) *NullableToken {
	return &NullableToken{value: val, isSet: true}
}

func (v NullableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

