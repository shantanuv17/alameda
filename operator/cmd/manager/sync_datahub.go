package main

import (
	"time"

	datahub_client_application "github.com/containers-ai/alameda/operator/datahub/client/application"
	datahub_client_cluster "github.com/containers-ai/alameda/operator/datahub/client/cluster"
	datahub_client_controller "github.com/containers-ai/alameda/operator/datahub/client/controller"
	datahub_client_kafka "github.com/containers-ai/alameda/operator/datahub/client/kafka"
	datahub_client_namespace "github.com/containers-ai/alameda/operator/datahub/client/namespace"
	datahub_client_node "github.com/containers-ai/alameda/operator/datahub/client/node"
	datahub_client_pod "github.com/containers-ai/alameda/operator/datahub/client/pod"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	k8sutils "github.com/containers-ai/alameda/pkg/utils/kubernetes"
	"google.golang.org/grpc"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func syncResourcesWithDatahub(
	client client.Client, datahubConn *grpc.ClientConn,
	datahubClient *datahubpkg.Client, enabledDA bool) {
	for {
		clusterUID, err := k8sutils.GetClusterUID(client)
		if err == nil {
			scope.Infof("Get cluster UID %s successfully, and then try synchronzing resources with datahub.", clusterUID)
			break
		} else {
			scope.Infof("Sync resources with datahub failed. %s", err.Error())
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
	go func() {
		if err := datahub_client_application.SyncWithDatahub(client,
			datahubClient); err != nil {
			scope.Errorf("sync application failed at start due to %s", err.Error())
		}
	}()
	if !enabledDA {
		go func() {
			if err := datahub_client_namespace.SyncWithDatahub(client,
				datahubClient); err != nil {
				scope.Errorf("sync namespace failed at start due to %s", err.Error())
			}
		}()
		go func() {
			if err := datahub_client_node.SyncWithDatahub(client,
				datahubClient); err != nil {
				scope.Errorf("sync node failed at start due to %s", err.Error())
			}
		}()
		go func() {
			if err := datahub_client_cluster.SyncWithDatahub(client,
				datahubConn); err != nil {
				scope.Errorf("sync cluster failed at start due to %s", err.Error())
			}
		}()
		go func() {
			if err := datahub_client_controller.SyncWithDatahub(client,
				datahubConn); err != nil {
				scope.Errorf("sync controller failed at start due to %s", err.Error())
			}
		}()
		go func() {
			if err := datahub_client_pod.SyncWithDatahub(client,
				datahubConn); err != nil {
				scope.Errorf("sync pod failed at start due to %s", err.Error())
			}
		}()
		go func() {
			if err := datahub_client_kafka.SyncWithDatahub(client, datahubClient); err != nil {
				scope.Errorf("sync kafka failed at start due to %s", err.Error())
			}
		}()
	}
}
