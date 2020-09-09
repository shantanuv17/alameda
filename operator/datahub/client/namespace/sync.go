package namespace

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	autoscalingv1alpha2 "github.com/containers-ai/alameda/operator/apis/autoscaling/v1alpha2"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	k8SUtils "github.com/containers-ai/alameda/pkg/utils"
	k8sutils "github.com/containers-ai/alameda/pkg/utils/kubernetes"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func SyncWithDatahub(client client.Client, datahubClient *datahubpkg.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	clusterUID, err := k8sutils.GetClusterUID(client)
	if err != nil {
		return errors.Wrap(err, "get cluster uid failed")
	}

	alamedaScalerList := autoscalingv1alpha2.AlamedaScalerList{}
	err = client.List(ctx, &alamedaScalerList)
	if err != nil {
		return errors.Wrap(err, "list alamedascaler for namespace sync error")
	}

	nsEntities := []entities.ResourceClusterStatusNamespace{}
	namespaceList := corev1.NamespaceList{}
	if err := client.List(ctx, &namespaceList); err != nil {
		return errors.Errorf(
			"Sync namespaces with datahub failed due to list namespaces from cluster failed: %s",
			err.Error())
	}
	for idx := range namespaceList.Items {
		if !IsNSExcluded(namespaceList.Items[idx].Name, alamedaScalerList.Items) {
			nsEntities = append(nsEntities, entities.ResourceClusterStatusNamespace{
				Name:        namespaceList.Items[idx].Name,
				ClusterName: clusterUID,
			})
		}
	}

	if err := datahubClient.Create(&nsEntities); err != nil {
		return fmt.Errorf(
			"Sync namespaces with datahub failed due to register namespace failed: %s", err.Error())
	}

	// Clean up unexisting namespaces from Datahub
	existingNamespaceMap := make(map[string]bool)
	for _, namespace := range namespaceList.Items {
		existingNamespaceMap[namespace.GetName()] = true
	}

	existingNSEntities := []entities.ResourceClusterStatusNamespace{}
	err = datahubClient.List(&existingNSEntities)
	if err != nil {
		return fmt.Errorf(
			"Sync namespaces with datahub failed due to list namespaces from datahub failed: %s", err.Error())
	}
	namespacesNeedDeleting := []entities.ResourceClusterStatusNamespace{}
	for _, n := range existingNSEntities {
		if IsNSExcluded(n.Name, alamedaScalerList.Items) {
			namespacesNeedDeleting = append(namespacesNeedDeleting, n)
			continue
		}
		if _, exist := existingNamespaceMap[n.Name]; exist {
			continue
		}
		namespacesNeedDeleting = append(namespacesNeedDeleting, n)
	}
	if len(namespacesNeedDeleting) > 0 {
		err = datahubClient.Delete(&namespacesNeedDeleting)
		if err != nil {
			return errors.Wrap(err, "delete namespaces from Datahub failed")
		}
	}

	return nil
}

func IsNSExcluded(ns string, allScalerIns []autoscalingv1alpha2.AlamedaScaler) bool {

	for _, app := range allScalerIns {
		if app.GetNamespace() == ns {
			return false
		}
	}

	if ns == k8SUtils.GetRunningNamespace() {
		return true
	}

	excludeNamespaces := viper.GetStringSlice("namespace_exclusion.namespaces")
	excludeNSRegs := viper.GetStringSlice("namespace_exclusion.namespace_regs")
	for _, excludeNSReg := range excludeNSRegs {
		matched, _ := regexp.MatchString(excludeNSReg, ns)
		if matched {
			return true
		}
	}
	for _, excludeNamespace := range excludeNamespaces {
		if excludeNamespace == ns {
			return true
		}
	}
	return false
}
