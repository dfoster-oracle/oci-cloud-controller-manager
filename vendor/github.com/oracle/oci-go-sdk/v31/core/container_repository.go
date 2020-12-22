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

// ContainerRepository Container repository metadata.
type ContainerRepository struct {

	// The OCID of the compartment in which the container repository exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The id of the user or principal that created the resource.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The container repository name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container repository.
	// Example: `ocid1.containerrepo.oc1..exampleuniqueID`
	Id *string `mandatory:"true" json:"id"`

	// Total number of images.
	ImageCount *int `mandatory:"true" json:"imageCount"`

	// Whether the repository is immutable. Images cannot be overwritten in an immutable repository.
	IsImmutable *bool `mandatory:"true" json:"isImmutable"`

	// Whether the repository is public. A public repository allows unauthenticated access.
	IsPublic *bool `mandatory:"true" json:"isPublic"`

	// Total number of layers.
	LayerCount *int `mandatory:"true" json:"layerCount"`

	// Total storage in bytes consumed by layers.
	LayersSizeInBytes *int64 `mandatory:"true" json:"layersSizeInBytes"`

	// The current state of the container repository.
	LifecycleState ContainerRepositoryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// An RFC 3339 timestamp indicating when the repository was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// An RFC 3339 timestamp indicating when an image was last pushed to the repository.
	TimeLastPushed *common.SDKTime `mandatory:"true" json:"timeLastPushed"`

	Readme *ContainerRepositoryReadme `mandatory:"false" json:"readme"`
}

func (m ContainerRepository) String() string {
	return common.PointerString(m)
}

// ContainerRepositoryLifecycleStateEnum Enum with underlying type: string
type ContainerRepositoryLifecycleStateEnum string

// Set of constants representing the allowable values for ContainerRepositoryLifecycleStateEnum
const (
	ContainerRepositoryLifecycleStateAvailable ContainerRepositoryLifecycleStateEnum = "AVAILABLE"
	ContainerRepositoryLifecycleStateDeleting  ContainerRepositoryLifecycleStateEnum = "DELETING"
	ContainerRepositoryLifecycleStateDeleted   ContainerRepositoryLifecycleStateEnum = "DELETED"
)

var mappingContainerRepositoryLifecycleState = map[string]ContainerRepositoryLifecycleStateEnum{
	"AVAILABLE": ContainerRepositoryLifecycleStateAvailable,
	"DELETING":  ContainerRepositoryLifecycleStateDeleting,
	"DELETED":   ContainerRepositoryLifecycleStateDeleted,
}

// GetContainerRepositoryLifecycleStateEnumValues Enumerates the set of values for ContainerRepositoryLifecycleStateEnum
func GetContainerRepositoryLifecycleStateEnumValues() []ContainerRepositoryLifecycleStateEnum {
	values := make([]ContainerRepositoryLifecycleStateEnum, 0)
	for _, v := range mappingContainerRepositoryLifecycleState {
		values = append(values, v)
	}
	return values
}
