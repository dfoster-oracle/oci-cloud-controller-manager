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

// ContainerRepositoryReadme Container repository readme.
type ContainerRepositoryReadme struct {

	// Readme content. Avoid entering confidential information.
	Content *string `mandatory:"true" json:"content"`

	// Readme format. Supported formats are text/plain and text/markdown.
	Format ContainerRepositoryReadmeFormatEnum `mandatory:"true" json:"format"`
}

func (m ContainerRepositoryReadme) String() string {
	return common.PointerString(m)
}

// ContainerRepositoryReadmeFormatEnum Enum with underlying type: string
type ContainerRepositoryReadmeFormatEnum string

// Set of constants representing the allowable values for ContainerRepositoryReadmeFormatEnum
const (
	ContainerRepositoryReadmeFormatMarkdown ContainerRepositoryReadmeFormatEnum = "TEXT_MARKDOWN"
	ContainerRepositoryReadmeFormatPlain    ContainerRepositoryReadmeFormatEnum = "TEXT_PLAIN"
)

var mappingContainerRepositoryReadmeFormat = map[string]ContainerRepositoryReadmeFormatEnum{
	"TEXT_MARKDOWN": ContainerRepositoryReadmeFormatMarkdown,
	"TEXT_PLAIN":    ContainerRepositoryReadmeFormatPlain,
}

// GetContainerRepositoryReadmeFormatEnumValues Enumerates the set of values for ContainerRepositoryReadmeFormatEnum
func GetContainerRepositoryReadmeFormatEnumValues() []ContainerRepositoryReadmeFormatEnum {
	values := make([]ContainerRepositoryReadmeFormatEnum, 0)
	for _, v := range mappingContainerRepositoryReadmeFormat {
		values = append(values, v)
	}
	return values
}
