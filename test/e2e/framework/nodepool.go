package framework

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	. "github.com/onsi/gomega"

	"github.com/oracle/oci-go-sdk/v49/common"
	oke "github.com/oracle/oci-go-sdk/v49/containerengine"
)

var (
	tenanciesBOAT = []string{"ocid1.tenancy.oc1..aaaaaaaagkbzgg6lpzrf47xzy4rjoxg4de6ncfiq2rncmjiujvy2hjgxvziq", "ocid1.tenancy.oc2..aaaaaaaagwtifvnhpr2zoymzwmqk364ycc5fspwx2zzc577cxmmnimazi6nq", "ocid1.tenancy.oc3..aaaaaaaatg7khloug5adtynfkhhhr6ysky5fxb57ghqp3ddpqwqvcavznnnq", "ocid1.tenancy.oc4..aaaaaaaak37nmbaszvdjdrmkvcvlypax53ila3yajff5tgdffk5njsm2czsa", "ocid1.tenancy.oc5..aaaaaaaalfjjthxqwuoxh6ps4aqx62zc46w3aj5n425y3dpvqlqwkant5gda", "ocid1.tenancy.oc6..aaaaaaaalfjjthxqwuoxh6ps4aqx62zc46w3aj5n425y3dpvqlqwkant5gda", "ocid1.tenancy.oc7..aaaaaaaalfjjthxqwuoxh6ps4aqx62zc46w3aj5n425y3dpvqlqwkant5gda"}
)

// NodePoolCreateConfig contains values that can be specified when creating a NodePool.
type NodePoolCreateConfig struct {
	ClusterID string

	CompartmentID string

	NodeImageName string

	NodeShape string

	QuantityPerSubnet *int

	Subnets []string

	KubeVersion string

	NodeConfigDetails *oke.CreateNodePoolNodeConfigDetails

	NodeSourceDetails oke.NodeSourceDetails

	NodeShapeConfig oke.CreateNodeShapeConfigDetails

	// Options contain the test related options that are to be used when doing Cluster create operations.
	Options TestOptions
}

// GetNodePoolOptions return the specified nodepool.
func (f *Framework) GetNodePoolOptions(nodePoolOptionId string) oke.NodePoolOptions {
	ctx, cancel := context.WithTimeout(f.context, f.timeout)
	defer cancel()
	// TODO: remove in the future when something other than "all" is allowed
	id := nodePoolOptionId
	if id != "all" {
		Logf("GetNodePoolOptions : nodePoolOptionId : '%s' overridden with 'all'", nodePoolOptionId)
		id = "all"
	}
	// Check if this is a cross tenancy request and set compartment1
	var compartmentId string
	crossTenancy := f.requestHeaders["x-cross-tenancy-request"]
	if f.authType == ServiceAuth && crossTenancy != "" {
		compartmentId = f.Compartment1
	}
	// if BOAT tenancy, set the compartmentId in the request
	if isBOATTenancy(f.Tenancy) {
		compartmentId = f.Compartment1
	}

	request := oke.GetNodePoolOptionsRequest{
		NodePoolOptionId: &id,
	}

	if len(compartmentId) > 0 {
		request.CompartmentId = &compartmentId
	}

	response, err := f.clustersClient.GetNodePoolOptions(ctx, request)

	Logf("nodePoolOptions : '%#v'", response.NodePoolOptions)

	Expect(err).NotTo(HaveOccurred())
	return response.NodePoolOptions
}

func isBOATTenancy(tenancy string) bool {
	for _, bTenancy := range tenanciesBOAT {
		if tenancy == bTenancy {
			return true
		}
	}
	return false
}

// ListNodePoolShapes return the set of instance shapes available to nodepools.
func (f *Framework) ListNodePoolShapes() []string {
	shapes := f.GetNodePoolOptions("all").Shapes
	Expect(len(shapes) > 0).To(BeTrue())
	return shapes
}

// ListNodePoolImages return the set of images available to nodepools.
func (f *Framework) ListNodePoolImages() map[string]string {
	nodePoolOptions := f.GetNodePoolOptions("all")
	sources := nodePoolOptions.Sources
	Expect(len(sources) > 0).To(BeTrue())
	images := make(map[string]string)
	for _, source := range sources {
		sourceName := *source.(oke.NodeSourceViaImageOption).SourceName
		imageId := *source.(oke.NodeSourceViaImageOption).ImageId
		images[sourceName] = imageId
	}
	return images
}

func (f *Framework) PickNonGPUImageWithAMDCompatibility(images map[string]string) (string, string, bool) {
	for sourceName, imageId := range images {
		if !strings.Contains(sourceName, "GPU") && !strings.Contains(sourceName, "-aarch64") {
			return imageId, sourceName, true
		}
	}
	return "", "", false
}

func (f *Framework) PickArmCompatibleImage(images map[string]string) (string, string, bool) {
	for sourceName, imageId := range images {
		// Hardcoding OL 7.x for now. Needs to be filtered for FIPS compatibility
		if strings.Contains(sourceName, "7.9-aarch64") {
			return imageId, sourceName, true
		}
	}
	return "", "", false
}

// IsValidNodePoolShape return true if the specified nodeShape is valid.
func (f *Framework) IsValidNodePoolShape(shape string) bool {
	nodePoolShapes := f.ListNodePoolShapes()
	for _, nodePoolShape := range nodePoolShapes {
		if nodePoolShape == shape {
			return true
		}
	}
	return false
}

func (f *Framework) ListNodePoolsPaged(clusterID string, compartmentId string, limit int, page *string) (oke.ListNodePoolsResponse, error) {
	ctx, cancel := context.WithTimeout(f.context, f.timeout)
	defer cancel()

	return f.clustersClient.ListNodePools(ctx, oke.ListNodePoolsRequest{
		CompartmentId: &compartmentId,
		ClusterId:     &clusterID,
		Limit:         &limit,
		Page:          page,
	})
}

// ListNodePools lists all nodepools in the specified cluster.
func (f *Framework) ListNodePools(clusterID string) []oke.NodePoolSummary {
	ctx, cancel := context.WithTimeout(f.context, f.timeout)
	limit := 5000
	defer cancel()

	var tmperr error
	start := time.Now()
	var page *string
	resp := make([]oke.NodePoolSummary, 0)
	timeout := 2 * time.Minute
	for {
		response, err := f.clustersClient.ListNodePools(ctx, oke.ListNodePoolsRequest{
			CompartmentId: &f.Compartment1,
			ClusterId:     &clusterID,
			Limit:         &limit,
			Page:          page,
		})
		if err != nil {
			Logf("Error received on ListNodePools: '%s', Retrying", err)
			tmperr = err
		} else {
			for _, np := range response.Items {
				resp = append(resp, np)
			}

			page = response.OpcNextPage
			if page == nil {
				break
			} else {
				Logf("received page token, continue calling ListNodePools()")
			}
		}

		if time.Since(start) > timeout {
			Logf("Timeout waiting for ListNodePools \n")
			Expect(tmperr).NotTo(HaveOccurred())
		}
	}

	return resp
}

// GetNodePoolSummary return the specified nodepool summary.
func (f *Framework) GetNodePoolSummary(clusterID string, nodepoolID string) *oke.NodePoolSummary {
	nodepools := f.ListNodePools(clusterID)
	for _, nodepool := range nodepools {
		if *nodepool.Id == nodepoolID {
			return &nodepool
		}
	}
	return nil
}

// GetNodePool return the specified nodepool.
// Retries added to catch rare connectivity errors
func (f *Framework) GetNodePool(nodepoolID string) *oke.NodePool {
	ctx, cancel := context.WithTimeout(f.context, f.timeout)
	defer cancel()
	// retries added because of the occasional "net/http: request canceled (Client.Timeout exceeded while awaiting headers)"
	timeout := 10 * time.Minute
	for start := time.Now(); time.Since(start) < timeout; time.Sleep(Poll) {
		response, err := f.clustersClient.GetNodePool(ctx, oke.GetNodePoolRequest{
			NodePoolId: &nodepoolID,
		})
		if err == nil {
			Expect(err).NotTo(HaveOccurred())
			return &response.NodePool
		} else {
			Logf("Error received on Get nodePool : '%s', Retrying", err)
		}
	}
	Failf("Timeout waiting for get nodePool '%s'\n", nodepoolID)
	return nil
}

// DeleteNodePool deletes the specified nodepool.
func (f *Framework) DeleteNodePool(nodepoolID string, waitForDeleted bool) {
	ctx, cancel := context.WithTimeout(f.context, f.timeout)
	defer cancel()

	response, err := f.clustersClient.DeleteNodePool(ctx, oke.DeleteNodePoolRequest{
		NodePoolId: &nodepoolID,
	})

	Expect(err).NotTo(HaveOccurred())
	workRequestID := *response.OpcWorkRequestId
	Logf("DeleteNodePool workRequestID : '%s'", workRequestID)

	// If specified wait for 'DELETED'.
	if waitForDeleted {
		wrSucceded := false
		timeout := 10 * time.Minute
		for start := time.Now(); time.Since(start) < timeout; time.Sleep(Poll) {
			if !wrSucceded {
				wrResponse := f.GetWorkRequest(workRequestID)
				Expect(wrResponse.WorkRequest.OperationType).To(Equal(oke.WorkRequestOperationTypeNodepoolDelete))
				Logf("Waiting for delete nodepool '%s', WorkRequest.Status: '%s'",
					nodepoolID, wrResponse.WorkRequest.Status)
				if wrResponse.WorkRequest.Status == oke.WorkRequestStatusSucceeded {
					wrSucceded = true
					return
				} else {
					Expect(wrResponse.WorkRequest.Status).To(SatisfyAny(
						Equal(oke.WorkRequestStatusAccepted),
						Equal(oke.WorkRequestStatusInProgress)))
					Expect(wrResponse.WorkRequest.Status).NotTo(SatisfyAny(
						Equal(oke.WorkRequestStatusFailed),
						Equal(oke.WorkRequestStatusCanceling),
						Equal(oke.WorkRequestStatusCanceled)))
				}
			}
		}
		Failf("Timeout waiting for nodepool '%s' workRequest to SUCCEED\n", nodepoolID)
	}
}

// CreateNodePool creates a nodepool with the specified characteristics. This function
// will block until the nodepool is provisioned and active. Defaults to latest cluster k8s version
func (f *Framework) CreateNodePool(clusterID, compartmentID string, nodeImageName, nodeShape string, size int,
	k8sversion string, subnets []string,
	nodeShapeConfig oke.CreateNodeShapeConfigDetails) (np *oke.NodePool) {

	if k8sversion == "" {
		k8sversion = f.OkeNodePoolK8sVersion
	}
	if os.Getenv("USE_REGIONALSUBNET") == "true" {
		return f.CreateNodePoolInRgnSubnetWithVersion(clusterID, compartmentID, nodeShape, &size,
			subnets[0], k8sversion, nil,
			nodeShapeConfig)
	}

	return f.CreateNodePoolWithVersion(clusterID, nodeImageName, nodeShape, size/(len(subnets)-1), subnets[1:], k8sversion, nil)
}

// CreateNodePoolInRgnSubnetWithVersion use the NodeConfigDetails property
func (f *Framework) CreateNodePoolInRgnSubnetWithVersion(clusterID, compartmentID string, nodeShape string,
	size *int, rgnSubnet string, kubeVersion string,
	expectedError common.ServiceError, nodeShapeConfig oke.CreateNodeShapeConfigDetails) *oke.NodePool {
	ctx, cancel := context.WithTimeout(f.context, f.timeout)
	defer cancel()
	ads := f.ListADs()
	var poolSize *int
	var imageId string
	var imageName string

	if size == nil {
		poolSize = common.Int(len(ads))
	} else {
		poolSize = size
	}
	nodeConfigDetails := oke.CreateNodePoolNodeConfigDetails{
		PlacementConfigs: make([]oke.NodePoolPlacementConfigDetails, 0, len(ads)),
		Size:             poolSize,
	}

	for _, ad := range ads {
		nodeConfigDetails.PlacementConfigs = append(nodeConfigDetails.PlacementConfigs,
			oke.NodePoolPlacementConfigDetails{
				AvailabilityDomain: ad.Name,
				SubnetId:           &rgnSubnet,
			})
	}

	imageId = ""
	imageName = ""
	var nonGPUImageFound = false
	var armCompatibleImageFound = false
	if f.Architecture == "AMD" {
		imageId, imageName, nonGPUImageFound = f.PickNonGPUImageWithAMDCompatibility(f.ListNodePoolImages())
		Expect(nonGPUImageFound).To(BeTrue())
	} else {
		imageId, imageName, armCompatibleImageFound = f.PickArmCompatibleImage(f.ListNodePoolImages())
		Expect(armCompatibleImageFound).To(BeTrue())
	}

	nodeSourceViaImageDetails := &oke.NodeSourceViaImageDetails{
		ImageId: common.String(imageId),
	}

	cfg := &NodePoolCreateConfig{
		ClusterID:         clusterID,
		CompartmentID:     compartmentID,
		NodeImageName:     imageName,
		NodeShape:         nodeShape,
		KubeVersion:       kubeVersion,
		NodeConfigDetails: &nodeConfigDetails,
		NodeSourceDetails: nodeSourceViaImageDetails,
		Options:           TestOptions{ExpectedError: expectedError},
	}

	if f.Architecture == "ARM" {
		cfg.NodeShapeConfig = nodeShapeConfig
	}

	response, _ := f.createNodePoolWithConfig(cfg, ctx)
	pool, done := f.waitForNodePool(response)
	if done {
		// compare to poolSize because size may be nil
		Expect(*pool.NodeConfigDetails.Size).To(Equal(*poolSize))
		return pool
	}
	return nil
}

// CreateNodePool creates a nodepool with the specified characteristics. This function
// will block until the nodepool is provisioned and active. Supply version of f.K8sVersion1, f.K8sVersion2 or f.K8sVersion3
func (f *Framework) CreateNodePoolWithVersion(clusterID, nodeImageName, nodeShape string, quantityPerSubnet int,
	subnets []string, kubeVersion string, expectedError common.ServiceError) *oke.NodePool {
	ctx, cancel := context.WithTimeout(f.context, f.timeout)
	defer cancel()
	cfg := &NodePoolCreateConfig{
		ClusterID:         clusterID,
		CompartmentID:     f.Compartment1,
		NodeImageName:     nodeImageName,
		NodeShape:         nodeShape,
		KubeVersion:       kubeVersion,
		QuantityPerSubnet: &quantityPerSubnet,
		Subnets:           subnets,
		Options:           TestOptions{ExpectedError: expectedError},
	}
	response, _ := f.createNodePoolWithConfig(cfg, ctx)
	pool, done := f.waitForNodePool(response)
	if done {
		Expect(*pool.QuantityPerSubnet).To(Equal(quantityPerSubnet))
		return pool
	}
	return nil
}

//CreateNodePoolWithResponse creates a nodepool with the provide config and returns the raw CreateNodePool response.
func (f *Framework) CreateNodePoolWithResponse(cfg *NodePoolCreateConfig) (response oke.CreateNodePoolResponse, err error) {
	ctx, cancel := context.WithTimeout(f.context, f.timeout)
	defer cancel()
	return f.createNodePoolWithConfig(cfg, ctx)
}

func (f *Framework) createNodePoolWithConfig(cfg *NodePoolCreateConfig, ctx context.Context) (response oke.CreateNodePoolResponse, err error) {
	nodepoolName := f.createNodePoolName(cfg.KubeVersion)
	request := oke.CreateNodePoolRequest{
		CreateNodePoolDetails: oke.CreateNodePoolDetails{
			Name:              &nodepoolName,
			KubernetesVersion: &cfg.KubeVersion,
			ClusterId:         &cfg.ClusterID,
			NodeShape:         &cfg.NodeShape,
			NodeImageName:     &cfg.NodeImageName,
			InitialNodeLabels: f.setNodeLabels(nodepoolName),
			SshPublicKey:      &f.PubSSHKey,
			CompartmentId:     &cfg.CompartmentID,
			SubnetIds:         cfg.Subnets,
			QuantityPerSubnet: cfg.QuantityPerSubnet,
			NodeConfigDetails: cfg.NodeConfigDetails,
			NodeSourceDetails: cfg.NodeSourceDetails,
		},
	}
	if f.Architecture == "ARM" {
		request.CreateNodePoolDetails.NodeShapeConfig = &cfg.NodeShapeConfig
	}
	requestB, _ := json.MarshalIndent(request, "", "  ")
	Logf("Creating nodepool, name: '%s' with request %s", nodepoolName, requestB)
	response, err = f.clustersClient.CreateNodePool(ctx, request)
	if err != nil && cfg.Options.ExpectedError != nil {
		Logf("Nodepool creation error detected: '%v'", err)
		serviceError, isServiceError := common.IsServiceError(err)
		if isServiceError {
			Expect(serviceError.GetHTTPStatusCode()).To(Equal(cfg.Options.ExpectedError.GetHTTPStatusCode()))
			Expect(serviceError.GetMessage()).To(Equal(cfg.Options.ExpectedError.GetMessage()))
			Expect(serviceError.GetCode()).To(Equal(cfg.Options.ExpectedError.GetCode()))

		} else {
			Failf("Unknown error detected on nodepool creation.\n")
		}
		return
	} else {
		Expect(err).NotTo(HaveOccurred())
	}
	return
}

func (f *Framework) setNodeLabels(nodepoolName string) []oke.KeyValue {
	key := fmt.Sprintf("%s-Key1", nodepoolName)
	value := fmt.Sprintf("%s-Value1", nodepoolName)
	nodeLabels := []oke.KeyValue{oke.KeyValue{Key: &key, Value: &value}}
	return nodeLabels
}

func (f *Framework) createNodePoolName(kubeVersion string) string {
	nodepoolName := "TestNodePool-" + UniqueID()
	// TODO Remove when nodepool name lengths > 32 cause an error
	if len(nodepoolName) > 32 {
		Logf("The nodePoolName: '%s' has len %d.\n", nodepoolName, len(nodepoolName))
		nodepoolName = fmt.Sprintf("%.32s", nodepoolName)
		Logf("Truncated nodePoolName: '%s' has len %d.\n", nodepoolName, len(nodepoolName))
	}
	Logf("Creating nodepool, name: '%s' with K8SVersion %s", nodepoolName, kubeVersion)
	return nodepoolName
}

func (f *Framework) waitForNodePool(response oke.CreateNodePoolResponse) (*oke.NodePool, bool) {
	nodePoolID := ""
	workRequestID := *response.OpcWorkRequestId
	Logf("CreateNodePoolRequest workRequestID : '%s'", workRequestID)
	// Waits for the nodepool to become ready. This is done by first waiting for the
	// workRequest.Status to reach the "SUCCEEDED" state, then waiting for all nodes in
	// the nodePool to enter the "ACTIVE" state.
	wrSucceded := false
	timeout := 35 * time.Minute
	for start := time.Now(); time.Since(start) < timeout; time.Sleep(Poll) {
		// NodePool WorkRequest - its finished as soon as it sets the desired config (about 5-10s)
		// - the tenant-agent takes over and takes action to ensure the desired state matches the
		// desired config.
		if !wrSucceded {
			wrResponse := f.GetWorkRequest(workRequestID)
			Expect(wrResponse.WorkRequest.OperationType).To(Equal(oke.WorkRequestOperationTypeNodepoolCreate))
			Logf("Waiting for create nodepool WorkRequest.Status: '%s'", wrResponse.WorkRequest.Status)
			if wrResponse.WorkRequest.Status == oke.WorkRequestStatusSucceeded {
				wrSucceded = true
				for _, resource := range wrResponse.Resources {
					Logf("EntityType:  '%s'", *resource.EntityType)
					if *resource.EntityType == "nodepool" {
						nodePoolID = *resource.Identifier
						Logf("Setting nodePoolID: '%s'", nodePoolID)
					}
				}
			} else {
				Expect(wrResponse.WorkRequest.Status).To(SatisfyAny(
					Equal(oke.WorkRequestStatusAccepted),
					Equal(oke.WorkRequestStatusInProgress)))
				Expect(wrResponse.WorkRequest.Status).NotTo(SatisfyAny(
					Equal(oke.WorkRequestStatusFailed),
					Equal(oke.WorkRequestStatusCanceling),
					Equal(oke.WorkRequestStatusCanceled)))
			}
		} else {
			Logf("Calling GetNodePool, nodePoolID: '%s'", nodePoolID)
			nodepool := *f.GetNodePool(nodePoolID)
			numActive := countNodesInState(nodepool, oke.NodeLifecycleStateActive)
			hasExpectedActive := hasExpectedNodesInState(nodepool, oke.NodeLifecycleStateActive)
			Logf("Waiting for create nodepool id %s numActiveNodes: '%d', hasExpectedActive: '%v'", nodePoolID, numActive, hasExpectedActive)
			if hasExpectedActive == false {
				for _, node := range nodepool.Nodes {
					detailsStr := "(nil)"
					if node.LifecycleDetails != nil {
						detailsBS, _ := json.Marshal(node.LifecycleDetails)
						detailsStr = string(detailsBS)
					}
					Logf("\tNode ID '%s', State '%s', Details '%s', Error '%v'", *node.Id, node.LifecycleState, detailsStr, node.NodeError)
				}
			} else {
				return f.GetNodePool(nodePoolID), true
			}
		}
	}
	Failf("Timeout waiting for create nodepool '%s' to reach all nodes ACTIVE state\n", nodePoolID)
	return nil, false
}

// WaitForActiveStateInNodePool checks and waits for all nodes in the node pool to be active.
// This is used specifically when calling 'update' against an existing cluster that causes OKE
// components to be updated in each of the nodes.
func (f *Framework) WaitForActiveStateInNodePool(nodePoolID string) *oke.NodePool {
	// Waits for all nodes in the nodepool to be in "ACTIVE" state.
	timeout := 25 * time.Minute
	for start := time.Now(); time.Since(start) < timeout; time.Sleep(Poll) {
		Logf("Calling GetNodePool, nodePoolID: '%s'", nodePoolID)
		nodepool := *f.GetNodePool(nodePoolID)
		numActive := countNodesInState(nodepool, oke.NodeLifecycleStateActive)
		hasExpectedActive := hasExpectedNodesInState(nodepool, oke.NodeLifecycleStateActive)
		Logf("Waiting for nodepool '%s', numActiveNodes: '%d', hasExpectedActive: '%v'",
			*nodepool.Name, numActive, hasExpectedActive)
		if hasExpectedActive == false {
			for _, node := range nodepool.Nodes {
				Logf("\tNode ID '%s', State '%s'", *node.Id, node.LifecycleState)
			}
		} else {
			return f.GetNodePool(nodePoolID)
		}
	}
	Failf("Timeout waiting for nodepool '%s' to reach all nodes ACTIVE state\n", nodePoolID)
	return nil
}

// version2 means use Nodeconfigdetails to scale.
func IsVersion2NodePool(pool *oke.NodePool) bool {
	return pool.QuantityPerSubnet == nil || *pool.QuantityPerSubnet == 0
}
