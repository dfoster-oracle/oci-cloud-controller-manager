// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// VnicaasVnicShapeConfig Shape configuration of the service VNIC. This is used to allocate resources when the VNIC is attached to an instance.
type VnicaasVnicShapeConfig struct {

	// The percentage of concurrent connections that can be tracked to the VNIC.
	PercentageOfConnTrack *int `mandatory:"false" json:"percentageOfConnTrack"`

	// The bandwidth in Mbps that the shape can use.
	AggregateBandwidthBps *int64 `mandatory:"false" json:"aggregateBandwidthBps"`

	// It defines the internetBandwidth limit.
	InternetBandwidthBps *int64 `mandatory:"false" json:"internetBandwidthBps"`

	// Fleet name in which the servie VNIC will be attached
	FleetName *string `mandatory:"false" json:"fleetName"`
}

func (m VnicaasVnicShapeConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VnicaasVnicShapeConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
