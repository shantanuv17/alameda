package clusterstatus

import (
	"fmt"
	EntityInfluxCluster "github.com/containers-ai/alameda/datahub/pkg/dao/entities/influxdb/clusterstatus"
	DaoClusterTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	RepoInflux "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb"
	Metadata "github.com/containers-ai/alameda/datahub/pkg/kubernetes/metadata"
	DBCommon "github.com/containers-ai/alameda/pkg/database/common"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	InfluxModels "github.com/containers-ai/alameda/pkg/database/influxdb/models"
	Utils "github.com/containers-ai/alameda/pkg/utils"
	ApiResources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
	"strings"
)

type PodRepository struct {
	influxDB *InfluxDB.InfluxClient
}

func NewPodRepository(influxDBCfg InfluxDB.Config) *PodRepository {
	return &PodRepository{
		influxDB: &InfluxDB.InfluxClient{
			Address:  influxDBCfg.Address,
			Username: influxDBCfg.Username,
			Password: influxDBCfg.Password,
		},
	}
}

func (p *PodRepository) IsTag(column string) bool {
	for _, tag := range EntityInfluxCluster.PodTags {
		if column == string(tag) {
			return true
		}
	}
	return false
}

func (p *PodRepository) CreatePods(pods []*DaoClusterTypes.Pod) error {
	points := make([]*InfluxClient.Point, 0)

	for _, pod := range pods {
		entity := pod.BuildEntity()

		// Add to influx point list
		if pt, err := entity.BuildInfluxPoint(string(Pod)); err == nil {
			points = append(points, pt)
		} else {
			scope.Error(err.Error())
		}
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

func (p *PodRepository) ListPods(request *DaoClusterTypes.ListPodsRequest) ([]*DaoClusterTypes.Pod, error) {
	pods := make([]*DaoClusterTypes.Pod, 0)

	statement := InfluxDB.Statement{
		QueryCondition: &request.QueryCondition,
		Measurement:    Pod,
		GroupByTags:    []string{string(EntityInfluxCluster.PodNamespace), string(EntityInfluxCluster.PodNodeName), string(EntityInfluxCluster.PodClusterName)},
	}

	// Build influx query command
	for _, objectMeta := range request.ObjectMeta {
		conditionList := make([]string, 0)

		metaCondition := p.genObjectMetaCondition(objectMeta, ApiResources.Kind(ApiResources.Kind_value[request.Kind]))
		if metaCondition != "" {
			conditionList = append(conditionList, metaCondition)
		}

		createCondition := p.genCreatePeriodCondition(request.QueryCondition)
		if createCondition != "" {
			conditionList = append(conditionList, createCondition)
		}

		if request.AlamedaScalerName != "" {
			conditionList = append(conditionList, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodAlamedaSpecScalerName, request.AlamedaScalerName))
		}

		if request.AlamedaScalerNamespace != "" {
			conditionList = append(conditionList, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodAlamedaSpecScalerNamespace, request.AlamedaScalerNamespace))
		}

		if request.TopControllerName != "" {
			conditionList = append(conditionList, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodTopControllerName, request.TopControllerName))
		}

		if request.ScalingTool != "" && request.ScalingTool != ApiResources.ScalingTool_name[0] {
			conditionList = append(conditionList, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodAlamedaSpecScalerScalingTool, request.ScalingTool))
		}

		condition := strings.Join(conditionList, " AND ")
		if condition != "" {
			condition = "(" + condition + ")"
		}
		statement.AppendWhereClauseDirectly("OR", condition)
	}
	if len(request.ObjectMeta) == 0 {
		if request.Kind != "" && request.Kind != ApiResources.Kind_name[0] {
			statement.AppendWhereClauseDirectly("AND", fmt.Sprintf(`("%s"='%s')`, EntityInfluxCluster.PodTopControllerKind, request.Kind))
		}
		if request.ScalingTool != "" && request.ScalingTool != ApiResources.ScalingTool_name[0] {
			statement.AppendWhereClauseDirectly("AND", fmt.Sprintf(`("%s"='%s')`, EntityInfluxCluster.PodAlamedaSpecScalerScalingTool, request.ScalingTool))
		}
		statement.AppendWhereClauseDirectly("AND", p.genCreatePeriodCondition(request.QueryCondition))
	}
	statement.SetOrderClauseFromQueryCondition()
	statement.SetLimitClauseFromQueryCondition()
	cmd := statement.BuildQueryCmd()

	response, err := p.influxDB.QueryDB(cmd, string(RepoInflux.ClusterStatus))
	if err != nil {
		return make([]*DaoClusterTypes.Pod, 0), errors.Wrap(err, "failed to list pods")
	}

	results := InfluxModels.NewInfluxResults(response)
	for _, result := range results {
		for i := 0; i < result.GetGroupNum(); i++ {
			group := result.GetGroup(i)
			for j := 0; j < group.GetRowNum(); j++ {
				row := group.GetRow(j)
				pod := DaoClusterTypes.NewPod(EntityInfluxCluster.NewPodEntity(row))
				pods = append(pods, pod)
			}
		}
	}

	return pods, nil
}

func (p *PodRepository) DeletePods(request *DaoClusterTypes.DeletePodsRequest) error {
	statement := InfluxDB.Statement{
		Measurement: Pod,
	}

	if !p.influxDB.MeasurementExist(string(RepoInflux.ClusterStatus), string(Pod)) {
		return nil
	}

	// Build influx drop command
	for _, podObjectMeta := range request.PodObjectMeta {
		keyList := make([]string, 0)
		valueList := make([]string, 0)

		if podObjectMeta.ObjectMeta != nil {
			keyList = podObjectMeta.ObjectMeta.GenerateKeyList()
			valueList = podObjectMeta.ObjectMeta.GenerateValueList()
		}

		if podObjectMeta.TopController != nil {
			keyList = append(keyList, string(EntityInfluxCluster.PodTopControllerName))
			valueList = append(valueList, podObjectMeta.TopController.Name)

			if !Utils.SliceContains(keyList, string(EntityInfluxCluster.PodNamespace)) {
				keyList = append(keyList, string(EntityInfluxCluster.PodNamespace))
				valueList = append(valueList, podObjectMeta.TopController.Namespace)
			}

			if !Utils.SliceContains(keyList, string(EntityInfluxCluster.PodClusterName)) {
				keyList = append(keyList, string(EntityInfluxCluster.PodClusterName))
				valueList = append(valueList, podObjectMeta.TopController.ClusterName)
			}
		}

		if podObjectMeta.AlamedaScaler != nil {
			if podObjectMeta.AlamedaScaler.Name != "" {
				keyList = append(keyList, string(EntityInfluxCluster.PodAlamedaSpecScalerName))
				valueList = append(valueList, podObjectMeta.AlamedaScaler.Name)
			}

			if podObjectMeta.AlamedaScaler.Namespace != "" {
				keyList = append(keyList, string(EntityInfluxCluster.PodAlamedaSpecScalerNamespace))
				valueList = append(valueList, podObjectMeta.AlamedaScaler.Namespace)
			}

			if !Utils.SliceContains(keyList, string(EntityInfluxCluster.PodClusterName)) {
				keyList = append(keyList, string(EntityInfluxCluster.PodClusterName))
				valueList = append(valueList, podObjectMeta.AlamedaScaler.ClusterName)
			}
		}

		if podObjectMeta.Kind != "" && podObjectMeta.Kind != ApiResources.Kind_name[0] {
			keyList = append(keyList, string(EntityInfluxCluster.PodTopControllerKind))
			valueList = append(valueList, podObjectMeta.Kind)
		}

		if podObjectMeta.ScalingTool != "" && podObjectMeta.ScalingTool != ApiResources.ScalingTool_name[0] {
			keyList = append(keyList, string(EntityInfluxCluster.PodAlamedaSpecScalerScalingTool))
			valueList = append(valueList, podObjectMeta.ScalingTool)
		}

		condition := statement.GenerateCondition(keyList, valueList, "AND")
		statement.AppendWhereClauseDirectly("OR", condition)
	}
	cmd := statement.BuildDropCmd()

	_, err := p.influxDB.QueryDB(cmd, string(RepoInflux.ClusterStatus))
	if err != nil {
		return errors.Wrap(err, "failed to delete pods")
	}

	return nil
}

func (p *PodRepository) genObjectMetaCondition(objectMeta *Metadata.ObjectMeta, kind ApiResources.Kind) string {
	conditions := make([]string, 0)

	switch kind {
	case ApiResources.Kind_KIND_UNDEFINED:
		if objectMeta.Namespace != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodNamespace, objectMeta.Namespace))
		}
		if objectMeta.Name != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodName, objectMeta.Name))
		}
		if objectMeta.ClusterName != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodClusterName, objectMeta.ClusterName))
		}
	case ApiResources.Kind_DEPLOYMENT:
		conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodTopControllerKind, ApiResources.Kind_name[int32(kind)]))
		if objectMeta.Namespace != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodNamespace, objectMeta.Namespace))
		}
		if objectMeta.Name != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodTopControllerName, objectMeta.Name))
		}
		if objectMeta.ClusterName != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodClusterName, objectMeta.ClusterName))
		}
	case ApiResources.Kind_DEPLOYMENTCONFIG:
		conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodTopControllerKind, ApiResources.Kind_name[int32(kind)]))
		if objectMeta.Namespace != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodNamespace, objectMeta.Namespace))
		}
		if objectMeta.Name != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodTopControllerName, objectMeta.Name))
		}
		if objectMeta.ClusterName != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodClusterName, objectMeta.ClusterName))
		}
	case ApiResources.Kind_ALAMEDASCALER:
		if objectMeta.Namespace != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodNamespace, objectMeta.Namespace))
		}
		if objectMeta.Name != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodAlamedaSpecScalerName, objectMeta.Name))
		}
		if objectMeta.ClusterName != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodClusterName, objectMeta.ClusterName))
		}
	case ApiResources.Kind_STATEFULSET:
		conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodTopControllerKind, ApiResources.Kind_name[int32(kind)]))
		if objectMeta.Namespace != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodNamespace, objectMeta.Namespace))
		}
		if objectMeta.Name != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodTopControllerName, objectMeta.Name))
		}
		if objectMeta.ClusterName != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodClusterName, objectMeta.ClusterName))
		}
	default:
		if objectMeta.Namespace != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodNamespace, objectMeta.Namespace))
		}
		if objectMeta.Name != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodName, objectMeta.Name))
		}
		if objectMeta.ClusterName != "" {
			conditions = append(conditions, fmt.Sprintf(`"%s"='%s'`, EntityInfluxCluster.PodClusterName, objectMeta.ClusterName))
		}
	}

	if len(conditions) > 0 {
		return strings.Join(conditions, " AND ")
	}

	return ""
}

func (p *PodRepository) genCreatePeriodCondition(query DBCommon.QueryCondition) string {
	if query.StartTime != nil && query.EndTime != nil {
		return fmt.Sprintf("\"%s\">=%d AND \"%s\"<%d", EntityInfluxCluster.PodCreateTime, query.StartTime.Unix(), EntityInfluxCluster.PodCreateTime, query.EndTime.Unix())
	} else if query.StartTime != nil && query.EndTime == nil {
		return fmt.Sprintf("\"%s\">=%d", EntityInfluxCluster.PodCreateTime, query.StartTime.Unix())
	} else if query.StartTime == nil && query.EndTime != nil {
		return fmt.Sprintf("\"%s\"<%d", EntityInfluxCluster.PodCreateTime, query.EndTime.Unix())
	} else {
		return ""
	}
}
