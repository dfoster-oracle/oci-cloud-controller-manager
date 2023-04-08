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

// DrgUrtRouteRuleDetails Details of Routing information needed by VCN DP to route DRG traffic.
type DrgUrtRouteRuleDetails struct {

	// The label of the attachment.
	DrgAttachmentLabel *int64 `mandatory:"true" json:"drgAttachmentLabel"`

	// Routing information needed by VCN DP to route DRG traffic.
	Rules []DrgUrtRouteRule `mandatory:"true" json:"rules"`

	// For debugging purposes.
	VrfLabel *int64 `mandatory:"false" json:"vrfLabel"`

	// Not supported now, reserved for future use.
	// The date and time the routes were last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	LastModified *common.SDKTime `mandatory:"false" json:"lastModified"`

	// Not supported now, reserved for future use
	PaginationToken *string `mandatory:"false" json:"paginationToken"`
}

func (m DrgUrtRouteRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrgUrtRouteRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
