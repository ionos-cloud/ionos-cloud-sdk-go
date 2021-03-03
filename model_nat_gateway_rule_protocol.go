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
	"fmt"
)

// NatGatewayRuleProtocol the model 'NatGatewayRuleProtocol'
type NatGatewayRuleProtocol string

// List of NatGatewayRuleProtocol
const (
	TCP NatGatewayRuleProtocol = "TCP"
	UDP NatGatewayRuleProtocol = "UDP"
	ICMP NatGatewayRuleProtocol = "ICMP"
	ALL NatGatewayRuleProtocol = "ALL"
)

func (v *NatGatewayRuleProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NatGatewayRuleProtocol(value)
	for _, existing := range []NatGatewayRuleProtocol{ "TCP", "UDP", "ICMP", "ALL",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NatGatewayRuleProtocol", value)
}

// Ptr returns reference to NatGatewayRuleProtocol value
func (v NatGatewayRuleProtocol) Ptr() *NatGatewayRuleProtocol {
	return &v
}

type NullableNatGatewayRuleProtocol struct {
	value *NatGatewayRuleProtocol
	isSet bool
}

func (v NullableNatGatewayRuleProtocol) Get() *NatGatewayRuleProtocol {
	return v.value
}

func (v *NullableNatGatewayRuleProtocol) Set(val *NatGatewayRuleProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableNatGatewayRuleProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableNatGatewayRuleProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNatGatewayRuleProtocol(val *NatGatewayRuleProtocol) *NullableNatGatewayRuleProtocol {
	return &NullableNatGatewayRuleProtocol{value: val, isSet: true}
}

func (v NullableNatGatewayRuleProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNatGatewayRuleProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

