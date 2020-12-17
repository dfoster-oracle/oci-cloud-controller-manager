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

// DefaultDrgRouteTables The default DRG Route Tables for this DRG. Each network type
// has a default DRG Route Table.
// You can update a network type to use a different DRG Route Table, but
// each network type must have a default DRG Route Table. You cannot delete
// a default route table.
// A DRG Route Table with a static route rule or import distribution for an RPC attachment cannot be listed as
// the default Route Table for Virtual Circuit and IPSec Tunnels.
type DefaultDrgRouteTables struct {

	// The OCID of the default DRG Route table to be assigned to DRG Attachments
	// of type VCN on creation.
	Vcn *string `mandatory:"false" json:"vcn"`

	// The OCID of the default DRG Route table to be assigned to DRG Attachments
	// of type IPSec on creation.
	IpsecTunnel *string `mandatory:"false" json:"ipsecTunnel"`

	// The OCID of the default DRG Route table to be assigned to DRG Attachments
	// of type RPC on creation.
	VirtualCircuit *string `mandatory:"false" json:"virtualCircuit"`

	// The OCID of the default DRG Route table to be assigned to DRG Attachments
	// of type VC on creation.
	RemotePeeringConnection *string `mandatory:"false" json:"remotePeeringConnection"`
}

func (m DefaultDrgRouteTables) String() string {
	return common.PointerString(m)
}
