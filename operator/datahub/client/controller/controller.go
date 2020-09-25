package controller

import (
	"context"

	appsapi_v1 "github.com/openshift/api/apps/v1"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	"prophetstor.com/alameda/operator/datahub/client"
	datahubpkg "prophetstor.com/alameda/pkg/datahub"
	datahub_resources "prophetstor.com/api/datahub/resources"
)

type ControllerRepository struct {
	datahubClient *datahubpkg.Client
	clusterUID    string
}

// NewControllerRepository return ControllerRepository instance
func NewControllerRepository(datahubClient *datahubpkg.Client, clusterUID string) *ControllerRepository {
	return &ControllerRepository{
		datahubClient: datahubClient,

		clusterUID: clusterUID,
	}
}

func (repo *ControllerRepository) ListControllers() ([]*datahub_resources.Controller, error) {
	req := datahub_resources.ListControllersRequest{
		ObjectMeta: []*datahub_resources.ObjectMeta{
			{
				ClusterName: repo.clusterUID,
			},
		},
	}

	resp, err := repo.datahubClient.ListControllers(&req)
	if err != nil {
		return nil, errors.Wrap(err, "list controllers from datahub failed")
	} else if resp == nil {
		return nil, errors.Errorf("list controllers from Datahub failed, receive nil response")
	} else if _, err := client.IsResponseStatusOK(resp.Status); err != nil {
		return nil, errors.Wrap(err, "list controllers from Datahub failed")
	}
	return resp.Controllers, nil
}

func (repo *ControllerRepository) ListControllersByApplication(ctx context.Context, namespace, name string) ([]*datahub_resources.Controller, error) {
	req := datahub_resources.ListControllersRequest{
		ObjectMeta: []*datahub_resources.ObjectMeta{
			{
				Namespace:   namespace,
				ClusterName: repo.clusterUID,
			},
		},
	}

	resp, err := repo.datahubClient.ListControllers(&req)
	if err != nil {
		return nil, errors.Wrap(err, "list controllers from datahub failed")
	} else if resp == nil {
		return nil, errors.Errorf("list controllers from Datahub failed, receive nil response")
	} else if _, err := client.IsResponseStatusOK(resp.Status); err != nil {
		return nil, errors.Wrap(err, "list controllers from Datahub failed")
	}
	controllers := make([]*datahub_resources.Controller, 0)
	for _, controller := range resp.Controllers {
		if controller != nil && repo.isControllerHasApplicationInfo(controller, namespace, name) {
			controllers = append(controllers, controller)
		}
	}
	return controllers, nil
}

// DeleteControllers delete controllers from datahub
func (repo *ControllerRepository) DeleteControllers(ctx context.Context, arg interface{}, kindIf interface{}) error {
	objMeta := []*datahub_resources.ObjectMeta{}
	kind := datahub_resources.Kind_KIND_UNDEFINED

	switch v := arg.(type) {
	case []*appsv1.Deployment:
		kind = datahub_resources.Kind_DEPLOYMENT
		for _, controller := range v {
			objMeta = append(objMeta, &datahub_resources.ObjectMeta{
				Name:        controller.GetName(),
				Namespace:   controller.GetNamespace(),
				ClusterName: repo.clusterUID,
			})
		}
	case []*appsv1.StatefulSet:
		kind = datahub_resources.Kind_STATEFULSET
		for _, controller := range v {
			objMeta = append(objMeta, &datahub_resources.ObjectMeta{
				Name:        controller.GetName(),
				Namespace:   controller.GetNamespace(),
				ClusterName: repo.clusterUID,
			})
		}
	case []*appsapi_v1.DeploymentConfig:
		kind = datahub_resources.Kind_DEPLOYMENTCONFIG
		for _, controller := range v {
			objMeta = append(objMeta, &datahub_resources.ObjectMeta{
				Name:        controller.GetName(),
				Namespace:   controller.GetNamespace(),
				ClusterName: repo.clusterUID,
			})
		}
	case []*datahub_resources.Controller:
		for _, controller := range v {
			kind = controller.GetKind()
			objMeta = append(objMeta, &datahub_resources.ObjectMeta{
				Name:        controller.GetObjectMeta().GetName(),
				Namespace:   controller.GetObjectMeta().GetNamespace(),
				ClusterName: repo.clusterUID,
			})
		}
	case []*datahub_resources.ObjectMeta:
		if theKind, ok := kindIf.(datahub_resources.Kind); ok {
			kind = theKind
			objMeta = v
		}
	default:
		return errors.Errorf("not supported type(%T)", v)
	}

	req := datahub_resources.DeleteControllersRequest{
		ObjectMeta: objMeta,
		Kind:       kind,
	}

	resp, err := repo.datahubClient.DeleteControllers(&req)
	if err != nil {
		return errors.Wrap(err, "delete controllers from Datahub failed")
	} else if _, err := client.IsResponseStatusOK(resp); err != nil {
		return errors.Wrap(err, "delete controllers from Datahub failed")
	}
	return nil
}

func (repo *ControllerRepository) isControllerHasApplicationInfo(controller *datahub_resources.Controller, appNamespace, appName string) bool {
	// TODO: Might compare namespace if Datahub return non empty controller.AlamedaControllerSpec.AlamedaScaler.Namespace
	// if controller.AlamedaControllerSpec != nil && controller.AlamedaControllerSpec.AlamedaScaler != nil &&
	// 	controller.AlamedaControllerSpec.AlamedaScaler.Namespace == appNamespace && controller.AlamedaControllerSpec.AlamedaScaler.Name == appName {
	// 	return true
	// }

	if controller.AlamedaControllerSpec != nil && controller.AlamedaControllerSpec.AlamedaScaler != nil && controller.AlamedaControllerSpec.AlamedaScaler.Name == appName {
		return true
	}
	return false
}
