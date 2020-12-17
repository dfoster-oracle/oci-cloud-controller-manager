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

// InstanceAgentPluginConfigDetails Instance Agent plugin configuration.
type InstanceAgentPluginConfigDetails struct {

	// The plugin name.
	Name *string `mandatory:"true" json:"name"`

	// Whether the plugin should be enabled or disabled.
	DesiredState InstanceAgentPluginConfigDetailsDesiredStateEnum `mandatory:"true" json:"desiredState"`
}

func (m InstanceAgentPluginConfigDetails) String() string {
	return common.PointerString(m)
}

// InstanceAgentPluginConfigDetailsDesiredStateEnum Enum with underlying type: string
type InstanceAgentPluginConfigDetailsDesiredStateEnum string

// Set of constants representing the allowable values for InstanceAgentPluginConfigDetailsDesiredStateEnum
const (
	InstanceAgentPluginConfigDetailsDesiredStateEnabled  InstanceAgentPluginConfigDetailsDesiredStateEnum = "ENABLED"
	InstanceAgentPluginConfigDetailsDesiredStateDisabled InstanceAgentPluginConfigDetailsDesiredStateEnum = "DISABLED"
)

var mappingInstanceAgentPluginConfigDetailsDesiredState = map[string]InstanceAgentPluginConfigDetailsDesiredStateEnum{
	"ENABLED":  InstanceAgentPluginConfigDetailsDesiredStateEnabled,
	"DISABLED": InstanceAgentPluginConfigDetailsDesiredStateDisabled,
}

// GetInstanceAgentPluginConfigDetailsDesiredStateEnumValues Enumerates the set of values for InstanceAgentPluginConfigDetailsDesiredStateEnum
func GetInstanceAgentPluginConfigDetailsDesiredStateEnumValues() []InstanceAgentPluginConfigDetailsDesiredStateEnum {
	values := make([]InstanceAgentPluginConfigDetailsDesiredStateEnum, 0)
	for _, v := range mappingInstanceAgentPluginConfigDetailsDesiredState {
		values = append(values, v)
	}
	return values
}
