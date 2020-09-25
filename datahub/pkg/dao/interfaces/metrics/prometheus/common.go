package prometheus

import (
	"fmt"

	"github.com/pkg/errors"
	DaoMetricTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/types"
	"prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
	"prophetstor.com/alameda/pkg/utils/log"
	ApiResources "prophetstor.com/api/datahub/resources"
)

var (
	scope = log.RegisterScope("dao_metrics_prometheu", "metrics dao implementing with Prometheus", 0)
)

func filterObjectMetaByClusterUID(clusterUID string, objectMetas []metadata.ObjectMeta) []metadata.ObjectMeta {
	newObjectMetas := make([]metadata.ObjectMeta, 0, len(objectMetas))
	for _, objectMeta := range objectMetas {
		if objectMeta.ClusterName == clusterUID {
			newObjectMetas = append(newObjectMetas, objectMeta)
		}
	}
	return newObjectMetas
}

func listPodNamesRegExpByControllerObjectMetas(objectMetas []DaoMetricTypes.ControllerObjectMeta) ([]string, error) {

	controllerKindToRegExpMap := map[string]string{
		ApiResources.Kind_DEPLOYMENTCONFIG.String(): "%s-[0-9]{0,7}-[0-9a-z]*",
		ApiResources.Kind_DEPLOYMENT.String():       "%s-[0-9a-z]{8,10}-[0-9a-z]*",
		ApiResources.Kind_STATEFULSET.String():      "%s-[0-9]+",
	}

	podNameRegExps := make([]string, len(objectMetas))
	for i, objectMeta := range objectMetas {
		format, exist := controllerKindToRegExpMap[objectMeta.Kind]
		if !exist {
			return nil, errors.Errorf(`not supported kind: %s`, objectMeta.Kind)
		}
		podNameRegExps[i] = fmt.Sprintf(format, objectMeta.Name)
	}
	return podNameRegExps, nil
}
