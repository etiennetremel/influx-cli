/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ResourceMembersLinks struct for ResourceMembersLinks
type ResourceMembersLinks struct {
	Self *string `json:"self,omitempty"`
}

// NewResourceMembersLinks instantiates a new ResourceMembersLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceMembersLinks() *ResourceMembersLinks {
	this := ResourceMembersLinks{}
	return &this
}

// NewResourceMembersLinksWithDefaults instantiates a new ResourceMembersLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceMembersLinksWithDefaults() *ResourceMembersLinks {
	this := ResourceMembersLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ResourceMembersLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceMembersLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ResourceMembersLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *ResourceMembersLinks) SetSelf(v string) {
	o.Self = &v
}

func (o ResourceMembersLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	return json.Marshal(toSerialize)
}

type NullableResourceMembersLinks struct {
	value *ResourceMembersLinks
	isSet bool
}

func (v NullableResourceMembersLinks) Get() *ResourceMembersLinks {
	return v.value
}

func (v *NullableResourceMembersLinks) Set(val *ResourceMembersLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceMembersLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceMembersLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceMembersLinks(val *ResourceMembersLinks) *NullableResourceMembersLinks {
	return &NullableResourceMembersLinks{value: val, isSet: true}
}

func (v NullableResourceMembersLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceMembersLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}