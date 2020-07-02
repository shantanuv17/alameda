package namespace

import (
	"context"
	"fmt"
	"regexp"
	"time"

	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	k8SUtils "github.com/containers-ai/alameda/pkg/utils"
	k8sutils "github.com/containers-ai/alameda/pkg/utils/kubernetes"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func SyncWithDatahub(client client.Client, conn *grpc.ClientConn) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	clusterUID, err := k8sutils.GetClusterUID(client)
	if err != nil {
		return errors.Wrap(err, "get cluster uid failed")
	}

	datahubNamespaceRepo := NewNamespaceRepository(conn, clusterUID)
	alamedaScalerList := autoscalingv1alpha1.AlamedaScalerList{}
	err = client.List(ctx, &alamedaScalerList)
	if err != nil {
		return errors.Wrap(err, "list alamedascaler for namespace sync error")
	}

	nsEntities := []corev1.Namespace{}
	namespaceList := corev1.NamespaceList{}
	if err := client.List(ctx, &namespaceList); err != nil {
		return errors.Errorf(
			"Sync namespaces with datahub failed due to list namespaces from cluster failed: %s",
			err.Error())
	}
	for idx := range namespaceList.Items {
		if !IsNSExcluded(namespaceList.Items[idx].Name, alamedaScalerList.Items) {
			nsEntities = append(nsEntities, namespaceList.Items[idx])
		}
	}

	if len(nsEntities) > 0 {
		if err := datahubNamespaceRepo.CreateNamespaces(nsEntities); err != nil {
			return fmt.Errorf(
				"Sync namespaces with datahub failed due to register namespace failed: %s", err.Error())
		}
	}

	// Clean up unexisting namespaces from Datahub
	existingNamespaceMap := make(map[string]bool)
	for _, namespace := range namespaceList.Items {
		existingNamespaceMap[namespace.GetName()] = true
	}

	namespacesFromDatahub, err := datahubNamespaceRepo.ListNamespaces()
	if err != nil {
		return fmt.Errorf(
			"Sync namespaces with datahub failed due to list namespaces from datahub failed: %s", err.Error())
	}
	namespacesNeedDeleting := []*datahub_resources.Namespace{}
	for _, n := range namespacesFromDatahub {
		if IsNSExcluded(n.GetObjectMeta().GetName(), alamedaScalerList.Items) {
			namespacesNeedDeleting = append(namespacesNeedDeleting, n)
			continue
		}
		if _, exist := existingNamespaceMap[n.GetObjectMeta().GetName()]; exist {
			continue
		}
		namespacesNeedDeleting = append(namespacesNeedDeleting, n)
	}
	if len(namespacesNeedDeleting) > 0 {
		err = datahubNamespaceRepo.DeleteNamespaces(namespacesNeedDeleting)
		if err != nil {
			return errors.Wrap(err, "delete namespaces from Datahub failed")
		}
	}

	return nil
}

func IsNSExcluded(ns string, allScalerIns []autoscalingv1alpha1.AlamedaScaler) bool {

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
