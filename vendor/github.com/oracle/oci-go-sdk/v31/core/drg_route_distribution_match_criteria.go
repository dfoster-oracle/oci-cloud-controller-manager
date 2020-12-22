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

// DrgRouteDistributionMatchCriteria A matchCriteria in a Route Distribution statement.The matchCriteria outlines which routes
// should be imported or exported.
type DrgRouteDistributionMatchCriteria interface {
}

type drgroutedistributionmatchcriteria struct {
	JsonData  []byte
	MatchType string `json:"matchType"`
}

// UnmarshalJSON unmarshals json
func (m *drgroutedistributionmatchcriteria) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdrgroutedistributionmatchcriteria drgroutedistributionmatchcriteria
	s := struct {
		Model Unmarshalerdrgroutedistributionmatchcriteria
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MatchType = s.Model.MatchType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *drgroutedistributionmatchcriteria) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MatchType {
	case "DRG_ATTACHMENT_ID":
		mm := DrgAttachmentIdDrgRouteDistributionMatchCriteria{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DRG_ATTACHMENT_TYPE":
		mm := DrgAttachmentTypeDrgRouteDistributionMatchCriteria{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m drgroutedistributionmatchcriteria) String() string {
	return common.PointerString(m)
}

// DrgRouteDistributionMatchCriteriaMatchTypeEnum Enum with underlying type: string
type DrgRouteDistributionMatchCriteriaMatchTypeEnum string

// Set of constants representing the allowable values for DrgRouteDistributionMatchCriteriaMatchTypeEnum
const (
	DrgRouteDistributionMatchCriteriaMatchTypeType DrgRouteDistributionMatchCriteriaMatchTypeEnum = "DRG_ATTACHMENT_TYPE"
	DrgRouteDistributionMatchCriteriaMatchTypeId   DrgRouteDistributionMatchCriteriaMatchTypeEnum = "DRG_ATTACHMENT_ID"
)

var mappingDrgRouteDistributionMatchCriteriaMatchType = map[string]DrgRouteDistributionMatchCriteriaMatchTypeEnum{
	"DRG_ATTACHMENT_TYPE": DrgRouteDistributionMatchCriteriaMatchTypeType,
	"DRG_ATTACHMENT_ID":   DrgRouteDistributionMatchCriteriaMatchTypeId,
}

// GetDrgRouteDistributionMatchCriteriaMatchTypeEnumValues Enumerates the set of values for DrgRouteDistributionMatchCriteriaMatchTypeEnum
func GetDrgRouteDistributionMatchCriteriaMatchTypeEnumValues() []DrgRouteDistributionMatchCriteriaMatchTypeEnum {
	values := make([]DrgRouteDistributionMatchCriteriaMatchTypeEnum, 0)
	for _, v := range mappingDrgRouteDistributionMatchCriteriaMatchType {
		values = append(values, v)
	}
	return values
}
