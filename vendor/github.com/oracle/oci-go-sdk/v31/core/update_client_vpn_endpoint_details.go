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

// UpdateClientVpnEndpointDetails A request to update clientVpnEndpoint.
type UpdateClientVpnEndpointDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A subnet for openVPN clients to access servers. Default is 172.27.224.0/20
	ClientSubnetCidr *string `mandatory:"false" json:"clientSubnetCidr"`

	// A list of accessible subnets from this clientVpnEnpoint.
	AccessibleSubnetCidrs []string `mandatory:"false" json:"accessibleSubnetCidrs"`

	// Whether re-route Internet traffic or not.
	IsRerouteInternetTraffic *bool `mandatory:"false" json:"isRerouteInternetTraffic"`

	// Allowed values:
	//   * `NAT`: NAT mode supports one-way access. In NAT mode, client can access the Internet from server endpoint
	//   but server endpoint cannot access the Internet from client.
	//   * `ROUTING`: ROUTING mode supports two-way access. In ROUTING mode, client and server endpoint can access the
	//   Internet to each other.
	AddressingMode UpdateClientVpnEndpointDetailsAddressingModeEnum `mandatory:"false" json:"addressingMode,omitempty"`

	RadiusConfig *RadiusConfigDetails `mandatory:"false" json:"radiusConfig"`

	LdapConfig *LdapConfigDetails `mandatory:"false" json:"ldapConfig"`

	DnsConfig *DnsConfigDetails `mandatory:"false" json:"dnsConfig"`

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

func (m UpdateClientVpnEndpointDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateClientVpnEndpointDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName              *string                                          `json:"displayName"`
		ClientSubnetCidr         *string                                          `json:"clientSubnetCidr"`
		AccessibleSubnetCidrs    []string                                         `json:"accessibleSubnetCidrs"`
		IsRerouteInternetTraffic *bool                                            `json:"isRerouteInternetTraffic"`
		AddressingMode           UpdateClientVpnEndpointDetailsAddressingModeEnum `json:"addressingMode"`
		RadiusConfig             *RadiusConfigDetails                             `json:"radiusConfig"`
		LdapConfig               *LdapConfigDetails                               `json:"ldapConfig"`
		DnsConfig                *DnsConfigDetails                                `json:"dnsConfig"`
		SslCertContent           sslcertcontentdetails                            `json:"sslCertContent"`
		DefinedTags              map[string]map[string]interface{}                `json:"definedTags"`
		FreeformTags             map[string]string                                `json:"freeformTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.ClientSubnetCidr = model.ClientSubnetCidr

	m.AccessibleSubnetCidrs = make([]string, len(model.AccessibleSubnetCidrs))
	for i, n := range model.AccessibleSubnetCidrs {
		m.AccessibleSubnetCidrs[i] = n
	}

	m.IsRerouteInternetTraffic = model.IsRerouteInternetTraffic

	m.AddressingMode = model.AddressingMode

	m.RadiusConfig = model.RadiusConfig

	m.LdapConfig = model.LdapConfig

	m.DnsConfig = model.DnsConfig

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

	return
}

// UpdateClientVpnEndpointDetailsAddressingModeEnum Enum with underlying type: string
type UpdateClientVpnEndpointDetailsAddressingModeEnum string

// Set of constants representing the allowable values for UpdateClientVpnEndpointDetailsAddressingModeEnum
const (
	UpdateClientVpnEndpointDetailsAddressingModeNat     UpdateClientVpnEndpointDetailsAddressingModeEnum = "NAT"
	UpdateClientVpnEndpointDetailsAddressingModeRouting UpdateClientVpnEndpointDetailsAddressingModeEnum = "ROUTING"
)

var mappingUpdateClientVpnEndpointDetailsAddressingMode = map[string]UpdateClientVpnEndpointDetailsAddressingModeEnum{
	"NAT":     UpdateClientVpnEndpointDetailsAddressingModeNat,
	"ROUTING": UpdateClientVpnEndpointDetailsAddressingModeRouting,
}

// GetUpdateClientVpnEndpointDetailsAddressingModeEnumValues Enumerates the set of values for UpdateClientVpnEndpointDetailsAddressingModeEnum
func GetUpdateClientVpnEndpointDetailsAddressingModeEnumValues() []UpdateClientVpnEndpointDetailsAddressingModeEnum {
	values := make([]UpdateClientVpnEndpointDetailsAddressingModeEnum, 0)
	for _, v := range mappingUpdateClientVpnEndpointDetailsAddressingMode {
		values = append(values, v)
	}
	return values
}
