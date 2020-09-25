package utils

import (
	"fmt"
	"os"
	"strings"

	OpenshiftApiApps "github.com/openshift/api/apps"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	ApiResources "prophetstor.com/api/datahub/resources"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

// GetNamespacedNameKey returns string "namespaced/name"
func GetNamespacedNameKey(namespace, name string) string {
	return fmt.Sprintf("%s/%s", namespace, name)
}

//ParseResourceLinkForTopController parses resourcelink string to get top controller information
func ParseResourceLinkForTopController(resourceLink string) (*ApiResources.Controller, error) {
	res := strings.Split(resourceLink, "/")
	if len(res) >= 5 {
		kind := ApiResources.Kind_KIND_UNDEFINED
		switch res[3] {
		case "deployments":
			kind = ApiResources.Kind_DEPLOYMENT
		case "deploymentconfigs":
			kind = ApiResources.Kind_DEPLOYMENTCONFIG
		case "statefulsets":
			kind = ApiResources.Kind_STATEFULSET
		default:
			kind = ApiResources.Kind_KIND_UNDEFINED
		}
		return &ApiResources.Controller{
			ObjectMeta: &ApiResources.ObjectMeta{
				Namespace: res[2],
				Name:      res[4],
			},
			Kind: kind,
		}, nil
	}
	return &ApiResources.Controller{}, fmt.Errorf("resource link format is not correct")
}

func GetNodeInfoDefaultStorageSizeBytes() string {
	return os.Getenv("ALAMEDA_OPERATOR_DEFAULT_NODEINFO_STORAGESIZE_BYTES")
}

var (
	hasOpenshiftAPIAppsV1 *bool
)

// ServerHasOpenshiftAPIAppsV1 returns true if the api-server has apiGroup named in "apps.openshift.io"
func ServerHasOpenshiftAPIAppsV1() (bool, error) {

	if hasOpenshiftAPIAppsV1 == nil {
		if exist, err := serverHasAPIGroup(OpenshiftApiApps.GroupName); err != nil {
			return false, err
		} else {
			hasOpenshiftAPIAppsV1 = &exist
		}
	}

	return *hasOpenshiftAPIAppsV1, nil
}

func serverHasAPIGroup(apiGroupName string) (bool, error) {

	config, err := config.GetConfig()
	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return false, err
	}
	apiGroups, err := k8sClient.ServerGroups()
	if err != nil {
		return false, err
	}
	for _, apiGroup := range apiGroups.Groups {
		if apiGroup.Name == apiGroupName {
			return true, nil
		}
	}
	return false, nil
}

func GetTotalResourceFromContainers(containers []corev1.Container) corev1.ResourceRequirements {
	total := corev1.ResourceRequirements{
		Limits:   corev1.ResourceList{},
		Requests: corev1.ResourceList{},
	}
	for _, c := range containers {
		for resourceName, quantity := range c.Resources.Limits {
			q := total.Limits[resourceName]
			q.Add(quantity)
			total.Limits[resourceName] = q
		}
		for resourceName, quantity := range c.Resources.Requests {
			q := total.Requests[resourceName]
			q.Add(quantity)
			total.Requests[resourceName] = q
		}
	}

	return total
}
