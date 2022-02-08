module github.com/oracle/oci-cloud-controller-manager

go 1.15

replace (
	github.com/docker/docker => github.com/docker/engine v0.0.0-20181106193140-f5749085e9cb
	github.com/oracle/oci-go-sdk/v31 => bitbucket.oci.oraclecorp.com/sdk/oci-go-sdk/v31 v31.0.0-20201215183620-aed686bb60a8
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.11.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	k8s.io/api => k8s.io/api v0.19.12
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.12
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.12
	k8s.io/apiserver => k8s.io/apiserver v0.19.12
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.19.12
	k8s.io/client-go => k8s.io/client-go v0.19.12
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.19.12
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.19.12
	k8s.io/code-generator => k8s.io/code-generator v0.19.12
	k8s.io/component-base => k8s.io/component-base v0.19.12
	k8s.io/cri-api => k8s.io/cri-api v0.19.12
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.19.12
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.19.12
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.19.12
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.19.12
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.19.12
	k8s.io/kubectl => k8s.io/kubectl v0.19.12
	k8s.io/kubelet => k8s.io/kubelet v0.19.12
	k8s.io/kubernetes => k8s.io/kubernetes v1.19.12
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.19.12
	k8s.io/metrics => k8s.io/metrics v0.19.12
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.19.12
	oracle.com/oci/httpiam => bitbucket.oci.oraclecorp.com/goiam/httpiam.git v0.0.0-00000000000000-973dbb679788e9727a86d30ea0cccadcc0fe33d6 // 0.14.
	oracle.com/oci/httpsigner => bitbucket.oci.oraclecorp.com/goiam/httpsigner.git v0.0.0-00000000000000-e8cb27ebf4409946b295b9e22e511a52fc967e91 // 0.17.1
	oracle.com/oci/ociauthz => bitbucket.oci.oraclecorp.com/goiam/ociauthz.git v0.0.0-00000000000000-b00a4280e2092ac2c220111731965f49392734c1
	oracle.com/oci/ocihttpiam => bitbucket.oci.oraclecorp.com/goiam/ocihttpiam.git v0.0.0-00000000000000-996aa4a919d9e80238807c1c63c385980a0302a8
	oracle.com/oci/tagging => bitbucket.oci.oraclecorp.com/GOPLEX/tagging.git v0.0.0-00000000000000-20a2e48911da14e503935718f66588ab14aad8d4
	oracle.com/oke/oci-go-common => bitbucket.oci.oraclecorp.com/oke/oci-go-common.git v0.0.0-00000000000000-f93927b2b66cb1de2a10cf0f9f0d7e349bc0ae27

)

require (
	bitbucket.oci.oraclecorp.com/oke/oke-common v1.0.1-0.20201218050159-057951da132b
	github.com/NYTimes/gziphandler v1.0.1 // indirect
	github.com/container-storage-interface/spec v1.2.0
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2
	github.com/kubernetes-csi/csi-lib-utils v0.8.1
	github.com/kubernetes-csi/external-attacher v0.0.0-20201106010650-6d1beabd0fad //v3.0.2
	github.com/kubernetes-csi/external-provisioner v0.0.0-20210409185916-86c2ba950e76 // v2.0.5
	github.com/kubernetes-csi/external-resizer v1.0.1 // v1.0.1
	github.com/kubernetes-csi/external-snapshotter/client/v2 v2.2.0-rc3
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/oracle/oci-go-sdk/v31 v31.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.9.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.3
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20210520170846-37e1c6afe023
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.38.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.19.12
	k8s.io/apimachinery v0.19.12
	k8s.io/apiserver v0.19.12
	k8s.io/client-go v0.19.12
	k8s.io/cloud-provider v0.19.12
	k8s.io/component-base v0.19.12
	k8s.io/csi-translation-lib v0.19.12
	k8s.io/klog v1.0.0
	k8s.io/kubelet v0.19.12
	k8s.io/kubernetes v1.19.12
	k8s.io/utils v0.0.0-20210305010621-2afb4311ab10
	sigs.k8s.io/sig-storage-lib-external-provisioner/v6 v6.3.0
)
