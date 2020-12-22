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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v31/common"
)

// CreateClientVpnEndpointDetails A request to create clientVpnEndpoint.
type CreateClientVpnEndpointDetails struct {

	// The OCID of the compartment that user sent request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the attachedSubnet (VNIC) in customer tenancy.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A list of accessible subnets from this clientVpnEnpoint.
	AccessibleSubnetCidrs []string `mandatory:"true" json:"accessibleSubnetCidrs"`

	DnsConfig *DnsConfigDetails `mandatory:"true" json:"dnsConfig"`

	// Whether re-route Internet traffic or not.
	IsRerouteInternetTraffic *bool `mandatory:"true" json:"isRerouteInternetTraffic"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A limit that allows the maximum number of VPN concurrent connections.
	MaxConnections *int `mandatory:"false" json:"maxConnections"`

	// A subnet for openVPN clients to access servers. Default is 172.27.224.0/20
	ClientSubnetCidr *string `mandatory:"false" json:"clientSubnetCidr"`

	// Allowed values:
	//   * `NAT`: NAT mode supports the one-way access. In NAT mode, client can access the Internet from server endpoint
	//   but server endpoint cannot access the Internet from client.
	//   * `ROUTING`: ROUTING mode supports the two-way access. In ROUTING mode, client and server endpoint can access the
	//   Internet to each other.
	AddressingMode CreateClientVpnEndpointDetailsAddressingModeEnum `mandatory:"false" json:"addressingMode,omitempty"`

	// Allowed values:
	//   * `LOCAL`: Local authentication mode that applies users and password to get authentication through the server.
	//   * `RADIUS`: RADIUS authentication mode applies users and password to get authentication through the RADIUS server.
	//   * `LDAP`: LDAP authentication mode that applies users and passwords to get authentication through the LDAP server.
	AuthenticationMode CreateClientVpnEndpointDetailsAuthenticationModeEnum `mandatory:"false" json:"authenticationMode,omitempty"`

	RadiusConfig *RadiusConfigDetails `mandatory:"false" json:"radiusConfig"`

	LdapConfig *LdapConfigDetails `mandatory:"false" json:"ldapConfig"`

	SslCertContent SslCertContentDetails `mandatory:"false" json:"sslCertContent"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateClientVpnEndpointDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateClientVpnEndpointDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                                              `json:"displayName"`
		MaxConnections           *int                                                 `json:"maxConnections"`
		ClientSubnetCidr         *string                                              `json:"clientSubnetCidr"`
		AddressingMode           CreateClientVpnEndpointDetailsAddressingModeEnum     `json:"addressingMode"`
		AuthenticationMode       CreateClientVpnEndpointDetailsAuthenticationModeEnum `json:"authenticationMode"`
		RadiusConfig             *RadiusConfigDetails                                 `json:"radiusConfig"`
		LdapConfig               *LdapConfigDetails                                   `json:"ldapConfig"`
		SslCertContent           sslcertcontentdetails                                `json:"sslCertContent"`
		DefinedTags              map[string]map[string]interface{}                    `json:"definedTags"`
		FreeformTags             map[string]string                                    `json:"freeformTags"`
		CompartmentId            *string                                              `json:"compartmentId"`
		SubnetId                 *string                                              `json:"subnetId"`
		AccessibleSubnetCidrs    []string                                             `json:"accessibleSubnetCidrs"`
		DnsConfig                *DnsConfigDetails                                    `json:"dnsConfig"`
		IsRerouteInternetTraffic *bool                                                `json:"isRerouteInternetTraffic"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.MaxConnections = model.MaxConnections

	m.ClientSubnetCidr = model.ClientSubnetCidr

	m.AddressingMode = model.AddressingMode

	m.AuthenticationMode = model.AuthenticationMode

	m.RadiusConfig = model.RadiusConfig

	m.LdapConfig = model.LdapConfig

	nn, e = model.SslCertContent.UnmarshalPolymorphicJSON(model.SslCertContent.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SslCertContent = nn.(SslCertContentDetails)
	} else {
		m.SslCertContent = nil
	}

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.CompartmentId = model.CompartmentId

	m.SubnetId = model.SubnetId

	m.AccessibleSubnetCidrs = make([]string, len(model.AccessibleSubnetCidrs))
	for i, n := range model.AccessibleSubnetCidrs {
		m.AccessibleSubnetCidrs[i] = n
	}

	m.DnsConfig = model.DnsConfig

	m.IsRerouteInternetTraffic = model.IsRerouteInternetTraffic

	return
}

// CreateClientVpnEndpointDetailsAddressingModeEnum Enum with underlying type: string
type CreateClientVpnEndpointDetailsAddressingModeEnum string

// Set of constants representing the allowable values for CreateClientVpnEndpointDetailsAddressingModeEnum
const (
	CreateClientVpnEndpointDetailsAddressingModeNat     CreateClientVpnEndpointDetailsAddressingModeEnum = "NAT"
	CreateClientVpnEndpointDetailsAddressingModeRouting CreateClientVpnEndpointDetailsAddressingModeEnum = "ROUTING"
)

var mappingCreateClientVpnEndpointDetailsAddressingMode = map[string]CreateClientVpnEndpointDetailsAddressingModeEnum{
	"NAT":     CreateClientVpnEndpointDetailsAddressingModeNat,
	"ROUTING": CreateClientVpnEndpointDetailsAddressingModeRouting,
}

// GetCreateClientVpnEndpointDetailsAddressingModeEnumValues Enumerates the set of values for CreateClientVpnEndpointDetailsAddressingModeEnum
func GetCreateClientVpnEndpointDetailsAddressingModeEnumValues() []CreateClientVpnEndpointDetailsAddressingModeEnum {
	values := make([]CreateClientVpnEndpointDetailsAddressingModeEnum, 0)
	for _, v := range mappingCreateClientVpnEndpointDetailsAddressingMode {
		values = append(values, v)
	}
	return values
}

// CreateClientVpnEndpointDetailsAuthenticationModeEnum Enum with underlying type: string
type CreateClientVpnEndpointDetailsAuthenticationModeEnum string

// Set of constants representing the allowable values for CreateClientVpnEndpointDetailsAuthenticationModeEnum
const (
	CreateClientVpnEndpointDetailsAuthenticationModeLocal  CreateClientVpnEndpointDetailsAuthenticationModeEnum = "LOCAL"
	CreateClientVpnEndpointDetailsAuthenticationModeRadius CreateClientVpnEndpointDetailsAuthenticationModeEnum = "RADIUS"
	CreateClientVpnEndpointDetailsAuthenticationModeLdap   CreateClientVpnEndpointDetailsAuthenticationModeEnum = "LDAP"
)

var mappingCreateClientVpnEndpointDetailsAuthenticationMode = map[string]CreateClientVpnEndpointDetailsAuthenticationModeEnum{
	"LOCAL":  CreateClientVpnEndpointDetailsAuthenticationModeLocal,
	"RADIUS": CreateClientVpnEndpointDetailsAuthenticationModeRadius,
	"LDAP":   CreateClientVpnEndpointDetailsAuthenticationModeLdap,
}

// GetCreateClientVpnEndpointDetailsAuthenticationModeEnumValues Enumerates the set of values for CreateClientVpnEndpointDetailsAuthenticationModeEnum
func GetCreateClientVpnEndpointDetailsAuthenticationModeEnumValues() []CreateClientVpnEndpointDetailsAuthenticationModeEnum {
	values := make([]CreateClientVpnEndpointDetailsAuthenticationModeEnum, 0)
	for _, v := range mappingCreateClientVpnEndpointDetailsAuthenticationMode {
		values = append(values, v)
	}
	return values
}
