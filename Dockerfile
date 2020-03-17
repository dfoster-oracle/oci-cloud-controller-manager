FROM oraclelinux:7-slim

COPY dist/* /usr/local/bin/

RUN yum install -y iscsi-initiator-utils-6.2.0.874-10.0.5.el7 \
 && yum install -y e2fsprogs \
 && yum clean all
