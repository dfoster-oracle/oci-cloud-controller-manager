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

// AddDrgRouteRuleDetails Details for adding a DRG Route Rule. You cannot add a static route rule
// for an RPC attachment to a table that is assigned to a Virtual Circuit or IPSec attachment.
// You cannot create a static route for a Virtual Circuit or IPSec Tunnel attachment.
type AddDrgRouteRuleDetails struct {

	// Type of destination for the rule. Required if `direction` = `EGRESS`.
	// Allowed values:
	//   * `CIDR_BLOCK`: If the rule's `destination` is an IP address range in CIDR notation.
	DestinationType AddDrgRouteRuleDetailsDestinationTypeEnum `mandatory:"true" json:"destinationType"`

	// Conceptually, this is the range of IP addresses used for matching when routing
	// traffic. Only CIDR_BLOCK values allowed
	// Potential values:
	//   * IP address range in CIDR notation. Can be an IPv4 or IPv6 CIDR. For example: `192.168.1.0/24`
	//   or `2001:0db8:0123:45::/56`.
	Destination *string `mandatory:"true" json:"destination"`

	// The OCID of the next hop DRG Attachment. The next hop DRG Attachment is responsible
	// for reaching the network destination.
	NextHopDrgAttachmentId *string `mandatory:"true" json:"nextHopDrgAttachmentId"`
}

func (m AddDrgRouteRuleDetails) String() string {
	return common.PointerString(m)
}

// AddDrgRouteRuleDetailsDestinationTypeEnum Enum with underlying type: string
type AddDrgRouteRuleDetailsDestinationTypeEnum string

// Set of constants representing the allowable values for AddDrgRouteRuleDetailsDestinationTypeEnum
const (
	AddDrgRouteRuleDetailsDestinationTypeCidrBlock AddDrgRouteRuleDetailsDestinationTypeEnum = "CIDR_BLOCK"
)

var mappingAddDrgRouteRuleDetailsDestinationType = map[string]AddDrgRouteRuleDetailsDestinationTypeEnum{
	"CIDR_BLOCK": AddDrgRouteRuleDetailsDestinationTypeCidrBlock,
}

// GetAddDrgRouteRuleDetailsDestinationTypeEnumValues Enumerates the set of values for AddDrgRouteRuleDetailsDestinationTypeEnum
func GetAddDrgRouteRuleDetailsDestinationTypeEnumValues() []AddDrgRouteRuleDetailsDestinationTypeEnum {
	values := make([]AddDrgRouteRuleDetailsDestinationTypeEnum, 0)
	for _, v := range mappingAddDrgRouteRuleDetailsDestinationType {
		values = append(values, v)
	}
	return values
}
