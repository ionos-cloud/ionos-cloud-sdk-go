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

// BackupUnitProperties struct for BackupUnitProperties
type BackupUnitProperties struct {
	// A name of that resource (only alphanumeric characters are acceptable)
	Name *string `json:"name"`
	// the password associated to that resource
	Password *string `json:"password,omitempty"`
	// The email associated with the backup unit. Bear in mind that this email does not be the same email as of the user.
	Email *string `json:"email,omitempty"`
}



// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BackupUnitProperties) GetName() *string {
	if o == nil {
		return nil
	}


	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupUnitProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Name, true
}

// SetName sets field value
func (o *BackupUnitProperties) SetName(v string) {


	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *BackupUnitProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}



// GetPassword returns the Password field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BackupUnitProperties) GetPassword() *string {
	if o == nil {
		return nil
	}


	return o.Password

}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupUnitProperties) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Password, true
}

// SetPassword sets field value
func (o *BackupUnitProperties) SetPassword(v string) {


	o.Password = &v

}

// HasPassword returns a boolean if a field has been set.
func (o *BackupUnitProperties) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}



// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BackupUnitProperties) GetEmail() *string {
	if o == nil {
		return nil
	}


	return o.Email

}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupUnitProperties) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}


	return o.Email, true
}

// SetEmail sets field value
func (o *BackupUnitProperties) SetEmail(v string) {


	o.Email = &v

}

// HasEmail returns a boolean if a field has been set.
func (o *BackupUnitProperties) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}


func (o BackupUnitProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	

	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	

	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	
	return json.Marshal(toSerialize)
}

type NullableBackupUnitProperties struct {
	value *BackupUnitProperties
	isSet bool
}

func (v NullableBackupUnitProperties) Get() *BackupUnitProperties {
	return v.value
}

func (v *NullableBackupUnitProperties) Set(val *BackupUnitProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupUnitProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupUnitProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupUnitProperties(val *BackupUnitProperties) *NullableBackupUnitProperties {
	return &NullableBackupUnitProperties{value: val, isSet: true}
}

func (v NullableBackupUnitProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupUnitProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


