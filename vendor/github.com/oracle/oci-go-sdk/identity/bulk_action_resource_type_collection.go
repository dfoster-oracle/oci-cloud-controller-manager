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

// BulkActionResourceTypeCollection Collection of resource types supported by bulk action.
type BulkActionResourceTypeCollection struct {

	// Collection of resource types supported by bulk action.
	Items []BulkActionResourceType `mandatory:"true" json:"items"`
}

func (m BulkActionResourceTypeCollection) String() string {
	return common.PointerString(m)
}
