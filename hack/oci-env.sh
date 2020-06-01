#!/bin/sh
# Source this script to export env vars from your oci config

CONFIG=${1:-~/.oci/config}

export OCI_TENANCY=$(awk -F= '/^tenancy/ {print $2;exit}' $CONFIG)
export OCI_REGION=$(awk -F= '/^region/ {print $2;exit}' $CONFIG)
export OCI_USER=$(awk -F= '/^user/ {print $2;exit}' $CONFIG)
export OCI_FINGERPRINT=$(awk -F= '/^fingerprint/ {print $2;exit}' $CONFIG)
export OCI_KEY=$(cat ${HOME}/.oci/oci_api_key.pem | base64)
