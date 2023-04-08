// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//NetworkLoadBalancerClient a client for NetworkLoadBalancer
type NetworkLoadBalancerClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewNetworkLoadBalancerClientWithConfigurationProvider Creates a new default NetworkLoadBalancer client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewNetworkLoadBalancerClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client NetworkLoadBalancerClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newNetworkLoadBalancerClientFromBaseClient(baseClient, provider)
}

// NewNetworkLoadBalancerClientWithOboToken Creates a new default NetworkLoadBalancer client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewNetworkLoadBalancerClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client NetworkLoadBalancerClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newNetworkLoadBalancerClientFromBaseClient(baseClient, configProvider)
}

func newNetworkLoadBalancerClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client NetworkLoadBalancerClient, err error) {
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = NetworkLoadBalancerClient{BaseClient: baseClient}
	client.BasePath = "20200501"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *NetworkLoadBalancerClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("networkloadbalancer", "https://network-load-balancer-api.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *NetworkLoadBalancerClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *NetworkLoadBalancerClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AttachNlbToPod Updates the network load balancer for NLB live migration and creates new mappings with new pod id and slot id
// A default retry strategy applies to this operation AttachNlbToPod()
func (client NetworkLoadBalancerClient) AttachNlbToPod(ctx context.Context, request AttachNlbToPodRequest) (response AttachNlbToPodResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.attachNlbToPod, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AttachNlbToPodResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AttachNlbToPodResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AttachNlbToPodResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AttachNlbToPodResponse")
	}
	return
}

// attachNlbToPod implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) attachNlbToPod(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkLoadBalancers/{networkLoadBalancerId}/actions/attachNlbToPod", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AttachNlbToPodResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancer/AttachNlbToPod"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "AttachNlbToPod", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeNetworkLoadBalancerCompartment Moves a network load balancer into a different compartment within the same tenancy. For information about moving resources
// between compartments, see Moving Resources to a Different Compartment (https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingcompartments.htm#moveRes).
// A default retry strategy applies to this operation ChangeNetworkLoadBalancerCompartment()
func (client NetworkLoadBalancerClient) ChangeNetworkLoadBalancerCompartment(ctx context.Context, request ChangeNetworkLoadBalancerCompartmentRequest) (response ChangeNetworkLoadBalancerCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeNetworkLoadBalancerCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeNetworkLoadBalancerCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeNetworkLoadBalancerCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeNetworkLoadBalancerCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeNetworkLoadBalancerCompartmentResponse")
	}
	return
}

// changeNetworkLoadBalancerCompartment implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) changeNetworkLoadBalancerCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkLoadBalancers/{networkLoadBalancerId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeNetworkLoadBalancerCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancer/ChangeNetworkLoadBalancerCompartment"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "ChangeNetworkLoadBalancerCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBackend Adds a backend server to a backend set.
// A default retry strategy applies to this operation CreateBackend()
func (client NetworkLoadBalancerClient) CreateBackend(ctx context.Context, request CreateBackendRequest) (response CreateBackendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createBackend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBackendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBackendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBackendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBackendResponse")
	}
	return
}

// createBackend implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) createBackend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}/backends", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBackendResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/Backend/CreateBackend"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "CreateBackend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateBackendSet Adds a backend set to a network load balancer.
// A default retry strategy applies to this operation CreateBackendSet()
func (client NetworkLoadBalancerClient) CreateBackendSet(ctx context.Context, request CreateBackendSetRequest) (response CreateBackendSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createBackendSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateBackendSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateBackendSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateBackendSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateBackendSetResponse")
	}
	return
}

// createBackendSet implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) createBackendSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateBackendSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/BackendSet/CreateBackendSet"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "CreateBackendSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateListener Adds a listener to a network load balancer.
// A default retry strategy applies to this operation CreateListener()
func (client NetworkLoadBalancerClient) CreateListener(ctx context.Context, request CreateListenerRequest) (response CreateListenerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createListener, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateListenerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateListenerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateListenerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateListenerResponse")
	}
	return
}

// createListener implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) createListener(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkLoadBalancers/{networkLoadBalancerId}/listeners", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateListenerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/Listener/CreateListener"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "CreateListener", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateNetworkLoadBalancer Creates a network load balancer.
// A default retry strategy applies to this operation CreateNetworkLoadBalancer()
func (client NetworkLoadBalancerClient) CreateNetworkLoadBalancer(ctx context.Context, request CreateNetworkLoadBalancerRequest) (response CreateNetworkLoadBalancerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createNetworkLoadBalancer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateNetworkLoadBalancerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateNetworkLoadBalancerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateNetworkLoadBalancerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateNetworkLoadBalancerResponse")
	}
	return
}

// createNetworkLoadBalancer implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) createNetworkLoadBalancer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkLoadBalancers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateNetworkLoadBalancerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancer/CreateNetworkLoadBalancer"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "CreateNetworkLoadBalancer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBackend Removes a backend server from a given network load balancer and backend set.
// A default retry strategy applies to this operation DeleteBackend()
func (client NetworkLoadBalancerClient) DeleteBackend(ctx context.Context, request DeleteBackendRequest) (response DeleteBackendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBackend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBackendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBackendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBackendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBackendResponse")
	}
	return
}

// deleteBackend implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) deleteBackend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}/backends/{backendName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBackendResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/Backend/DeleteBackend"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "DeleteBackend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteBackendSet Deletes the specified backend set. Note that deleting a backend set removes its backend servers from the network load balancer.
// Before you can delete a backend set, you must remove it from any active listeners.
// A default retry strategy applies to this operation DeleteBackendSet()
func (client NetworkLoadBalancerClient) DeleteBackendSet(ctx context.Context, request DeleteBackendSetRequest) (response DeleteBackendSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteBackendSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteBackendSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteBackendSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteBackendSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteBackendSetResponse")
	}
	return
}

// deleteBackendSet implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) deleteBackendSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteBackendSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/BackendSet/DeleteBackendSet"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "DeleteBackendSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteListener Deletes a listener from a network load balancer.
// A default retry strategy applies to this operation DeleteListener()
func (client NetworkLoadBalancerClient) DeleteListener(ctx context.Context, request DeleteListenerRequest) (response DeleteListenerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteListener, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteListenerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteListenerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteListenerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteListenerResponse")
	}
	return
}

// deleteListener implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) deleteListener(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkLoadBalancers/{networkLoadBalancerId}/listeners/{listenerName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteListenerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/Listener/DeleteListener"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "DeleteListener", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteNetworkLoadBalancer Deletes a network load balancer resource by identifier.
// A default retry strategy applies to this operation DeleteNetworkLoadBalancer()
func (client NetworkLoadBalancerClient) DeleteNetworkLoadBalancer(ctx context.Context, request DeleteNetworkLoadBalancerRequest) (response DeleteNetworkLoadBalancerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteNetworkLoadBalancer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteNetworkLoadBalancerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteNetworkLoadBalancerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteNetworkLoadBalancerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteNetworkLoadBalancerResponse")
	}
	return
}

// deleteNetworkLoadBalancer implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) deleteNetworkLoadBalancer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/networkLoadBalancers/{networkLoadBalancerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteNetworkLoadBalancerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancer/DeleteNetworkLoadBalancer"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "DeleteNetworkLoadBalancer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachNlbFromDestinationPod Update mappings for a provisioned network load balancer for the provided pod Id in case of rollback.
// A default retry strategy applies to this operation DetachNlbFromDestinationPod()
func (client NetworkLoadBalancerClient) DetachNlbFromDestinationPod(ctx context.Context, request DetachNlbFromDestinationPodRequest) (response DetachNlbFromDestinationPodResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.detachNlbFromDestinationPod, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachNlbFromDestinationPodResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachNlbFromDestinationPodResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachNlbFromDestinationPodResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachNlbFromDestinationPodResponse")
	}
	return
}

// detachNlbFromDestinationPod implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) detachNlbFromDestinationPod(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkLoadBalancers/{networkLoadBalancerId}/actions/detachNlbFromDestinationPod", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachNlbFromDestinationPodResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancer/DetachNlbFromDestinationPod"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "DetachNlbFromDestinationPod", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DetachNlbFromSourcePod Delete mappings for a provisioned network load balancer for the provided pod Id.
// A default retry strategy applies to this operation DetachNlbFromSourcePod()
func (client NetworkLoadBalancerClient) DetachNlbFromSourcePod(ctx context.Context, request DetachNlbFromSourcePodRequest) (response DetachNlbFromSourcePodResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.detachNlbFromSourcePod, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DetachNlbFromSourcePodResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DetachNlbFromSourcePodResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DetachNlbFromSourcePodResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DetachNlbFromSourcePodResponse")
	}
	return
}

// detachNlbFromSourcePod implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) detachNlbFromSourcePod(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkLoadBalancers/{networkLoadBalancerId}/actions/detachNlbFromSourcePod", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DetachNlbFromSourcePodResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancer/DetachNlbFromSourcePod"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "DetachNlbFromSourcePod", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBackend Retrieves the configuration information for the specified backend server.
// A default retry strategy applies to this operation GetBackend()
func (client NetworkLoadBalancerClient) GetBackend(ctx context.Context, request GetBackendRequest) (response GetBackendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBackend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBackendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBackendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBackendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBackendResponse")
	}
	return
}

// getBackend implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) getBackend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}/backends/{backendName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBackendResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/Backend/GetBackend"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "GetBackend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBackendHealth Retrieves the current health status of the specified backend server.
// A default retry strategy applies to this operation GetBackendHealth()
func (client NetworkLoadBalancerClient) GetBackendHealth(ctx context.Context, request GetBackendHealthRequest) (response GetBackendHealthResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBackendHealth, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBackendHealthResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBackendHealthResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBackendHealthResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBackendHealthResponse")
	}
	return
}

// getBackendHealth implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) getBackendHealth(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}/backends/{backendName}/health", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBackendHealthResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/BackendHealth/GetBackendHealth"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "GetBackendHealth", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBackendSet Retrieves the configuration information for the specified backend set.
// A default retry strategy applies to this operation GetBackendSet()
func (client NetworkLoadBalancerClient) GetBackendSet(ctx context.Context, request GetBackendSetRequest) (response GetBackendSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBackendSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBackendSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBackendSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBackendSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBackendSetResponse")
	}
	return
}

// getBackendSet implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) getBackendSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBackendSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/BackendSet/GetBackendSet"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "GetBackendSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetBackendSetHealth Retrieves the health status for the specified backend set.
// A default retry strategy applies to this operation GetBackendSetHealth()
func (client NetworkLoadBalancerClient) GetBackendSetHealth(ctx context.Context, request GetBackendSetHealthRequest) (response GetBackendSetHealthResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getBackendSetHealth, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetBackendSetHealthResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetBackendSetHealthResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetBackendSetHealthResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetBackendSetHealthResponse")
	}
	return
}

// getBackendSetHealth implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) getBackendSetHealth(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}/health", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetBackendSetHealthResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/BackendSetHealth/GetBackendSetHealth"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "GetBackendSetHealth", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetHealthChecker Retrieves the health check policy information for a given network load balancer and backend set.
// A default retry strategy applies to this operation GetHealthChecker()
func (client NetworkLoadBalancerClient) GetHealthChecker(ctx context.Context, request GetHealthCheckerRequest) (response GetHealthCheckerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.getHealthChecker, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetHealthCheckerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetHealthCheckerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetHealthCheckerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetHealthCheckerResponse")
	}
	return
}

// getHealthChecker implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) getHealthChecker(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}/healthChecker", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetHealthCheckerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/HealthChecker/GetHealthChecker"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "GetHealthChecker", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetListener Retrieves listener properties associated with a given network load balancer and listener name.
// A default retry strategy applies to this operation GetListener()
func (client NetworkLoadBalancerClient) GetListener(ctx context.Context, request GetListenerRequest) (response GetListenerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getListener, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetListenerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetListenerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetListenerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetListenerResponse")
	}
	return
}

// getListener implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) getListener(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/{networkLoadBalancerId}/listeners/{listenerName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetListenerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/Listener/GetListener"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "GetListener", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLiveMigration List history of live migration operation and its status for a particular nlb.
// A default retry strategy applies to this operation GetLiveMigration()
func (client NetworkLoadBalancerClient) GetLiveMigration(ctx context.Context, request GetLiveMigrationRequest) (response GetLiveMigrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLiveMigration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLiveMigrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLiveMigrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLiveMigrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLiveMigrationResponse")
	}
	return
}

// getLiveMigration implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) getLiveMigration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/{networkLoadBalancerId}/actions/liveMigrations/{liveMigrationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetLiveMigrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/LiveMigration/GetLiveMigration"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "GetLiveMigration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNetworkLoadBalancer Retrieves network load balancer configuration information by identifier.
// A default retry strategy applies to this operation GetNetworkLoadBalancer()
func (client NetworkLoadBalancerClient) GetNetworkLoadBalancer(ctx context.Context, request GetNetworkLoadBalancerRequest) (response GetNetworkLoadBalancerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNetworkLoadBalancer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNetworkLoadBalancerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNetworkLoadBalancerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNetworkLoadBalancerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNetworkLoadBalancerResponse")
	}
	return
}

// getNetworkLoadBalancer implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) getNetworkLoadBalancer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/{networkLoadBalancerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNetworkLoadBalancerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancer/GetNetworkLoadBalancer"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "GetNetworkLoadBalancer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNetworkLoadBalancerHealth Retrieves the health status for the specified network load balancer.
// A default retry strategy applies to this operation GetNetworkLoadBalancerHealth()
func (client NetworkLoadBalancerClient) GetNetworkLoadBalancerHealth(ctx context.Context, request GetNetworkLoadBalancerHealthRequest) (response GetNetworkLoadBalancerHealthResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNetworkLoadBalancerHealth, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNetworkLoadBalancerHealthResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNetworkLoadBalancerHealthResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNetworkLoadBalancerHealthResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNetworkLoadBalancerHealthResponse")
	}
	return
}

// getNetworkLoadBalancerHealth implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) getNetworkLoadBalancerHealth(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/{networkLoadBalancerId}/health", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetNetworkLoadBalancerHealthResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancerHealth/GetNetworkLoadBalancerHealth"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "GetNetworkLoadBalancerHealth", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Retrieves the details of the work request with the given identifier.
// A default retry strategy applies to this operation GetWorkRequest()
func (client NetworkLoadBalancerClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) getWorkRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/WorkRequest/GetWorkRequest"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "GetWorkRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBackendSets Lists all backend sets associated with a given network load balancer.
// A default retry strategy applies to this operation ListBackendSets()
func (client NetworkLoadBalancerClient) ListBackendSets(ctx context.Context, request ListBackendSetsRequest) (response ListBackendSetsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBackendSets, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBackendSetsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBackendSetsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBackendSetsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBackendSetsResponse")
	}
	return
}

// listBackendSets implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) listBackendSets(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBackendSetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/BackendSetSummary/ListBackendSets"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "ListBackendSets", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListBackends Lists the backend servers for a given network load balancer and backend set.
// A default retry strategy applies to this operation ListBackends()
func (client NetworkLoadBalancerClient) ListBackends(ctx context.Context, request ListBackendsRequest) (response ListBackendsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listBackends, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListBackendsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListBackendsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListBackendsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListBackendsResponse")
	}
	return
}

// listBackends implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) listBackends(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}/backends", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListBackendsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/BackendSummary/ListBackends"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "ListBackends", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListListeners Lists all listeners associated with a given network load balancer.
// A default retry strategy applies to this operation ListListeners()
func (client NetworkLoadBalancerClient) ListListeners(ctx context.Context, request ListListenersRequest) (response ListListenersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listListeners, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListListenersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListListenersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListListenersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListListenersResponse")
	}
	return
}

// listListeners implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) listListeners(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/{networkLoadBalancerId}/listeners", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListListenersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/ListenerSummary/ListListeners"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "ListListeners", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLiveMigrations List history of live migration operation and its status for a particular nlb.
// A default retry strategy applies to this operation ListLiveMigrations()
func (client NetworkLoadBalancerClient) ListLiveMigrations(ctx context.Context, request ListLiveMigrationsRequest) (response ListLiveMigrationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLiveMigrations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLiveMigrationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLiveMigrationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLiveMigrationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLiveMigrationsResponse")
	}
	return
}

// listLiveMigrations implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) listLiveMigrations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/{networkLoadBalancerId}/actions/liveMigrations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListLiveMigrationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/LiveMigrationSummary/ListLiveMigrations"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "ListLiveMigrations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNetworkLoadBalancerHealths Lists the summary health statuses for all network load balancers in the specified compartment.
// A default retry strategy applies to this operation ListNetworkLoadBalancerHealths()
func (client NetworkLoadBalancerClient) ListNetworkLoadBalancerHealths(ctx context.Context, request ListNetworkLoadBalancerHealthsRequest) (response ListNetworkLoadBalancerHealthsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNetworkLoadBalancerHealths, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNetworkLoadBalancerHealthsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNetworkLoadBalancerHealthsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNetworkLoadBalancerHealthsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNetworkLoadBalancerHealthsResponse")
	}
	return
}

// listNetworkLoadBalancerHealths implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) listNetworkLoadBalancerHealths(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers/health", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNetworkLoadBalancerHealthsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancerHealth/ListNetworkLoadBalancerHealths"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "ListNetworkLoadBalancerHealths", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNetworkLoadBalancers Returns a list of network load balancers.
// A default retry strategy applies to this operation ListNetworkLoadBalancers()
func (client NetworkLoadBalancerClient) ListNetworkLoadBalancers(ctx context.Context, request ListNetworkLoadBalancersRequest) (response ListNetworkLoadBalancersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNetworkLoadBalancers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNetworkLoadBalancersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNetworkLoadBalancersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNetworkLoadBalancersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNetworkLoadBalancersResponse")
	}
	return
}

// listNetworkLoadBalancers implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) listNetworkLoadBalancers(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancers", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNetworkLoadBalancersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancer/ListNetworkLoadBalancers"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "ListNetworkLoadBalancers", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNetworkLoadBalancersPolicies Lists the available network load balancer policies.
// A default retry strategy applies to this operation ListNetworkLoadBalancersPolicies()
func (client NetworkLoadBalancerClient) ListNetworkLoadBalancersPolicies(ctx context.Context, request ListNetworkLoadBalancersPoliciesRequest) (response ListNetworkLoadBalancersPoliciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNetworkLoadBalancersPolicies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNetworkLoadBalancersPoliciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNetworkLoadBalancersPoliciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNetworkLoadBalancersPoliciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNetworkLoadBalancersPoliciesResponse")
	}
	return
}

// listNetworkLoadBalancersPolicies implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) listNetworkLoadBalancersPolicies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancersPolicies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNetworkLoadBalancersPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancingPolicy/ListNetworkLoadBalancersPolicies"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "ListNetworkLoadBalancersPolicies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNetworkLoadBalancersProtocols This API has been deprecated so it won't return the updated list of supported protocls.
// Lists all supported traffic protocols.
// A default retry strategy applies to this operation ListNetworkLoadBalancersProtocols()
func (client NetworkLoadBalancerClient) ListNetworkLoadBalancersProtocols(ctx context.Context, request ListNetworkLoadBalancersProtocolsRequest) (response ListNetworkLoadBalancersProtocolsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNetworkLoadBalancersProtocols, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNetworkLoadBalancersProtocolsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNetworkLoadBalancersProtocolsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNetworkLoadBalancersProtocolsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNetworkLoadBalancersProtocolsResponse")
	}
	return
}

// listNetworkLoadBalancersProtocols implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) listNetworkLoadBalancersProtocols(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/networkLoadBalancersProtocols", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListNetworkLoadBalancersProtocolsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/ListenerProtocols/ListNetworkLoadBalancersProtocols"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "ListNetworkLoadBalancersProtocols", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
// A default retry strategy applies to this operation ListWorkRequestErrors()
func (client NetworkLoadBalancerClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/errors", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/WorkRequestError/ListWorkRequestErrors"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "ListWorkRequestErrors", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Returns a (paginated) list of logs for a given work request.
// A default retry strategy applies to this operation ListWorkRequestLogs()
func (client NetworkLoadBalancerClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests/{workRequestId}/logs", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/WorkRequestLogEntry/ListWorkRequestLogs"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "ListWorkRequestLogs", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists all work requests.
// A default retry strategy applies to this operation ListWorkRequests()
func (client NetworkLoadBalancerClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) listWorkRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/workRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/WorkRequest/ListWorkRequests"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "ListWorkRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// LiveMigration Perform the live migration operation by interacting with multiple nlb and vnic interfaces.
// A default retry strategy applies to this operation LiveMigration()
func (client NetworkLoadBalancerClient) LiveMigration(ctx context.Context, request LiveMigrationRequest) (response LiveMigrationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.liveMigration, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = LiveMigrationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = LiveMigrationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(LiveMigrationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into LiveMigrationResponse")
	}
	return
}

// liveMigration implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) liveMigration(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/networkLoadBalancers/{networkLoadBalancerId}/actions/liveMigrations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response LiveMigrationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancer/LiveMigration"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "LiveMigration", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBackend Updates the configuration of a backend server within the specified backend set.
// A default retry strategy applies to this operation UpdateBackend()
func (client NetworkLoadBalancerClient) UpdateBackend(ctx context.Context, request UpdateBackendRequest) (response UpdateBackendResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateBackend, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBackendResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBackendResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBackendResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBackendResponse")
	}
	return
}

// updateBackend implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) updateBackend(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}/backends/{backendName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBackendResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/Backend/UpdateBackend"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "UpdateBackend", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateBackendSet Updates a backend set.
// A default retry strategy applies to this operation UpdateBackendSet()
func (client NetworkLoadBalancerClient) UpdateBackendSet(ctx context.Context, request UpdateBackendSetRequest) (response UpdateBackendSetResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateBackendSet, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateBackendSetResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateBackendSetResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateBackendSetResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateBackendSetResponse")
	}
	return
}

// updateBackendSet implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) updateBackendSet(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateBackendSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/BackendSet/UpdateBackendSet"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "UpdateBackendSet", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateHealthChecker Updates the health check policy for a given network load balancer and backend set.
// A default retry strategy applies to this operation UpdateHealthChecker()
func (client NetworkLoadBalancerClient) UpdateHealthChecker(ctx context.Context, request UpdateHealthCheckerRequest) (response UpdateHealthCheckerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateHealthChecker, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateHealthCheckerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateHealthCheckerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateHealthCheckerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateHealthCheckerResponse")
	}
	return
}

// updateHealthChecker implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) updateHealthChecker(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkLoadBalancers/{networkLoadBalancerId}/backendSets/{backendSetName}/healthChecker", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateHealthCheckerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/HealthChecker/UpdateHealthChecker"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "UpdateHealthChecker", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateListener Updates a listener for a given network load balancer.
// A default retry strategy applies to this operation UpdateListener()
func (client NetworkLoadBalancerClient) UpdateListener(ctx context.Context, request UpdateListenerRequest) (response UpdateListenerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateListener, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateListenerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateListenerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateListenerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateListenerResponse")
	}
	return
}

// updateListener implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) updateListener(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkLoadBalancers/{networkLoadBalancerId}/listeners/{listenerName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateListenerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/Listener/UpdateListener"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "UpdateListener", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateNetworkLoadBalancer Updates the network load balancer
// A default retry strategy applies to this operation UpdateNetworkLoadBalancer()
func (client NetworkLoadBalancerClient) UpdateNetworkLoadBalancer(ctx context.Context, request UpdateNetworkLoadBalancerRequest) (response UpdateNetworkLoadBalancerResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateNetworkLoadBalancer, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateNetworkLoadBalancerResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateNetworkLoadBalancerResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateNetworkLoadBalancerResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateNetworkLoadBalancerResponse")
	}
	return
}

// updateNetworkLoadBalancer implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) updateNetworkLoadBalancer(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkLoadBalancers/{networkLoadBalancerId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateNetworkLoadBalancerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancer/UpdateNetworkLoadBalancer"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "UpdateNetworkLoadBalancer", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateNetworkSecurityGroups Updates the network security groups associated with the specified network load balancer.
// A default retry strategy applies to this operation UpdateNetworkSecurityGroups()
func (client NetworkLoadBalancerClient) UpdateNetworkSecurityGroups(ctx context.Context, request UpdateNetworkSecurityGroupsRequest) (response UpdateNetworkSecurityGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.updateNetworkSecurityGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateNetworkSecurityGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateNetworkSecurityGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateNetworkSecurityGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateNetworkSecurityGroupsResponse")
	}
	return
}

// updateNetworkSecurityGroups implements the OCIOperation interface (enables retrying operations)
func (client NetworkLoadBalancerClient) updateNetworkSecurityGroups(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/networkLoadBalancers/{networkLoadBalancerId}/networkSecurityGroups", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateNetworkSecurityGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/networkloadbalancer/20200501/NetworkLoadBalancer/UpdateNetworkSecurityGroups"
		err = common.PostProcessServiceError(err, "NetworkLoadBalancer", "UpdateNetworkSecurityGroups", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
