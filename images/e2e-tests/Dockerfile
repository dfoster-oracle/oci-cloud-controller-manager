FROM golang:1.10

ARG OCI_CLI_VERSION
ARG KUBECTL_VERSION

RUN apt-get update && apt-get install -y python3 openssl locales
RUN locale-gen C.UTF-8
ENV LC_ALL="C.UTF-8"
ENV LANG="C.UTF-8"
# Install test runner
RUN go get github.com/onsi/ginkgo/ginkgo

# Install OCI cli
RUN mkdir -p oci-cli-install && \
    curl -L "https://raw.githubusercontent.com/oracle/oci-cli/${OCI_CLI_VERSION}/scripts/install/install.sh" > oci-cli-install/install.sh && \
    chmod u+x ./oci-cli-install/install.sh && \
    ./oci-cli-install/install.sh --install-dir /usr/local/oci --accept-all-defaults && \
    ln -s /root/bin/oci /usr/local/bin/oci

# Install kubectl
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl && \
    chmod +x ./kubectl && \
    mv ./kubectl /usr/local/bin/kubectl

# Install sops to decrypt secrets
RUN curl -L https://artifactory.oci.oraclecorp.com/atd-release-generic-local/sops/refs/heads/master/7/sops-linux-amd64 > /usr/local/bin/sops && \
    chmod +x /usr/local/bin/sops


RUN mkdir -p $GOPATH/src/github.com/oracle/oci-cloud-controller-manager
COPY . $GOPATH/src/github.com/oracle/oci-cloud-controller-manager/
RUN chmod +x $GOPATH/src/github.com/oracle/oci-cloud-controller-manager/images/e2e-tests/e2e_pop_run.sh

RUN mv $GOPATH/src/github.com/oracle/oci-cloud-controller-manager/images/e2e-tests/create_oci_config.sh /usr/local/bin/create_oci_config.sh && chmod +x /usr/local/bin/create_oci_config.sh
