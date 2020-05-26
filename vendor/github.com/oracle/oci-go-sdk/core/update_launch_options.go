// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateLaunchOptions Options for tuning compatibility and performance of VM shapes.
type UpdateLaunchOptions struct {

	// Emulation type for volume.
	// * `ISCSI` - ISCSI attached block storage device. This is the default for Boot Volumes and Remote Block
	// Storage volumes on Oracle provided images.
	// * `PARAVIRTUALIZED` - Paravirtualized disk.
	BootVolumeType UpdateLaunchOptionsBootVolumeTypeEnum `mandatory:"false" json:"bootVolumeType,omitempty"`

	// Emulation type for the physical network interface card (NIC).
	// * `VFIO` - Direct attached Virtual Function network controller. This is the networking type
	// when you launch an instance using hardware-assisted (SR-IOV) networking.
	// * `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using virtio drivers.
	NetworkType UpdateLaunchOptionsNetworkTypeEnum `mandatory:"false" json:"networkType,omitempty"`

	// Whether to enable in-transit encryption for the boot volume's paravirtualized attachment. The default value is false.
	IsPvEncryptionInTransitEnabled *bool `mandatory:"false" json:"isPvEncryptionInTransitEnabled"`

	// Whether to enable consistent volume naming feature. Defaults to false.
	IsConsistentVolumeNamingEnabled *bool `mandatory:"false" json:"isConsistentVolumeNamingEnabled"`
}

func (m UpdateLaunchOptions) String() string {
	return common.PointerString(m)
}

// UpdateLaunchOptionsBootVolumeTypeEnum Enum with underlying type: string
type UpdateLaunchOptionsBootVolumeTypeEnum string

// Set of constants representing the allowable values for UpdateLaunchOptionsBootVolumeTypeEnum
const (
	UpdateLaunchOptionsBootVolumeTypeIscsi           UpdateLaunchOptionsBootVolumeTypeEnum = "ISCSI"
	UpdateLaunchOptionsBootVolumeTypeParavirtualized UpdateLaunchOptionsBootVolumeTypeEnum = "PARAVIRTUALIZED"
)

var mappingUpdateLaunchOptionsBootVolumeType = map[string]UpdateLaunchOptionsBootVolumeTypeEnum{
	"ISCSI":           UpdateLaunchOptionsBootVolumeTypeIscsi,
	"PARAVIRTUALIZED": UpdateLaunchOptionsBootVolumeTypeParavirtualized,
}

// GetUpdateLaunchOptionsBootVolumeTypeEnumValues Enumerates the set of values for UpdateLaunchOptionsBootVolumeTypeEnum
func GetUpdateLaunchOptionsBootVolumeTypeEnumValues() []UpdateLaunchOptionsBootVolumeTypeEnum {
	values := make([]UpdateLaunchOptionsBootVolumeTypeEnum, 0)
	for _, v := range mappingUpdateLaunchOptionsBootVolumeType {
		values = append(values, v)
	}
	return values
}

// UpdateLaunchOptionsNetworkTypeEnum Enum with underlying type: string
type UpdateLaunchOptionsNetworkTypeEnum string

// Set of constants representing the allowable values for UpdateLaunchOptionsNetworkTypeEnum
const (
	UpdateLaunchOptionsNetworkTypeVfio            UpdateLaunchOptionsNetworkTypeEnum = "VFIO"
	UpdateLaunchOptionsNetworkTypeParavirtualized UpdateLaunchOptionsNetworkTypeEnum = "PARAVIRTUALIZED"
)

var mappingUpdateLaunchOptionsNetworkType = map[string]UpdateLaunchOptionsNetworkTypeEnum{
	"VFIO":            UpdateLaunchOptionsNetworkTypeVfio,
	"PARAVIRTUALIZED": UpdateLaunchOptionsNetworkTypeParavirtualized,
}

// GetUpdateLaunchOptionsNetworkTypeEnumValues Enumerates the set of values for UpdateLaunchOptionsNetworkTypeEnum
func GetUpdateLaunchOptionsNetworkTypeEnumValues() []UpdateLaunchOptionsNetworkTypeEnum {
	values := make([]UpdateLaunchOptionsNetworkTypeEnum, 0)
	for _, v := range mappingUpdateLaunchOptionsNetworkType {
		values = append(values, v)
	}
	return values
}
