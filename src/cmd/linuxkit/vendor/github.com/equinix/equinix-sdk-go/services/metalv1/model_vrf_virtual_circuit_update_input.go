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
)

// checks if the VrfVirtualCircuitUpdateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VrfVirtualCircuitUpdateInput{}

// VrfVirtualCircuitUpdateInput struct for VrfVirtualCircuitUpdateInput
type VrfVirtualCircuitUpdateInput struct {
	// An IPv4 address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used.
	CustomerIp  *string `json:"customer_ip,omitempty"`
	Description *string `json:"description,omitempty"`
	// The plaintext BGP peering password shared by neighbors as an MD5 checksum: * must be 10-20 characters long * may not include punctuation * must be a combination of numbers and letters * must contain at least one lowercase, uppercase, and digit character
	Md5 *string `json:"md5,omitempty"`
	// An IPv4 address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used.
	MetalIp *string `json:"metal_ip,omitempty"`
	Name    *string `json:"name,omitempty"`
	// The peer ASN that will be used with the VRF on the Virtual Circuit.
	PeerAsn *int64 `json:"peer_asn,omitempty"`
	// Speed can be changed only if it is an interconnection on a Dedicated Port
	Speed *string `json:"speed,omitempty"`
	// The /30 or /31 IPv4 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP.
	Subnet *string `json:"subnet,omitempty"`
	// The /126 or /127 IPv6 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IPv6 and Customer IPv6 must be IPs from this subnet. For /126 subnets, the network and broadcast IPs cannot be used as the Metal IPv6 or Customer IPv6. The subnet specified must be contained within an already-defined IP Range for the VRF.
	SubnetIpv6 *string `json:"subnet_ipv6,omitempty"`
	// An IPv6 address from the subnet IPv6 that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet IPv6 as the Metal IPv6. By default, the last usable IP address in the subnet IPv6 will be used.
	CustomerIpv6 *string `json:"customer_ipv6,omitempty"`
	// An IPv6 address from the subnet IPv6 that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IPv6 address in the subnet IPv6 as the Customer IP. By default, the first usable IPv6 address in the subnet IPv6 will be used.
	MetalIpv6            *string  `json:"metal_ipv6,omitempty"`
	Tags                 []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VrfVirtualCircuitUpdateInput VrfVirtualCircuitUpdateInput

// NewVrfVirtualCircuitUpdateInput instantiates a new VrfVirtualCircuitUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVrfVirtualCircuitUpdateInput() *VrfVirtualCircuitUpdateInput {
	this := VrfVirtualCircuitUpdateInput{}
	return &this
}

// NewVrfVirtualCircuitUpdateInputWithDefaults instantiates a new VrfVirtualCircuitUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVrfVirtualCircuitUpdateInputWithDefaults() *VrfVirtualCircuitUpdateInput {
	this := VrfVirtualCircuitUpdateInput{}
	return &this
}

// GetCustomerIp returns the CustomerIp field value if set, zero value otherwise.
func (o *VrfVirtualCircuitUpdateInput) GetCustomerIp() string {
	if o == nil || IsNil(o.CustomerIp) {
		var ret string
		return ret
	}
	return *o.CustomerIp
}

// GetCustomerIpOk returns a tuple with the CustomerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitUpdateInput) GetCustomerIpOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerIp) {
		return nil, false
	}
	return o.CustomerIp, true
}

// HasCustomerIp returns a boolean if a field has been set.
func (o *VrfVirtualCircuitUpdateInput) HasCustomerIp() bool {
	if o != nil && !IsNil(o.CustomerIp) {
		return true
	}

	return false
}

// SetCustomerIp gets a reference to the given string and assigns it to the CustomerIp field.
func (o *VrfVirtualCircuitUpdateInput) SetCustomerIp(v string) {
	o.CustomerIp = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VrfVirtualCircuitUpdateInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitUpdateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VrfVirtualCircuitUpdateInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VrfVirtualCircuitUpdateInput) SetDescription(v string) {
	o.Description = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *VrfVirtualCircuitUpdateInput) GetMd5() string {
	if o == nil || IsNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitUpdateInput) GetMd5Ok() (*string, bool) {
	if o == nil || IsNil(o.Md5) {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *VrfVirtualCircuitUpdateInput) HasMd5() bool {
	if o != nil && !IsNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *VrfVirtualCircuitUpdateInput) SetMd5(v string) {
	o.Md5 = &v
}

// GetMetalIp returns the MetalIp field value if set, zero value otherwise.
func (o *VrfVirtualCircuitUpdateInput) GetMetalIp() string {
	if o == nil || IsNil(o.MetalIp) {
		var ret string
		return ret
	}
	return *o.MetalIp
}

// GetMetalIpOk returns a tuple with the MetalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitUpdateInput) GetMetalIpOk() (*string, bool) {
	if o == nil || IsNil(o.MetalIp) {
		return nil, false
	}
	return o.MetalIp, true
}

// HasMetalIp returns a boolean if a field has been set.
func (o *VrfVirtualCircuitUpdateInput) HasMetalIp() bool {
	if o != nil && !IsNil(o.MetalIp) {
		return true
	}

	return false
}

// SetMetalIp gets a reference to the given string and assigns it to the MetalIp field.
func (o *VrfVirtualCircuitUpdateInput) SetMetalIp(v string) {
	o.MetalIp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VrfVirtualCircuitUpdateInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitUpdateInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VrfVirtualCircuitUpdateInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VrfVirtualCircuitUpdateInput) SetName(v string) {
	o.Name = &v
}

// GetPeerAsn returns the PeerAsn field value if set, zero value otherwise.
func (o *VrfVirtualCircuitUpdateInput) GetPeerAsn() int64 {
	if o == nil || IsNil(o.PeerAsn) {
		var ret int64
		return ret
	}
	return *o.PeerAsn
}

// GetPeerAsnOk returns a tuple with the PeerAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitUpdateInput) GetPeerAsnOk() (*int64, bool) {
	if o == nil || IsNil(o.PeerAsn) {
		return nil, false
	}
	return o.PeerAsn, true
}

// HasPeerAsn returns a boolean if a field has been set.
func (o *VrfVirtualCircuitUpdateInput) HasPeerAsn() bool {
	if o != nil && !IsNil(o.PeerAsn) {
		return true
	}

	return false
}

// SetPeerAsn gets a reference to the given int64 and assigns it to the PeerAsn field.
func (o *VrfVirtualCircuitUpdateInput) SetPeerAsn(v int64) {
	o.PeerAsn = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VrfVirtualCircuitUpdateInput) GetSpeed() string {
	if o == nil || IsNil(o.Speed) {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitUpdateInput) GetSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VrfVirtualCircuitUpdateInput) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *VrfVirtualCircuitUpdateInput) SetSpeed(v string) {
	o.Speed = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *VrfVirtualCircuitUpdateInput) GetSubnet() string {
	if o == nil || IsNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitUpdateInput) GetSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Subnet) {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *VrfVirtualCircuitUpdateInput) HasSubnet() bool {
	if o != nil && !IsNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *VrfVirtualCircuitUpdateInput) SetSubnet(v string) {
	o.Subnet = &v
}

// GetSubnetIpv6 returns the SubnetIpv6 field value if set, zero value otherwise.
func (o *VrfVirtualCircuitUpdateInput) GetSubnetIpv6() string {
	if o == nil || IsNil(o.SubnetIpv6) {
		var ret string
		return ret
	}
	return *o.SubnetIpv6
}

// GetSubnetIpv6Ok returns a tuple with the SubnetIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitUpdateInput) GetSubnetIpv6Ok() (*string, bool) {
	if o == nil || IsNil(o.SubnetIpv6) {
		return nil, false
	}
	return o.SubnetIpv6, true
}

// HasSubnetIpv6 returns a boolean if a field has been set.
func (o *VrfVirtualCircuitUpdateInput) HasSubnetIpv6() bool {
	if o != nil && !IsNil(o.SubnetIpv6) {
		return true
	}

	return false
}

// SetSubnetIpv6 gets a reference to the given string and assigns it to the SubnetIpv6 field.
func (o *VrfVirtualCircuitUpdateInput) SetSubnetIpv6(v string) {
	o.SubnetIpv6 = &v
}

// GetCustomerIpv6 returns the CustomerIpv6 field value if set, zero value otherwise.
func (o *VrfVirtualCircuitUpdateInput) GetCustomerIpv6() string {
	if o == nil || IsNil(o.CustomerIpv6) {
		var ret string
		return ret
	}
	return *o.CustomerIpv6
}

// GetCustomerIpv6Ok returns a tuple with the CustomerIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitUpdateInput) GetCustomerIpv6Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomerIpv6) {
		return nil, false
	}
	return o.CustomerIpv6, true
}

// HasCustomerIpv6 returns a boolean if a field has been set.
func (o *VrfVirtualCircuitUpdateInput) HasCustomerIpv6() bool {
	if o != nil && !IsNil(o.CustomerIpv6) {
		return true
	}

	return false
}

// SetCustomerIpv6 gets a reference to the given string and assigns it to the CustomerIpv6 field.
func (o *VrfVirtualCircuitUpdateInput) SetCustomerIpv6(v string) {
	o.CustomerIpv6 = &v
}

// GetMetalIpv6 returns the MetalIpv6 field value if set, zero value otherwise.
func (o *VrfVirtualCircuitUpdateInput) GetMetalIpv6() string {
	if o == nil || IsNil(o.MetalIpv6) {
		var ret string
		return ret
	}
	return *o.MetalIpv6
}

// GetMetalIpv6Ok returns a tuple with the MetalIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitUpdateInput) GetMetalIpv6Ok() (*string, bool) {
	if o == nil || IsNil(o.MetalIpv6) {
		return nil, false
	}
	return o.MetalIpv6, true
}

// HasMetalIpv6 returns a boolean if a field has been set.
func (o *VrfVirtualCircuitUpdateInput) HasMetalIpv6() bool {
	if o != nil && !IsNil(o.MetalIpv6) {
		return true
	}

	return false
}

// SetMetalIpv6 gets a reference to the given string and assigns it to the MetalIpv6 field.
func (o *VrfVirtualCircuitUpdateInput) SetMetalIpv6(v string) {
	o.MetalIpv6 = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VrfVirtualCircuitUpdateInput) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VrfVirtualCircuitUpdateInput) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VrfVirtualCircuitUpdateInput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VrfVirtualCircuitUpdateInput) SetTags(v []string) {
	o.Tags = v
}

func (o VrfVirtualCircuitUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VrfVirtualCircuitUpdateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerIp) {
		toSerialize["customer_ip"] = o.CustomerIp
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Md5) {
		toSerialize["md5"] = o.Md5
	}
	if !IsNil(o.MetalIp) {
		toSerialize["metal_ip"] = o.MetalIp
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PeerAsn) {
		toSerialize["peer_asn"] = o.PeerAsn
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !IsNil(o.SubnetIpv6) {
		toSerialize["subnet_ipv6"] = o.SubnetIpv6
	}
	if !IsNil(o.CustomerIpv6) {
		toSerialize["customer_ipv6"] = o.CustomerIpv6
	}
	if !IsNil(o.MetalIpv6) {
		toSerialize["metal_ipv6"] = o.MetalIpv6
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VrfVirtualCircuitUpdateInput) UnmarshalJSON(data []byte) (err error) {
	varVrfVirtualCircuitUpdateInput := _VrfVirtualCircuitUpdateInput{}

	err = json.Unmarshal(data, &varVrfVirtualCircuitUpdateInput)

	if err != nil {
		return err
	}

	*o = VrfVirtualCircuitUpdateInput(varVrfVirtualCircuitUpdateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_ip")
		delete(additionalProperties, "description")
		delete(additionalProperties, "md5")
		delete(additionalProperties, "metal_ip")
		delete(additionalProperties, "name")
		delete(additionalProperties, "peer_asn")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "subnet")
		delete(additionalProperties, "subnet_ipv6")
		delete(additionalProperties, "customer_ipv6")
		delete(additionalProperties, "metal_ipv6")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVrfVirtualCircuitUpdateInput struct {
	value *VrfVirtualCircuitUpdateInput
	isSet bool
}

func (v NullableVrfVirtualCircuitUpdateInput) Get() *VrfVirtualCircuitUpdateInput {
	return v.value
}

func (v *NullableVrfVirtualCircuitUpdateInput) Set(val *VrfVirtualCircuitUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfVirtualCircuitUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfVirtualCircuitUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfVirtualCircuitUpdateInput(val *VrfVirtualCircuitUpdateInput) *NullableVrfVirtualCircuitUpdateInput {
	return &NullableVrfVirtualCircuitUpdateInput{value: val, isSet: true}
}

func (v NullableVrfVirtualCircuitUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfVirtualCircuitUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
