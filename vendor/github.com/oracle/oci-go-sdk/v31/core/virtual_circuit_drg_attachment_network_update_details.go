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

// VirtualCircuitDrgAttachmentNetworkUpdateDetails Specifies the Update details for the Virtual Circuit Attachment
type VirtualCircuitDrgAttachmentNetworkUpdateDetails struct {

	// The BGP ASN to use for the Virtual Circuit's route target
	RegionalOciAsn *string `mandatory:"false" json:"regionalOciAsn"`

	// Whether the Fast Connect exists through an edge pop region.
	// Example: `true`
	IsEdgePop *bool `mandatory:"false" json:"isEdgePop"`

	// The OCI region name
	RegionName *string `mandatory:"false" json:"regionName"`
}

func (m VirtualCircuitDrgAttachmentNetworkUpdateDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m VirtualCircuitDrgAttachmentNetworkUpdateDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVirtualCircuitDrgAttachmentNetworkUpdateDetails VirtualCircuitDrgAttachmentNetworkUpdateDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeVirtualCircuitDrgAttachmentNetworkUpdateDetails
	}{
		"VIRTUAL_CIRCUIT",
		(MarshalTypeVirtualCircuitDrgAttachmentNetworkUpdateDetails)(m),
	}

	return json.Marshal(&s)
}
