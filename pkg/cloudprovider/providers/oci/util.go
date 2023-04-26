// Copyright 2017 Oracle and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package oci

import (
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	api "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	listersv1 "k8s.io/client-go/listers/core/v1"
)

const (
	virtualNodeOcidPrefix      = "ocid1.virtualnode."
	virtualNodeOcidDevPrefix   = "ocid1.virtualnodedev."
	virtualNodeOcidIntegPrefix = "ocid1.virtualnodeinteg."
)

// Protects Load Balancers against multiple updates in parallel
type loadBalancerLocks struct {
	locks sets.String
	mux   sync.Mutex
}

func NewLoadBalancerLocks() *loadBalancerLocks {
	return &loadBalancerLocks{
		locks: sets.NewString(),
	}
}

func (lbl *loadBalancerLocks) TryAcquire(lbname string) bool {
	lbl.mux.Lock()
	defer lbl.mux.Unlock()
	if lbl.locks.Has(lbname) {
		return false
	}
	lbl.locks.Insert(lbname)
	return true
}

func (lbl *loadBalancerLocks) Release(lbname string) {
	lbl.mux.Lock()
	defer lbl.mux.Unlock()
	lbl.locks.Delete(lbname)
}

// MapProviderIDToResourceID parses the provider id and returns the instance ocid.
func MapProviderIDToResourceID(providerID string) (string, error) {
	if providerID == "" {
		return providerID, errors.New("provider ID is empty")
	}
	if strings.HasPrefix(providerID, providerPrefix) {
		return strings.TrimPrefix(providerID, providerPrefix), nil
	}
	return providerID, nil
}

// NodeInternalIP returns the nodes internal ip
// A node managed by the CCM will always have an internal ip
// since it's not possible to deploy an instance without a private ip.
func NodeInternalIP(node *api.Node) string {
	for _, addr := range node.Status.Addresses {
		if addr.Type == api.NodeInternalIP {
			return addr.Address
		}
	}
	return ""
}

// RemoveDuplicatesFromList takes Slice and returns new Slice with no duplicate elements
// (e.g. if given list is {"a", "b", "a"}, function returns new slice with {"a", "b"}
func RemoveDuplicatesFromList(list []string) []string {
	return sets.NewString(list...).List()
}

// DeepEqualLists diffs two slices and returns bool if the slices are equal/not-equal.
// the duplicates and order of items in both lists is ignored.
func DeepEqualLists(listA, listB []string) bool {
	return sets.NewString(listA...).Equal(sets.NewString(listB...))
}

// IsVirtualNodeId Returns true if providerId is a Virtual Node OCID
func IsVirtualNodeId(resourceId string) bool {
	return strings.HasPrefix(resourceId, virtualNodeOcidPrefix) || strings.HasPrefix(resourceId, virtualNodeOcidDevPrefix) || strings.HasPrefix(resourceId, virtualNodeOcidIntegPrefix)
}

// IsVirtualNode returns true if a node object corresponds to a Virtual Node
func IsVirtualNode(node *api.Node) bool {
	resourceId, err := MapProviderIDToResourceID(node.Spec.ProviderID)
	if err != nil {
		// OVK ensures the providerId is set on a Virtual Node
		// If the providerId is empty it is safe to assume the node is not virtual
		return false
	}
	return IsVirtualNodeId(resourceId)
}

// VirtualNodeExists returns true if a virtual node exists in the cluster
func VirtualNodeExists(nodeLister listersv1.NodeLister) (bool, error) {
	nodeList, err := nodeLister.List(labels.Everything())
	if err != nil {
		return false, err
	}
	for _, node := range nodeList {
		if IsVirtualNode(node) {
			return true, nil
		} else {
			//TODO: Change this when clusters with mixed node pools are introduced, we will need to check every node
			return false, nil
		}
	}
	return false, nil
}

func GetIsFeatureEnabledFromEnv(logger *zap.SugaredLogger, featureName string, defaultValue bool) bool {
	enableFeature := defaultValue
	enableFeatureEnvVar, ok := os.LookupEnv(featureName)
	if ok {
		var err error
		enableFeature, err = strconv.ParseBool(enableFeatureEnvVar)
		if err != nil {
			logger.With(zap.Error(err)).Errorf("failed to parse %s envvar, defaulting to %t", featureName, defaultValue)
			return defaultValue
		}
	}
	return enableFeature
}
