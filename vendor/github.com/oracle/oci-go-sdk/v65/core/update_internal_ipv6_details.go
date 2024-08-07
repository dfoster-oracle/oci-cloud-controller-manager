// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateInternalIpv6Details Attributes availble for updating IPv6. Users can move the IPv6 by specifying the target VNIC ID.
// Internet access can also be enabled/disabled via isInternetAccessAllowed flag.
type UpdateInternalIpv6Details struct {

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The hostname of the IPv6 address. Only the hostname label, not the FQDN.
	Hostname *string `mandatory:"false" json:"hostname"`

	// Whether IPv6 is usable for intenet communication. Internet access via IPv6 will not be allowed for
	// private subnet the same way as IPv4. Internet access will be enabled by default for a public subnet.
	// If VCN has IPv6 enabled with a custom IPv6 prefix, a different public IPv6 address will be assigned
	// for a particular IPv6.
	IsInternetAccessAllowed *bool `mandatory:"false" json:"isInternetAccessAllowed"`

	// Lifetime of the IP address.
	// There are two types of IPv6 IPs:
	//  - Ephemeral
	//  - Reserved
	Lifetime UpdateInternalIpv6DetailsLifetimeEnum `mandatory:"false" json:"lifetime,omitempty"`

	// The OCID of the VNIC to reassign the IPv6 to. The VNIC must
	// be in the same subnet as the current VNIC.
	VnicId *string `mandatory:"false" json:"vnicId"`

	NextHop *PrivateIpNextHopConfiguration `mandatory:"false" json:"nextHop"`
}

func (m UpdateInternalIpv6Details) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInternalIpv6Details) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateInternalIpv6DetailsLifetimeEnum(string(m.Lifetime)); !ok && m.Lifetime != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Lifetime: %s. Supported values are: %s.", m.Lifetime, strings.Join(GetUpdateInternalIpv6DetailsLifetimeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateInternalIpv6DetailsLifetimeEnum Enum with underlying type: string
type UpdateInternalIpv6DetailsLifetimeEnum string

// Set of constants representing the allowable values for UpdateInternalIpv6DetailsLifetimeEnum
const (
	UpdateInternalIpv6DetailsLifetimeEphemeral UpdateInternalIpv6DetailsLifetimeEnum = "EPHEMERAL"
	UpdateInternalIpv6DetailsLifetimeReserved  UpdateInternalIpv6DetailsLifetimeEnum = "RESERVED"
)

var mappingUpdateInternalIpv6DetailsLifetimeEnum = map[string]UpdateInternalIpv6DetailsLifetimeEnum{
	"EPHEMERAL": UpdateInternalIpv6DetailsLifetimeEphemeral,
	"RESERVED":  UpdateInternalIpv6DetailsLifetimeReserved,
}

var mappingUpdateInternalIpv6DetailsLifetimeEnumLowerCase = map[string]UpdateInternalIpv6DetailsLifetimeEnum{
	"ephemeral": UpdateInternalIpv6DetailsLifetimeEphemeral,
	"reserved":  UpdateInternalIpv6DetailsLifetimeReserved,
}

// GetUpdateInternalIpv6DetailsLifetimeEnumValues Enumerates the set of values for UpdateInternalIpv6DetailsLifetimeEnum
func GetUpdateInternalIpv6DetailsLifetimeEnumValues() []UpdateInternalIpv6DetailsLifetimeEnum {
	values := make([]UpdateInternalIpv6DetailsLifetimeEnum, 0)
	for _, v := range mappingUpdateInternalIpv6DetailsLifetimeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateInternalIpv6DetailsLifetimeEnumStringValues Enumerates the set of values in String for UpdateInternalIpv6DetailsLifetimeEnum
func GetUpdateInternalIpv6DetailsLifetimeEnumStringValues() []string {
	return []string{
		"EPHEMERAL",
		"RESERVED",
	}
}

// GetMappingUpdateInternalIpv6DetailsLifetimeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateInternalIpv6DetailsLifetimeEnum(val string) (UpdateInternalIpv6DetailsLifetimeEnum, bool) {
	enum, ok := mappingUpdateInternalIpv6DetailsLifetimeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
