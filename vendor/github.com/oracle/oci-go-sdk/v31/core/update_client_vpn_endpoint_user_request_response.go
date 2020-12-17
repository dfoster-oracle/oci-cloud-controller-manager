// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/v31/common"
	"net/http"
)

// UpdateClientVpnEndpointUserRequest wrapper for the UpdateClientVpnEndpointUser operation
type UpdateClientVpnEndpointUserRequest struct {

	// The OCID of the ClientVpnEndpoint.
	ClientVpnEndpointId *string `mandatory:"true" contributesTo:"path" name:"clientVpnEndpointId"`

	// The username of the ClientVpnEndpointUser.
	UserName *string `mandatory:"true" contributesTo:"path" name:"userName"`

	// Detail for updating `ClientVpnEndpointUser`.
	UpdateClientVpnEndpointUserDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource. The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateClientVpnEndpointUserRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateClientVpnEndpointUserRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateClientVpnEndpointUserRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateClientVpnEndpointUserResponse wrapper for the UpdateClientVpnEndpointUser operation
type UpdateClientVpnEndpointUserResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ClientVpnEndpointUser instance
	ClientVpnEndpointUser `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateClientVpnEndpointUserResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateClientVpnEndpointUserResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
