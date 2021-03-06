## AlamedaScaler Custom Resource Definition

After Alameda is installed, it does not orchestrate any pod resources by default.
Alameda use _alamedascaler_ [CRD](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) as a channel for users to tell Alameda which pods needs autoscaling services and what policy to follow.

Here is an example _alamedascaler_ CR:

```
  apiVersion: autoscaling.containers.ai/v1alpha1
  kind: AlamedaScaler
  metadata:
    name: alameda
    namespace: webapp
    labels:
      app.federator.ai/name: db
      app.federator.ai/part-of: wordpress
  spec:
    policy: stable
    scalingTool:
      type: hpa
    enableExecution: false
    selector:
      matchLabels:
        app: mysql
```

In this example, it creates an _AlamedaScaler_ CR with name _alameda_ in namespace _webapp_. With this CR, Alameda will look for K8s deployment and deploymentconfig resource objects with label _app_ equals to _nginx_ in the same _webapp_ namespace. Any containers derivated from the found objects will be managed for their resource usages by Alameda.
Please also note that the label `app.federator.ai/name` of this CR tells Alameda that the selected objects by this CR belongs to application _nginx_ and the label `app.federator.ai/part-of` says they are also part of the higher application called _wordpress_. In another word, the _wordpress_ application is built by _nginx_ and maybe also other components, and this AlamedaScaler CR is created to autoscale the _nginx_ deployment.  
For detailed _AlamedaScaler_ schema, check out the remaining sections of this document.

> **Note:** The supported K8s api objects are created by resource _kind_:
- ```Deployment``` of _groupversion_ ```apps/v1```, ```apps/v1beta1```, ```apps/v1beta2```, ```extentions/v1beta1``` and
- ```DeploymentConfig``` of _groupversion_ ```apps.openshift.io/v1```

When an _AlamedaScaler_ CR is created, Alameda will process it and add the selected K8s api objects information. For example, you can see from the `status` field of the following _AlamedaScaler_ CR to know what K8s resources are watched:
```
$ kubectl get alamedascaler -n alameda -o yaml
apiVersion: v1
items:
- apiVersion: autoscaling.containers.ai/v1alpha1
  kind: AlamedaScaler
  metadata:
    annotations:
      kubectl.kubernetes.io/last-applied-configuration: |
        {"apiVersion":"autoscaling.containers.ai/v1alpha1","kind":"AlamedaScaler","metadata":{"annotations":{},"name":"as","namespace":"alameda"},"spec":{"enable":true,"policy":"stable","selector":{"matchLabels":{"app.kubernetes.io/name":"alameda-ai"}}}}
    creationTimestamp: "2019-03-05T05:51:34Z"
    generation: 2
    name: as
    namespace: alameda
    resourceVersion: "1232719"
    selfLink: /apis/autoscaling.containers.ai/v1alpha1/namespaces/alameda/alamedascalers/as
    uid: bb9e1b3f-3f0a-11e9-b062-08606e0a1cbb
  spec:
    enableExecution: false
    policy: stable
    scalingTool:
      type: hpa
    selector:
      matchLabels:
        app.kubernetes.io/name: alameda-ai
  status:
    alamedaController:
      deploymentConfigs: {}
      deployments:
        alameda/alameda-ai:
          name: alameda-ai
          namespace: alameda
          pods:
            alameda/alameda-ai-7f5b6b6d8-8fqrv:
              containers:
              - name: alameda-ai
                resources: {}
              name: alameda-ai-7f5b6b6d8-8fqrv
              namespace: alameda
              uid: 2eb43d4c-3eee-11e9-b062-08606e0a1cbb
          uid: 28c96445-39b7-11e9-b062-08606e0a1cbb
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""

```

The `status` field shows no _deploymentconfigs_ resource is selected and one _deployment_ called _alameda-ai_ is seleted.

## Schema of AlamedaScaler

- Field: metadata
  - type: ObjectMeta
  - description: This follows the ObjectMeta definition in [Kubernetes API Reference](https://kubernetes.io/docs/reference/#api-reference).  
One special note is Alameda uses labels `app.federator.ai/name` and `app.federator.ai/part-of` to correlate whether selected objects of different AlamedaScaler CRs belong to the same application. If `app.federator.ai/name` label is not given, Alameda will use "`metadata.namespace`-`metadata.name`" as its value. If `app.federator.ai/part-of` is not given, Alameda will use the value of `app.federator.ai/name` as its value.
  > **NOTE**: Do not set metadata name to `alamedaSelfDriving`

  > **NOTE**: The application label lookup is cross namespaces. If two CRs are created in different namespaces with the same `app.kubernetes.io/part-of` value, the selected `deployment`/`deploymentconfig` objects are considered be part of the same application.

  The following is a table to describe the two special labels.

Key                       | Descripion                                                                                                  | Example   | Type
--------------------------|-------------------------------------------------------------------------------------------------------------|-----------|-----
app.federator.ai/name     | The name of the application that is built by the selected objects of this CR                                | mysql     | string 
app.federator.ai/part-of  | The name of a higher level application this one is part of                                                  | wordpress | string

- Field: spec
  - type: [AlamedaScalerSpec](#alamedascalerspec)
  - description: Spec of AlamedaScaler.

### AlamedaScalerSpec

- Field: type
  - type: string
  - description:Which type is this AlamedaScaler, currently support _default_ and _kafka_.
- Field: selector
  - type: LabelSelector
  - description: This field is only effective whe type equals to _default_. This follows the _LabelSelector_ definition in [Kubernetes API Reference](https://kubernetes.io/docs/reference/#api-reference) except that Alameda only processes the `matchLabels` field of `LabelSelector`.
- Field: kafkaSpec
  - type: [KafkaSpec](#kafkaspec)
  - description: This field is only effective whe type equals to _kafka_. AlamedaScaler will use this field to decide which topics and consumer-groups need to be monitored.
- Field: policy
  - type: string
  - description: Policy used by Alameda for resource recommendations. _stable_ and _compact_ are supported. Default is _stable_.
- Field: scalingTool
  - type: [ScalingToolSpec](#scalingtoolspec)
  - description: Scaling tool configuration.
- Field: enableExecution
  - type: boolean
  - description: Set to _true_ to enable recommendation execution for api objects selected by this AlamedaScaler. Default is _false_.

### ScalingToolSpec

- Field: type
  - type: string
  - description: Type of scaling tool that will be used in recommendations. Currently supported tools are _N/A_ and _hpa_.
_N/A_ means Alameda will only produces resource predictions.
_hpa_ means Alameda will make recommendations to change the _replicas_ of a managed _Deployment_/_DeploymentConfig_ object. Default is _N/A_.

### KafkaSpec

- Field: exporterNamespace
  - type: string
  - description: Defines which namespace does Kafka-Exported resides in.
- Field: topics
  - type: array of string
  - description: Defines which topics that the consumer-groups might be consuming.
- Field: consumerGroups
  - type:  array of [KafkaConsumerGroupSpec](#kafkaconsumergroupspec) 
  - description: Defines which consumer-groups need to be monitored.  | [][KafkaConsumerGroupSpec](#kafkaconsumergroupspec)

### KafkaConsumerGroupSpec

- Field: name
  - type: string
  - description: Name of the consumer-group. 
- Field: majorTopic
  - type: string
  - description: The topic name that user want Alameda choose to predict, if it is empty, Alameda will traverses AlamedaScaler.Spec.Kafka.Topics and choose the first topic to be predicted if the consumer-group is currently consuming the topic in Kafka.
- Field: minReplicas
  - type: integer
  - description: **At least** how many replicas need to running when autoscaling. Default is _1_. 
- Field: maxReplicas
  - type: integer
  - description: **At most** how many replicas can run when autoscaling. Default value will be decided dynamically by the count of topic's partitions that Alameda choose to predict.
- Field: resource
  - type: [KafkaConsumerGroupResourceSpec](#kafkaconsumergroupresourcespec)
  - description: Way to map the consumer-group to the workload.

### KafkaConsumerGroupResourceSpec

- Field: custom
  - type: string
  - description: A name to map to the consumer-group when it cannot be directly mapped to a resource in Kubernetes.
- Field: kubernetes
  - type: [KubernetesResourceSpec](#kubernetesresourcespec)
  - description: This filed provides user to tell Alameda what's the resource in Kubernetes belongs to the consumer-group.

### KubernetesResourceSpec

- Field: selector
  - type: LabelSelector
  - description: This follows the _LabelSelector_ definition in [Kubernetes API Reference](https://kubernetes.io/docs/reference/#api-reference) except that Alameda only processes the `matchLabels` field of `LabelSelector`.