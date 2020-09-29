# Developers Guide


## How to add new application controller into Alameda-Operator?
1. Define new struct and implement business logic in alameda/operator/controller/new_controller.go. The new struct might has follow procedures. 
- 1.1. Function(e.g. function "Reconcile") will be triggered by events(e.g. user inputs, states change) to reconcile the desired state.
- 1.2. Function prepares items that needs to be synchornized with remote(e.g. Datahub).
- 1.3. Function synchornizes items with remote.
- 1.4. Function collects result/state of last synchornization.
- 1.5. Function updates result/state for trigger inspecting.(e.g. Update AlamedaScaler status or send events to Datahub)
2. Declare and add the new controller into manager,(i.e. function "addControllersToManager" in [file](../cmd/manager/main.go))
3. Delete redudant data from Datahub while initializing.
- 3.1. Add logic into function [syncResourcesWithDatahub](../cmd/manager/sync_datahub.go)

## How to add new configuration into Alameda-Operator?
1. Define new configuration's Golang structure.
2. Add new field with the new type into [Config](../config.go).
3. Add default value or comments for the configuration into [file](../etc/operator.toml) which will be parsed when Alameda-Operator initializing.

## How to mock interface?
Use util "mockgen" to generate the mock implementation of interface.
Example using mockgen to generate mock of interface Kafka.Client.
```
mockgen --source=${GOPATH}/src/prophetstor.com/alameda/internal/pkg/message-queue/kafka/interface.go -destination=${GOPATH}/src/prophetstor.com/alameda/internal/pkg/message-queue/kafka/mock/mock.go -self_package=prophetstor.com/alameda/internal/pkg/message-queue/kafka/mock
``` 

## How to generate CRD?
Command:
```
cd ${GOPATH}/src/prophetstor.com/alameda/operator
make manifests
```
Reference of validataion markers:
- https://kubebuilder.io/reference/markers/crd.html
- https://kubebuilder.io/reference/markers/crd-validation.html

## Code generation for CustomResource
To generate code that implements function "DeepCopy" for custom resource's structure, execute targe "code-gen" in Makefile.
For further usage of k8s code-gen tools, inspect [doc](#https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/).