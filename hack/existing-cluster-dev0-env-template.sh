#!/bin/bash

##################################################################################################
# This template can be used to tweak the environment variables needed to run the E2E tests locally #
# Default behavior:
# Runs test on an existing cluster in dev0-iad

# To run the tests:
# 1. Change the FOCUS variable here to specify the subset of E2E tests to run
# 2. Set CLUSTER_KUBECONFIG and CLOUD_CONFIG if needed
# 3. run 'source existing-cluster-dev0-env-template.sh' to set the variables
# 4. run 'make run-ccm-e2e-tests-local`
##################################################################################################

# The test suites to run (can replace or add tags)
export FOCUS="\[cloudprovider\]"

# This variable tells the test not to install oci cli and wipe out your .oci/config
export LOCAL_RUN=1
export TC_BUILD=0

# This allows you to use your existing cluster
export ENABLE_CREATE_CLUSTER=false

# Set path to kubeconfig of existing cluster if it does not exist in default path. Defaults to $HOME/.kube/config_*
export CLUSTER_KUBECONFIG_AMD=$HOME/.kube/config_amd
export CLUSTER_KUBECONFIG_ARM=$HOME/.kube/config_arm

# Set path to cloud_config of existing cluster if it does not exist in default path. Defaults to $HOME/cloudconfig_*
export CLOUD_CONFIG_AMD=$HOME/cloudconfig_amd
export CLOUD_CONFIG_ARM=$HOME/cloudconfig_arm


export IMAGE_PULL_REPO="iad.ocir.io/okedev/e2e-tests/"
export ADLOCATION="IqDk:US-ASHBURN-AD-1"

#KMS key for CMEK testing
export CMEK_KMS_KEY="ocid1.key.oc1.iad.bbpvrcsaaaeuk.abuwcljsav7rilbt6bnu3dqoakpzdtxhfk27uixzdz3yk7jrwngptfwg5u5a"

#NSG Network security group created in cluster's VCN
export NSG_OCIDS=""

#Reserved IP created in e2e test compartment
export RESERVED_IP="144.25.98.32"

#Architecture to run tests on
export ARCHITECTURE_AMD="AMD"
export ARCHITECTURE_ARM="ARM"

#Focus the tests : ARM, AMD or BOTH
export SCOPE="BOTH"

# For debugging the tests in existing cluster, do not turn it off by default.
# export DELETE_NAMESPACE=false

# FSS volume handle
# format is FileSystemOCID:serverIP:path
export FSS_VOLUME_HANDLE=""
export FSS_VOLUME_HANDLE_ARM=""
export LUSTRE_VOLUME_HANDLE=""
export LUSTRE_VOLUME_HANDLE_ARM=""
export LUSTRE_SUBNET_CIDR=""
export MNT_TARGET_ID=""
export MNT_TARGET_SUBNET_ID=
export MNT_TARGET_COMPARTMENT_ID=

export STATIC_SNAPSHOT_COMPARTMENT_ID=
export CREATE_UHP_NODEPOOL=FALSE

# Workload Identity Principal Feature only available for ENHANCED_CLUSTER
export CLUSTER_TYPE="ENHANCED_CLUSTER"

# For SKE node, node_info, node_lifecycle controller tests against PDE
# To setup PDE and point your localhost:25000 to the PDE CP API refer: Refer: https://bitbucket.oci.oraclecorp.com/projects/OKE/repos/oke-control-plane/browse/personal-environments/README.md
# export CE_ENDPOINT_OVERRIDE="http://localhost:25000"

# Ip family of cluster to create cluster as per required ip stack
export CLUSTER_IP_FAMILY="IPv4"
export NP_IMAGE_OS="Oracle-Linux-8"
export SKIP_CLUSTER_DELETION="false"

export FSS_VOLUME_HANDLE_IPV6=""
export MNT_TARGET_ID_IPV6=""
export MNT_TARGET_SUBNET_ID_IPV6=""
export MNT_TARGET_SUBNET_ID_DUAL_STACK=""
export OCI_K8SSUBNET_IPV6=""
export OCI_K8SSUBNET_DUAL_STACK=""
export OCI_NODESUBNET_IPV6=""
export OCI_NODESUBNET_DUAL_STACK=""
export LBRGNSUBNET_IPV6=""
export LBRGNSUBNET_DUAL_STACK=""
export EXISTING_CLUSTER_OCID=""
