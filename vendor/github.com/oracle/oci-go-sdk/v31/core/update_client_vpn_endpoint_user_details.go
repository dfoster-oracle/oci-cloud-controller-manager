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

// UpdateClientVpnEndpointUserDetails The request to update clientVpnEndpointUser.
type UpdateClientVpnEndpointUserDetails struct {

	// The password of the user want to create.
	Password *string `mandatory:"false" json:"password"`

	// Whether to log in the user by cert-authentication only or not.
	IsCertificateAuthenticationOnly *bool `mandatory:"false" json:"isCertificateAuthenticationOnly"`

	// Whether to enable auto-login feature or not.
	IsAutomaticLoginEnabled *bool `mandatory:"false" json:"isAutomaticLoginEnabled"`
}

func (m UpdateClientVpnEndpointUserDetails) String() string {
	return common.PointerString(m)
}
