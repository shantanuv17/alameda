# Architecture

Alameda Operator leverages kubebuilder to create controller to reconcile kubernetes resources. Currently, there are 6 types of controller instances running in the code, reside under [folder](./pkg/controller). These controller instances are listed below.
*   AlamedaScaler
*   Deployment
*   DeploymentConfig
*   Node
*   StatefulSet

Alameda Operator lists **Pods** owned by each k8s/openshift **Workload Controller** (Deployment, DeploymentConfig and StatefulSet) in **AlamedaScaler.Status**. To simplify and centralize this process, Alameda Operator only updates **AlamedaScaler.Status** in AlamedaScaler controller instance, other controller instances (Deployment, DeploymentConfig and StatefulSet) just try to trigger the process by updating the **AlamedaScaler.Spec.CustomResourceVersion**.

## Deployment

The deployment controller watches Deployment type, it tries to update the CustomResourceVersion in the spec of the AlamedaScaler which is found monitoring the Deployment.

## DeploymentConfig

The deploymentConfig controller watches DeploymentConfig type, it tries to update the CustomResourceVersion in the spec of the AlamedaScaler which is found monitoring the DeploymentConfig.

## Node 

The node controller watches Node type. It provides the detail information of the node to **Alameda Datahub**.

## StatefulSet

The statefulSet controller watches StatefulSet type, it tries to update the CustomResourceVersion in the spec of the AlamedaScaler which is found monitoring the StatefulSet.