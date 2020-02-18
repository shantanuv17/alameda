package kafka

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/autoscaling/v1alpha1"
	"github.com/containers-ai/alameda/operator/pkg/kafka"
	k8sutils "github.com/containers-ai/alameda/pkg/utils/kubernetes"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r KafkaRepository) SyncWithDatahub(ctx context.Context, k8sClient client.Client, conn *grpc.ClientConn) error {
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
		consumerGroups, err := r.ListConsumerGroups(ctx, ListConsumerGroupsOption{
			ClusterName: clusterUID,
		})
		if err != nil {
			return errors.Wrap(err, "list consumerGroups from Datahub failed")
		}

		consumerGroupsToDelete := make([]kafka.ConsumerGroup, 0)
		for _, consumerGroup := range consumerGroups {
			alamedaScalerNamespace := consumerGroup.AlamedaScalerNamespace
			alamedaScalerName := consumerGroup.AlamedaScalerName

			exist := true
			if nameSet, ok := alamedaScalerSet[alamedaScalerNamespace]; ok {
				if _, ok := nameSet[alamedaScalerName]; !ok {
					exist = false
				}
			} else {
				exist = false
			}

			if !exist {
				consumerGroupsToDelete = append(consumerGroupsToDelete, consumerGroup)
			}
		}

		if err := r.DeleteConsumerGroups(ctx, consumerGroupsToDelete); err != nil {
			return errors.Wrap(err, "delete consumerGroups from Datahub failed")
		}
		return nil
	})

	wg.Go(func() error {
		topics, err := r.ListTopics(ctx, ListTopicsOption{
			ClusterName: clusterUID,
		})
		if err != nil {
			return errors.Wrap(err, "list topics from Datahub failed")
		}

		topicsToDelete := make([]kafka.Topic, 0)
		for _, topic := range topics {
			alamedaScalerNamespace := topic.AlamedaScalerNamespace
			alamedaScalerName := topic.AlamedaScalerName

			exist := true
			if nameSet, ok := alamedaScalerSet[alamedaScalerNamespace]; ok {
				if _, ok := nameSet[alamedaScalerName]; !ok {
					exist = false
				}
			} else {
				exist = false
			}

			if !exist {
				topicsToDelete = append(topicsToDelete, topic)
			}
		}

		if err := r.DeleteTopics(ctx, topicsToDelete); err != nil {
			return errors.Wrap(err, "delete topics from Datahub failed")
		}
		return nil
	})

	return wg.Wait()
}
