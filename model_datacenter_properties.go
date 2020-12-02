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

// DatacenterProperties struct for DatacenterProperties
type DatacenterProperties struct {
	// A name of that resource
	Name *string `json:"name,omitempty"`
	// A description for the datacenter, e.g. staging, production
	Description *string `json:"description,omitempty"`
	// The physical location where the datacenter will be created. This will be where all of your servers live. Property cannot be modified after datacenter creation (disallowed in update requests)
	Location *string `json:"location"`
	// The version of that Data Center. Gets incremented with every change
	Version *int32 `json:"version,omitempty"`
	// List of features supported by the location this data center is part of
	Features *[]string `json:"features,omitempty"`
	// Boolean value representing if the data center requires extra protection e.g. two factor protection
	SecAuthProtection *bool `json:"secAuthProtection,omitempty"`
}



// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DatacenterProperties) GetName() *string {
	if o == nil {
		return nil
	}


	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatacenterProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Name, true
}

// SetName sets field value
func (o *DatacenterProperties) SetName(v string) {


	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *DatacenterProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}



// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DatacenterProperties) GetDescription() *string {
	if o == nil {
		return nil
	}


	return o.Description

}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatacenterProperties) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Description, true
}

// SetDescription sets field value
func (o *DatacenterProperties) SetDescription(v string) {


	o.Description = &v

}

// HasDescription returns a boolean if a field has been set.
func (o *DatacenterProperties) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}



// GetLocation returns the Location field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DatacenterProperties) GetLocation() *string {
	if o == nil {
		return nil
	}


	return o.Location

}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatacenterProperties) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Location, true
}

// SetLocation sets field value
func (o *DatacenterProperties) SetLocation(v string) {


	o.Location = &v

}

// HasLocation returns a boolean if a field has been set.
func (o *DatacenterProperties) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}



// GetVersion returns the Version field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DatacenterProperties) GetVersion() *int32 {
	if o == nil {
		return nil
	}


	return o.Version

}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatacenterProperties) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}


	return o.Version, true
}

// SetVersion sets field value
func (o *DatacenterProperties) SetVersion(v int32) {


	o.Version = &v

}

// HasVersion returns a boolean if a field has been set.
func (o *DatacenterProperties) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}



// GetFeatures returns the Features field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *DatacenterProperties) GetFeatures() *[]string {
	if o == nil {
		return nil
	}


	return o.Features

}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatacenterProperties) GetFeaturesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Features, true
}

// SetFeatures sets field value
func (o *DatacenterProperties) SetFeatures(v []string) {


	o.Features = &v

}

// HasFeatures returns a boolean if a field has been set.
func (o *DatacenterProperties) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}



// GetSecAuthProtection returns the SecAuthProtection field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *DatacenterProperties) GetSecAuthProtection() *bool {
	if o == nil {
		return nil
	}


	return o.SecAuthProtection

}

// GetSecAuthProtectionOk returns a tuple with the SecAuthProtection field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatacenterProperties) GetSecAuthProtectionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}


	return o.SecAuthProtection, true
}

// SetSecAuthProtection sets field value
func (o *DatacenterProperties) SetSecAuthProtection(v bool) {


	o.SecAuthProtection = &v

}

// HasSecAuthProtection returns a boolean if a field has been set.
func (o *DatacenterProperties) HasSecAuthProtection() bool {
	if o != nil && o.SecAuthProtection != nil {
		return true
	}

	return false
}


func (o DatacenterProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	

	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	

	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	

	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	

	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	

	if o.SecAuthProtection != nil {
		toSerialize["secAuthProtection"] = o.SecAuthProtection
	}
	
	return json.Marshal(toSerialize)
}

type NullableDatacenterProperties struct {
	value *DatacenterProperties
	isSet bool
}

func (v NullableDatacenterProperties) Get() *DatacenterProperties {
	return v.value
}

func (v *NullableDatacenterProperties) Set(val *DatacenterProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableDatacenterProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableDatacenterProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatacenterProperties(val *DatacenterProperties) *NullableDatacenterProperties {
	return &NullableDatacenterProperties{value: val, isSet: true}
}

func (v NullableDatacenterProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatacenterProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


