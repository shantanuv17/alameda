package clusterstatus

import (
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
	EntityInfluxCluster "prophetstor.com/alameda/datahub/pkg/dao/entities/influxdb/clusterstatus"
	DaoClusterTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	RepoInflux "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	InfluxModels "prophetstor.com/alameda/pkg/database/influxdb/models"
	Utils "prophetstor.com/alameda/pkg/utils"
	ApiResources "prophetstor.com/api/datahub/resources"
)

type ControllerRepository struct {
	influxDB *InfluxDB.InfluxClient
}

func NewControllerRepository(influxDBCfg InfluxDB.Config) *ControllerRepository {
	return &ControllerRepository{
		influxDB: &InfluxDB.InfluxClient{
			Address:  influxDBCfg.Address,
			Username: influxDBCfg.Username,
			Password: influxDBCfg.Password,
		},
	}
}

func (p *ControllerRepository) CreateControllers(controllers []*DaoClusterTypes.Controller) error {
	points := make([]*InfluxClient.Point, 0)

	for _, controller := range controllers {
		entity := controller.BuildEntity()

		// Add to influx point list
		pt, err := entity.BuildInfluxPoint(string(Controller))
		if err != nil {
			scope.Error(err.Error())
			return errors.Wrap(err, "failed to instance influxdb data point")
		}
		points = append(points, pt)
	}

	// Batch write influxdb data points
	err := p.influxDB.WritePoints(points, InfluxClient.BatchPointsConfig{
		Database: string(RepoInflux.ClusterStatus),
	})
	if err != nil {
		scope.Error(err.Error())
		return errors.Wrap(err, "failed to batch write influxdb data points")
	}

	return nil
}

func (p *ControllerRepository) ListControllers(request *DaoClusterTypes.ListControllersRequest) ([]*DaoClusterTypes.Controller, error) {
	controllers := make([]*DaoClusterTypes.Controller, 0)

	statement := InfluxDB.Statement{
		QueryCondition: &request.QueryCondition,
		Measurement:    Controller,
		GroupByTags:    []string{string(EntityInfluxCluster.ControllerNamespace), string(EntityInfluxCluster.ControllerClusterName)},
	}

	// Build influx query command
	for _, controllerObjectMeta := range request.ControllerObjectMeta {
		keyList := make([]string, 0)
		valueList := make([]string, 0)

		if controllerObjectMeta.ObjectMeta != nil {
			keyList = controllerObjectMeta.ObjectMeta.GenerateKeyList()
			valueList = controllerObjectMeta.ObjectMeta.GenerateValueList()
		}

		if controllerObjectMeta.AlamedaScaler != nil {
			if controllerObjectMeta.AlamedaScaler.Name != "" {
				keyList = append(keyList, string(EntityInfluxCluster.ControllerAlamedaSpecScalerName))
				valueList = append(valueList, controllerObjectMeta.AlamedaScaler.Name)
			}
			if controllerObjectMeta.AlamedaScaler.Namespace != "" {
				keyList = append(keyList, string(EntityInfluxCluster.ControllerAlamedaSpecScalerNamespace))
				valueList = append(valueList, controllerObjectMeta.AlamedaScaler.Namespace)
			}
			if controllerObjectMeta.AlamedaScaler.ClusterName != "" {
				keyList = append(keyList, string(EntityInfluxCluster.ControllerClusterName))
				valueList = append(valueList, controllerObjectMeta.AlamedaScaler.ClusterName)
			}
		}

		if controllerObjectMeta.Kind != "" && controllerObjectMeta.Kind != ApiResources.Kind_name[0] {
			keyList = append(keyList, string(EntityInfluxCluster.ControllerKind))
			valueList = append(valueList, controllerObjectMeta.Kind)
		}

		if controllerObjectMeta.ScalingTool != "" && controllerObjectMeta.ScalingTool != ApiResources.ScalingTool_name[0] {
			keyList = append(keyList, string(EntityInfluxCluster.ControllerAlamedaSpecScalerScalingTool))
			valueList = append(valueList, controllerObjectMeta.ScalingTool)
		}

		condition := statement.GenerateCondition(keyList, valueList, "AND")
		statement.AppendWhereClauseDirectly("OR", condition)
	}
	statement.SetOrderClauseFromQueryCondition()
	statement.SetLimitClauseFromQueryCondition()
	cmd := statement.BuildQueryCmd()

	response, err := p.influxDB.QueryDB(cmd, string(RepoInflux.ClusterStatus))
	if err != nil {
		return make([]*DaoClusterTypes.Controller, 0), errors.Wrap(err, "failed to list controllers")
	}

	results := InfluxModels.NewInfluxResults(response)
	for _, result := range results {
		for i := 0; i < result.GetGroupNum(); i++ {
			group := result.GetGroup(i)
			for j := 0; j < group.GetRowNum(); j++ {
				row := group.GetRow(j)
				controller := DaoClusterTypes.NewController(EntityInfluxCluster.NewControllerEntity(row))
				controllers = append(controllers, controller)
			}
		}
	}

	return controllers, nil
}

func (p *ControllerRepository) DeleteControllers(request *DaoClusterTypes.DeleteControllersRequest) error {
	statement := InfluxDB.Statement{
		Measurement: Controller,
	}

	if !p.influxDB.MeasurementExist(string(RepoInflux.ClusterStatus), string(Controller)) {
		return nil
	}

	// Build influx drop command
	for _, controllerObjectMeta := range request.ControllerObjectMeta {
		keyList := make([]string, 0)
		valueList := make([]string, 0)

		if controllerObjectMeta.ObjectMeta != nil {
			keyList = controllerObjectMeta.ObjectMeta.GenerateKeyList()
			valueList = controllerObjectMeta.ObjectMeta.GenerateValueList()
		}

		if controllerObjectMeta.AlamedaScaler != nil {
			if controllerObjectMeta.AlamedaScaler.Name != "" {
				keyList = append(keyList, string(EntityInfluxCluster.ControllerAlamedaSpecScalerName))
				valueList = append(valueList, controllerObjectMeta.AlamedaScaler.Name)
			}

			if controllerObjectMeta.AlamedaScaler.Namespace != "" {
				keyList = append(keyList, string(EntityInfluxCluster.ControllerAlamedaSpecScalerNamespace))
				valueList = append(valueList, controllerObjectMeta.AlamedaScaler.Namespace)
			}

			if !Utils.SliceContains(keyList, string(EntityInfluxCluster.ControllerClusterName)) {
				if controllerObjectMeta.AlamedaScaler.ClusterName != "" {
					keyList = append(keyList, string(EntityInfluxCluster.ControllerClusterName))
					valueList = append(valueList, controllerObjectMeta.AlamedaScaler.ClusterName)
				}
			}
		}

		if controllerObjectMeta.Kind != "" && controllerObjectMeta.Kind != ApiResources.Kind_name[0] {
			keyList = append(keyList, string(EntityInfluxCluster.ControllerKind))
			valueList = append(valueList, controllerObjectMeta.Kind)
		}

		if controllerObjectMeta.ScalingTool != "" && controllerObjectMeta.ScalingTool != ApiResources.ScalingTool_name[0] {
			keyList = append(keyList, string(EntityInfluxCluster.ControllerAlamedaSpecScalerScalingTool))
			valueList = append(valueList, controllerObjectMeta.ScalingTool)
		}

		condition := statement.GenerateCondition(keyList, valueList, "AND")
		statement.AppendWhereClauseDirectly("OR", condition)
	}
	cmd := statement.BuildDropCmd()

	_, err := p.influxDB.QueryDB(cmd, string(RepoInflux.ClusterStatus))
	if err != nil {
		scope.Error(err.Error())
		return errors.Wrap(err, "failed to delete controllers")
	}

	return nil
}
