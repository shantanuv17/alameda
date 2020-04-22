package nginx

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/autoscaling/v1alpha1"
	"github.com/containers-ai/alameda/operator/pkg/nginx"
	k8sutils "github.com/containers-ai/alameda/pkg/utils/kubernetes"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r NginxRepository) SyncWithDatahub(ctx context.Context, k8sClient client.Client, conn *grpc.ClientConn) error {
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
		nginxs, err := r.ListNginxs(ctx, ListNginxsOption{
			ClusterName: clusterUID,
		})
		if err != nil {
			return errors.Wrap(err, "list nginxs from Datahub failed")
		}

		nginxsToDelete := make([]nginx.Nginx, 0)
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

		if err := r.DeleteNginxs(ctx, nginxsToDelete); err != nil {
			return errors.Wrap(err, "delete nginxs from Datahub failed")
		}
		return nil
	})

	return wg.Wait()
}
