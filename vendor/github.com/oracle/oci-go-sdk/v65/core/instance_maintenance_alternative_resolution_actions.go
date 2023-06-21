// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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
	"strings"
)

// InstanceMaintenanceAlternativeResolutionActionsEnum Enum with underlying type: string
type InstanceMaintenanceAlternativeResolutionActionsEnum string

// Set of constants representing the allowable values for InstanceMaintenanceAlternativeResolutionActionsEnum
const (
	InstanceMaintenanceAlternativeResolutionActionsRebootMigration InstanceMaintenanceAlternativeResolutionActionsEnum = "REBOOT_MIGRATION"
	InstanceMaintenanceAlternativeResolutionActionsTerminate       InstanceMaintenanceAlternativeResolutionActionsEnum = "TERMINATE"
)

var mappingInstanceMaintenanceAlternativeResolutionActionsEnum = map[string]InstanceMaintenanceAlternativeResolutionActionsEnum{
	"REBOOT_MIGRATION": InstanceMaintenanceAlternativeResolutionActionsRebootMigration,
	"TERMINATE":        InstanceMaintenanceAlternativeResolutionActionsTerminate,
}

var mappingInstanceMaintenanceAlternativeResolutionActionsEnumLowerCase = map[string]InstanceMaintenanceAlternativeResolutionActionsEnum{
	"reboot_migration": InstanceMaintenanceAlternativeResolutionActionsRebootMigration,
	"terminate":        InstanceMaintenanceAlternativeResolutionActionsTerminate,
}

// GetInstanceMaintenanceAlternativeResolutionActionsEnumValues Enumerates the set of values for InstanceMaintenanceAlternativeResolutionActionsEnum
func GetInstanceMaintenanceAlternativeResolutionActionsEnumValues() []InstanceMaintenanceAlternativeResolutionActionsEnum {
	values := make([]InstanceMaintenanceAlternativeResolutionActionsEnum, 0)
	for _, v := range mappingInstanceMaintenanceAlternativeResolutionActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceMaintenanceAlternativeResolutionActionsEnumStringValues Enumerates the set of values in String for InstanceMaintenanceAlternativeResolutionActionsEnum
func GetInstanceMaintenanceAlternativeResolutionActionsEnumStringValues() []string {
	return []string{
		"REBOOT_MIGRATION",
		"TERMINATE",
	}
}

// GetMappingInstanceMaintenanceAlternativeResolutionActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceMaintenanceAlternativeResolutionActionsEnum(val string) (InstanceMaintenanceAlternativeResolutionActionsEnum, bool) {
	enum, ok := mappingInstanceMaintenanceAlternativeResolutionActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
