// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPrivateEndpointAssociationsRequest wrapper for the ListPrivateEndpointAssociations operation
type ListPrivateEndpointAssociationsRequest struct {

	// The endpoint service's OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	EndpointServiceId *string `mandatory:"true" contributesTo:"path" name:"endpointServiceId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListPrivateEndpointAssociationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListPrivateEndpointAssociationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPrivateEndpointAssociationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPrivateEndpointAssociationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPrivateEndpointAssociationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPrivateEndpointAssociationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPrivateEndpointAssociationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPrivateEndpointAssociationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPrivateEndpointAssociationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPrivateEndpointAssociationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPrivateEndpointAssociationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPrivateEndpointAssociationsResponse wrapper for the ListPrivateEndpointAssociations operation
type ListPrivateEndpointAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []PrivateEndpointAssociation instances
	Items []PrivateEndpointAssociation `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListPrivateEndpointAssociationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPrivateEndpointAssociationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPrivateEndpointAssociationsSortByEnum Enum with underlying type: string
type ListPrivateEndpointAssociationsSortByEnum string

// Set of constants representing the allowable values for ListPrivateEndpointAssociationsSortByEnum
const (
	ListPrivateEndpointAssociationsSortByTimecreated ListPrivateEndpointAssociationsSortByEnum = "TIMECREATED"
	ListPrivateEndpointAssociationsSortByDisplayname ListPrivateEndpointAssociationsSortByEnum = "DISPLAYNAME"
)

var mappingListPrivateEndpointAssociationsSortByEnum = map[string]ListPrivateEndpointAssociationsSortByEnum{
	"TIMECREATED": ListPrivateEndpointAssociationsSortByTimecreated,
	"DISPLAYNAME": ListPrivateEndpointAssociationsSortByDisplayname,
}

var mappingListPrivateEndpointAssociationsSortByEnumLowerCase = map[string]ListPrivateEndpointAssociationsSortByEnum{
	"timecreated": ListPrivateEndpointAssociationsSortByTimecreated,
	"displayname": ListPrivateEndpointAssociationsSortByDisplayname,
}

// GetListPrivateEndpointAssociationsSortByEnumValues Enumerates the set of values for ListPrivateEndpointAssociationsSortByEnum
func GetListPrivateEndpointAssociationsSortByEnumValues() []ListPrivateEndpointAssociationsSortByEnum {
	values := make([]ListPrivateEndpointAssociationsSortByEnum, 0)
	for _, v := range mappingListPrivateEndpointAssociationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivateEndpointAssociationsSortByEnumStringValues Enumerates the set of values in String for ListPrivateEndpointAssociationsSortByEnum
func GetListPrivateEndpointAssociationsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListPrivateEndpointAssociationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivateEndpointAssociationsSortByEnum(val string) (ListPrivateEndpointAssociationsSortByEnum, bool) {
	enum, ok := mappingListPrivateEndpointAssociationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPrivateEndpointAssociationsSortOrderEnum Enum with underlying type: string
type ListPrivateEndpointAssociationsSortOrderEnum string

// Set of constants representing the allowable values for ListPrivateEndpointAssociationsSortOrderEnum
const (
	ListPrivateEndpointAssociationsSortOrderAsc  ListPrivateEndpointAssociationsSortOrderEnum = "ASC"
	ListPrivateEndpointAssociationsSortOrderDesc ListPrivateEndpointAssociationsSortOrderEnum = "DESC"
)

var mappingListPrivateEndpointAssociationsSortOrderEnum = map[string]ListPrivateEndpointAssociationsSortOrderEnum{
	"ASC":  ListPrivateEndpointAssociationsSortOrderAsc,
	"DESC": ListPrivateEndpointAssociationsSortOrderDesc,
}

var mappingListPrivateEndpointAssociationsSortOrderEnumLowerCase = map[string]ListPrivateEndpointAssociationsSortOrderEnum{
	"asc":  ListPrivateEndpointAssociationsSortOrderAsc,
	"desc": ListPrivateEndpointAssociationsSortOrderDesc,
}

// GetListPrivateEndpointAssociationsSortOrderEnumValues Enumerates the set of values for ListPrivateEndpointAssociationsSortOrderEnum
func GetListPrivateEndpointAssociationsSortOrderEnumValues() []ListPrivateEndpointAssociationsSortOrderEnum {
	values := make([]ListPrivateEndpointAssociationsSortOrderEnum, 0)
	for _, v := range mappingListPrivateEndpointAssociationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPrivateEndpointAssociationsSortOrderEnumStringValues Enumerates the set of values in String for ListPrivateEndpointAssociationsSortOrderEnum
func GetListPrivateEndpointAssociationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPrivateEndpointAssociationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPrivateEndpointAssociationsSortOrderEnum(val string) (ListPrivateEndpointAssociationsSortOrderEnum, bool) {
	enum, ok := mappingListPrivateEndpointAssociationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
