# E2E Tests

These tests are adapted from the [Service test suite][1] in the Kubernetes core
E2E tests.

## Pre-requisite
You need to have [ginkgo][2] installed and configured.

## Running

1. Set `$VERSION` of the CCM you want to deploy CCM to your cluster and set `$INSTALL_ENABLED=true`.
2. If you want to run tests on already existing CCM deployment in your cluster, you need not do step 1. 
3. Set `$KUBECONFIG` of the cluster which the tests should access. [kubeconfig][3]
4. Set `$CLOUDCONFIG` to the path that points to cloud-provider.yaml for your cluster. 
   An example of cloud-provider.yaml - See [provider-config-example.yaml](../../../manifests/provider-config-example.yaml)

```bash
make run-ccm-e2e-tests
```
Alternatively you can pass above arguments along with the run command below. 
```bash
KUBECONFIG=xxxx CLOUDCONFIG=xxx make run-ccm-e2e-tests
```

NOTE: Test suite will fail if executed behind a `$HTTP_PROXY` that returns a
200 OK response upon failure to connect.

## Additional Options

Additional seclist count based sanity checks can be applied during e2e testing
by providing the appropriate seclist ocids. Both must be supplied.

```bash
export CCM_SECLIST_ID="ocid1.securitylist.$ccmloadblancerid"
export K8S_SECLIST_ID="ocid1.securitylist.$k8sworkerid"
```

If you wish to use the tests in debug mode and want to look at the namespaces created by tests you can provide the option to keep the namespaces 
```bash
--delete-namespaces=false
```
By default the namespaces will be deleted. So if you use above option, then you should delete the resources manually after investigation is complete.
Alternatively, these values can be specified as command line parameters.

```bash
$ ginkgo -v -progress test/e2e/cloud-controller-manager -- \
    --kubeconfig=${HOME}/.kube/config \
    --delete-namespace=false \
    --cloud-config=$CLOUDCONFIG \
    --ccm-seclist-id=ocid1.securitylist.$ccmloadblancerid \
    --k8s-seclist-id=ocid1.securitylist.$k8sworkerid
```


[1]: https://github.com/kubernetes/kubernetes/blob/0cb15453dae92d8be66cf42e6c1b04e21a2d0fb6/test/e2e/network/service.go
[2]: https://onsi.github.io/ginkgo/
[3]: https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/
