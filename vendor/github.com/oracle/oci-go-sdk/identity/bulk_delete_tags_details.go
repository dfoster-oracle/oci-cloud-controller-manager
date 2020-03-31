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

// BulkDeleteTagsDetails Properties for deleting tags in bulk
type BulkDeleteTagsDetails struct {

	// The OCIDs of the tag definitions to delete
	TagDefinitionIds []string `mandatory:"true" json:"tagDefinitionIds"`
}

func (m BulkDeleteTagsDetails) String() string {
	return common.PointerString(m)
}
