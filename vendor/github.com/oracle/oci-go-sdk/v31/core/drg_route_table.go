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

// DrgRouteTable The DRG Route Table is the data table stored in the DRG that manages routing inside the DRG.
// The DRG Route Table lists routes to particular network destinations. Drg Route Tables are assigned to DRG Attachments. Routing decisions
// for packets entering on a DRG Attachment will be determined based on the DRG Attachment's DRG Route Table. Routing decisions are made by
// doing an LPM match on the destination address.
// Each DRG Attachment will inject routes in the table given there is an distribution for that DRG Attachment specified in the Route Table's
// importDrgRouteDistribution. A customer can specify static routes to insert routes in the DRG Route Tables. A customer will continue to manage routing
// through the VCN by specifying static routes in the VCN ingress table associated with the VCN Attachment as they do today.
// The DRG Route Table is created in the compartment of the DRG.
// There must always be a default DRG Route Table for each attachment type.
type DrgRouteTable struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the
	// DRG Route Table.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment the DRG Route Table belongs to. The DRG Route Table
	// always lives in the DRG's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DRG the route table belongs to.
	DrgId *string `mandatory:"true" json:"drgId"`

	// The date and time the DRG Route Table was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The DRG Route Table's current state.
	LifecycleState DrgRouteTableLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// If a user wants traffic to be routed using ECMP across their Virtual Circuits or IPSec Tunnels to
	// their on-prem networks, the user should enable ECMP on the route table to which these attachments
	// import routes.
	IsEcmpEnabled *bool `mandatory:"true" json:"isEcmpEnabled"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID of the import route distribution used to specify how incoming route advertisements through
	// referenced attachements are inserted into the DRG route table.
	ImportDrgRouteDistributionId *string `mandatory:"false" json:"importDrgRouteDistributionId"`
}

func (m DrgRouteTable) String() string {
	return common.PointerString(m)
}

// DrgRouteTableLifecycleStateEnum Enum with underlying type: string
type DrgRouteTableLifecycleStateEnum string

// Set of constants representing the allowable values for DrgRouteTableLifecycleStateEnum
const (
	DrgRouteTableLifecycleStateProvisioning DrgRouteTableLifecycleStateEnum = "PROVISIONING"
	DrgRouteTableLifecycleStateAvailable    DrgRouteTableLifecycleStateEnum = "AVAILABLE"
	DrgRouteTableLifecycleStateTerminating  DrgRouteTableLifecycleStateEnum = "TERMINATING"
	DrgRouteTableLifecycleStateTerminated   DrgRouteTableLifecycleStateEnum = "TERMINATED"
)

var mappingDrgRouteTableLifecycleState = map[string]DrgRouteTableLifecycleStateEnum{
	"PROVISIONING": DrgRouteTableLifecycleStateProvisioning,
	"AVAILABLE":    DrgRouteTableLifecycleStateAvailable,
	"TERMINATING":  DrgRouteTableLifecycleStateTerminating,
	"TERMINATED":   DrgRouteTableLifecycleStateTerminated,
}

// GetDrgRouteTableLifecycleStateEnumValues Enumerates the set of values for DrgRouteTableLifecycleStateEnum
func GetDrgRouteTableLifecycleStateEnumValues() []DrgRouteTableLifecycleStateEnum {
	values := make([]DrgRouteTableLifecycleStateEnum, 0)
	for _, v := range mappingDrgRouteTableLifecycleState {
		values = append(values, v)
	}
	return values
}
