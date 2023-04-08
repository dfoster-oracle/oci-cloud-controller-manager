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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateInternalDrgRouteTableDetails Details to create and update internal Drg route table. Partitioned DRG Route Tables are supported when specifying the sharding information.
type CreateInternalDrgRouteTableDetails struct {

	// The label of the drg attachment.
	DrgAttachmentLabel *int64 `mandatory:"true" json:"drgAttachmentLabel"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG which contains this route table.
	DrgId *string `mandatory:"true" json:"drgId"`

	// The collection of rules which will be used by VCN Dataplane to route DRG traffic.
	Rules []InternalDrgRouteRule `mandatory:"true" json:"rules"`

	// The sequence number for the DRG Route Table update (version of the DRG Route Table). Only supported for partitioned route tables.
	SequenceNumber *int64 `mandatory:"false" json:"sequenceNumber"`

	// The total number of shards/partitions for the specified DRG Route Table. Only supported for partitioned route tables.
	ShardsTotal *int64 `mandatory:"false" json:"shardsTotal"`

	// The shard number for the DRG Route Table shard. Only supported for partitioned route tables.
	ShardId *int64 `mandatory:"false" json:"shardId"`

	// The DRG Route Table partitions's physical availability domain. This attribute will be null if this is a non-partitioned DRG Route Table.
	// Example: `PHX-AD-1`
	InternalAvailabilityDomain *string `mandatory:"false" json:"internalAvailabilityDomain"`
}

func (m CreateInternalDrgRouteTableDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInternalDrgRouteTableDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
