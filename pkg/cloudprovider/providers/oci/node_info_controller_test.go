package oci

import (
	"fmt"
	"reflect"
	"testing"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/containerengine"
	"github.com/oracle/oci-go-sdk/v65/core"
)

var (
	instanceCompID    = "instanceCompID"
	instanceFD        = "instanceFD"
	instanceID        = "instanceID"
	virtualNodeId     = "ocid1.virtualnode.oc1.iad.default"
	virtualNodePoolId = "vnpId"
	virtualNodeFD     = "virtualNodeFD"
)

func TestGetNodePatchBytes(t *testing.T) {
	testCases := map[string]struct {
		node               *v1.Node
		instance           *core.Instance
		expectedPatchBytes []byte
	}{
		"FD label and CompartmentID annotation already present": {
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						CompartmentIDAnnotation: "compID",
					},
					Labels: map[string]string{
						FaultDomainLabel: "FD",
					},
				},
			},
			instance: &core.Instance{
				CompartmentId: &instanceCompID,
				FaultDomain:   &instanceFD,
			},
			expectedPatchBytes: nil,
		},
		"Only FD label present": {
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						FaultDomainLabel: "FD",
					},
				},
			},
			instance: &core.Instance{
				CompartmentId: &instanceCompID,
				FaultDomain:   &instanceFD,
			},
			expectedPatchBytes: []byte("{\"metadata\": {\"annotations\": {\"oci.oraclecloud.com/compartment-id\":\"instanceCompID\"}}}"),
		},
		"Only CompartmentID annotation present": {
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						CompartmentIDAnnotation: "compID",
					},
				},
			},
			instance: &core.Instance{
				CompartmentId: &instanceCompID,
				FaultDomain:   &instanceFD,
			},
			expectedPatchBytes: []byte("{\"metadata\": {\"labels\": {\"oci.oraclecloud.com/fault-domain\":\"instanceFD\"}}}"),
		},
		"none present": {
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{},
			},
			instance: &core.Instance{
				CompartmentId: &instanceCompID,
				FaultDomain:   &instanceFD,
			},
			expectedPatchBytes: []byte("{\"metadata\": {\"labels\": {\"oci.oraclecloud.com/fault-domain\":\"instanceFD\"},\"annotations\": {\"oci.oraclecloud.com/compartment-id\":\"instanceCompID\"}}}"),
		},
	}
	logger := zap.L()
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			patchedBytes := getNodePatchBytes(tc.node, tc.instance, logger.Sugar())
			if !reflect.DeepEqual(patchedBytes, tc.expectedPatchBytes) {
				t.Errorf("Expected PatchBytes \n%+v\nbut got\n%+v", tc.expectedPatchBytes, patchedBytes)
			}
		})
	}
}

func TestGetInstanceByNode(t *testing.T) {
	testCases := map[string]struct {
		node             *v1.Node
		nic              *NodeInfoController
		expectedInstance *core.Instance
	}{
		"Get Instance": {
			node: &v1.Node{
				Spec: v1.NodeSpec{
					ProviderID: instanceID,
				},
			},
			nic: &NodeInfoController{
				ociClient: MockOCIClient{},
			},
			expectedInstance: &core.Instance{
				AvailabilityDomain: common.String("NWuj:PHX-AD-1"),
				CompartmentId:      common.String("default"),
				Id:                 &instanceID,
				Region:             common.String("PHX"),
				Shape:              common.String("VM.Standard1.2"),
			},
		},
		"Get Instance when providerID is prefixed with providerName": {
			node: &v1.Node{
				Spec: v1.NodeSpec{
					ProviderID: providerPrefix + instanceID,
				},
			},
			nic: &NodeInfoController{
				ociClient: MockOCIClient{},
			},
			expectedInstance: &core.Instance{
				AvailabilityDomain: common.String("NWuj:PHX-AD-1"),
				CompartmentId:      common.String("default"),
				Id:                 &instanceID,
				Region:             common.String("PHX"),
				Shape:              common.String("VM.Standard1.2"),
			},
		},
	}

	logger := zap.L()
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			instance, err := getInstanceByNode(tc.node, tc.nic, logger.Sugar())
			if err != nil {
				t.Fatalf("%s unexpected service add error: %v", name, err)
			}
			if !reflect.DeepEqual(instance, tc.expectedInstance) {
				t.Errorf("Expected instance \n%+v\nbut got\n%+v", tc.expectedInstance, instanceID)
			}
		})
	}
}

func TestGetVirtualNode(t *testing.T) {
	testCases := map[string]struct {
		node     *v1.Node
		nic      *NodeInfoController
		expected *containerengine.VirtualNode
	}{
		"Get Instance": {
			node: &v1.Node{
				Spec: v1.NodeSpec{
					ProviderID: virtualNodeId,
				},
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						VirtualNodePoolIdAnnotation: virtualNodePoolId,
					},
				},
			},
			nic: &NodeInfoController{
				ociClient: MockOCIClient{},
			},
			expected: &containerengine.VirtualNode{
				Id:                &virtualNodeId,
				VirtualNodePoolId: &virtualNodePoolId,
			},
		},
		"Get Instance when providerID is prefixed with providerName": {
			node: &v1.Node{
				Spec: v1.NodeSpec{
					ProviderID: providerPrefix + virtualNodeId,
				},
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						VirtualNodePoolIdAnnotation: virtualNodePoolId,
					},
				},
			},
			nic: &NodeInfoController{
				ociClient: MockOCIClient{},
			},
			expected: &containerengine.VirtualNode{
				Id:                &virtualNodeId,
				VirtualNodePoolId: &virtualNodePoolId,
			},
		},
	}

	logger := zap.L()
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			virtualNode, err := getVirtualNode(tc.node, tc.nic, logger.Sugar())
			if err != nil {
				t.Fatalf("%s unexpected service add error: %v", name, err)
			}
			if !reflect.DeepEqual(virtualNode, tc.expected) {
				t.Errorf("Expected virtual node \n%+v\nbut got\n%+v", tc.expected, virtualNodeId)
			}
		})
	}
}

func TestGetVirtualNodePatchBytes(t *testing.T) {
	testCases := map[string]struct {
		node               *v1.Node
		virtualNode        *containerengine.VirtualNode
		expectedPatchBytes []byte
	}{
		"FD and Node-Role label not present": {
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{},
			},
			virtualNode: &containerengine.VirtualNode{
				FaultDomain: &virtualNodeFD,
			},
			expectedPatchBytes: []byte(fmt.Sprintf("{\"metadata\": {\"labels\": {\"%s\":\"%s\", \"%s\":\"%s\"}}}", FaultDomainLabel, virtualNodeFD, VirtualNodeRoleLabel, VirtualNodeRoleLabelValue)),
		},
		"Node-Role label not present": {
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						FaultDomainLabel: virtualNodeFD,
					},
				},
			},
			virtualNode: &containerengine.VirtualNode{
				FaultDomain: &virtualNodeFD,
			},
			expectedPatchBytes: []byte(fmt.Sprintf("{\"metadata\": {\"labels\": {\"%s\":\"%s\", \"%s\":\"%s\"}}}", FaultDomainLabel, virtualNodeFD, VirtualNodeRoleLabel, "")),
		},
		"Fault Domain label not present": {
			node: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						VirtualNodeRoleLabel: "",
					},
				},
			},
			virtualNode: &containerengine.VirtualNode{
				FaultDomain: &virtualNodeFD,
			},
			expectedPatchBytes: []byte(fmt.Sprintf("{\"metadata\": {\"labels\": {\"%s\":\"%s\", \"%s\":\"%s\"}}}", FaultDomainLabel, virtualNodeFD, VirtualNodeRoleLabel, "")),
		},
	}
	logger := zap.L()
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			patchedBytes := getVirtualNodePatchBytes(tc.virtualNode, logger.Sugar())
			if !reflect.DeepEqual(patchedBytes, tc.expectedPatchBytes) {
				t.Errorf("Expected PatchBytes \n%+v\nbut got\n%+v", string(tc.expectedPatchBytes), string(patchedBytes))
			}
		})
	}
}
