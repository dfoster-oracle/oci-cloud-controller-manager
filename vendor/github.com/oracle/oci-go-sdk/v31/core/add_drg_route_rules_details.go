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

// AddDrgRouteRulesDetails Details request to add static routes to a DRG Route Table.
type AddDrgRouteRulesDetails struct {

	// The collection of static rules used to insert into the DRG Route Table. You cannot add a static route rule
	// for an RPC attachment to a table that is assigned to a Virtual Circuit or IPSec attachment.
	// You cannot create a static route for a Virtual Circuit or IPSec Tunnel attachment.
	RouteRules []AddDrgRouteRuleDetails `mandatory:"false" json:"routeRules"`
}

func (m AddDrgRouteRulesDetails) String() string {
	return common.PointerString(m)
}
