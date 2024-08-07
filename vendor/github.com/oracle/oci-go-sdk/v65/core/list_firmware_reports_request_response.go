// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFirmwareReportsRequest wrapper for the ListFirmwareReports operation
type ListFirmwareReportsRequest struct {

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" contributesTo:"query" name:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// FirmwareReport lifecycle state.
	LifecycleState FirmwareReportLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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
	SortBy ListFirmwareReportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListFirmwareReportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFirmwareReportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFirmwareReportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFirmwareReportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request ListFirmwareReportsRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
	if mandatoryParamMap["availabilityDomain"] != nil {
		templateParam := mandatoryParamMap["availabilityDomain"]
		for _, template := range templateParam {
			replacementParam := *request.AvailabilityDomain
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
	if mandatoryParamMap["compartmentId"] != nil {
		templateParam := mandatoryParamMap["compartmentId"]
		for _, template := range templateParam {
			replacementParam := *request.CompartmentId
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFirmwareReportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFirmwareReportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFirmwareReportLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetFirmwareReportLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFirmwareReportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFirmwareReportsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFirmwareReportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFirmwareReportsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFirmwareReportsResponse wrapper for the ListFirmwareReports operation
type ListFirmwareReportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FirmwareReportCollection instances
	FirmwareReportCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListFirmwareReportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFirmwareReportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFirmwareReportsSortByEnum Enum with underlying type: string
type ListFirmwareReportsSortByEnum string

// Set of constants representing the allowable values for ListFirmwareReportsSortByEnum
const (
	ListFirmwareReportsSortByTimecreated ListFirmwareReportsSortByEnum = "TIMECREATED"
	ListFirmwareReportsSortByDisplayname ListFirmwareReportsSortByEnum = "DISPLAYNAME"
)

var mappingListFirmwareReportsSortByEnum = map[string]ListFirmwareReportsSortByEnum{
	"TIMECREATED": ListFirmwareReportsSortByTimecreated,
	"DISPLAYNAME": ListFirmwareReportsSortByDisplayname,
}

var mappingListFirmwareReportsSortByEnumLowerCase = map[string]ListFirmwareReportsSortByEnum{
	"timecreated": ListFirmwareReportsSortByTimecreated,
	"displayname": ListFirmwareReportsSortByDisplayname,
}

// GetListFirmwareReportsSortByEnumValues Enumerates the set of values for ListFirmwareReportsSortByEnum
func GetListFirmwareReportsSortByEnumValues() []ListFirmwareReportsSortByEnum {
	values := make([]ListFirmwareReportsSortByEnum, 0)
	for _, v := range mappingListFirmwareReportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFirmwareReportsSortByEnumStringValues Enumerates the set of values in String for ListFirmwareReportsSortByEnum
func GetListFirmwareReportsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListFirmwareReportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFirmwareReportsSortByEnum(val string) (ListFirmwareReportsSortByEnum, bool) {
	enum, ok := mappingListFirmwareReportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFirmwareReportsSortOrderEnum Enum with underlying type: string
type ListFirmwareReportsSortOrderEnum string

// Set of constants representing the allowable values for ListFirmwareReportsSortOrderEnum
const (
	ListFirmwareReportsSortOrderAsc  ListFirmwareReportsSortOrderEnum = "ASC"
	ListFirmwareReportsSortOrderDesc ListFirmwareReportsSortOrderEnum = "DESC"
)

var mappingListFirmwareReportsSortOrderEnum = map[string]ListFirmwareReportsSortOrderEnum{
	"ASC":  ListFirmwareReportsSortOrderAsc,
	"DESC": ListFirmwareReportsSortOrderDesc,
}

var mappingListFirmwareReportsSortOrderEnumLowerCase = map[string]ListFirmwareReportsSortOrderEnum{
	"asc":  ListFirmwareReportsSortOrderAsc,
	"desc": ListFirmwareReportsSortOrderDesc,
}

// GetListFirmwareReportsSortOrderEnumValues Enumerates the set of values for ListFirmwareReportsSortOrderEnum
func GetListFirmwareReportsSortOrderEnumValues() []ListFirmwareReportsSortOrderEnum {
	values := make([]ListFirmwareReportsSortOrderEnum, 0)
	for _, v := range mappingListFirmwareReportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFirmwareReportsSortOrderEnumStringValues Enumerates the set of values in String for ListFirmwareReportsSortOrderEnum
func GetListFirmwareReportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFirmwareReportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFirmwareReportsSortOrderEnum(val string) (ListFirmwareReportsSortOrderEnum, bool) {
	enum, ok := mappingListFirmwareReportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
