#!/bin/bash

##################################################################################################
# This template can be used to tweak the environment variables needed to run the E2E tests locally #
# Default behavior:
# * ENDPOINT: your dev0 endpoint - eg. api-mp-apiserver.dev.api.us-ashburn-1.clusters.oci.oc-test.com
# * USER/TENANCY: BOAT User and TenancyId (ocid1.tenancy.oc1..aaaaaaaagkbzgg6lpzrf47xzy4rjoxg4de6ncfiq2rncmjiujvy2hjgxvziq)
# * RESOURCES:  Compartments: ccm-e2e-tests in okedev tenancy
#               VCN, Subnets: oke-vcn-quick-ccm-e2e-test-c4cf8d39e in compartment ccm-e2e-tests

# To run the tests:
# 1. Change the FOCUS valiable here to specify the subset of E2E tests to run
# 2. Set OKE_DEPLOYMENT_ID and OCI_KEY_PASSPHRASE if needed
# 3. run 'source cluster-create-dev0-env-template.sh' to set the variables
# 4. run 'make run-ccm-e2e-tests-local`
##################################################################################################

# ------------------------------------------------------------------------------------------------
# Test vars:  these variables set up how the e2e tests behave
# ------------------------------------------------------------------------------------------------

# The test suites to run (can replace or add tags)
export FOCUS="\[cloudprovider\]"

# Specify the passphrase if you have one set up for your oci key, else leave blank
export OCI_KEY_PASSPHRASE=

# Set this variable to override the deployment/TM_ID to be other than the combination of the first letter of firstname + lastname as per your oracle.com email address
# OKE_DEPLOYMENT_ID = xyz

# ------------------------------------------------------------------------------------------------
# Test vars:  these variables set up how the e2e tests behave - optional to tweak
# ------------------------------------------------------------------------------------------------

# This variable tells the test not to install oci cli and wipe out your .oci/config
export LOCAL_RUN=1
export TC_BUILD=0

# The number of nodes to use for running tests
export E2E_NODE_COUNT=1

export IMAGE_PULL_REPO="iad.ocir.io/okedev/e2e-tests/"

CONFIG=${1:-~/.oci/config}
# reads DEFAULT or OC1 config (as dev envs are set up in ashburn)
OC1_CONFIG=$(awk '/oc1/ || /OC1/ || /DEFAULT/ {for(i=1; i<=5; i++) {getline; print}}' $CONFIG)

export OCI_TENANCY=$(awk -F= '/^tenancy/ {print $2;exit}' <<< "$OC1_CONFIG")
export OCI_USER=$(awk -F= '/^user/ {print $2;exit}' <<< "$OC1_CONFIG")
export OCI_REGION=$(awk -F= '/^region/ {print $2;exit}' <<< "$OC1_CONFIG")
export OCI_FINGERPRINT=$(awk -F= '/^fingerprint/ {print $2;exit}' <<< "$OC1_CONFIG")
OCI_KEY_FILE=$(awk -F= '/^key_file/ {print $2;exit}' <<< "$OC1_CONFIG")
export OCI_KEY=$(cat ${OCI_KEY_FILE} | base64)

# The contents of your public ssh key (if it is located here)
export PUB_SSHKEY="$(cat ~/.ssh/id_rsa.pub)"
export TM_ID=$(hack/gen-deploymentID)
if [ -z $TM_ID ]; then
    echo "User ID not set."
    return
fi

export OKE_ENDPOINT=api-${TM_ID}-apiserver.dev.api.${OCI_REGION}.clusters.oci.oc-test.com

# The path to your okei/secrets directory
export SECRETS_LOCAL=${GOPATH}/src/bitbucket.oci.oraclecorp.com/okei/secrets

# ------------------------------------------------------------------------------------------------
# Target resources:  these variables are related to the resources you want new clusters associated with
# Specify overrides below, if different from default values defined above.
# ------------------------------------------------------------------------------------------------

# - Compartment

# The ocid for the compartment in which the cluster will be created
# Typically this is the compartment named with your initials in okedev tenancy
export COMPARTMENT="ocid1.compartment.oc1..aaaaaaaa6pfueflc6fc364vopfw3yielvcq4cvzl7ddlf6xuq6uiuoaiv5sa"

# - VCN

# The ocid for the VCN in your the above compartment that cluster will use
export VCN="ocid1.vcn.oc1.iad.amaaaaaa2ahbgkyaqfn6l2tpwak6o4cyfdlzw6imwapkjhrl6ejm5hh4hdba"

# - Subnets

# The flag to use regional subnets (should be true)
export USE_REGIONALSUBNET=true

# The ocid for the lb subnet created in your vcn
export LBRGNSUBNET=ocid1.subnet.oc1.iad.aaaaaaaanf2waxfnrjhhes4xgi7e2jfrr7p3cuaufukkaznzsgtgy3i6ykqq

# The following are ignored for USE_REGIONALSUBNET=true, but may be required to be defined for validation checks
export LBSUBNET1=dummy
export LBSUBNET2=dummy

# The ocid for the non-lb subnet created in your vcn
export OCI_RGNSUBNET=ocid1.subnet.oc1.iad.aaaaaaaawxtdowo3evyuo7djwveulq2d3ecrfvihgendviqlgwy4wtvqx2xq

# The following are ignored for USE_REGIONALSUBNET=true, but may be required to be defined for validation checks
export OCI_SUBNET1=dummy
export OCI_SUBNET2=dummy
export OCI_SUBNET3=dummy

# The shape of nodes to create
export NODE_SHAPE=VM.Standard1.2

# The secrets to use
export REGION_SECRETS=dev0-iad

export ADLOCATION="IqDk:US-ASHBURN-AD-1"

#KMS key for CMEK testing
export CMEK_KMS_KEY="ocid1.key.oc1.iad.bbpvrcsaaaeuk.abuwcljsav7rilbt6bnu3dqoakpzdtxhfk27uixzdz3yk7jrwngptfwg5u5a"

# ------------------------------------------------------------------------------------------------
# Bonus: other stuff to tweak if you know what they mean
# ------------------------------------------------------------------------------------------------
export DELEGATION_GROUP_ID=NA

export OKE_WAIT_TIME=120

# Will affect the version of k8s that is installed
# export OKE_CLUSTER_K8S_VERSION_INDEX=-2
# export OKE_CLUSTER_K8S_VERSION_INDEX_UPGRADE_TO=-1

# Will affect the version of k8s in the nodepool
# export OKE_NODEPOOL_K8S_VERSION_INDEX=-2
# export OKE_NODEPOOL_K8S_VERSION_INDEX_UPGRADE_TO=-1


