package nginx

import (
	"context"
	"time"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	k8sutils "github.com/containers-ai/alameda/pkg/utils/kubernetes"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func SyncWithDatahub(k8sClient client.Client, datahubClient *datahubpkg.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	clusterUID, err := k8sutils.GetClusterUID(k8sClient)
	if err != nil {
		return errors.Wrap(err, "get cluster uid failed")
	}

	empty := struct{}{}

	alamedaScalerSet := make(map[string]map[string]struct{})
	alamedaScalerList := autoscalingv1alpha1.AlamedaScalerList{}
	err = k8sClient.List(ctx, &alamedaScalerList)
	if err != nil {
		return errors.Wrap(err, "list AlamedaScaler failed")
	}
	for _, item := range alamedaScalerList.Items {
		namesapce := item.Namespace
		name := item.Name
		if _, exist := alamedaScalerSet[namesapce]; !exist {
			alamedaScalerSet[namesapce] = make(map[string]struct{})
		}
		alamedaScalerSet[namesapce][name] = empty
	}

	wg := errgroup.Group{}
	wg.Go(func() error {
		nginxs := []entities.ApplicationNginx{}
		err := datahubClient.List(&nginxs, datahubpkg.Option{
			Entity: entities.ApplicationNginx{
				ClusterName: clusterUID,
			},
			Fields: []string{"ClusterName"},
		})
		if err != nil {
			return errors.Wrap(err, "list nginxs from Datahub failed")
		}

		nginxsToDelete := make([]entities.ApplicationNginx, 0)
		for _, nginx := range nginxs {
			alamedaScalerNamespace := nginx.AlamedaScalerNamespace
			alamedaScalerName := nginx.AlamedaScalerName

			exist := true
			if nameSet, ok := alamedaScalerSet[alamedaScalerNamespace]; ok {
				if _, ok := nameSet[alamedaScalerName]; !ok {
					exist = false
				}
			} else {
				exist = false
			}

			if !exist {
				nginxsToDelete = append(nginxsToDelete, nginx)
			}
		}

		if err := datahubClient.Delete(&nginxsToDelete); err != nil {
			return errors.Wrap(err, "delete nginxs from Datahub failed")
		}
		return nil
	})

	return wg.Wait()
}
