# kubetest2-plugins

This project contains the [kubetest2](https://github.com/kubernetes-sigs/kubetest2) plugins for deploying k8s on different ppc64le cloud and run the tests on it. This plugin predominantly uses terraform for infrastructure provisioning and ansible for setting up k8s on the deployed infrastructure.

## kubetest2-powervs

kubetest2-powervs is a deployer created for deploying on [IBM Cloud Power Virtual Server](https://www.ibm.com/in-en/cloud/power-virtual-server) infrastructure.

## Development
```shell
$ MODE=dev ./hack/build.sh
$ export TF_DATA=`pwd`/data
$ ./bin/kubetest2-tf
```

### Update the Submodule
```shell script
$ UPDATE_SUBMODULE=true ./hack/build.sh
```
