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

// CreateDrgRouteDistributionDetails Details request to create a Route Distribution.
type CreateDrgRouteDistributionDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DRG the DRG Route Table belongs to.
	DrgId *string `mandatory:"true" json:"drgId"`

	// Whether this distribution defines how routes get imported into route tables or exported through DRG Attachments
	DistributionType CreateDrgRouteDistributionDetailsDistributionTypeEnum `mandatory:"true" json:"distributionType"`

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

func (m CreateDrgRouteDistributionDetails) String() string {
	return common.PointerString(m)
}

// CreateDrgRouteDistributionDetailsDistributionTypeEnum Enum with underlying type: string
type CreateDrgRouteDistributionDetailsDistributionTypeEnum string

// Set of constants representing the allowable values for CreateDrgRouteDistributionDetailsDistributionTypeEnum
const (
	CreateDrgRouteDistributionDetailsDistributionTypeImport CreateDrgRouteDistributionDetailsDistributionTypeEnum = "IMPORT"
)

var mappingCreateDrgRouteDistributionDetailsDistributionType = map[string]CreateDrgRouteDistributionDetailsDistributionTypeEnum{
	"IMPORT": CreateDrgRouteDistributionDetailsDistributionTypeImport,
}

// GetCreateDrgRouteDistributionDetailsDistributionTypeEnumValues Enumerates the set of values for CreateDrgRouteDistributionDetailsDistributionTypeEnum
func GetCreateDrgRouteDistributionDetailsDistributionTypeEnumValues() []CreateDrgRouteDistributionDetailsDistributionTypeEnum {
	values := make([]CreateDrgRouteDistributionDetailsDistributionTypeEnum, 0)
	for _, v := range mappingCreateDrgRouteDistributionDetailsDistributionType {
		values = append(values, v)
	}
	return values
}
