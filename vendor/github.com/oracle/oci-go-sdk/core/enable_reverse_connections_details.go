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

// EnableReverseConnectionsDetails Details for enabling reverse connections for a private endpoint.
type EnableReverseConnectionsDetails struct {

	// A list of IP addresses in the customer VCN to be used as the source IPs for reverse connection packets
	// traveling from the service's VCN to the customer's VCN. If no list is specified or
	// an empty list is provided, an IP address will be chosen from the customer subnet's CIDR.
	ReverseConnectionsSourceIps []ReverseConnectionsSourceIpDetails `mandatory:"false" json:"reverseConnectionsSourceIps"`

	// Whether a DNS proxy should be configured for the reverse connection. If the service
	// does not intend to use DNS FQDN to communicate to customer endpoints, set this to `false`.
	// Example: `false`
	IsProxyEnabled *bool `mandatory:"false" json:"isProxyEnabled"`

	// The IP address in the service VCN to be used to reach the DNS proxy that resolves the
	// customer FQDN for reverse connections. If no value is provided, an available IP address will
	// be chosen from the service subnet's CIDR.
	DnsProxyIp *string `mandatory:"false" json:"dnsProxyIp"`

	// The context in which the DNS proxy will resolve the DNS queries. The default is `SERVICE`.
	// For example: if the service does not know the specific DNS zones for the customer VCNs, set
	// this to `CUSTOMER`, and set `excludedDnsZones` to the list of DNS zones in your service
	// provider VCN.
	// Allowed values:
	//  * `SERVICE`: All DNS queries will be resolved within the service VCN's DNS context,
	//    unless the FQDN belongs to one of zones in the `excludedDnsZones` list.
	//  * `CUSTOMER`: All DNS queries will be resolved within the customer VCN's DNS context,
	//    unless the FQDN belongs to one of zones in the `excludedDnsZones` list.
	DefaultDnsResolutionContext EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum `mandatory:"false" json:"defaultDnsResolutionContext,omitempty"`

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

func (m EnableReverseConnectionsDetails) String() string {
	return common.PointerString(m)
}

// EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum Enum with underlying type: string
type EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum string

// Set of constants representing the allowable values for EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum
const (
	EnableReverseConnectionsDetailsDefaultDnsResolutionContextService  EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum = "SERVICE"
	EnableReverseConnectionsDetailsDefaultDnsResolutionContextCustomer EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum = "CUSTOMER"
)

var mappingEnableReverseConnectionsDetailsDefaultDnsResolutionContext = map[string]EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum{
	"SERVICE":  EnableReverseConnectionsDetailsDefaultDnsResolutionContextService,
	"CUSTOMER": EnableReverseConnectionsDetailsDefaultDnsResolutionContextCustomer,
}

// GetEnableReverseConnectionsDetailsDefaultDnsResolutionContextEnumValues Enumerates the set of values for EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum
func GetEnableReverseConnectionsDetailsDefaultDnsResolutionContextEnumValues() []EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum {
	values := make([]EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum, 0)
	for _, v := range mappingEnableReverseConnectionsDetailsDefaultDnsResolutionContext {
		values = append(values, v)
	}
	return values
}
