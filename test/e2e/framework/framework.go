package framework

import (
	"context"
	"flag"
	"fmt"
	"github.com/oracle/oci-cloud-controller-manager/pkg/cloudprovider/providers/oci/config"
	"io/ioutil"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/tools/record"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"time"

	. "github.com/onsi/gomega"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/common/auth"
	oke "github.com/oracle/oci-go-sdk/containerengine"
	"github.com/oracle/oci-go-sdk/core"
	"github.com/oracle/oci-go-sdk/identity"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	// Poll is the default polling period when checking lifecycle status.
	Poll = 15 * time.Second
	// Poll defines how regularly to poll kubernetes resources.
	K8sResourcePoll = 2 * time.Second
    DefaultClusterKubeconfig = "/tmp/clusterkubeconfig"
	DefaultCloudConfig = "/tmp/cloudconfig"
)

var (
	okeendpoint                        string
	defaultOCIUser                     string
	defaultOCIUserFingerprint          string
	defaultOCIUserKeyFile              string
	defaultOCIUserPassphrase           string
	defaultOCITenancy                  string
	defaultOCIRegion                   string
	compartment1                       string
	compartment2                       string
	vcn                                string
	lbsubnet1                          string
	lbsubnet2                          string
	lbrgnsubnet                        string
	nodeshape                          string
	subnet1                            string
	subnet2                            string
	subnet3                            string
	rgnsubnet                          string
	okeClusterK8sVersionIndex          int
	okeNodePoolK8sVersionIndex         int
	pubsshkey                           string
	flexvolumedrivertestversion         string
	volumeprovisionertestversion string
	secretsDir                   string
	defaultKubeConfig            string
	delegateGroupID              string
	instanceCfg                  *DelegationPrincipalConfig
	enableCreateCluster          bool
	kmsKeyID                     string
	adlocation                   string
	clusterkubeconfig            string // path to kubeconfig file
	deleteNamespace              bool   // whether or not to delete test namespaces
	cloudConfigFile              string // path to cloud provider config file
	nodePortTest                 bool   // whether or not to test the connectivity of node ports.
	ccmSeclistID                 string // The ocid of the loadbalancer subnet seclist. Optional.
	k8sSeclistID                 string // The ocid of the k8s worker subnet seclist. Optional.
)

func init() {
	flag.StringVar(&okeendpoint, "okeendpoint", "", "OKE endpoint to test against.")
	flag.StringVar(&defaultOCIUser, "ociuser", "", "OCI user OCID.")
	flag.StringVar(&defaultOCIUserFingerprint, "ocifingerprint", "", "OCI user fingerprint.")
	flag.StringVar(&defaultOCIUserKeyFile, "ocikeyfile", "", "OCI key file.")
	flag.StringVar(&defaultOCIUserPassphrase, "ocikeypassphrase", "", "OCI key passphrase.")
	flag.StringVar(&defaultOCITenancy, "ocitenancy", "", "OCI tenancy.")
	flag.StringVar(&defaultOCIRegion, "ociregion", "", "OCI region.")
	flag.StringVar(&compartment1, "compartment1", "", "OCID of the compartment1 in which to manage clusters.")
	flag.StringVar(&compartment2, "compartment2", "", "OCID of the compartment2 in which to manage clusters.")
	flag.StringVar(&vcn, "vcn", "", "OCID of the VCN in which to create clusters.")
	flag.StringVar(&lbsubnet1, "lbsubnet1", "", "OCID of the 1st subnet in which to create load balancers.")
	flag.StringVar(&lbsubnet2, "lbsubnet2", "", "OCID of the 2nd subnet in which to create load balancers.")
	flag.StringVar(&lbrgnsubnet, "lbrgnsubnet", "", "OCID of the regional subnet in which to create load balancers.")
	flag.StringVar(&nodeshape, "nodeshape", "VM.Standard1.2", "node shape of the nodepool.")
	flag.StringVar(&subnet1, "subnet1", "", "OCID of the 1st worker subnet.")
	flag.StringVar(&subnet2, "subnet2", "", "OCID of the 2nd worker subnet.")
	flag.StringVar(&subnet3, "subnet3", "", "OCID of the 3rd worker subnet.")
	flag.StringVar(&rgnsubnet, "rgnsubnet", "", "OCID of the regional subnet worker subnet.")
	flag.IntVar(&okeClusterK8sVersionIndex, "okeClusterK8sVersionIndex", -1, "The index of k8s versionList (0 means the 1st version, 1 means the 2nd version. -1 means the latest version. versionList is like ['1.10.11', 1.11.8', '1.12.6']) used when create cluster")
	flag.IntVar(&okeNodePoolK8sVersionIndex, "okeNodePoolK8sVersionIndex", -1, "The index of k8s versionList (0 means the 1st version, 1 means the 2nd version. -1 means the latest version. versionList is like ['1.10.11', 1.11.8', '1.12.6']) used when create nodepool")
	flag.StringVar(&pubsshkey, "pubsshkey", "", "Public SSH Key for node access.")
	flag.StringVar(&flexvolumedrivertestversion, "flexvolumedrivertestversion", "", "FlexVolumeDriver test version.")
	flag.StringVar(&volumeprovisionertestversion, "volumeprovisionertestversion", "", "VolumeProvisioner test version.")
	flag.StringVar(&secretsDir, "secrets-dir", "/secrets", "path to the root secrets directory")
	flag.StringVar(&defaultKubeConfig, "kubeconfig_file", "", "the path to the kubeconfig file for the regional Primordial cluster")
	flag.StringVar(&delegateGroupID, "delegate-group-id", "", "the ocid of the dynamic group to used for fetching delegation tokens")

	flag.BoolVar(&enableCreateCluster, "enable-create-cluster", true, "Whether or not to enable creating a cluster before test execution")

	flag.StringVar(&kmsKeyID, "kms-key-id", "ocid1.key.oc1.iad.annnb3f4aacuu.abuwcljsj6vlwxrtm2bzjmre3ynqfnhkfmcayia3mq47opjp5f2joozrrdaa", "The KMS Key OCID used for testing KMS integration")

	flag.StringVar(&adlocation, "adlocation", "zkJl:US-ASHBURN-AD-1", "Default Ad Location.")

	//Below two flags need to be provided if test cluster already exists.
	flag.StringVar(&clusterkubeconfig, "cluster-kubeconfig", DefaultClusterKubeconfig, "Path to Cluster's Kubeconfig file with authorization and master location information. Only provide if test cluster exists.")
	flag.StringVar(&cloudConfigFile, "cloud-config", DefaultCloudConfig, "The path to the cloud provider configuration file. Empty string for no configuration file. Only provide if test cluster exists.")

	flag.BoolVar(&nodePortTest, "nodeport-test", false, "If true test will include 'nodePort' connectectivity tests.")
	flag.StringVar(&ccmSeclistID, "ccm-seclist-id", "", "The ocid of the loadbalancer subnet seclist. Enables additional seclist rule tests. If specified the 'k8s-seclist-id parameter' is also required.")
	flag.StringVar(&k8sSeclistID, "k8s-seclist-id", "", "The ocid of the k8s worker subnet seclist. Enables additional seclist rule tests. If specified the 'ccm-seclist-id parameter' is also required.")
	flag.BoolVar(&deleteNamespace, "delete-namespace", true, "If true tests will delete namespace after completion. It is only designed to make debugging easier, DO NOT turn it off by default.")

	flag.Parse()
}

func getDefaultOCIUser() OCIUser {
	// defaultUser details are read from flags and there are cases when this function gets invoked before
	// flag.parse has happened resulting in the variables not being initialized. In such case return an empty user.
	if defaultOCIUser == "" {
		return OCIUser{}
	}
	privateKey, err := ioutil.ReadFile(defaultOCIUserKeyFile)
	if err != nil {
		Failf("Error reading oci private key file '%s' : %s", defaultOCIUserKeyFile, err)
	}
	return OCIUser{
		OCID:                 defaultOCIUser,
		Fingerprint:          defaultOCIUserFingerprint,
		PrivateKey:           privateKey,
		PrivateKeyPassPhrase: defaultOCIUserPassphrase,
	}
}

type AuthType string

const (
	UserAuth      AuthType = "user"
	ServiceAuth   AuthType = "service"
	DelegatedAuth AuthType = "delegated"
)

// Framework is the context of the text execution.
type Framework struct {
	//Note: Reusing framework will not work for OBO requests(check interceptor) with changing users and hence it is recommended to use new framework for each context for obo scenarios.
	computeClient  *core.ComputeClient
	clustersClient *oke.ContainerEngineClient
	identityClient *identity.IdentityClient

	context context.Context
	timeout time.Duration

	authType AuthType

	RegionalKubeConfig string

	// Set of headers those would be injected in to an outgoing HTTP request through an interceptor.
	requestHeaders map[string]string

	// The OCI tenancy.
	Tenancy string
	// The OCI region.
	Region string

	// The OCI test user.
	User OCIUser

	// The compartment1 the cluster is running in.
	Compartment1 string
	// The compartment2 the cluster is running in.
	Compartment2 string
	// The VCN the cluster is running in.
	Vcn string
	// Loadbalancer subnet 1.
	LbSubnet1 string
	// Loadbalancer subnet 2.
	LbSubnet2 string
	// Loadbalancer regional subnet
	Lbrgnsubnet string
	// NodePool node shape.
	NodeShape string
	// NodePool subnet 1.
	Subnet1 string
	// NodePool subnet 2.
	Subnet2 string
	// NodePool subnet 3.
	Subnet3 string
	// NodePool regional subnet
	Rgnsubnet string

	//k8s version value (eg. v1.10.11, v1.11.8) used when create cluster
	OkeClusterK8sVersion string
	//k8s version value (eg. v1.10.11, v1.11.8) used when create nodepool
	OkeNodePoolK8sVersion string

	// The Public SSH key to use when accessing nodes.
	PubSSHKey string

	// The lower of the Kubernetes versions supported.
	K8sVersion1 string
	// The middle Kubernetes versions supported.
	K8sVersion2 string
	// The higher of the Kubernetes versions supported.
	K8sVersion3 string
	// The lower of the nodePool Kubernetes versions supported.
	NodePoolK8sVersion1 string
	// The middle nodePool Kubernetes versions supported.
	NodePoolK8sVersion2 string
	// The higher of the nodePool Kubernetes versions supported.
	NodePoolK8sVersion3 string
	// If true, then delete operations will block until the target resource has been deleted.
	WaitForDeleted bool
	// If true, then validation operations should cascade and deleted any dependent child resources.
	ValidateChildResources bool

	// OKEEndpoint is the endpoint url for the TM API we'll be performing tests against.
	OKEEndpoint string

	// Temporary variable to denote which version of the flexvolume driver tests to use.
	// TODO: Dynamically determine which version of the tests to run based on the OKE environment.
	//       Which release version of the oci-flex-volume-driver is runnning in the OKE environment?
	FlexVolumeDriverTestVersion string
	// Temporary variable to denote which version of the volume provisioner tests to use.
	// TODO: Dynamically determine which version of the tests to run based on the OKE environment.
	//       Which release version of the oci-volume-provisioner is runnning in the OKE environment?
	VolumeProvisionerTestVersion string
	//DynamicGroup for auth delegation testing
	DynamicGroup       string
	instanceConfigFile string
	// Target services for obtaining delegation token.
	DelegationTargetServices string

	// Default adLocation
	AdLocation string

	//is cluster creation required
	EnableCreateCluster bool

	ClusterKubeconfigPath string

	CloudConfigPath string

}

// New creates a new a framework that holds the context of the test
// execution.
func New() *Framework {
	return NewWithConfig(&FrameworkConfig{})

}

//FrameworkConfig helps in passing the configuration options while creating a Framework instance.
type FrameworkConfig struct {
	// authType specifies the framework's authType to be used before initialization.
	AuthType AuthType
	// User specifies the framework's OCIUser to be used before initialization.
	User *OCIUser

	// InstanceConfigFile specifies the framework's  InstanceConfigFile to be used before initialization.
	InstanceConfigFile string
	// True, for a cross tenancy test
	CrossTenancy bool
	// Used as target services to obtain a delegation token.
	DelegationTargetServices string
}

// OCIUser contains details about an OCI user along with their API key details.
type OCIUser struct {
	// OCID of the User
	OCID string
	// Fingerprint of the public key attached to the User
	Fingerprint string
	// User's API Private key
	PrivateKey []byte
	// PassPhrase of the API Private key. If API Private key has no passphrase then set it to ""
	PrivateKeyPassPhrase string
}

// NewWithConfig creates a new Framework instance and configures the instance as per the configuration options in the given config.
func NewWithConfig(config *FrameworkConfig) *Framework {
	rand.Seed(time.Now().UTC().UnixNano())

	f := &Framework{
		// Default to user auth for new frameworks.
		authType: UserAuth,

		computeClient:  nil,
		clustersClient: nil,

		Tenancy:      defaultOCITenancy,
		Region:       defaultOCIRegion,
		User:         getDefaultOCIUser(),
		Compartment1: compartment1,
		Compartment2: compartment2,

		RegionalKubeConfig: defaultKubeConfig,
		requestHeaders:     map[string]string{},

		timeout:                    3 * time.Minute,
		context:                    context.Background(),
		WaitForDeleted:             false,
		ValidateChildResources:     true,
		PubSSHKey:                  pubsshkey,
		Vcn:                        vcn,
		LbSubnet1:                  lbsubnet1,
		LbSubnet2:                  lbsubnet2,
		Lbrgnsubnet:                lbrgnsubnet,
		Subnet1:                    subnet1,
		Subnet2:                    subnet2,
		Subnet3:                    subnet3,
		Rgnsubnet:                  rgnsubnet,
		NodeShape:                  nodeshape,
		DelegationTargetServices:   "oke",
	}

	f.EnableCreateCluster = enableCreateCluster

	if enableCreateCluster {
		if config.AuthType != "" {
			f.authType = config.AuthType
		}

		if config.InstanceConfigFile != "" {
			f.instanceConfigFile = config.InstanceConfigFile
		}

		if config.User != nil {
			f.User = *config.User
		}

		if config.CrossTenancy {
			f.requestHeaders["x-cross-tenancy-request"] = f.Tenancy
		}

		f.OKEEndpoint = okeendpoint
		if !strings.HasPrefix(okeendpoint, "https://") {
			f.OKEEndpoint = "https://" + okeendpoint
		}

		f.FlexVolumeDriverTestVersion = flexvolumedrivertestversion
		f.VolumeProvisionerTestVersion = volumeprovisionertestversion
	}

	f.CloudConfigPath = cloudConfigFile
	f.ClusterKubeconfigPath = clusterkubeconfig

	f.Initialize()

	return f
}

func (f *Framework) AddRequestHeader(name string, value string) *Framework {
	f.requestHeaders[name] = value
	return f
}

// NewInstancePrincipalKubeClient creates a new kubernetes client that points to the
// primordial dev cluster. This is only to run the instance principal pod tests against
// the OKE endpoint specified for the e2e test so it really doesn't matter which cluster it runs in.
func (f *Framework) NewInstancePrincipalKubeClient() *KubeClient {
	b, err := ioutil.ReadFile(f.RegionalKubeConfig)
	Expect(err).To(BeNil())

	return NewKubeClient(string(b))
}

func (f *Framework) newOCIRequestSigner(authType AuthType) (common.HTTPRequestSigner, error) {

	switch authType {
	case UserAuth:
		cfgProvider := common.NewRawConfigurationProvider(
			f.Tenancy,
			f.User.OCID,
			f.Region,
			f.User.Fingerprint,
			string(f.User.PrivateKey),
			&f.User.PrivateKeyPassPhrase,
		)
		return common.DefaultRequestSigner(cfgProvider), nil
	case DelegatedAuth:
		if instanceCfg == nil {
			cfgBytes, err := ioutil.ReadFile(f.instanceConfigFile)
			if err != nil {
				return nil, err
			}

			var cfg SecretValuesConfig

			if err := yaml.Unmarshal(cfgBytes, &cfg); err != nil {
				return nil, err
			}

			instanceCfg = &cfg.BMCSCredentials.DelegationPrincipalConfig
		}
		// We use the test Instance Principal certificates from https://confluence.oci.oraclecorp.com/pages/viewpage.action?pageId=56468176#HowtotestyourservicewithOBO/Delegation?-Certificates
		cfgProvider, err := auth.InstancePrincipalConfigurationWithCerts(common.Region(instanceCfg.Region), []byte(instanceCfg.Cert), []byte(instanceCfg.KeyPassphrase), []byte(instanceCfg.Key), [][]byte{[]byte(instanceCfg.Intermediate)})
		if err != nil {
			return nil, err
		}
		return common.RequestSigner(cfgProvider, append(common.DefaultGenericHeaders(), "opc-obo-token"), common.DefaultBodyHeaders()), nil
	default:
		return nil, fmt.Errorf("unknown auth type %q: cannot construct signing client", f.authType)
	}
}

// BeforeEach will be executed before each Ginkgo test is executed.
func (f *Framework) Initialize() {
	f.EnableCreateCluster = enableCreateCluster
	if !enableCreateCluster {
		f.ClusterKubeconfigPath = clusterkubeconfig
		f.CloudConfigPath = cloudConfigFile
		return
	}
	Logf("initializing framework")
	Logf("Auth Type: %v", f.authType)
	f.PubSSHKey = pubsshkey
	Logf("Public SSHKey : %s", f.PubSSHKey)
	f.Compartment1 = compartment1
	Logf("OCI compartment1 OCID: %s", f.Compartment1)
	f.Compartment2 = compartment2
	Logf("OCI compartment1 OCID: %s", f.Compartment2)
	f.Vcn = vcn
	Logf("OCI VCN OCID: %s", f.Vcn)
	f.LbSubnet1 = lbsubnet1
	Logf("OCI lbSubnet1 OCID: %s", f.LbSubnet1)
	f.LbSubnet2 = lbsubnet2
	Logf("OCI lbSubnet2 OCID: %s", f.LbSubnet2)
	f.Lbrgnsubnet = lbrgnsubnet
	Logf("OCI lbrgnsubnet OCID: %s", f.Lbrgnsubnet)
	f.Subnet1 = subnet1
	Logf("OCI Subnet1 OCID: %s", f.Subnet1)
	f.Subnet2 = subnet2
	Logf("OCI Subnet2 OCID: %s", f.Subnet2)
	f.Subnet3 = subnet3
	Logf("OCI Subnet3 OCID: %s", f.Subnet3)
	f.Rgnsubnet = rgnsubnet
	Logf("OCI Rgnsubnet OCID: %s", f.Rgnsubnet)
	f.NodeShape = nodeshape
	Logf("Nodepool NodeShape: %s", f.NodeShape)

	f.OKEEndpoint = okeendpoint
	if !strings.HasPrefix(okeendpoint, "https://") {
		f.OKEEndpoint = "https://" + okeendpoint
	}

	f.FlexVolumeDriverTestVersion = flexvolumedrivertestversion
	f.VolumeProvisionerTestVersion = volumeprovisionertestversion

	Logf("OCI User: %v", f.User.OCID)
	Logf("OCI finger print: %v", f.User.Fingerprint)

	if f.clustersClient == nil {
		Logf("Creating OCI client")
		//if it's a delegation obo scenario
		if f.instanceConfigFile != "" {
			Logf("instanceConfigFile is %s", f.instanceConfigFile)
			cfgBytes, err := ioutil.ReadFile(f.instanceConfigFile)
			if err != nil {
				Failf("Error reading '%s' : %s", f.instanceConfigFile, err)
			}

			var cfg SecretValuesConfig

			if err := yaml.Unmarshal(cfgBytes, &cfg); err != nil {
				Failf("Error ubmarshaling '%s' : %s", f.instanceConfigFile, err)
			}

			instanceCfg = &cfg.BMCSCredentials.DelegationPrincipalConfig
			f.DynamicGroup = instanceCfg.DynamicGroup
		} else {
			f.DynamicGroup = delegateGroupID
		}

		userConfigProvider := common.NewRawConfigurationProvider(
			f.Tenancy, f.User.OCID, f.Region, f.User.Fingerprint,
			string(f.User.PrivateKey), &f.User.PrivateKeyPassPhrase)

		clustersClient, err := oke.NewContainerEngineClientWithConfigurationProvider(userConfigProvider)

		Expect(err).NotTo(HaveOccurred())
		f.clustersClient = &clustersClient
		f.clustersClient.Host = okeendpoint
		f.clustersClient.BasePath = "20180222"
		f.clustersClient.Interceptor = nil

		requestSigner, err := f.newOCIRequestSigner(f.authType)
		if err != nil {
			Failf("f.authType: %q., unable to create request signer: %v", f.authType, err)
		}
		f.clustersClient.BaseClient.Signer = requestSigner

		Logf("oke endpoint: '%s'", okeendpoint)
		computeClient, err := core.NewComputeClientWithConfigurationProvider(userConfigProvider)
		Expect(err).NotTo(HaveOccurred())
		f.computeClient = &computeClient

		identityClient, err := identity.NewIdentityClientWithConfigurationProvider(userConfigProvider)
		Expect(err).NotTo(HaveOccurred())
		f.identityClient = &identityClient

		// Determine which cluster k8s versions are supported for this OKE Release
		clusterOptions := f.GetClusterOptions("all")
		versions := filterVersionsToLatestMinor(clusterOptions.KubernetesVersions)
		f.OkeClusterK8sVersion = getK8sVersionValue(versions, okeClusterK8sVersionIndex)
		//below K8sVersion* would be deprecated if OkeClusterK8sVersion is used in test code
		numVersions := len(versions)
		if numVersions < 2 {
			Failf("less than two supported cluster k8sVersions!!!  '%#v'", versions)
		}
		f.K8sVersion1 = versions[0]
		f.K8sVersion2 = versions[1]
		if numVersions >= 3 {
			f.K8sVersion3 = versions[2]
		} else {
			f.K8sVersion3 = ""
		}

		// Determine which cluster k8s versions are supported for this OKE Release
		nodePoolOptions := f.GetNodePoolOptions("all")
		nodePoolVersions := filterVersionsToLatestMinor(nodePoolOptions.KubernetesVersions)
		f.OkeNodePoolK8sVersion = getK8sVersionValue(nodePoolVersions, okeNodePoolK8sVersionIndex)

		Logf("OkeClusterK8sVersion=%v", f.OkeClusterK8sVersion)
		Logf("OkeNodePoolK8sVersion=%v", f.OkeNodePoolK8sVersion)

		//below NodePoolK8sVersion* would be deprecated if OkeNodePoolK8sVersion is used in test code
		numNodePoolVersions := len(nodePoolVersions)
		if numNodePoolVersions < 2 {
			Failf("less than two supported node k8sVersions!!!  '%#v'", nodePoolVersions)
		}
		f.NodePoolK8sVersion1 = nodePoolVersions[0]
		f.NodePoolK8sVersion2 = nodePoolVersions[1]
		if numVersions >= 3 {
			f.NodePoolK8sVersion3 = nodePoolVersions[2]
		} else {
			f.NodePoolK8sVersion3 = ""
		}
	}
}

//getK8sVersionValue returns the version value according to version index
//eg. versions=["1.10.11", "1.11.5"]
//return 1.10.11 if index=0,
//return 1.11.5 if index=1, or -1
func getK8sVersionValue(versions []string, index int) string {
	numVersions := len(versions)
	if numVersions == 0 {
		Failf("No supported k8sVersions for cluster or nodepool!!! ")
	}

	if index < 0 && numVersions+index >= 0 {
		return versions[numVersions+index]
	}

	if index >= 0 && index < numVersions {
		return versions[index]
	}

	Failf("Index=%d is not valid for current version List=%v", index, versions)
	return ""
}

func filterVersionsToLatestMinor(versions []string) []string {
	var filtered []string
	for _, v := range versions {
		if len(filtered) == 0 {
			filtered = append(filtered, v)
			continue
		}

		idx := -1
		for i, existing := range filtered {
			if stripPatchVersion(existing) == stripPatchVersion(v) {
				idx = i
				break
			}
		}

		if idx != -1 {
			filtered[idx] = v
		} else {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func stripPatchVersion(version string) string {
	parts := strings.Split(version, ".")
	return parts[0] + "." + parts[1]
}

func (f *Framework) CleanAllWithoutWait() []byte {
	f.CleanAll(false)
	return []byte{}
}

// CleanAll will attempt to clean the tenancy and compartment1 of all
// clusters and nodepools. Use with caution.
func (f *Framework) CleanAll(waitForDeleted bool) {
	Logf("Cleaning all OKE clusters in compartment1 '%s' ", f.Compartment1)
	for _, cluster := range f.ListClusters() {
		if cluster.LifecycleState == oke.ClusterSummaryLifecycleStateActive || cluster.LifecycleState == oke.ClusterSummaryLifecycleStateFailed {
			f.DeleteCluster(*cluster.Id, waitForDeleted)
		}
	}
}

// CreateKubeClient creates a kubeClient for the specified cluster.
func (f *Framework) CreateKubeClient(clusterID string) *KubeClient {
	Logf("Creating kube client, clusterID: %s", clusterID)
	kubeConfig := f.CreateClusterKubeconfigContent(clusterID)
	return NewKubeClient(kubeConfig)
}

func (f *Framework) RunKuberang(clusterID string) {
	// clever trick to get kuberang based on where we are running from. Since we package everything
	// into the repository and don't build a binary this works fine.
	_, filename, _, _ := runtime.Caller(0)
	kuberangBinary := path.Join(path.Dir(filename), fmt.Sprintf("../kuberang/%s/%s/kuberang", runtime.GOOS, runtime.GOARCH))

	Logf(kuberangBinary)
	kubeConfig := f.CreateClusterKubeconfigContent(clusterID)
	Expect(kubeConfig).ToNot(BeEmpty())

	kubeconfigFile := "/tmp/" + clusterID
	err := ioutil.WriteFile(kubeconfigFile, []byte(kubeConfig), 0644)
	Expect(err).To(BeNil())

	defer func() {
		Expect(os.Remove(kubeconfigFile)).To(BeNil())
	}()

	cmd := exec.Command(kuberangBinary, "--kubeconfig", kubeconfigFile, "--registry-url", "iad.ocir.io/odx-oke/oke-public")
	cmd.Env = os.Environ()
	output, err := cmd.CombinedOutput()

	Logf("Kuberang output:\n", string(output))
	Expect(err).To(BeNil())
}

// CreateClusterKubeconfigContent gets a valid 'kubeconfig' file for the target
// cluster.
func (f *Framework) CreateClusterKubeconfigContent(clusterID string) string {
	Logf("Getting cluster kubeconfig, id: %s", clusterID)
	ctx, cancel := context.WithTimeout(f.context, f.timeout)

	// pass in the token version to request a v2 kubeconfig
	tokenVersion := "2.0.0"

	defer cancel()
	response, err := f.clustersClient.CreateKubeconfig(ctx, oke.CreateKubeconfigRequest{
		ClusterId:                             &clusterID,
		CreateClusterKubeconfigContentDetails: oke.CreateClusterKubeconfigContentDetails{TokenVersion: &tokenVersion},
	})

	if err != nil {
		Logf("CreateKubeconfig error:  %s; retrying...", err)
		response, err = f.clustersClient.CreateKubeconfig(ctx, oke.CreateKubeconfigRequest{
			ClusterId: &clusterID,
		})
		if err != nil {
			Logf("CreateKubeconfig error:  %s; returning empty content", err)
			return ""
		}
	}

	content, err := ioutil.ReadAll(response.Content)
	if err != nil {
		Logf("ReadContent error:  %s; continuing", err)
	}
	//	Expect(err).NotTo(HaveOccurred())
	return string(content)
}

// CreateClusterKubeconfigContentWithKubeconfigRequest gets a valid 'kubeconfig' file for the target
// cluster using the kubeconfig request passed in by the caller.
func (f *Framework) CreateClusterKubeconfigContentWithKubeconfigRequest(clusterID string, kubeconfigRequest oke.CreateKubeconfigRequest) string {
	Logf("Getting cluster kubeconfig, id: %s", clusterID)

	ctx, cancel := context.WithTimeout(f.context, f.timeout)
	defer cancel()

	response, err := f.clustersClient.CreateKubeconfig(ctx, kubeconfigRequest)
	if err != nil {
		Logf("CreateKubeconfig error:  %s; continuing", err)
	}

	content, err := ioutil.ReadAll(response.Content)
	if err != nil {
		Logf("ReadContent error:  %s; continuing", err)
	}

	return string(content)
}

// GetWorkRequest gets the specified WorkRequest for an initialising cluster.
// Retries added to catch rare connectivity errors
func (f *Framework) GetWorkRequest(id string) oke.GetWorkRequestResponse {
	// retries added because of the occasional connection errors
	timeout := 10 * time.Minute
	ctx, cancel := context.WithTimeout(f.context, f.timeout)
	defer cancel()
	response, err := f.clustersClient.GetWorkRequest(ctx, oke.GetWorkRequestRequest{
		WorkRequestId: &id,
	})
	if err == nil {
		Expect(err).NotTo(HaveOccurred())
		return response
	}

	Logf("Error received on Get workRequest : '%s', Retrying", err)

	for start := time.Now(); time.Since(start) < timeout; time.Sleep(Poll) {
		response, err = f.clustersClient.GetWorkRequest(ctx, oke.GetWorkRequestRequest{
			WorkRequestId: &id,
		})
		if err == nil {
			Expect(err).NotTo(HaveOccurred())
			return response
		}

		Logf("Error received on Get workRequest : '%s', Retrying", err)
	}
	Failf("Timeout waiting for get workRequest '%s'", id)
	return response
}

// AquireRunLock blocks until the test run lock is required or a timeout
// elapses. A lock is required as only one test run can safely be executed on
// the same cluster at any given time.
func AquireRunLock(client clientset.Interface, lockName string) error {
	lec, err := makeLeaderElectionConfig(client, lockName)
	if err != nil {
		return err
	}

	readyCh := make(chan struct{})
	lec.Callbacks = leaderelection.LeaderCallbacks{
		OnStartedLeading: func(ctx context.Context) {
			Logf("Test run lock aquired")
			readyCh <- struct{}{}
		},
		OnStoppedLeading: func() {
			Failf("Lost test run lock unexpectedly")
		},
	}

	le, err := leaderelection.NewLeaderElector(*lec)
	if err != nil {
		return err
	}

	go le.Run(context.Background())

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-readyCh:
			return nil
		case <-ticker.C:
			Logf("Waiting to aquire test run lock. %q currently has it.", le.GetLeader())
		case <-time.After(2 * time.Minute):
			return errors.New("timed out trying to aquire test run lock")
		}
	}
	panic("unreachable")
}

func makeLeaderElectionConfig(client clientset.Interface, lockName string) (*leaderelection.LeaderElectionConfig, error) {
	eventBroadcaster := record.NewBroadcaster()
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: lockName})

	id := os.Getenv("WERCKER_STEP_ID")
	if id == "" {
		id = UniqueID()
	}

	Logf("Test run lock id: %q", id)

	rl, err := resourcelock.New(
		resourcelock.ConfigMapsResourceLock,
		"kube-system",
		lockName,
		client.CoreV1(),
		client.CoordinationV1(),
		resourcelock.ResourceLockConfig{
			Identity:      id,
			EventRecorder: recorder,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("couldn't create resource lock: %v", err)
	}

	return &leaderelection.LeaderElectionConfig{
		Lock:          rl,
		LeaseDuration: 60 * time.Second,
		RenewDeadline: 15 * time.Second,
		RetryPeriod:   2 * time.Second,
	}, nil
}

// CreateAndAwaitDaemonSet creates/updates the given DaemonSet and waits for it
// to be ready.
func CreateAndAwaitDaemonSet(client clientset.Interface, desired *appsv1.DaemonSet) error {
	actual, err := client.AppsV1().DaemonSets(desired.Namespace).Create(desired)
	if err != nil {
		if !apierrors.IsAlreadyExists(err) {
			return errors.Wrapf(err, "failed to create %q DaemonSet", desired.Name)
		}
		Logf("%q DaemonSet already exists. Updating.", desired.Name)
		actual, err = client.AppsV1().DaemonSets(desired.Namespace).Update(desired)
		if err != nil {
			return errors.Wrapf(err, "updating DaemonSet %q", desired.Name)
		}
	} else {
		Logf("Created DaemonSet %q in namespace %q", actual.Name, actual.Namespace)
	}

	return wait.PollImmediate(5*time.Second, 5*time.Minute, func() (bool, error) {
		actual, err := client.AppsV1().DaemonSets(actual.Namespace).Get(actual.Name, metav1.GetOptions{})
		if err != nil {
			return false, errors.Wrap(err, "waiting for DaemonSet to be ready")
		}

		if actual.Status.DesiredNumberScheduled != 0 && actual.Status.NumberReady == actual.Status.DesiredNumberScheduled {
			return true, nil
		}

		Logf("%q DaemonSet not yet ready (desired=%d, ready=%d). Waiting...",
			actual.Name, actual.Status.DesiredNumberScheduled, actual.Status.NumberReady)
		return false, nil
	})
}

// CreateAndAwaitDeployment creates/updates the given Deployment and waits for
// it to be ready.
func CreateAndAwaitDeployment(client clientset.Interface, desired *appsv1.Deployment) error {
	actual, err := client.AppsV1().Deployments(desired.Namespace).Create(desired)
	if err != nil {
		if !apierrors.IsAlreadyExists(err) {
			return errors.Wrapf(err, "failed to create Deployment %q", desired.Name)
		}
		Logf("Deployment %q already exists. Updating.", desired.Name)
		actual, err = client.AppsV1().Deployments(desired.Namespace).Update(desired)
		if err != nil {
			return errors.Wrapf(err, "updating Deployment %q", desired.Name)
		}
	} else {
		Logf("Created Deployment %q in namespace %q", actual.Name, actual.Namespace)
	}

	return wait.PollImmediate(5*time.Second, 5*time.Minute, func() (bool, error) {
		actual, err := client.AppsV1().Deployments(actual.Namespace).Get(actual.Name, metav1.GetOptions{})
		if err != nil {
			return false, errors.Wrap(err, "waiting for Deployment to be ready")
		}
		if actual.Status.Replicas != 0 && actual.Status.Replicas == actual.Status.ReadyReplicas {
			return true, nil
		}
		Logf("%s Deployment not yet ready (replicas=%d, readyReplicas=%d). Waiting...",
			actual.Name, actual.Status.Replicas, actual.Status.ReadyReplicas)
		return false, nil
	})
}

func (f *Framework) CreateCloudConfig() config.Config {
	cloudConfig := config.Config{
		Auth:                  config.AuthConfig{
			Region:                f.Region,
			TenancyID:             f.Tenancy,
			UserID:                f.User.OCID,
			PrivateKey:            string(f.User.PrivateKey),
			Fingerprint:           f.User.Fingerprint,
			Passphrase:            f.User.PrivateKeyPassPhrase,
			RegionKey:             "",
			UseInstancePrincipals: false,
			CompartmentID:         "",
			PrivateKeyPassphrase:  "",
		},
		LoadBalancer:          &config.LoadBalancerConfig{
			DisableSecurityListManagement: true,
			Subnet1: f.LbSubnet1,
			Subnet2: f.LbSubnet2,
		},
		RateLimiter:           nil,
		RegionKey:             "",
		UseInstancePrincipals: false,
		CompartmentID:         f.Compartment1,
		VCNID:                 f.Vcn,
		UseServicePrincipals:  false,
	}
	return cloudConfig
}

func (f *Framework) SaveCloudConfig(cloudConfig config.Config) error {
	cloudConfigBytes, err := yaml.Marshal(&cloudConfig)
	if err != nil {
		return err
	}
	cloudConfigFile, err := os.Create(f.CloudConfigPath)
	if err != nil {
		return err
	}
	defer cloudConfigFile.Close()
	bytesWritten, err := cloudConfigFile.Write(cloudConfigBytes)
	if err != nil {
		return err
	}
	Logf("Cloud Config File bytes written %d at location %s", bytesWritten, f.CloudConfigPath)
	return nil
}

func (f *Framework) SaveKubeConfig(kubeconfig string) error {
	kubeconfigFile, err := os.Create(f.ClusterKubeconfigPath)
	if err != nil {
		return err
	}
	defer kubeconfigFile.Close()
	bytesWritten, err := kubeconfigFile.WriteString(kubeconfig)
	if err != nil {
		return err
	}
	Logf("Kubeconfig File bytes written %d at location %s", bytesWritten, f.ClusterKubeconfigPath)
	return nil
}
