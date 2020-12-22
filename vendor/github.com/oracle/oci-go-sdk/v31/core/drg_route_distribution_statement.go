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

// DrgRouteDistributionStatement A statement within a Route Distribution. All matchCriteria in a statement must be met
// in order for the action to take place.
type DrgRouteDistributionStatement struct {

	// The action is applied only if all of the matchCriteria is met.
	// If there are no matchCriteria in a statement, that implies match ALL.
	MatchCriteria []DrgRouteDistributionMatchCriteria `mandatory:"true" json:"matchCriteria"`

	// Accept: import/export the route "as is"
	Action DrgRouteDistributionStatementActionEnum `mandatory:"true" json:"action"`

	// This field is used to specify the priority of each statement in a route distribution.
	// The priority will be represented as a number between 0 and 65535 where a lower number
	// indicates a higher priority. When a route is processed, statements are applied in the order
	// defined by their priority. The first matching rule dictates the action that will be taken
	// on the route.
	Priority *int `mandatory:"true" json:"priority"`

	// The Oracle-assigned ID of the Route Distribution Statement.
	Id *string `mandatory:"true" json:"id"`
}

func (m DrgRouteDistributionStatement) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DrgRouteDistributionStatement) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		MatchCriteria []drgroutedistributionmatchcriteria     `json:"matchCriteria"`
		Action        DrgRouteDistributionStatementActionEnum `json:"action"`
		Priority      *int                                    `json:"priority"`
		Id            *string                                 `json:"id"`
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

	m.Action = model.Action

	m.Priority = model.Priority

	m.Id = model.Id

	return
}

// DrgRouteDistributionStatementActionEnum Enum with underlying type: string
type DrgRouteDistributionStatementActionEnum string

// Set of constants representing the allowable values for DrgRouteDistributionStatementActionEnum
const (
	DrgRouteDistributionStatementActionAccept DrgRouteDistributionStatementActionEnum = "ACCEPT"
)

var mappingDrgRouteDistributionStatementAction = map[string]DrgRouteDistributionStatementActionEnum{
	"ACCEPT": DrgRouteDistributionStatementActionAccept,
}

// GetDrgRouteDistributionStatementActionEnumValues Enumerates the set of values for DrgRouteDistributionStatementActionEnum
func GetDrgRouteDistributionStatementActionEnumValues() []DrgRouteDistributionStatementActionEnum {
	values := make([]DrgRouteDistributionStatementActionEnum, 0)
	for _, v := range mappingDrgRouteDistributionStatementAction {
		values = append(values, v)
	}
	return values
}
