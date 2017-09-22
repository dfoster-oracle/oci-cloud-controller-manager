// Copyright 2017 The Oracle Kubernetes Cloud Controller Manager Authors
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

// Package bmcs implements an external Kubernetes cloud-provider for Oracle Bare
// Metal Cloud.
package bmcs

import (
	"fmt"
	"io"

	"k8s.io/kubernetes/pkg/client/clientset_generated/clientset"

	"github.com/golang/glog"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/cloudprovider"
	"k8s.io/kubernetes/pkg/controller"

	"github.com/oracle/kubernetes-cloud-controller-manager/pkg/bmcs/client"
)

// ProviderName uniquely identifies the Oracle Bare Metal Cloud Services (BMCS)
// cloud-provider.
const ProviderName = "bmcs"

// CloudProvider is an implementation of the cloud-provider interface for BMCS.
type CloudProvider struct {
	client              client.Interface
	kubeclient          clientset.Interface
	securityListManager securityListManager
	config              *client.Config
}

// Compile time check that CloudProvider implements the cloudprovider.Interface
// interface.
var _ cloudprovider.Interface = &CloudProvider{}

// NewCloudProvider creates a new baremetal.CloudProvider.
func NewCloudProvider(cfg *client.Config) (cloudprovider.Interface, error) {
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	c = c.Compartment(cfg.Global.CompartmentOCID)

	err = c.Validate()
	if err != nil {
		glog.Errorf("Validate failed to communicate with bmc: %v", err)
		return nil, err
	}

	var securityListManager securityListManager
	if cfg.Global.DisableSecurityListManagement {
		securityListManager = newSecurityListManagerNOOP()
	} else {
		securityListManager = newSecurityListManager(c)
	}

	return &CloudProvider{
		client:              c,
		config:              cfg,
		securityListManager: securityListManager,
	}, nil
}

func init() {
	cloudprovider.RegisterCloudProvider(ProviderName, func(config io.Reader) (cloudprovider.Interface, error) {
		cfg, err := client.ReadConfig(config)
		if err != nil {
			return nil, err
		}
		return NewCloudProvider(cfg)
	})
}

// Initialize passes a Kubernetes clientBuilder interface to the cloud provider.
func (cp *CloudProvider) Initialize(clientBuilder controller.ControllerClientBuilder) {
	var err error
	cp.kubeclient, err = clientBuilder.Client("cloud-controller-manager")
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("failed to create kubeclient: %v", err))
	}
}

// ProviderName returns the cloud-provider ID.
func (cp *CloudProvider) ProviderName() string {
	return ProviderName
}

// LoadBalancer returns a balancer interface. Also returns true if the interface
// is supported, false otherwise.
func (cp *CloudProvider) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	glog.V(6).Info("Claiming to support Load Balancers")
	return cp, true
}

// Instances returns an instances interface. Also returns true if the interface
// is supported, false otherwise.
func (cp *CloudProvider) Instances() (cloudprovider.Instances, bool) {
	glog.V(6).Info("Claiming to support instances")
	return cp, true
}

// Zones returns a zones interface. Also returns true if the interface is
// supported, false otherwise.
func (cp *CloudProvider) Zones() (cloudprovider.Zones, bool) {
	glog.V(6).Info("Claiming *not* to support Zones")
	return cp, false
}

// Clusters returns a clusters interface.  Also returns true if the interface is
// supported, false otherwise.
func (cp *CloudProvider) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

// Routes returns a routes interface along with whether the interface is
// supported.
func (cp *CloudProvider) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

// ScrubDNS provides an opportunity for cloud-provider-specific code to process
// DNS settings for pods.
func (cp *CloudProvider) ScrubDNS(nameservers, searches []string) (nsOut, srchOut []string) {
	return nameservers, searches
}

// HasClusterID returns true if the cluster has a clusterID.
func (cp *CloudProvider) HasClusterID() bool {
	return true
}