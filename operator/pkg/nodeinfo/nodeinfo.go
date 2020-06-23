package nodeinfo

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	operatorutils "github.com/containers-ai/alameda/operator/pkg/utils"
	"github.com/containers-ai/alameda/pkg/provider"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

type role = string

const (
	masterRole role = "master"
	workerRole role = "worker"

	defaultNodeStorageSize = "100Gi"
)

var (
	roleMap = map[string]role{
		"node-role.kubernetes.io/master": masterRole,
	}
)

// NewNodeInfo creates node from k8s node
func NewNodeInfo(k8sNode corev1.Node, clusterName string) (
	entities.ResourceClusterStatusNode, error) {
	node := entities.ResourceClusterStatusNode{
		ClusterName: clusterName,
		Name:        k8sNode.Name,
	}

	rv := reflect.ValueOf(&node).Elem()
	keyStrs := []string{
		"Role", "Region", "Zone", "InstanceType", "Os",
		"Provider", "InstanceId", "StorageSize",
	}
	for idx := range keyStrs {
		key := keyStrs[idx]
		// parse node label information
		for labelKey, labelV := range k8sNode.Labels {
			if strings.Contains(labelKey, "stackpoint.") && strings.Contains(labelKey, "stackpoint.io/role") == false {
				continue
			}
			value := parseKeyValue(labelKey, key, labelV)
			if len(value) > 0 {
				rv.FieldByName(fmt.Sprintf("IO%s", key)).SetString(string(labelV))
				break
			}
		}
	}
	node.IOStorageSize = k8sNode.Status.Capacity.StorageEphemeral().Value()
	if node.IORole == "" {
		found := false
		for key, role := range roleMap {
			if _, exist := k8sNode.Labels[key]; exist {
				found = true
				node.IORole = role
				break
			}
		}
		if !found {
			node.IORole = workerRole
		}
	}

	if len(k8sNode.Spec.ProviderID) > 0 {
		provider, _, instanceID := parseProviderID(k8sNode.Spec.ProviderID)
		node.IOProvider = provider
		node.IOInstanceId = instanceID
	}

	// Below ard original convert logic
	node.CreateTime = k8sNode.ObjectMeta.GetCreationTimestamp().Unix()

	cpuCores, ok := k8sNode.Status.Capacity.Cpu().AsInt64()
	if !ok {
		return entities.ResourceClusterStatusNode{}, errors.Errorf("cannot convert cpu capacity from k8s Node")
	}
	node.NodeCPUCores = cpuCores

	memoryBytes, ok := k8sNode.Status.Capacity.Memory().AsInt64()
	if !ok {
		return entities.ResourceClusterStatusNode{}, errors.Errorf("cannot convert memory capacity from k8s Node")
	}
	node.NodeMemoryBytes = memoryBytes

	if regionMap, exist := provider.ProviderRegionMap[node.IOProvider]; exist {
		if region, exist := regionMap[node.IORegion]; exist {
			node.IORegion = region
		}
	}

	// set default storage size
	storageSize := operatorutils.GetNodeInfoDefaultStorageSizeBytes()
	if storageSize == "" {
		storageSize = defaultNodeStorageSize
	}
	defaultNodeStorageQuantity := resource.MustParse(storageSize)
	if node.IOStorageSize == 0 {
		node.IOStorageSize = defaultNodeStorageQuantity.Value()
	}

	return node, nil
}

func parseKeyValue(strParse string, key string, value string) string {
	pattern, err := regexp.Compile(strings.ToLower(fmt.Sprintf("/%s$", key)))
	if err != nil {
		return ""
	}
	if len(pattern.FindString(strings.Replace(strParse, "-", "", -1))) > 0 {
		return value
	}
	return ""
}

func parseProviderID(providerID string) (string, string, string) {
	var provider string
	var region string
	var instanceID string
	rex, err := regexp.Compile("([^\\:/]+)")
	if err != nil {
		fmt.Println(err)
		return "", "", ""
	}
	res := rex.FindAllString(providerID, -1)
	if res == nil || len(res) == 0 {
		return "", "", ""
	}
	for i := 0; i < len(res) && i < 3; i++ {
		switch i {
		case 0:
			provider = res[i]
		case 1:
			region = res[i]
		case 2:
			instanceID = res[i]
		}
	}
	return provider, region, instanceID
}
