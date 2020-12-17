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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v31/common"
)

// DrgAttachmentNetworkUpdateDetails The representation of DrgAttachmentNetworkUpdateDetails
type DrgAttachmentNetworkUpdateDetails interface {
}

type drgattachmentnetworkupdatedetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *drgattachmentnetworkupdatedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdrgattachmentnetworkupdatedetails drgattachmentnetworkupdatedetails
	s := struct {
		Model Unmarshalerdrgattachmentnetworkupdatedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *drgattachmentnetworkupdatedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "INTERNET":
		mm := InternetDrgAttachmentNetworkUpdateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VIRTUAL_CIRCUIT":
		mm := VirtualCircuitDrgAttachmentNetworkUpdateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VCN":
		mm := VcnDrgAttachmentNetworkUpdateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m drgattachmentnetworkupdatedetails) String() string {
	return common.PointerString(m)
}

// DrgAttachmentNetworkUpdateDetailsTypeEnum Enum with underlying type: string
type DrgAttachmentNetworkUpdateDetailsTypeEnum string

// Set of constants representing the allowable values for DrgAttachmentNetworkUpdateDetailsTypeEnum
const (
	DrgAttachmentNetworkUpdateDetailsTypeVcn            DrgAttachmentNetworkUpdateDetailsTypeEnum = "VCN"
	DrgAttachmentNetworkUpdateDetailsTypeVirtualCircuit DrgAttachmentNetworkUpdateDetailsTypeEnum = "VIRTUAL_CIRCUIT"
	DrgAttachmentNetworkUpdateDetailsTypeInternet       DrgAttachmentNetworkUpdateDetailsTypeEnum = "INTERNET"
)

var mappingDrgAttachmentNetworkUpdateDetailsType = map[string]DrgAttachmentNetworkUpdateDetailsTypeEnum{
	"VCN":             DrgAttachmentNetworkUpdateDetailsTypeVcn,
	"VIRTUAL_CIRCUIT": DrgAttachmentNetworkUpdateDetailsTypeVirtualCircuit,
	"INTERNET":        DrgAttachmentNetworkUpdateDetailsTypeInternet,
}

// GetDrgAttachmentNetworkUpdateDetailsTypeEnumValues Enumerates the set of values for DrgAttachmentNetworkUpdateDetailsTypeEnum
func GetDrgAttachmentNetworkUpdateDetailsTypeEnumValues() []DrgAttachmentNetworkUpdateDetailsTypeEnum {
	values := make([]DrgAttachmentNetworkUpdateDetailsTypeEnum, 0)
	for _, v := range mappingDrgAttachmentNetworkUpdateDetailsType {
		values = append(values, v)
	}
	return values
}
