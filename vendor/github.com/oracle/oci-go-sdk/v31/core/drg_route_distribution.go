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

// DrgRouteDistribution A Route Distribution establishes how routes get imported into DRG Route Tables and exported through the DRG attachments.
// A route distribution is a list of statements. Each statement consists of a set of matches, all of which must be True in order for
// the statement's action to take place. Each statement determines which routes are propagated.
// The user can assign a Route Distribution as a route table's import distribution. The statements in an import route distribution
// specify how how incoming route advertisements through a referenced attachement or all attachments of a certain type are inserted
// into the route table.
// The user can assign a Route Distribution as a DRG Attachment's export distribution. Export route distribution statements specify how routes
// DRG Attachment's assigned table are advertised out through the attachment. When a DRG Attachment is assigned to a table, an Route Distribution
// is created with a single statement, match (ALL→Accept).
// The two auto-generated DRG Route Tables (one as the default for VCN attachments, and the other for all other types of attachments)
// are each assigned an auto generated import Route Distribution. The default VCN table's import distribution has a statement to import routes from
// each DRG Attachment type (4 statements in total). The other table's import distribution has a statement to import routes from VCN attachments.
// The Route Distribution is created in the compartment of the DRG.
type DrgRouteDistribution struct {

	// The OCID of the DRG that contains this Route Distribution
	DrgId *string `mandatory:"true" json:"drgId"`

	// The OCID of the compartment containing the Route Distribution.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Route Distribution's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The Route Distribution's current state.
	LifecycleState DrgRouteDistributionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the Route Distribution was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Whether this distribution defines how routes get imported into route tables or exported through DRG Attachments
	DistributionType DrgRouteDistributionDistributionTypeEnum `mandatory:"true" json:"distributionType"`

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
}

func (m DrgRouteDistribution) String() string {
	return common.PointerString(m)
}

// DrgRouteDistributionLifecycleStateEnum Enum with underlying type: string
type DrgRouteDistributionLifecycleStateEnum string

// Set of constants representing the allowable values for DrgRouteDistributionLifecycleStateEnum
const (
	DrgRouteDistributionLifecycleStateProvisioning DrgRouteDistributionLifecycleStateEnum = "PROVISIONING"
	DrgRouteDistributionLifecycleStateAvailable    DrgRouteDistributionLifecycleStateEnum = "AVAILABLE"
	DrgRouteDistributionLifecycleStateTerminating  DrgRouteDistributionLifecycleStateEnum = "TERMINATING"
	DrgRouteDistributionLifecycleStateTerminated   DrgRouteDistributionLifecycleStateEnum = "TERMINATED"
)

var mappingDrgRouteDistributionLifecycleState = map[string]DrgRouteDistributionLifecycleStateEnum{
	"PROVISIONING": DrgRouteDistributionLifecycleStateProvisioning,
	"AVAILABLE":    DrgRouteDistributionLifecycleStateAvailable,
	"TERMINATING":  DrgRouteDistributionLifecycleStateTerminating,
	"TERMINATED":   DrgRouteDistributionLifecycleStateTerminated,
}

// GetDrgRouteDistributionLifecycleStateEnumValues Enumerates the set of values for DrgRouteDistributionLifecycleStateEnum
func GetDrgRouteDistributionLifecycleStateEnumValues() []DrgRouteDistributionLifecycleStateEnum {
	values := make([]DrgRouteDistributionLifecycleStateEnum, 0)
	for _, v := range mappingDrgRouteDistributionLifecycleState {
		values = append(values, v)
	}
	return values
}

// DrgRouteDistributionDistributionTypeEnum Enum with underlying type: string
type DrgRouteDistributionDistributionTypeEnum string

// Set of constants representing the allowable values for DrgRouteDistributionDistributionTypeEnum
const (
	DrgRouteDistributionDistributionTypeImport DrgRouteDistributionDistributionTypeEnum = "IMPORT"
	DrgRouteDistributionDistributionTypeExport DrgRouteDistributionDistributionTypeEnum = "EXPORT"
)

var mappingDrgRouteDistributionDistributionType = map[string]DrgRouteDistributionDistributionTypeEnum{
	"IMPORT": DrgRouteDistributionDistributionTypeImport,
	"EXPORT": DrgRouteDistributionDistributionTypeExport,
}

// GetDrgRouteDistributionDistributionTypeEnumValues Enumerates the set of values for DrgRouteDistributionDistributionTypeEnum
func GetDrgRouteDistributionDistributionTypeEnumValues() []DrgRouteDistributionDistributionTypeEnum {
	values := make([]DrgRouteDistributionDistributionTypeEnum, 0)
	for _, v := range mappingDrgRouteDistributionDistributionType {
		values = append(values, v)
	}
	return values
}
