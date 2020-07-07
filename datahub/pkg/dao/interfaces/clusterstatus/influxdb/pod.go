package influxdb

import (
	DaoClusterTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	RepoInfluxCluster "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/clusterstatus"
	"github.com/containers-ai/alameda/datahub/pkg/kubernetes/metadata"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
)

// Implement ContainerOperation interface
type Pod struct {
	InfluxDBConfig InternalInflux.Config
}

func NewPodWithConfig(config InternalInflux.Config) DaoClusterTypes.PodDAO {
	return &Pod{InfluxDBConfig: config}
}

func (p *Pod) CreatePods(pods []*DaoClusterTypes.Pod) error {
	delPodReq := DaoClusterTypes.NewDeletePodsRequest()
	for _, pod := range pods {
		podMeta := DaoClusterTypes.PodObjectMeta{}
		if pod.ObjectMeta != nil {
			podMeta.ObjectMeta = pod.ObjectMeta
		}
		if pod.TopController != nil && pod.TopController.ObjectMeta != nil {
			podMeta.TopController = pod.TopController.ObjectMeta
		}
		if pod.AlamedaPodSpec != nil && pod.AlamedaPodSpec.AlamedaScaler != nil {
			podMeta.AlamedaScaler = pod.AlamedaPodSpec.AlamedaScaler
		}
		delPodReq.PodObjectMeta = append(delPodReq.PodObjectMeta, &podMeta)
	}

	containerMap := make(map[string][]*DaoClusterTypes.Container)
	for _, pod := range pods {
		identifier := pod.ClusterNamespacePodName()
		containerMap[identifier] = make([]*DaoClusterTypes.Container, 0)
		for _, container := range pod.Containers {
			containerMap[identifier] = append(containerMap[identifier], container)
		}
	}

	// Do delete pods before creating them
	err := p.DeletePods(delPodReq)
	if err != nil {
		scope.Error("failed to delete pods first in influxdb before creating them")
		return err
	}

	// Create containers
	containerRepo := RepoInfluxCluster.NewContainerRepository(p.InfluxDBConfig)
	if err := containerRepo.CreateContainers(containerMap); err != nil {
		scope.Error(err.Error())
		return err
	}

	// Create pods
	podRepo := RepoInfluxCluster.NewPodRepository(p.InfluxDBConfig)
	if err := podRepo.CreatePods(pods); err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (p *Pod) ListPods(request *DaoClusterTypes.ListPodsRequest) ([]*DaoClusterTypes.Pod, error) {
	podRepo := RepoInfluxCluster.NewPodRepository(p.InfluxDBConfig)
	pods, err := podRepo.ListPods(request)
	if err != nil {
		scope.Error(err.Error())
		return make([]*DaoClusterTypes.Pod, 0), err
	}

	containerMap, err := p.listContainersByPods(pods)
	for clusterNamespaceNodeName, containers := range containerMap {
		for _, pod := range pods {
			if pod.ClusterNamespacePodName() == clusterNamespaceNodeName {
				for _, container := range containers {
					pod.Containers = append(pod.Containers, container)
				}
				break
			}
		}
	}

	controllers, err := p.listControllersByPods(pods)
	for _, pod := range pods {
		for _, controller := range controllers {
			if pod.TopController.ObjectMeta.Name == controller.ObjectMeta.Name && pod.TopController.Kind == controller.Kind {
				if pod.ObjectMeta.Namespace == controller.ObjectMeta.Namespace && pod.ObjectMeta.ClusterName == controller.ObjectMeta.ClusterName {
					pod.TopController = controller
					break
				}
			}
		}
	}

	return pods, nil
}

func (p *Pod) DeletePods(request *DaoClusterTypes.DeletePodsRequest) error {
	delContainerReq := DaoClusterTypes.NewDeleteContainersRequest()
	for _, podObjectMeta := range request.PodObjectMeta {
		containerMeta := DaoClusterTypes.ContainerObjectMeta{}
		containerMeta.TopControllerKind = podObjectMeta.Kind
		containerMeta.AlamedaScalerScalingTool = podObjectMeta.ScalingTool
		if podObjectMeta.ObjectMeta != nil {
			containerMeta.PodName = podObjectMeta.ObjectMeta.Name
			containerMeta.Namespace = podObjectMeta.ObjectMeta.Namespace
			containerMeta.NodeName = podObjectMeta.ObjectMeta.NodeName
			containerMeta.ClusterName = podObjectMeta.ObjectMeta.ClusterName
		}
		if podObjectMeta.TopController != nil {
			containerMeta.TopControllerName = podObjectMeta.TopController.Name
			if containerMeta.Namespace == "" {
				containerMeta.Namespace = podObjectMeta.TopController.Namespace
			}
			if containerMeta.ClusterName == "" {
				containerMeta.ClusterName = podObjectMeta.TopController.ClusterName
			}
		}
		if podObjectMeta.AlamedaScaler != nil {
			containerMeta.AlamedaScalerName = podObjectMeta.AlamedaScaler.Name
			if containerMeta.Namespace == "" {
				containerMeta.Namespace = podObjectMeta.AlamedaScaler.Namespace
			}
			if containerMeta.ClusterName == "" {
				containerMeta.ClusterName = podObjectMeta.AlamedaScaler.ClusterName
			}
		}
		delContainerReq.ContainerObjectMeta = append(delContainerReq.ContainerObjectMeta, &containerMeta)
	}

	// Delete pods
	podRepo := RepoInfluxCluster.NewPodRepository(p.InfluxDBConfig)
	if err := podRepo.DeletePods(request); err != nil {
		scope.Error(err.Error())
		return err
	}

	// Delete containers
	containerRepo := RepoInfluxCluster.NewContainerRepository(p.InfluxDBConfig)
	if err := containerRepo.DeleteContainers(delContainerReq); err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (p *Pod) listControllersByPods(pods []*DaoClusterTypes.Pod) ([]*DaoClusterTypes.Controller, error) {
	controllerRepo := RepoInfluxCluster.NewControllerRepository(p.InfluxDBConfig)

	request := DaoClusterTypes.NewListControllersRequest()
	for _, pod := range pods {
		controllerMeta := DaoClusterTypes.NewControllerObjectMeta(nil, nil, pod.TopController.Kind, pod.AlamedaPodSpec.ScalingTool)
		controllerMeta.ObjectMeta = &metadata.ObjectMeta{}
		controllerMeta.ObjectMeta.Name = pod.TopController.ObjectMeta.Name
		controllerMeta.ObjectMeta.ClusterName = pod.ObjectMeta.ClusterName
		controllerMeta.ObjectMeta.Namespace = pod.ObjectMeta.Namespace

		request.ControllerObjectMeta = append(request.ControllerObjectMeta, controllerMeta)
	}

	return controllerRepo.ListControllers(request)
}

func (p *Pod) listContainersByPods(pods []*DaoClusterTypes.Pod) (map[string][]*DaoClusterTypes.Container, error) {
	containerRepo := RepoInfluxCluster.NewContainerRepository(p.InfluxDBConfig)

	request := DaoClusterTypes.NewListContainersRequest()
	for _, pod := range pods {
		containerMeta := DaoClusterTypes.ContainerObjectMeta{}
		containerMeta.PodName = pod.ObjectMeta.Name
		containerMeta.Namespace = pod.ObjectMeta.Namespace
		containerMeta.NodeName = pod.ObjectMeta.NodeName
		containerMeta.ClusterName = pod.ObjectMeta.ClusterName
		containerMeta.TopControllerName = pod.TopController.ObjectMeta.Name
		containerMeta.TopControllerKind = pod.TopController.Kind
		containerMeta.AlamedaScalerName = pod.AlamedaPodSpec.AlamedaScaler.Name
		containerMeta.AlamedaScalerScalingTool = pod.AlamedaPodSpec.ScalingTool
		request.ContainerObjectMeta = append(request.ContainerObjectMeta, &containerMeta)
	}

	return containerRepo.ListContainers(request)
}
