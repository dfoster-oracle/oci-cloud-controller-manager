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

// TargetOverridePairDetails Optional and valid only for ServiceGateway (SGW) or Private Access Gateway (PAGW) to support ADB-S colocation.
// In order to support ADB-S colocation, ADB-S traffic via SGW/PAGW needs to be routed to
// dedicated fleets. Thus the existing default routes will be overridden with the overrideTargets
// provided here by SGW/PAGW at the time of GGW creation.
// When overrideEcmpTargets are provided by SGW/PAGW/NATGW at the time of GGW creation, destinations in InternalEcmpGroup
// takes precedence over default routes (including routes for colocation in case of SGW and PAGW)
type TargetOverridePairDetails struct {

	// The destination CIDRs that needs to be overridden. The rule's `destination` is an IP address range in CIDR notation.
	OverrideDestination *string `mandatory:"false" json:"overrideDestination"`

	OverrideTarget *OverrideTargetDetails `mandatory:"false" json:"overrideTarget"`

	OverrideEcmpTarget *OverrideEcmpTargetDetails `mandatory:"false" json:"overrideEcmpTarget"`
}

func (m TargetOverridePairDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetOverridePairDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
