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

// ClientVpnEndpointSummary a summary of ClientVpnEndpoint.
type ClientVpnEndpointSummary struct {

	// The OCID of the compartment that user sent request.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID of clientVPNEndpoint.
	Id *string `mandatory:"false" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A limit that allows the maximum number of VPN concurrent connections.
	MaxConnections *int `mandatory:"false" json:"maxConnections"`

	// The current state of the ClientVpnEndpoint.
	LifecycleState ClientVpnEndpointLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A subnet for openVPN clients to access servers.
	ClientSubnetCidr *string `mandatory:"false" json:"clientSubnetCidr"`

	// The OCID of the attachedSubnet (VNIC) in customer tenancy.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The IP address in attached subnet.
	IpAddressInAttachedSubnet *string `mandatory:"false" json:"ipAddressInAttachedSubnet"`

	// Whether re-route Internet traffic or not.
	IsRerouteInternetTraffic *bool `mandatory:"false" json:"isRerouteInternetTraffic"`

	// Allowed values:
	//   * `NAT`: NAT mode supports the one-way access. In NAT mode, client can access the Internet from server endpoint
	//   but server endpoint cannot access the Internet from client.
	//   * `ROUTING`: ROUTING mode supports the two-way access. In ROUTING mode, client and server endpoint can access the
	//   Internet to each other.
	AddressingMode ClientVpnEndpointSummaryAddressingModeEnum `mandatory:"false" json:"addressingMode,omitempty"`

	// Allowed values:
	//   * `LOCAL`: Local authentication mode that applies users and password to get authentication through the server.
	//   * `RADIUS`: RADIUS authentication mode applies users and password to get authentication through the RADIUS server.
	//   * `LDAP`: LDAP authentication mode that applies users and passwords to get authentication through the LDAP server.
	AuthenticationMode ClientVpnEndpointSummaryAuthenticationModeEnum `mandatory:"false" json:"authenticationMode,omitempty"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m ClientVpnEndpointSummary) String() string {
	return common.PointerString(m)
}

// ClientVpnEndpointSummaryAddressingModeEnum Enum with underlying type: string
type ClientVpnEndpointSummaryAddressingModeEnum string

// Set of constants representing the allowable values for ClientVpnEndpointSummaryAddressingModeEnum
const (
	ClientVpnEndpointSummaryAddressingModeNat     ClientVpnEndpointSummaryAddressingModeEnum = "NAT"
	ClientVpnEndpointSummaryAddressingModeRouting ClientVpnEndpointSummaryAddressingModeEnum = "ROUTING"
)

var mappingClientVpnEndpointSummaryAddressingMode = map[string]ClientVpnEndpointSummaryAddressingModeEnum{
	"NAT":     ClientVpnEndpointSummaryAddressingModeNat,
	"ROUTING": ClientVpnEndpointSummaryAddressingModeRouting,
}

// GetClientVpnEndpointSummaryAddressingModeEnumValues Enumerates the set of values for ClientVpnEndpointSummaryAddressingModeEnum
func GetClientVpnEndpointSummaryAddressingModeEnumValues() []ClientVpnEndpointSummaryAddressingModeEnum {
	values := make([]ClientVpnEndpointSummaryAddressingModeEnum, 0)
	for _, v := range mappingClientVpnEndpointSummaryAddressingMode {
		values = append(values, v)
	}
	return values
}

// ClientVpnEndpointSummaryAuthenticationModeEnum Enum with underlying type: string
type ClientVpnEndpointSummaryAuthenticationModeEnum string

// Set of constants representing the allowable values for ClientVpnEndpointSummaryAuthenticationModeEnum
const (
	ClientVpnEndpointSummaryAuthenticationModeLocal  ClientVpnEndpointSummaryAuthenticationModeEnum = "LOCAL"
	ClientVpnEndpointSummaryAuthenticationModeRadius ClientVpnEndpointSummaryAuthenticationModeEnum = "RADIUS"
	ClientVpnEndpointSummaryAuthenticationModeLdap   ClientVpnEndpointSummaryAuthenticationModeEnum = "LDAP"
)

var mappingClientVpnEndpointSummaryAuthenticationMode = map[string]ClientVpnEndpointSummaryAuthenticationModeEnum{
	"LOCAL":  ClientVpnEndpointSummaryAuthenticationModeLocal,
	"RADIUS": ClientVpnEndpointSummaryAuthenticationModeRadius,
	"LDAP":   ClientVpnEndpointSummaryAuthenticationModeLdap,
}

// GetClientVpnEndpointSummaryAuthenticationModeEnumValues Enumerates the set of values for ClientVpnEndpointSummaryAuthenticationModeEnum
func GetClientVpnEndpointSummaryAuthenticationModeEnumValues() []ClientVpnEndpointSummaryAuthenticationModeEnum {
	values := make([]ClientVpnEndpointSummaryAuthenticationModeEnum, 0)
	for _, v := range mappingClientVpnEndpointSummaryAuthenticationMode {
		values = append(values, v)
	}
	return values
}
