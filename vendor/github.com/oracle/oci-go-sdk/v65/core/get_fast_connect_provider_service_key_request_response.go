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

// GetFastConnectProviderServiceKeyRequest wrapper for the GetFastConnectProviderServiceKey operation
type GetFastConnectProviderServiceKeyRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the provider service.
	ProviderServiceId *string `mandatory:"true" contributesTo:"path" name:"providerServiceId"`

	// The provider service key that the provider gives you when you set up a virtual circuit connection
	// from the provider to Oracle Cloud Infrastructure. You can set up that connection and get your
	// provider service key at the provider's website or portal. For the portal location, see the `description`
	// attribute of the FastConnectProviderService.
	ProviderServiceKeyName *string `mandatory:"true" contributesTo:"path" name:"providerServiceKeyName"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetFastConnectProviderServiceKeyRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetFastConnectProviderServiceKeyRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetFastConnectProviderServiceKeyRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request GetFastConnectProviderServiceKeyRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
	if mandatoryParamMap["providerServiceId"] != nil {
		templateParam := mandatoryParamMap["providerServiceId"]
		for _, template := range templateParam {
			replacementParam := *request.ProviderServiceId
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
	if mandatoryParamMap["providerServiceKeyName"] != nil {
		templateParam := mandatoryParamMap["providerServiceKeyName"]
		for _, template := range templateParam {
			replacementParam := *request.ProviderServiceKeyName
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetFastConnectProviderServiceKeyRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetFastConnectProviderServiceKeyRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetFastConnectProviderServiceKeyResponse wrapper for the GetFastConnectProviderServiceKey operation
type GetFastConnectProviderServiceKeyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The FastConnectProviderServiceKey instance
	FastConnectProviderServiceKey `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetFastConnectProviderServiceKeyResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetFastConnectProviderServiceKeyResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
