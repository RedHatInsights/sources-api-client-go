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
	"time"
)

// Authentication struct for Authentication
type Authentication struct {
	Authtype *string `json:"authtype,omitempty"`
	AvailabilityStatus *string `json:"availability_status,omitempty"`
	AvailabilityStatusError *string `json:"availability_status_error,omitempty"`
	Extra *AuthenticationExtra `json:"extra,omitempty"`
	// ID of the resource
	Id *string `json:"id,omitempty"`
	LastAvailableAt *time.Time `json:"last_available_at,omitempty"`
	LastCheckedAt *time.Time `json:"last_checked_at,omitempty"`
	Name *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
	// ID of the resource
	ResourceId *string `json:"resource_id,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
	// ID of the resource
	SourceId *string `json:"source_id,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusDetails *string `json:"status_details,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewAuthentication instantiates a new Authentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthentication() *Authentication {
	this := Authentication{}
	return &this
}

// NewAuthenticationWithDefaults instantiates a new Authentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationWithDefaults() *Authentication {
	this := Authentication{}
	return &this
}

// GetAuthtype returns the Authtype field value if set, zero value otherwise.
func (o *Authentication) GetAuthtype() string {
	if o == nil || o.Authtype == nil {
		var ret string
		return ret
	}
	return *o.Authtype
}

// GetAuthtypeOk returns a tuple with the Authtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetAuthtypeOk() (*string, bool) {
	if o == nil || o.Authtype == nil {
		return nil, false
	}
	return o.Authtype, true
}

// HasAuthtype returns a boolean if a field has been set.
func (o *Authentication) HasAuthtype() bool {
	if o != nil && o.Authtype != nil {
		return true
	}

	return false
}

// SetAuthtype gets a reference to the given string and assigns it to the Authtype field.
func (o *Authentication) SetAuthtype(v string) {
	o.Authtype = &v
}

// GetAvailabilityStatus returns the AvailabilityStatus field value if set, zero value otherwise.
func (o *Authentication) GetAvailabilityStatus() string {
	if o == nil || o.AvailabilityStatus == nil {
		var ret string
		return ret
	}
	return *o.AvailabilityStatus
}

// GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetAvailabilityStatusOk() (*string, bool) {
	if o == nil || o.AvailabilityStatus == nil {
		return nil, false
	}
	return o.AvailabilityStatus, true
}

// HasAvailabilityStatus returns a boolean if a field has been set.
func (o *Authentication) HasAvailabilityStatus() bool {
	if o != nil && o.AvailabilityStatus != nil {
		return true
	}

	return false
}

// SetAvailabilityStatus gets a reference to the given string and assigns it to the AvailabilityStatus field.
func (o *Authentication) SetAvailabilityStatus(v string) {
	o.AvailabilityStatus = &v
}

// GetAvailabilityStatusError returns the AvailabilityStatusError field value if set, zero value otherwise.
func (o *Authentication) GetAvailabilityStatusError() string {
	if o == nil || o.AvailabilityStatusError == nil {
		var ret string
		return ret
	}
	return *o.AvailabilityStatusError
}

// GetAvailabilityStatusErrorOk returns a tuple with the AvailabilityStatusError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetAvailabilityStatusErrorOk() (*string, bool) {
	if o == nil || o.AvailabilityStatusError == nil {
		return nil, false
	}
	return o.AvailabilityStatusError, true
}

// HasAvailabilityStatusError returns a boolean if a field has been set.
func (o *Authentication) HasAvailabilityStatusError() bool {
	if o != nil && o.AvailabilityStatusError != nil {
		return true
	}

	return false
}

// SetAvailabilityStatusError gets a reference to the given string and assigns it to the AvailabilityStatusError field.
func (o *Authentication) SetAvailabilityStatusError(v string) {
	o.AvailabilityStatusError = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *Authentication) GetExtra() AuthenticationExtra {
	if o == nil || o.Extra == nil {
		var ret AuthenticationExtra
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetExtraOk() (*AuthenticationExtra, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *Authentication) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given AuthenticationExtra and assigns it to the Extra field.
func (o *Authentication) SetExtra(v AuthenticationExtra) {
	o.Extra = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Authentication) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Authentication) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Authentication) SetId(v string) {
	o.Id = &v
}

// GetLastAvailableAt returns the LastAvailableAt field value if set, zero value otherwise.
func (o *Authentication) GetLastAvailableAt() time.Time {
	if o == nil || o.LastAvailableAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAvailableAt
}

// GetLastAvailableAtOk returns a tuple with the LastAvailableAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetLastAvailableAtOk() (*time.Time, bool) {
	if o == nil || o.LastAvailableAt == nil {
		return nil, false
	}
	return o.LastAvailableAt, true
}

// HasLastAvailableAt returns a boolean if a field has been set.
func (o *Authentication) HasLastAvailableAt() bool {
	if o != nil && o.LastAvailableAt != nil {
		return true
	}

	return false
}

// SetLastAvailableAt gets a reference to the given time.Time and assigns it to the LastAvailableAt field.
func (o *Authentication) SetLastAvailableAt(v time.Time) {
	o.LastAvailableAt = &v
}

// GetLastCheckedAt returns the LastCheckedAt field value if set, zero value otherwise.
func (o *Authentication) GetLastCheckedAt() time.Time {
	if o == nil || o.LastCheckedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastCheckedAt
}

// GetLastCheckedAtOk returns a tuple with the LastCheckedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetLastCheckedAtOk() (*time.Time, bool) {
	if o == nil || o.LastCheckedAt == nil {
		return nil, false
	}
	return o.LastCheckedAt, true
}

// HasLastCheckedAt returns a boolean if a field has been set.
func (o *Authentication) HasLastCheckedAt() bool {
	if o != nil && o.LastCheckedAt != nil {
		return true
	}

	return false
}

// SetLastCheckedAt gets a reference to the given time.Time and assigns it to the LastCheckedAt field.
func (o *Authentication) SetLastCheckedAt(v time.Time) {
	o.LastCheckedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Authentication) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Authentication) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Authentication) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Authentication) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *Authentication) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *Authentication) SetPassword(v string) {
	o.Password = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *Authentication) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *Authentication) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *Authentication) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *Authentication) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *Authentication) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *Authentication) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *Authentication) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *Authentication) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *Authentication) SetSourceId(v string) {
	o.SourceId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Authentication) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Authentication) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Authentication) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise.
func (o *Authentication) GetStatusDetails() string {
	if o == nil || o.StatusDetails == nil {
		var ret string
		return ret
	}
	return *o.StatusDetails
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetStatusDetailsOk() (*string, bool) {
	if o == nil || o.StatusDetails == nil {
		return nil, false
	}
	return o.StatusDetails, true
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *Authentication) HasStatusDetails() bool {
	if o != nil && o.StatusDetails != nil {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given string and assigns it to the StatusDetails field.
func (o *Authentication) SetStatusDetails(v string) {
	o.StatusDetails = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Authentication) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Authentication) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Authentication) SetUsername(v string) {
	o.Username = &v
}

func (o Authentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Authtype != nil {
		toSerialize["authtype"] = o.Authtype
	}
	if o.AvailabilityStatus != nil {
		toSerialize["availability_status"] = o.AvailabilityStatus
	}
	if o.AvailabilityStatusError != nil {
		toSerialize["availability_status_error"] = o.AvailabilityStatusError
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastAvailableAt != nil {
		toSerialize["last_available_at"] = o.LastAvailableAt
	}
	if o.LastCheckedAt != nil {
		toSerialize["last_checked_at"] = o.LastCheckedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.ResourceId != nil {
		toSerialize["resource_id"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.SourceId != nil {
		toSerialize["source_id"] = o.SourceId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusDetails != nil {
		toSerialize["status_details"] = o.StatusDetails
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableAuthentication struct {
	value *Authentication
	isSet bool
}

func (v NullableAuthentication) Get() *Authentication {
	return v.value
}

func (v *NullableAuthentication) Set(val *Authentication) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthentication(val *Authentication) *NullableAuthentication {
	return &NullableAuthentication{value: val, isSet: true}
}

func (v NullableAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


