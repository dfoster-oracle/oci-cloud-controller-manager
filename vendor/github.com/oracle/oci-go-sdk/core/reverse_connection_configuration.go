// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ReverseConnectionConfiguration Reverse connection configuration details for the private endpoint.
// When reverse connection support is enabled, the private endpoint allows reverse connections to be
// established to the customer VCN. The packets traveling from the service's VCN to the customer's
// VCN in a reverse connection use a different source IP address than the private endpoint's IP address.
type ReverseConnectionConfiguration struct {

	// The reverse connection configuration's current state.
	LifecycleState ReverseConnectionConfigurationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The list of IP addresses in the customer VCN to be used as a source IP for reverse connection packets
	// traveling from the service's VCN to the customer's VCN.
	ReverseConnectionsSourceIps []ReverseConnectionsSourceIpDetails `mandatory:"false" json:"reverseConnectionsSourceIps"`

	// Whether a DNS proxy is configured for the reverse connection. If the service
	// does not intend to use DNS FQDN to communicate to customer endpoints, set this to `false`.
	// Example: `false`
	IsProxyEnabled *bool `mandatory:"false" json:"isProxyEnabled"`

	// The IP address in the service VCN to be used to reach the DNS proxy that resolves the
	// customer FQDN for reverse connections.
	DnsProxyIp *string `mandatory:"false" json:"dnsProxyIp"`

	// The context in which the DNS proxy will resolve the DNS queries. The default is `SERVICE`.
	// For example: if the service does not know the specific DNS zones for the customer VCNs, set
	// this to `CUSTOMER`, and set `excludedDnsZones` to the list of DNS zones in your service
	// provider VCN.
	// Allowed values:
	//  * `SERVICE` : All DNS queries will be resolved within the service VCN's DNS context,
	//    unless the FQDN belongs to one of zones in the `excludedDnsZones` list.
	//  * `CUSTOMER` : All DNS queries will be resolved within the customer VCN's DNS context,
	//    unless the FQDN belongs to one of zones in the `excludedDnsZones` list.
	DefaultDnsResolutionContext ReverseConnectionConfigurationDefaultDnsResolutionContextEnum `mandatory:"false" json:"defaultDnsResolutionContext,omitempty"`

	// List of DNS zones to exclude from the default DNS resolution context.
	ExcludedDnsZones []string `mandatory:"false" json:"excludedDnsZones"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the service's subnet where
	// the DNS proxy endpoint will be created.
	ServiceSubnetId *string `mandatory:"false" json:"serviceSubnetId"`

	// A list of the OCIDs of the network security groups that the reverse connection's VNIC belongs to.
	// For more information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m ReverseConnectionConfiguration) String() string {
	return common.PointerString(m)
}

// ReverseConnectionConfigurationLifecycleStateEnum Enum with underlying type: string
type ReverseConnectionConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for ReverseConnectionConfigurationLifecycleStateEnum
const (
	ReverseConnectionConfigurationLifecycleStateProvisioning ReverseConnectionConfigurationLifecycleStateEnum = "PROVISIONING"
	ReverseConnectionConfigurationLifecycleStateAvailable    ReverseConnectionConfigurationLifecycleStateEnum = "AVAILABLE"
	ReverseConnectionConfigurationLifecycleStateUpdating     ReverseConnectionConfigurationLifecycleStateEnum = "UPDATING"
	ReverseConnectionConfigurationLifecycleStateTerminating  ReverseConnectionConfigurationLifecycleStateEnum = "TERMINATING"
	ReverseConnectionConfigurationLifecycleStateTerminated   ReverseConnectionConfigurationLifecycleStateEnum = "TERMINATED"
	ReverseConnectionConfigurationLifecycleStateFailed       ReverseConnectionConfigurationLifecycleStateEnum = "FAILED"
)

var mappingReverseConnectionConfigurationLifecycleState = map[string]ReverseConnectionConfigurationLifecycleStateEnum{
	"PROVISIONING": ReverseConnectionConfigurationLifecycleStateProvisioning,
	"AVAILABLE":    ReverseConnectionConfigurationLifecycleStateAvailable,
	"UPDATING":     ReverseConnectionConfigurationLifecycleStateUpdating,
	"TERMINATING":  ReverseConnectionConfigurationLifecycleStateTerminating,
	"TERMINATED":   ReverseConnectionConfigurationLifecycleStateTerminated,
	"FAILED":       ReverseConnectionConfigurationLifecycleStateFailed,
}

// GetReverseConnectionConfigurationLifecycleStateEnumValues Enumerates the set of values for ReverseConnectionConfigurationLifecycleStateEnum
func GetReverseConnectionConfigurationLifecycleStateEnumValues() []ReverseConnectionConfigurationLifecycleStateEnum {
	values := make([]ReverseConnectionConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingReverseConnectionConfigurationLifecycleState {
		values = append(values, v)
	}
	return values
}

// ReverseConnectionConfigurationDefaultDnsResolutionContextEnum Enum with underlying type: string
type ReverseConnectionConfigurationDefaultDnsResolutionContextEnum string

// Set of constants representing the allowable values for ReverseConnectionConfigurationDefaultDnsResolutionContextEnum
const (
	ReverseConnectionConfigurationDefaultDnsResolutionContextService  ReverseConnectionConfigurationDefaultDnsResolutionContextEnum = "SERVICE"
	ReverseConnectionConfigurationDefaultDnsResolutionContextCustomer ReverseConnectionConfigurationDefaultDnsResolutionContextEnum = "CUSTOMER"
)

var mappingReverseConnectionConfigurationDefaultDnsResolutionContext = map[string]ReverseConnectionConfigurationDefaultDnsResolutionContextEnum{
	"SERVICE":  ReverseConnectionConfigurationDefaultDnsResolutionContextService,
	"CUSTOMER": ReverseConnectionConfigurationDefaultDnsResolutionContextCustomer,
}

// GetReverseConnectionConfigurationDefaultDnsResolutionContextEnumValues Enumerates the set of values for ReverseConnectionConfigurationDefaultDnsResolutionContextEnum
func GetReverseConnectionConfigurationDefaultDnsResolutionContextEnumValues() []ReverseConnectionConfigurationDefaultDnsResolutionContextEnum {
	values := make([]ReverseConnectionConfigurationDefaultDnsResolutionContextEnum, 0)
	for _, v := range mappingReverseConnectionConfigurationDefaultDnsResolutionContext {
		values = append(values, v)
	}
	return values
}
