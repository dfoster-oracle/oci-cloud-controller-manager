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

// ClientVpnEndpoint The ClientVPNEnpoint is a server endpoint that allow customer get access to openVPN service.
type ClientVpnEndpoint struct {

	// The OCID of the compartment that user sent request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of clientVPNEndpoint.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the attachedSubnet (VNIC) in customer tenancy.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The IP address in attached subnet.
	IpAddressInAttachedSubnet *string `mandatory:"true" json:"ipAddressInAttachedSubnet"`

	// Whether re-route Internet traffic or not.
	IsRerouteInternetTraffic *bool `mandatory:"true" json:"isRerouteInternetTraffic"`

	// A list of accessible subnets from this clientVPNEndpoint.
	AccessibleSubnetCidrs []string `mandatory:"true" json:"accessibleSubnetCidrs"`

	DnsConfig *DnsConfigDetails `mandatory:"true" json:"dnsConfig"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The current state of the ClientVPNenpoint.
	LifecycleState ClientVpnEndpointLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A limit that allows the maximum number of VPN concurrent connections.
	MaxConnections *int `mandatory:"false" json:"maxConnections"`

	// A subnet for openVPN clients to access servers.
	ClientSubnetCidr *string `mandatory:"false" json:"clientSubnetCidr"`

	// Allowed values:
	//   * `NAT`: NAT mode supports the one-way access. In NAT mode, client can access the Internet from server endpoint
	//   but server endpoint cannot access the Internet from client.
	//   * `ROUTING`: ROUTING mode supports the two-way access. In ROUTING mode, client and server endpoint can access the
	//   Internet to each other.
	AddressingMode ClientVpnEndpointAddressingModeEnum `mandatory:"false" json:"addressingMode,omitempty"`

	// Allowed values:
	//   * `LOCAL`: Local authentication mode that applies users and password to get authentication through the server.
	//   * `RADIUS`: RADIUS authentication mode applies users and password to get authentication through the RADIUS server.
	//   * `LDAP`: LDAP authentication mode that applies users and passwords to get authentication through the LDAP server.
	AuthenticationMode ClientVpnEndpointAuthenticationModeEnum `mandatory:"false" json:"authenticationMode,omitempty"`

	RadiusConfig *RadiusConfigDetails `mandatory:"false" json:"radiusConfig"`

	LdapConfig *LdapConfigDetails `mandatory:"false" json:"ldapConfig"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m ClientVpnEndpoint) String() string {
	return common.PointerString(m)
}

// ClientVpnEndpointLifecycleStateEnum Enum with underlying type: string
type ClientVpnEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for ClientVpnEndpointLifecycleStateEnum
const (
	ClientVpnEndpointLifecycleStateCreating ClientVpnEndpointLifecycleStateEnum = "CREATING"
	ClientVpnEndpointLifecycleStateActive   ClientVpnEndpointLifecycleStateEnum = "ACTIVE"
	ClientVpnEndpointLifecycleStateInactive ClientVpnEndpointLifecycleStateEnum = "INACTIVE"
	ClientVpnEndpointLifecycleStateFailed   ClientVpnEndpointLifecycleStateEnum = "FAILED"
	ClientVpnEndpointLifecycleStateDeleted  ClientVpnEndpointLifecycleStateEnum = "DELETED"
	ClientVpnEndpointLifecycleStateDeleting ClientVpnEndpointLifecycleStateEnum = "DELETING"
	ClientVpnEndpointLifecycleStateUpdating ClientVpnEndpointLifecycleStateEnum = "UPDATING"
)

var mappingClientVpnEndpointLifecycleState = map[string]ClientVpnEndpointLifecycleStateEnum{
	"CREATING": ClientVpnEndpointLifecycleStateCreating,
	"ACTIVE":   ClientVpnEndpointLifecycleStateActive,
	"INACTIVE": ClientVpnEndpointLifecycleStateInactive,
	"FAILED":   ClientVpnEndpointLifecycleStateFailed,
	"DELETED":  ClientVpnEndpointLifecycleStateDeleted,
	"DELETING": ClientVpnEndpointLifecycleStateDeleting,
	"UPDATING": ClientVpnEndpointLifecycleStateUpdating,
}

// GetClientVpnEndpointLifecycleStateEnumValues Enumerates the set of values for ClientVpnEndpointLifecycleStateEnum
func GetClientVpnEndpointLifecycleStateEnumValues() []ClientVpnEndpointLifecycleStateEnum {
	values := make([]ClientVpnEndpointLifecycleStateEnum, 0)
	for _, v := range mappingClientVpnEndpointLifecycleState {
		values = append(values, v)
	}
	return values
}

// ClientVpnEndpointAddressingModeEnum Enum with underlying type: string
type ClientVpnEndpointAddressingModeEnum string

// Set of constants representing the allowable values for ClientVpnEndpointAddressingModeEnum
const (
	ClientVpnEndpointAddressingModeNat     ClientVpnEndpointAddressingModeEnum = "NAT"
	ClientVpnEndpointAddressingModeRouting ClientVpnEndpointAddressingModeEnum = "ROUTING"
)

var mappingClientVpnEndpointAddressingMode = map[string]ClientVpnEndpointAddressingModeEnum{
	"NAT":     ClientVpnEndpointAddressingModeNat,
	"ROUTING": ClientVpnEndpointAddressingModeRouting,
}

// GetClientVpnEndpointAddressingModeEnumValues Enumerates the set of values for ClientVpnEndpointAddressingModeEnum
func GetClientVpnEndpointAddressingModeEnumValues() []ClientVpnEndpointAddressingModeEnum {
	values := make([]ClientVpnEndpointAddressingModeEnum, 0)
	for _, v := range mappingClientVpnEndpointAddressingMode {
		values = append(values, v)
	}
	return values
}

// ClientVpnEndpointAuthenticationModeEnum Enum with underlying type: string
type ClientVpnEndpointAuthenticationModeEnum string

// Set of constants representing the allowable values for ClientVpnEndpointAuthenticationModeEnum
const (
	ClientVpnEndpointAuthenticationModeLocal  ClientVpnEndpointAuthenticationModeEnum = "LOCAL"
	ClientVpnEndpointAuthenticationModeRadius ClientVpnEndpointAuthenticationModeEnum = "RADIUS"
	ClientVpnEndpointAuthenticationModeLdap   ClientVpnEndpointAuthenticationModeEnum = "LDAP"
)

var mappingClientVpnEndpointAuthenticationMode = map[string]ClientVpnEndpointAuthenticationModeEnum{
	"LOCAL":  ClientVpnEndpointAuthenticationModeLocal,
	"RADIUS": ClientVpnEndpointAuthenticationModeRadius,
	"LDAP":   ClientVpnEndpointAuthenticationModeLdap,
}

// GetClientVpnEndpointAuthenticationModeEnumValues Enumerates the set of values for ClientVpnEndpointAuthenticationModeEnum
func GetClientVpnEndpointAuthenticationModeEnumValues() []ClientVpnEndpointAuthenticationModeEnum {
	values := make([]ClientVpnEndpointAuthenticationModeEnum, 0)
	for _, v := range mappingClientVpnEndpointAuthenticationMode {
		values = append(values, v)
	}
	return values
}
