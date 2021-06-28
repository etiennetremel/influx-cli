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

// LegacyAuthorizationPostRequestAllOf struct for LegacyAuthorizationPostRequestAllOf
type LegacyAuthorizationPostRequestAllOf struct {
	// ID of org that authorization is scoped to.
	OrgID *string `json:"orgID,omitempty" yaml:"orgID,omitempty"`
	// ID of user that authorization is scoped to.
	UserID *string `json:"userID,omitempty" yaml:"userID,omitempty"`
	// Token (name) of the authorization
	Token *string `json:"token,omitempty" yaml:"token,omitempty"`
	// List of permissions for an auth.  An auth must have at least one Permission.
	Permissions *[]Permission `json:"permissions,omitempty" yaml:"permissions,omitempty"`
}

// NewLegacyAuthorizationPostRequestAllOf instantiates a new LegacyAuthorizationPostRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyAuthorizationPostRequestAllOf() *LegacyAuthorizationPostRequestAllOf {
	this := LegacyAuthorizationPostRequestAllOf{}
	return &this
}

// NewLegacyAuthorizationPostRequestAllOfWithDefaults instantiates a new LegacyAuthorizationPostRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyAuthorizationPostRequestAllOfWithDefaults() *LegacyAuthorizationPostRequestAllOf {
	this := LegacyAuthorizationPostRequestAllOf{}
	return &this
}

// GetOrgID returns the OrgID field value if set, zero value otherwise.
func (o *LegacyAuthorizationPostRequestAllOf) GetOrgID() string {
	if o == nil || o.OrgID == nil {
		var ret string
		return ret
	}
	return *o.OrgID
}

// GetOrgIDOk returns a tuple with the OrgID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAuthorizationPostRequestAllOf) GetOrgIDOk() (*string, bool) {
	if o == nil || o.OrgID == nil {
		return nil, false
	}
	return o.OrgID, true
}

// HasOrgID returns a boolean if a field has been set.
func (o *LegacyAuthorizationPostRequestAllOf) HasOrgID() bool {
	if o != nil && o.OrgID != nil {
		return true
	}

	return false
}

// SetOrgID gets a reference to the given string and assigns it to the OrgID field.
func (o *LegacyAuthorizationPostRequestAllOf) SetOrgID(v string) {
	o.OrgID = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *LegacyAuthorizationPostRequestAllOf) GetUserID() string {
	if o == nil || o.UserID == nil {
		var ret string
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAuthorizationPostRequestAllOf) GetUserIDOk() (*string, bool) {
	if o == nil || o.UserID == nil {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *LegacyAuthorizationPostRequestAllOf) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given string and assigns it to the UserID field.
func (o *LegacyAuthorizationPostRequestAllOf) SetUserID(v string) {
	o.UserID = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *LegacyAuthorizationPostRequestAllOf) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAuthorizationPostRequestAllOf) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *LegacyAuthorizationPostRequestAllOf) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *LegacyAuthorizationPostRequestAllOf) SetToken(v string) {
	o.Token = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *LegacyAuthorizationPostRequestAllOf) GetPermissions() []Permission {
	if o == nil || o.Permissions == nil {
		var ret []Permission
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAuthorizationPostRequestAllOf) GetPermissionsOk() (*[]Permission, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *LegacyAuthorizationPostRequestAllOf) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []Permission and assigns it to the Permissions field.
func (o *LegacyAuthorizationPostRequestAllOf) SetPermissions(v []Permission) {
	o.Permissions = &v
}

func (o LegacyAuthorizationPostRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OrgID != nil {
		toSerialize["orgID"] = o.OrgID
	}
	if o.UserID != nil {
		toSerialize["userID"] = o.UserID
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableLegacyAuthorizationPostRequestAllOf struct {
	value *LegacyAuthorizationPostRequestAllOf
	isSet bool
}

func (v NullableLegacyAuthorizationPostRequestAllOf) Get() *LegacyAuthorizationPostRequestAllOf {
	return v.value
}

func (v *NullableLegacyAuthorizationPostRequestAllOf) Set(val *LegacyAuthorizationPostRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacyAuthorizationPostRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacyAuthorizationPostRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacyAuthorizationPostRequestAllOf(val *LegacyAuthorizationPostRequestAllOf) *NullableLegacyAuthorizationPostRequestAllOf {
	return &NullableLegacyAuthorizationPostRequestAllOf{value: val, isSet: true}
}

func (v NullableLegacyAuthorizationPostRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacyAuthorizationPostRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}