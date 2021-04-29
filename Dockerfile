
FROM iad.ocir.io/odx-oke/oke/golang-buildbox:1.13.15-fips as builder

ARG COMPONENT

ENV SRC /go/src/github.com/oracle/oci-cloud-controller-manager

ENV GOPATH /go/
RUN mkdir -p /go/bin $SRC
ADD . $SRC
WORKDIR $SRC

RUN COMPONENT=${COMPONENT} make clean build

FROM docker-remote.artifactory.oci.oraclecorp.com/oraclelinux:7-slim

RUN yum install -y util-linux \
  && yum install -y e2fsprogs \
  && yum clean all

COPY --from=0 /go/src/github.com/oracle/oci-cloud-controller-manager/dist/* /usr/local/bin/
