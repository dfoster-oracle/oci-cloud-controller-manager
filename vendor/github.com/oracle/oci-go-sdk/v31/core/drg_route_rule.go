// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm),
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm), and
// Registry (https://docs.cloud.oracle.com/iaas/Content/Registry/Concepts/registryoverview.htm) services.
// Use this API to manage resources such as virtual cloud networks (VCNs),
// compute instances, block storage volumes, and container images.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// DrgRouteRule Route Rule in the DRG Route Table. A DRG Route Rule is a mapping between a destination IP address range
// and a DRG Attachment to route matching packets. If there are multiple rules with identical destinations and different nextHopAttachments,
// and none of the rules are marked as "conflict", traffic will be routed across the attachments using ECMP.
type DrgRouteRule struct {

	// Conceptually, this is the range of IP addresses used for matching when routing
	// traffic.
	// Potential values:
	//   * IP address range in CIDR notation. Can be an IPv4 or IPv6 CIDR. For example: `192.168.1.0/24`
	//   or `2001:0db8:0123:45::/56`.
	//   * The `cidrBlock` value for a Service, if you're
	//     setting up a security rule for traffic destined for a particular `Service` through
	//     a service gateway. For example: `oci-phx-objectstorage`.
	Destination *string `mandatory:"true" json:"destination"`

	// Type of destination for the rule. Required if `direction` = `EGRESS`.
	// Allowed values:
	//   * `CIDR_BLOCK`: If the rule's `destination` is an IP address range in CIDR notation.
	//   * `SERVICE_CIDR_BLOCK`: If the rule's `destination` is the `cidrBlock` value for a
	//     Service (the rule is for traffic destined for a
	//     particular `Service` through a service gateway).
	DestinationType DrgRouteRuleDestinationTypeEnum `mandatory:"true" json:"destinationType"`

	// The OCID of the next hop DRG Attachment. The next hop DRG Attachment is responsible
	// for reaching the network destination.
	// value is "BLACKHOLE" if the route is blackholed
	NextHopDrgAttachmentId *string `mandatory:"true" json:"nextHopDrgAttachmentId"`

	// The Oracle-assigned ID of the DRG Route Rule.
	Id *string `mandatory:"true" json:"id"`

	// The earliest origin of a route. If a route is advertised to a DRG through an IPsec Tunnel attachment,
	// and is propagated to peered DRG(s) via RPC attachments, the route's provenance in the peered DRGs remains as IPSEC_TUNNEL,
	// because that is the earliest origin.
	// No routes with provenance IPSEC_TUNNEL or VIRTUAL_CIRCUIT will be exported to IPSEC_TUNNEL or VIRTUAL_CIRCUIT attachments,
	// regardless of the attachment's export distribution.
	RouteProvenance DrgRouteRuleRouteProvenanceEnum `mandatory:"true" json:"routeProvenance"`

	// Static routes are specified by the user through the DRG Route Table API.
	// Dynamic routes are learned by the DRG from the DRG Attachments through various routing protocols.
	RouteType DrgRouteRuleRouteTypeEnum `mandatory:"false" json:"routeType,omitempty"`

	// If the route was not imported due to conflict
	IsConflict *bool `mandatory:"false" json:"isConflict"`
}

func (m DrgRouteRule) String() string {
	return common.PointerString(m)
}

// DrgRouteRuleDestinationTypeEnum Enum with underlying type: string
type DrgRouteRuleDestinationTypeEnum string

// Set of constants representing the allowable values for DrgRouteRuleDestinationTypeEnum
const (
	DrgRouteRuleDestinationTypeCidrBlock        DrgRouteRuleDestinationTypeEnum = "CIDR_BLOCK"
	DrgRouteRuleDestinationTypeServiceCidrBlock DrgRouteRuleDestinationTypeEnum = "SERVICE_CIDR_BLOCK"
)

var mappingDrgRouteRuleDestinationType = map[string]DrgRouteRuleDestinationTypeEnum{
	"CIDR_BLOCK":         DrgRouteRuleDestinationTypeCidrBlock,
	"SERVICE_CIDR_BLOCK": DrgRouteRuleDestinationTypeServiceCidrBlock,
}

// GetDrgRouteRuleDestinationTypeEnumValues Enumerates the set of values for DrgRouteRuleDestinationTypeEnum
func GetDrgRouteRuleDestinationTypeEnumValues() []DrgRouteRuleDestinationTypeEnum {
	values := make([]DrgRouteRuleDestinationTypeEnum, 0)
	for _, v := range mappingDrgRouteRuleDestinationType {
		values = append(values, v)
	}
	return values
}

// DrgRouteRuleRouteTypeEnum Enum with underlying type: string
type DrgRouteRuleRouteTypeEnum string

// Set of constants representing the allowable values for DrgRouteRuleRouteTypeEnum
const (
	DrgRouteRuleRouteTypeStatic  DrgRouteRuleRouteTypeEnum = "STATIC"
	DrgRouteRuleRouteTypeDynamic DrgRouteRuleRouteTypeEnum = "DYNAMIC"
)

var mappingDrgRouteRuleRouteType = map[string]DrgRouteRuleRouteTypeEnum{
	"STATIC":  DrgRouteRuleRouteTypeStatic,
	"DYNAMIC": DrgRouteRuleRouteTypeDynamic,
}

// GetDrgRouteRuleRouteTypeEnumValues Enumerates the set of values for DrgRouteRuleRouteTypeEnum
func GetDrgRouteRuleRouteTypeEnumValues() []DrgRouteRuleRouteTypeEnum {
	values := make([]DrgRouteRuleRouteTypeEnum, 0)
	for _, v := range mappingDrgRouteRuleRouteType {
		values = append(values, v)
	}
	return values
}

// DrgRouteRuleRouteProvenanceEnum Enum with underlying type: string
type DrgRouteRuleRouteProvenanceEnum string

// Set of constants representing the allowable values for DrgRouteRuleRouteProvenanceEnum
const (
	DrgRouteRuleRouteProvenanceStatic         DrgRouteRuleRouteProvenanceEnum = "STATIC"
	DrgRouteRuleRouteProvenanceIpsecTunnel    DrgRouteRuleRouteProvenanceEnum = "IPSEC_TUNNEL"
	DrgRouteRuleRouteProvenanceVirtualCircuit DrgRouteRuleRouteProvenanceEnum = "VIRTUAL_CIRCUIT"
	DrgRouteRuleRouteProvenanceVcnCidr        DrgRouteRuleRouteProvenanceEnum = "VCN_CIDR"
	DrgRouteRuleRouteProvenanceVcnSubnet      DrgRouteRuleRouteProvenanceEnum = "VCN_SUBNET"
	DrgRouteRuleRouteProvenanceVcnRoute       DrgRouteRuleRouteProvenanceEnum = "VCN_ROUTE"
	DrgRouteRuleRouteProvenanceVcnVlan        DrgRouteRuleRouteProvenanceEnum = "VCN_VLAN"
)

var mappingDrgRouteRuleRouteProvenance = map[string]DrgRouteRuleRouteProvenanceEnum{
	"STATIC":          DrgRouteRuleRouteProvenanceStatic,
	"IPSEC_TUNNEL":    DrgRouteRuleRouteProvenanceIpsecTunnel,
	"VIRTUAL_CIRCUIT": DrgRouteRuleRouteProvenanceVirtualCircuit,
	"VCN_CIDR":        DrgRouteRuleRouteProvenanceVcnCidr,
	"VCN_SUBNET":      DrgRouteRuleRouteProvenanceVcnSubnet,
	"VCN_ROUTE":       DrgRouteRuleRouteProvenanceVcnRoute,
	"VCN_VLAN":        DrgRouteRuleRouteProvenanceVcnVlan,
}

// GetDrgRouteRuleRouteProvenanceEnumValues Enumerates the set of values for DrgRouteRuleRouteProvenanceEnum
func GetDrgRouteRuleRouteProvenanceEnumValues() []DrgRouteRuleRouteProvenanceEnum {
	values := make([]DrgRouteRuleRouteProvenanceEnum, 0)
	for _, v := range mappingDrgRouteRuleRouteProvenance {
		values = append(values, v)
	}
	return values
}
