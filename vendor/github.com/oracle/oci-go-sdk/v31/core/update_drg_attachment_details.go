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

// UpdateDrgAttachmentDetails The representation of UpdateDrgAttachmentDetails
type UpdateDrgAttachmentDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID of the DRG route table the attachment will use.
	// The DRG route table manages traffic inside the DRG.
	// If you don't specify a route table here, the DRG attachment will be assigned to a DRG route table according to
	// the DRG's route tables' defaultAssignmentTypes.
	// You cannot remove a DRG Route Table from a DRG Attachment, but you can change the assignment to a different
	// DRG Route Table.
	DrgRouteTableId *string `mandatory:"false" json:"drgRouteTableId"`

	NetworkDetails DrgAttachmentNetworkUpdateDetails `mandatory:"false" json:"networkDetails"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID of the export route distribution used to specify how routes in the assigned DRG route table
	// are advertised out through the attachment.
	// If this value is null, no routes are advertised through this attachment.
	ExportRouteDistributionId *string `mandatory:"false" json:"exportRouteDistributionId"`

	// The OCID of the route table the DRG attachment will use.
	// For information about why you would associate a route table with a DRG attachment, see:
	//   * Transit Routing: Access to Multiple VCNs in Same Region (https://docs.cloud.oracle.com/Content/Network/Tasks/transitrouting.htm)
	//   * Transit Routing: Private Access to Oracle Services (https://docs.cloud.oracle.com/Content/Network/Tasks/transitroutingoracleservices.htm)
	RouteTableId *string `mandatory:"false" json:"routeTableId"`
}

func (m UpdateDrgAttachmentDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDrgAttachmentDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName               *string                           `json:"displayName"`
		DrgRouteTableId           *string                           `json:"drgRouteTableId"`
		NetworkDetails            drgattachmentnetworkupdatedetails `json:"networkDetails"`
		DefinedTags               map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags              map[string]string                 `json:"freeformTags"`
		ExportRouteDistributionId *string                           `json:"exportRouteDistributionId"`
		RouteTableId              *string                           `json:"routeTableId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.DrgRouteTableId = model.DrgRouteTableId

	nn, e = model.NetworkDetails.UnmarshalPolymorphicJSON(model.NetworkDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NetworkDetails = nn.(DrgAttachmentNetworkUpdateDetails)
	} else {
		m.NetworkDetails = nil
	}

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.ExportRouteDistributionId = model.ExportRouteDistributionId

	m.RouteTableId = model.RouteTableId

	return
}
