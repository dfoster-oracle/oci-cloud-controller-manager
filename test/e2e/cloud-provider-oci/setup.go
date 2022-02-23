package e2e

import (
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	sharedfw "github.com/oracle/oci-cloud-controller-manager/test/e2e/framework"
	oke "github.com/oracle/oci-go-sdk/v49/containerengine"
)

var setupF *sharedfw.Framework

var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {

	setupF = sharedfw.New()

	sharedfw.Logf("CloudProviderFramework Setup")

	if setupF.EnableCreateCluster {
		sharedfw.Logf("Creating the cluster...")
		clusterOCID := setupF.CreateCluster()
		Expect(clusterOCID).ShouldNot(BeZero())
		sharedfw.Logf("Cluster OCID is %s", clusterOCID)

		kubeConfig := setupF.CreateClusterKubeconfigContent(clusterOCID)
		Expect(setupF.IsNotJsonFormatStr(kubeConfig)).To(BeTrue())
		Expect(kubeConfig != "").To(BeTrue())

		err := setupF.SaveKubeConfig(kubeConfig)
		sharedfw.Logf("Returned Kubeconfig: \n%s", kubeConfig)

		cloudConfig := setupF.CreateCloudConfig()
		Expect(cloudConfig).ShouldNot(BeNil())

		err = setupF.SaveCloudConfig(cloudConfig)
		Expect(err).Should(BeNil())

		var ocpus = float32(1.0)
		var memoryInGBs = float32(6.0)
		var NodeShapeConfig = oke.CreateNodeShapeConfigDetails{
			Ocpus:       &ocpus,
			MemoryInGBs: &memoryInGBs,
		}

		size := 3
		nodepool := setupF.CreateNodePool(clusterOCID, setupF.Compartment1, "Oracle-Linux-7.6",
			setupF.NodeShape, size, setupF.OkeNodePoolK8sVersion,
			[]string{setupF.NodeSubnet, setupF.NodeSubnet, setupF.NodeSubnet},
			NodeShapeConfig)
		Expect(nodepool).ShouldNot(BeNil())
		sharedfw.Logf(" Created cluster %s with nodepool %s ", clusterOCID, *nodepool.Id)
		setupF.CrossValidateCluster(clusterOCID, setupF.ValidateChildResources)
	} else {
		sharedfw.Logf("Cluster creation skipped. Running tests with existing cluster.")
	}
	return nil
}, func(data []byte) {})

var _ = ginkgo.SynchronizedAfterSuite(func() {
	sharedfw.Logf("Running AfterSuite actions on all node")
	sharedfw.RunCleanupActions()
	if setupF.EnableCreateCluster {
		setupF.CleanAllWithoutWait()
	}
}, func() {})
