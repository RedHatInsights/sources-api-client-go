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

// ApplicationTypesCollection struct for ApplicationTypesCollection
type ApplicationTypesCollection struct {
	Meta *CollectionMetadata `json:"meta,omitempty"`
	Links *CollectionLinks `json:"links,omitempty"`
	Data *[]ApplicationType `json:"data,omitempty"`
}

// NewApplicationTypesCollection instantiates a new ApplicationTypesCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationTypesCollection() *ApplicationTypesCollection {
	this := ApplicationTypesCollection{}
	return &this
}

// NewApplicationTypesCollectionWithDefaults instantiates a new ApplicationTypesCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationTypesCollectionWithDefaults() *ApplicationTypesCollection {
	this := ApplicationTypesCollection{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ApplicationTypesCollection) GetMeta() CollectionMetadata {
	if o == nil || o.Meta == nil {
		var ret CollectionMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTypesCollection) GetMetaOk() (*CollectionMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ApplicationTypesCollection) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CollectionMetadata and assigns it to the Meta field.
func (o *ApplicationTypesCollection) SetMeta(v CollectionMetadata) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApplicationTypesCollection) GetLinks() CollectionLinks {
	if o == nil || o.Links == nil {
		var ret CollectionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTypesCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApplicationTypesCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CollectionLinks and assigns it to the Links field.
func (o *ApplicationTypesCollection) SetLinks(v CollectionLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApplicationTypesCollection) GetData() []ApplicationType {
	if o == nil || o.Data == nil {
		var ret []ApplicationType
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationTypesCollection) GetDataOk() (*[]ApplicationType, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApplicationTypesCollection) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ApplicationType and assigns it to the Data field.
func (o *ApplicationTypesCollection) SetData(v []ApplicationType) {
	o.Data = &v
}

func (o ApplicationTypesCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationTypesCollection struct {
	value *ApplicationTypesCollection
	isSet bool
}

func (v NullableApplicationTypesCollection) Get() *ApplicationTypesCollection {
	return v.value
}

func (v *NullableApplicationTypesCollection) Set(val *ApplicationTypesCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationTypesCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationTypesCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationTypesCollection(val *ApplicationTypesCollection) *NullableApplicationTypesCollection {
	return &NullableApplicationTypesCollection{value: val, isSet: true}
}

func (v NullableApplicationTypesCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationTypesCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


