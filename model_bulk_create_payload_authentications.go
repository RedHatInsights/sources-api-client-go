/*
 * Sources
 *
 * Sources
 *
 * API version: 3.1.0
 * Contact: support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sourcesapi

import (
	"encoding/json"
)

// BulkCreatePayloadAuthentications struct for BulkCreatePayloadAuthentications
type BulkCreatePayloadAuthentications struct {
	Authtype *string `json:"authtype,omitempty"`
	Extra *map[string]interface{} `json:"extra,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	ResourceName *string `json:"resource_name,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
}

// NewBulkCreatePayloadAuthentications instantiates a new BulkCreatePayloadAuthentications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkCreatePayloadAuthentications() *BulkCreatePayloadAuthentications {
	this := BulkCreatePayloadAuthentications{}
	return &this
}

// NewBulkCreatePayloadAuthenticationsWithDefaults instantiates a new BulkCreatePayloadAuthentications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkCreatePayloadAuthenticationsWithDefaults() *BulkCreatePayloadAuthentications {
	this := BulkCreatePayloadAuthentications{}
	return &this
}

// GetAuthtype returns the Authtype field value if set, zero value otherwise.
func (o *BulkCreatePayloadAuthentications) GetAuthtype() string {
	if o == nil || o.Authtype == nil {
		var ret string
		return ret
	}
	return *o.Authtype
}

// GetAuthtypeOk returns a tuple with the Authtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreatePayloadAuthentications) GetAuthtypeOk() (*string, bool) {
	if o == nil || o.Authtype == nil {
		return nil, false
	}
	return o.Authtype, true
}

// HasAuthtype returns a boolean if a field has been set.
func (o *BulkCreatePayloadAuthentications) HasAuthtype() bool {
	if o != nil && o.Authtype != nil {
		return true
	}

	return false
}

// SetAuthtype gets a reference to the given string and assigns it to the Authtype field.
func (o *BulkCreatePayloadAuthentications) SetAuthtype(v string) {
	o.Authtype = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *BulkCreatePayloadAuthentications) GetExtra() map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreatePayloadAuthentications) GetExtraOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *BulkCreatePayloadAuthentications) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *BulkCreatePayloadAuthentications) SetExtra(v map[string]interface{}) {
	o.Extra = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *BulkCreatePayloadAuthentications) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreatePayloadAuthentications) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *BulkCreatePayloadAuthentications) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *BulkCreatePayloadAuthentications) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *BulkCreatePayloadAuthentications) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreatePayloadAuthentications) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *BulkCreatePayloadAuthentications) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *BulkCreatePayloadAuthentications) SetPassword(v string) {
	o.Password = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *BulkCreatePayloadAuthentications) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreatePayloadAuthentications) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *BulkCreatePayloadAuthentications) HasResourceName() bool {
	if o != nil && o.ResourceName != nil {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *BulkCreatePayloadAuthentications) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *BulkCreatePayloadAuthentications) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCreatePayloadAuthentications) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *BulkCreatePayloadAuthentications) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *BulkCreatePayloadAuthentications) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o BulkCreatePayloadAuthentications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Authtype != nil {
		toSerialize["authtype"] = o.Authtype
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	return json.Marshal(toSerialize)
}

type NullableBulkCreatePayloadAuthentications struct {
	value *BulkCreatePayloadAuthentications
	isSet bool
}

func (v NullableBulkCreatePayloadAuthentications) Get() *BulkCreatePayloadAuthentications {
	return v.value
}

func (v *NullableBulkCreatePayloadAuthentications) Set(val *BulkCreatePayloadAuthentications) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkCreatePayloadAuthentications) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkCreatePayloadAuthentications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkCreatePayloadAuthentications(val *BulkCreatePayloadAuthentications) *NullableBulkCreatePayloadAuthentications {
	return &NullableBulkCreatePayloadAuthentications{value: val, isSet: true}
}

func (v NullableBulkCreatePayloadAuthentications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkCreatePayloadAuthentications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


