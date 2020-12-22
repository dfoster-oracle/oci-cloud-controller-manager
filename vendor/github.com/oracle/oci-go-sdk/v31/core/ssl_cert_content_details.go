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

// SslCertContentDetails The encoded SslCertContent of ClientVpnEndpoint which is used in LDAP.
type SslCertContentDetails interface {
}

type sslcertcontentdetails struct {
	JsonData    []byte
	ContentType string `json:"contentType"`
}

// UnmarshalJSON unmarshals json
func (m *sslcertcontentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersslcertcontentdetails sslcertcontentdetails
	s := struct {
		Model Unmarshalersslcertcontentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ContentType = s.Model.ContentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sslcertcontentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ContentType {
	case "BASE64ENCODED":
		mm := Base64EncodedSslCertContentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m sslcertcontentdetails) String() string {
	return common.PointerString(m)
}

// SslCertContentDetailsContentTypeEnum Enum with underlying type: string
type SslCertContentDetailsContentTypeEnum string

// Set of constants representing the allowable values for SslCertContentDetailsContentTypeEnum
const (
	SslCertContentDetailsContentTypeBase64encoded SslCertContentDetailsContentTypeEnum = "BASE64ENCODED"
)

var mappingSslCertContentDetailsContentType = map[string]SslCertContentDetailsContentTypeEnum{
	"BASE64ENCODED": SslCertContentDetailsContentTypeBase64encoded,
}

// GetSslCertContentDetailsContentTypeEnumValues Enumerates the set of values for SslCertContentDetailsContentTypeEnum
func GetSslCertContentDetailsContentTypeEnumValues() []SslCertContentDetailsContentTypeEnum {
	values := make([]SslCertContentDetailsContentTypeEnum, 0)
	for _, v := range mappingSslCertContentDetailsContentType {
		values = append(values, v)
	}
	return values
}
