/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"time"
)

// checks if the OrganizationInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationInput{}

// OrganizationInput struct for OrganizationInput
type OrganizationInput struct {
	Address        *Address               `json:"address,omitempty"`
	BillingAddress *Address               `json:"billing_address,omitempty"`
	Customdata     map[string]interface{} `json:"customdata,omitempty"`
	Description    *string                `json:"description,omitempty"`
	// Force to all members to have enabled the two factor authentication after that date, unless the value is null
	Enforce2faAt *time.Time `json:"enforce_2fa_at,omitempty"`
	// The logo for the organization; must be base64-encoded image data
	Logo                 *string `json:"logo,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Twitter              *string `json:"twitter,omitempty"`
	Website              *string `json:"website,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationInput OrganizationInput

// NewOrganizationInput instantiates a new OrganizationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInput() *OrganizationInput {
	this := OrganizationInput{}
	return &this
}

// NewOrganizationInputWithDefaults instantiates a new OrganizationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInputWithDefaults() *OrganizationInput {
	this := OrganizationInput{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrganizationInput) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrganizationInput) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *OrganizationInput) SetAddress(v Address) {
	o.Address = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *OrganizationInput) GetBillingAddress() Address {
	if o == nil || IsNil(o.BillingAddress) {
		var ret Address
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetBillingAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *OrganizationInput) HasBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given Address and assigns it to the BillingAddress field.
func (o *OrganizationInput) SetBillingAddress(v Address) {
	o.BillingAddress = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *OrganizationInput) GetCustomdata() map[string]interface{} {
	if o == nil || IsNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *OrganizationInput) HasCustomdata() bool {
	if o != nil && !IsNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *OrganizationInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationInput) SetDescription(v string) {
	o.Description = &v
}

// GetEnforce2faAt returns the Enforce2faAt field value if set, zero value otherwise.
func (o *OrganizationInput) GetEnforce2faAt() time.Time {
	if o == nil || IsNil(o.Enforce2faAt) {
		var ret time.Time
		return ret
	}
	return *o.Enforce2faAt
}

// GetEnforce2faAtOk returns a tuple with the Enforce2faAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetEnforce2faAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Enforce2faAt) {
		return nil, false
	}
	return o.Enforce2faAt, true
}

// HasEnforce2faAt returns a boolean if a field has been set.
func (o *OrganizationInput) HasEnforce2faAt() bool {
	if o != nil && !IsNil(o.Enforce2faAt) {
		return true
	}

	return false
}

// SetEnforce2faAt gets a reference to the given time.Time and assigns it to the Enforce2faAt field.
func (o *OrganizationInput) SetEnforce2faAt(v time.Time) {
	o.Enforce2faAt = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *OrganizationInput) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *OrganizationInput) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *OrganizationInput) SetLogo(v string) {
	o.Logo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationInput) SetName(v string) {
	o.Name = &v
}

// GetTwitter returns the Twitter field value if set, zero value otherwise.
func (o *OrganizationInput) GetTwitter() string {
	if o == nil || IsNil(o.Twitter) {
		var ret string
		return ret
	}
	return *o.Twitter
}

// GetTwitterOk returns a tuple with the Twitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetTwitterOk() (*string, bool) {
	if o == nil || IsNil(o.Twitter) {
		return nil, false
	}
	return o.Twitter, true
}

// HasTwitter returns a boolean if a field has been set.
func (o *OrganizationInput) HasTwitter() bool {
	if o != nil && !IsNil(o.Twitter) {
		return true
	}

	return false
}

// SetTwitter gets a reference to the given string and assigns it to the Twitter field.
func (o *OrganizationInput) SetTwitter(v string) {
	o.Twitter = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *OrganizationInput) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInput) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *OrganizationInput) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *OrganizationInput) SetWebsite(v string) {
	o.Website = &v
}

func (o OrganizationInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.BillingAddress) {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if !IsNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enforce2faAt) {
		toSerialize["enforce_2fa_at"] = o.Enforce2faAt
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Twitter) {
		toSerialize["twitter"] = o.Twitter
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationInput) UnmarshalJSON(data []byte) (err error) {
	varOrganizationInput := _OrganizationInput{}

	err = json.Unmarshal(data, &varOrganizationInput)

	if err != nil {
		return err
	}

	*o = OrganizationInput(varOrganizationInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "billing_address")
		delete(additionalProperties, "customdata")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enforce_2fa_at")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "name")
		delete(additionalProperties, "twitter")
		delete(additionalProperties, "website")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationInput struct {
	value *OrganizationInput
	isSet bool
}

func (v NullableOrganizationInput) Get() *OrganizationInput {
	return v.value
}

func (v *NullableOrganizationInput) Set(val *OrganizationInput) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInput) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInput(val *OrganizationInput) *NullableOrganizationInput {
	return &NullableOrganizationInput{value: val, isSet: true}
}

func (v NullableOrganizationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
