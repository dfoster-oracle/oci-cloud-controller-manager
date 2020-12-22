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

// VcnDrgAttachmentNetworkUpdateDetails Specifies the Update details for the VCN Attachment
type VcnDrgAttachmentNetworkUpdateDetails struct {

	// The OCID of the route table the DRG attachment is using.
	// For information about why you would associate a route table with a DRG attachment, see
	// Advanced Scenario: Transit Routing (https://docs.cloud.oracle.com/Content/Network/Tasks/transitrouting.htm).
	// For information about why you would associate a route table with a DRG attachment, see:
	//   * Transit Routing: Access to Multiple VCNs in Same Region (https://docs.cloud.oracle.com/Content/Network/Tasks/transitrouting.htm)
	//   * Transit Routing: Private Access to Oracle Services (https://docs.cloud.oracle.com/Content/Network/Tasks/transitroutingoracleservices.htm)
	RouteTableId *string `mandatory:"false" json:"routeTableId"`
}

func (m VcnDrgAttachmentNetworkUpdateDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m VcnDrgAttachmentNetworkUpdateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVcnDrgAttachmentNetworkUpdateDetails VcnDrgAttachmentNetworkUpdateDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeVcnDrgAttachmentNetworkUpdateDetails
	}{
		"VCN",
		(MarshalTypeVcnDrgAttachmentNetworkUpdateDetails)(m),
	}

	return json.Marshal(&s)
}
