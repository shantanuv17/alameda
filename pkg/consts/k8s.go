package consts

const (
	K8S_KIND_REPLICASET            = "ReplicaSet"
	K8S_KIND_REPLICATIONCONTROLLER = "ReplicationController"
	K8S_KIND_DEPLOYMENT            = "Deployment"
	K8S_KIND_DEPLOYMENTCONFIG      = "DeploymentConfig"
	K8S_KIND_ALAMEDASCALER         = "AlamedaScaler"
	K8S_KIND_STATEFULSET           = "StatefulSet"
)

// Regular expression of pod naming rule
const (
	DeploymentPodFormat       = `%s-(([a-z0-9]+)-([a-z0-9]+))$`
	StatefulSetPodFormat      = `%s-([0-9]+)$`
	DeploymentConfigPodFormat = `%s-(([0-9]+)-([a-z0-9]+))$`
)
