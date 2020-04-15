// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Vlan For use with Oracle Cloud Infrastructure VCN.
// A virtual LAN (VLAN) is a broadcast domain that is created by partitioning and isolating a network at the data
// link layer more commonly referred to as layer 2 network. It is similar to a subnet which partitions a network
// at the IP layer ( Layer 3).
// VLAN's work by using vlan tags. These vlan tags are applied to network packets and the network switches route
// these packets based on the vlan tags
type Vlan struct {

	// The VLAN's CIDR block.
	// Example: `192.168.1.0/24`
	CidrBlock *string `mandatory:"true" json:"cidrBlock"`

	// The OCID of the compartment containing the subnet.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The VLAN's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The VLAN's current state.
	LifecycleState VlanLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the VCN the subnet is in.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The availability domain of the VLAN.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

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

	// A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more
	// information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The VLAN tag of this VLAN
	// Example: `100`
	VlanTag *int `mandatory:"false" json:"vlanTag"`

	// The OCID of the route table that the subnet uses.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// The date and time the subnet was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m Vlan) String() string {
	return common.PointerString(m)
}

// VlanLifecycleStateEnum Enum with underlying type: string
type VlanLifecycleStateEnum string

// Set of constants representing the allowable values for VlanLifecycleStateEnum
const (
	VlanLifecycleStateProvisioning VlanLifecycleStateEnum = "PROVISIONING"
	VlanLifecycleStateAvailable    VlanLifecycleStateEnum = "AVAILABLE"
	VlanLifecycleStateTerminating  VlanLifecycleStateEnum = "TERMINATING"
	VlanLifecycleStateTerminated   VlanLifecycleStateEnum = "TERMINATED"
	VlanLifecycleStateUpdating     VlanLifecycleStateEnum = "UPDATING"
)

var mappingVlanLifecycleState = map[string]VlanLifecycleStateEnum{
	"PROVISIONING": VlanLifecycleStateProvisioning,
	"AVAILABLE":    VlanLifecycleStateAvailable,
	"TERMINATING":  VlanLifecycleStateTerminating,
	"TERMINATED":   VlanLifecycleStateTerminated,
	"UPDATING":     VlanLifecycleStateUpdating,
}

// GetVlanLifecycleStateEnumValues Enumerates the set of values for VlanLifecycleStateEnum
func GetVlanLifecycleStateEnumValues() []VlanLifecycleStateEnum {
	values := make([]VlanLifecycleStateEnum, 0)
	for _, v := range mappingVlanLifecycleState {
		values = append(values, v)
	}
	return values
}
