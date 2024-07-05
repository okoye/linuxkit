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
	"fmt"
)

// PortVlanAssignmentBatchState the model 'PortVlanAssignmentBatchState'
type PortVlanAssignmentBatchState string

// List of PortVlanAssignmentBatch_state
const (
	PORTVLANASSIGNMENTBATCHSTATE_QUEUED      PortVlanAssignmentBatchState = "queued"
	PORTVLANASSIGNMENTBATCHSTATE_IN_PROGRESS PortVlanAssignmentBatchState = "in_progress"
	PORTVLANASSIGNMENTBATCHSTATE_COMPLETED   PortVlanAssignmentBatchState = "completed"
	PORTVLANASSIGNMENTBATCHSTATE_FAILED      PortVlanAssignmentBatchState = "failed"
)

// All allowed values of PortVlanAssignmentBatchState enum
var AllowedPortVlanAssignmentBatchStateEnumValues = []PortVlanAssignmentBatchState{
	"queued",
	"in_progress",
	"completed",
	"failed",
}

func (v *PortVlanAssignmentBatchState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortVlanAssignmentBatchState(value)
	for _, existing := range AllowedPortVlanAssignmentBatchStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortVlanAssignmentBatchState", value)
}

// NewPortVlanAssignmentBatchStateFromValue returns a pointer to a valid PortVlanAssignmentBatchState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortVlanAssignmentBatchStateFromValue(v string) (*PortVlanAssignmentBatchState, error) {
	ev := PortVlanAssignmentBatchState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortVlanAssignmentBatchState: valid values are %v", v, AllowedPortVlanAssignmentBatchStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortVlanAssignmentBatchState) IsValid() bool {
	for _, existing := range AllowedPortVlanAssignmentBatchStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortVlanAssignmentBatch_state value
func (v PortVlanAssignmentBatchState) Ptr() *PortVlanAssignmentBatchState {
	return &v
}

type NullablePortVlanAssignmentBatchState struct {
	value *PortVlanAssignmentBatchState
	isSet bool
}

func (v NullablePortVlanAssignmentBatchState) Get() *PortVlanAssignmentBatchState {
	return v.value
}

func (v *NullablePortVlanAssignmentBatchState) Set(val *PortVlanAssignmentBatchState) {
	v.value = val
	v.isSet = true
}

func (v NullablePortVlanAssignmentBatchState) IsSet() bool {
	return v.isSet
}

func (v *NullablePortVlanAssignmentBatchState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortVlanAssignmentBatchState(val *PortVlanAssignmentBatchState) *NullablePortVlanAssignmentBatchState {
	return &NullablePortVlanAssignmentBatchState{value: val, isSet: true}
}

func (v NullablePortVlanAssignmentBatchState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortVlanAssignmentBatchState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
