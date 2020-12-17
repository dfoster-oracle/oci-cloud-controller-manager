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

// UpdateDrgRouteDistributionStatementDetails Route distribution statements to be updated in the Route Distribution.
type UpdateDrgRouteDistributionStatementDetails struct {

	// The Oracle-assigned ID of each DRG Route Distribution statement to be updated.
	Id *string `mandatory:"true" json:"id"`

	// The action is applied only if all of the matchCriteria is met.
	// If there are no matchCriteria in a statement, that implies match ALL.
	MatchCriteria []DrgRouteDistributionMatchCriteria `mandatory:"false" json:"matchCriteria"`

	// The priority of the statement you'd like to update.
	Priority *int `mandatory:"false" json:"priority"`
}

func (m UpdateDrgRouteDistributionStatementDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDrgRouteDistributionStatementDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		MatchCriteria []drgroutedistributionmatchcriteria `json:"matchCriteria"`
		Priority      *int                                `json:"priority"`
		Id            *string                             `json:"id"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.MatchCriteria = make([]DrgRouteDistributionMatchCriteria, len(model.MatchCriteria))
	for i, n := range model.MatchCriteria {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.MatchCriteria[i] = nn.(DrgRouteDistributionMatchCriteria)
		} else {
			m.MatchCriteria[i] = nil
		}
	}

	m.Priority = model.Priority

	m.Id = model.Id

	return
}
