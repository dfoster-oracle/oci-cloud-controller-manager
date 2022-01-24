FROM odo-docker-signed-local.artifactory.oci.oraclecorp.com/odx-oke/oke/golang-buildbox:1.16.7-fips-08f10c614600a2b2477d254e19479978184fb88c-47 as builder

ARG COMPONENT

ENV SRC /go/src/github.com/oracle/oci-cloud-controller-manager

ENV GOPATH /go/
RUN mkdir -p /go/bin $SRC
ADD . $SRC
WORKDIR $SRC

RUN COMPONENT=${COMPONENT} make clean build

FROM docker-remote.artifactory.oci.oraclecorp.com/oraclelinux:7-slim

RUN yum-config-manager --disable \* && yum-config-manager --add-repo https://artifactory.oci.oraclecorp.com/io-ol7-latest-yum-local && yum repolist enabled

RUN yum install -y util-linux \
  && yum install -y e2fsprogs \
  && yum clean all

COPY --from=0 /go/src/github.com/oracle/oci-cloud-controller-manager/dist/* /usr/local/bin/
