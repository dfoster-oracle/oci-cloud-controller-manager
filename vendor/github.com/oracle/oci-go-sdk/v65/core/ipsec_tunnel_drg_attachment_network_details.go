// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IpsecTunnelDrgAttachmentNetworkDetails Specifies the IPSec tunnel attached to the DRG.
type IpsecTunnelDrgAttachmentNetworkDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network attached to the DRG.
	Id *string `mandatory:"true" json:"id"`

	// The IPSec connection that contains the attached IPSec tunnel.
	IpsecConnectionId *string `mandatory:"false" json:"ipsecConnectionId"`

	// Routes which may be imported from the attachment (subject to import policy) appear in the route reflectors
	// tagged with the attachment's import route target.
	ImportRouteTarget *string `mandatory:"false" json:"importRouteTarget"`

	// Routes which are exported to the attachment are exported to the route reflectors
	// with the route target set to the value of the attachment's export route target.
	ExportRouteTarget *string `mandatory:"false" json:"exportRouteTarget"`

	// The MPLS label of the DRG attachment.
	MplsLabel *int `mandatory:"false" json:"mplsLabel"`

	// The BGP ASN to use for the IPSec connection's route target.
	RegionalOciAsn *string `mandatory:"false" json:"regionalOciAsn"`

	// IPv4 address used to encapsulate ingress traffic to the DRG through this attachment
	IngressVip *string `mandatory:"false" json:"ingressVip"`

	// Whether traffic from this network is forwarded to the El Paso Gamma VIPs (or not)
	IsGammaDrg *bool `mandatory:"false" json:"isGammaDrg"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the virtual-circuit drg attachment.
	TransportAttachmentId *string `mandatory:"false" json:"transportAttachmentId"`
}

//GetId returns Id
func (m IpsecTunnelDrgAttachmentNetworkDetails) GetId() *string {
	return m.Id
}

func (m IpsecTunnelDrgAttachmentNetworkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IpsecTunnelDrgAttachmentNetworkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IpsecTunnelDrgAttachmentNetworkDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIpsecTunnelDrgAttachmentNetworkDetails IpsecTunnelDrgAttachmentNetworkDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeIpsecTunnelDrgAttachmentNetworkDetails
	}{
		"IPSEC_TUNNEL",
		(MarshalTypeIpsecTunnelDrgAttachmentNetworkDetails)(m),
	}

	return json.Marshal(&s)
}
