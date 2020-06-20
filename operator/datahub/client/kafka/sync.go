package kafka

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	k8sutils "github.com/containers-ai/alameda/pkg/utils/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func SyncWithDatahub(k8sClient client.Client,
	datahubClient *datahubpkg.Client) error {
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
		consumerGroups := []entities.ApplicationKafkaConsumerGroup{}
		err := datahubClient.List(&consumerGroups, datahubpkg.Option{
			Entity: entities.ApplicationKafkaConsumerGroup{
				ClusterName: clusterUID,
			},
			Fields: []string{"ClusterName"},
		})
		if err != nil {
			return errors.Wrap(err, "list consumerGroups from Datahub failed")
		}

		consumerGroupsToDelete := make([]entities.ApplicationKafkaConsumerGroup, 0)
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

		if err := datahubClient.Delete(&consumerGroupsToDelete); err != nil {
			return errors.Wrap(err, "delete consumerGroups from Datahub failed")
		}
		return nil
	})

	wg.Go(func() error {
		topics := []entities.ApplicationKafkaTopic{}
		err := datahubClient.List(&topics, datahubpkg.Option{
			Entity: entities.ApplicationKafkaTopic{
				ClusterName: clusterUID,
			},
			Fields: []string{"ClusterName"},
		})
		if err != nil {
			return errors.Wrap(err, "list topics from Datahub failed")
		}

		topicsToDelete := make([]entities.ApplicationKafkaTopic, 0)
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

		if err := datahubClient.Delete(&topicsToDelete); err != nil {
			return errors.Wrap(err, "delete topics from Datahub failed")
		}
		return nil
	})

	return wg.Wait()
}
