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

// ClientVpnEndpointStatus The status of ClientVpnEndpoint.
type ClientVpnEndpointStatus struct {

	// The number of current connections on given clientVpnEndpoint.
	CurrentConnections *int `mandatory:"true" json:"currentConnections"`

	// The list of active users.
	ActiveUsers []VpnActiveUser `mandatory:"true" json:"activeUsers"`

	// The current state of the ClientVPNendpoint.
	LifecycleState ClientVpnEndpointStatusLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

func (m ClientVpnEndpointStatus) String() string {
	return common.PointerString(m)
}

// ClientVpnEndpointStatusLifecycleStateEnum Enum with underlying type: string
type ClientVpnEndpointStatusLifecycleStateEnum string

// Set of constants representing the allowable values for ClientVpnEndpointStatusLifecycleStateEnum
const (
	ClientVpnEndpointStatusLifecycleStateCreating ClientVpnEndpointStatusLifecycleStateEnum = "CREATING"
	ClientVpnEndpointStatusLifecycleStateActive   ClientVpnEndpointStatusLifecycleStateEnum = "ACTIVE"
	ClientVpnEndpointStatusLifecycleStateInactive ClientVpnEndpointStatusLifecycleStateEnum = "INACTIVE"
	ClientVpnEndpointStatusLifecycleStateFailed   ClientVpnEndpointStatusLifecycleStateEnum = "FAILED"
	ClientVpnEndpointStatusLifecycleStateDeleted  ClientVpnEndpointStatusLifecycleStateEnum = "DELETED"
	ClientVpnEndpointStatusLifecycleStateDeleting ClientVpnEndpointStatusLifecycleStateEnum = "DELETING"
	ClientVpnEndpointStatusLifecycleStateUpdating ClientVpnEndpointStatusLifecycleStateEnum = "UPDATING"
)

var mappingClientVpnEndpointStatusLifecycleState = map[string]ClientVpnEndpointStatusLifecycleStateEnum{
	"CREATING": ClientVpnEndpointStatusLifecycleStateCreating,
	"ACTIVE":   ClientVpnEndpointStatusLifecycleStateActive,
	"INACTIVE": ClientVpnEndpointStatusLifecycleStateInactive,
	"FAILED":   ClientVpnEndpointStatusLifecycleStateFailed,
	"DELETED":  ClientVpnEndpointStatusLifecycleStateDeleted,
	"DELETING": ClientVpnEndpointStatusLifecycleStateDeleting,
	"UPDATING": ClientVpnEndpointStatusLifecycleStateUpdating,
}

// GetClientVpnEndpointStatusLifecycleStateEnumValues Enumerates the set of values for ClientVpnEndpointStatusLifecycleStateEnum
func GetClientVpnEndpointStatusLifecycleStateEnumValues() []ClientVpnEndpointStatusLifecycleStateEnum {
	values := make([]ClientVpnEndpointStatusLifecycleStateEnum, 0)
	for _, v := range mappingClientVpnEndpointStatusLifecycleState {
		values = append(values, v)
	}
	return values
}
