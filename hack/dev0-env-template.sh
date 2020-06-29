#!/bin/bash

##################################################################################################
# This template can be used to set the environment variables needed to run the E2E tests locally #
# After filling it in properly, run 'source ./dev0-env-template.sh' to set the variables        #
# Then run 'make run-ccm-e2e-tests-local`
##################################################################################################



# ------------------------------------------------------------------------------------------------
# Test vars:  these variables set up how the e2e tests behave
# ------------------------------------------------------------------------------------------------

# This variable tells the test not to install oci cli and wipe out your .oci/config
export LOCAL_RUN=1
export TC_BUILD=0

# The number of nodes to use for running tests
export E2E_NODE_COUNT=1

# The test suites to run (can remove unwanted ones)
export FOCUS="\[cloudprovider\]"

export IMAGE_PULL_REPO="iad.ocir.io/okedev/e2e-tests/"

# ------------------------------------------------------------------------------------------------
# dev: These variables apply if you are running against dev in iad and using okedev tenancy
# ------------------------------------------------------------------------------------------------

# The endpoint (dev) where apiserver is running
export OKE_ENDPOINT=dev.api.us-ashburn-1.clusters.oci.oc-test.com

# The region where apiserver is running
export OCI_REGION=us-ashburn-1

# The okedev tenancy
export OCI_TENANCY=ocid1.tenancy.oc1..aaaaaaaaizwb7xbe7nt3wfb2jdrnsehrbip53s64qtwi2hx4y3ydkdmaeywq


# ------------------------------------------------------------------------------------------------
# Auth:  in this section, I'm assuming you're using a user in the okedev tenancy (not BOAT)
#        If you don't have a user in the okedev tenancy, you can create one like this:
#            1. Open Console to the okedev tenancy and navigate to Identity -> Users, Create a User
#               (username is typically full email address here)
#            2. Add your user to the Administrators group (under the groups tab)
#            3. Add the public key of your public/private key pair to the API Keys section
#            4. Note your user ocid, and API Key fingerprint for use below
# ------------------------------------------------------------------------------------------------

# The ocid of your user in the okedev tenancy - note this is different from your user ocid in BOAT
export OCI_USER=ocid1.user.oc1..[your ocid]

# The fingerprint from your user's api key (step 3 above), usually looks like ab:cd:ef:12:...
export OCI_FINGERPRINT=[your fingerprint]

# The base64 encrypted contents of your private key (if it is located in the same place as mine)
export OCI_KEY=$(cat ${HOME}/.oci/oci_api_key.pem | base64)

# The passphrase for your private key (if it was created with one)
export OCI_KEY_PASSPHRASE=[your passphrase]

# The contents of your public key (if it is located here)
export PUB_SSHKEY="$(cat ~/.ssh/id_rsa.pub)"

# The path to your okei/secrets directory
export SECRETS_LOCAL=[your path]/bitbucket.oci.oraclecorp.com/okei/secrets


# ------------------------------------------------------------------------------------------------
# Target resources:  these variables are related to the resources you want new clusters associated with
# ------------------------------------------------------------------------------------------------

# - Compartment

# The ocid for the compartment in which the cluster will be created
# Typically this is the compartment named with your initials in okedev tenancy
export COMPARTMENT="ocid1.compartment.oc1..[your ocid]"

export COMPARTMENT2="ocid1.compartment.oc1..[your ocid]"

# The "initials" used to identify your compartment  ie ("jl" for Jake Lindholm)
export TM_ID=[your initials]

# - VCN

# The ocid for the VCN in your the above compartment that cluster will use
export VCN="ocid1.vcn.oc1.iad.[your ocid]"

# - Subnets

# The flag to use regional subnets (should be true)
export USE_REGIONALSUBNET=true

# The ocid for the lb subnet created in your vcn
export LBRGNSUBNET=ocid1.subnet.oc1.iad.[your ocid]

# The following are ignored for USE_REGIONALSUBNET=true, but may be required to be defined for validation checks
export LBSUBNET1=ocid1.subnet.oc1.iad.aaaaaaaatnl6z7wgcgcuwbw43j6cdkku3ubysoo6niau2iuowcqn5of53kbq
export LBSUBNET2=ocid1.subnet.oc1.iad.aaaaaaaatnl6z7wgcgcuwbw43j6cdkku3ubysoo6niau2iuowcqn5of53kbq

# The ocid for the non-lb subnet created in your vcn
export OCI_RGNSUBNET=ocid1.subnet.oc1.iad.[your ocid]

# The following are ignored for USE_REGIONALSUBNET=true, but may be required to be defined for validation checks
export OCI_SUBNET1=ocid1.subnet.oc1.iad.aaaaaaaatdrateffd4wqccadtfla2p4naydltn26xmmwbsg7thzg34q7c6dq
export OCI_SUBNET2=ocid1.subnet.oc1.iad.aaaaaaaatdrateffd4wqccadtfla2p4naydltn26xmmwbsg7thzg34q7c6dq
export OCI_SUBNET3=ocid1.subnet.oc1.iad.aaaaaaaatdrateffd4wqccadtfla2p4naydltn26xmmwbsg7thzg34q7c6dq

# The shape of nodes to create
export NODE_SHAPE=VM.Standard1.2

# The secrets to use
export REGION_SECRETS=dev0-iad

export ADLOCATION="IqDk:US-ASHBURN-AD-1"

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

# For running the tests without creating cluster
# export ENABLE_CREATE_CLUSTER=false
# export CLUSTER_KUBECONFIG=<path to kubeconfig of existing cluster where tests should run>
# export CLOUD_CONFIG=<path to cloud config for existing cluster>
# For debugging the tests in existing cluster, do not turn it off by default.
# export DELETE_NAMESPACE=false

