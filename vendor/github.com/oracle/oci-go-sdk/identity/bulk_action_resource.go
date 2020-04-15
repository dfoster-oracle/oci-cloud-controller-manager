// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// BulkActionResource The bulk action resource entity.
type BulkActionResource struct {

	// The resource type.
	EntityType *string `mandatory:"true" json:"entityType"`

	// The resource identifier.
	Identifier *string `mandatory:"false" json:"identifier"`

	// OCI region identifier.
	// See full region list: https://docs.cloud.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm
	// Default value is Tenancy's home region
	Region *string `mandatory:"false" json:"region"`

	// Additional information that helps to identity the resource.
	// Most OCI resources can be identified by a resource identifier(ocid).
	// For a resource type that can NOT be identified by an ocid,
	// use ListBulkActionResourceTypes API or refer to the resource's
	// public API document to find the required metadata keys.
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m BulkActionResource) String() string {
	return common.PointerString(m)
}
