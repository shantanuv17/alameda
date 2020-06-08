package execution

import (
	"context"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	"github.com/containers-ai/alameda/pkg/utils/datahub/consts"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func ListMachineSetExecutionByNamespacedName(
	datahubClient datahub.DatahubServiceClient, clusterName, machineNS,
	machineName string, limit uint64) ([]*entities.ExecutionClusterAutoscalerMachineset, error) {
	ctx, cancel := context.WithTimeout(context.Background(), consts.DatahubTimeout)
	defer cancel()

	machineSetExecutionList := []*entities.ExecutionClusterAutoscalerMachineset{}
	whereCondition := &common.Condition{
		Keys:      []string{"cluster_name", "namespace", "name"},
		Values:    []string{clusterName, machineNS, machineName},
		Operators: []string{"=", "=", "="},
		Types: []common.DataType{
			common.DataType_DATATYPE_STRING,
			common.DataType_DATATYPE_STRING,
			common.DataType_DATATYPE_STRING,
		},
	}
	queryCondition := &common.QueryCondition{
		WhereCondition: []*common.Condition{
			whereCondition,
		},
		Order: common.QueryCondition_DESC,
	}
	if limit > 0 {
		queryCondition.Limit = limit
	}
	req := &data.ReadDataRequest{
		SchemaMeta: &schemas.SchemaMeta{
			Scope:    schemas.Scope_SCOPE_EXECUTION,
			Category: consts.CategoryClusterAutoscaler,
			Type:     consts.TypeMachineSet,
		},
		ReadData: []*data.ReadData{
			{
				Measurement:    consts.MeasurementMachineset,
				QueryCondition: queryCondition,
			},
		},
	}
	resp, err := datahubClient.ReadData(ctx, req)
	for _, rd := range resp.GetData().GetRawdata() {
		for _, grp := range rd.GetGroups() {
			cols := grp.GetColumns()
			for _, row := range grp.GetRows() {
				entity := &entities.ExecutionClusterAutoscalerMachineset{}
				entity.Populate(entity, row.GetTime(), cols, row.GetValues())
				machineSetExecutionList = append(machineSetExecutionList, entity)
			}
		}
	}
	return machineSetExecutionList, err
}

func CreateMachineSetExecution(datahubClient datahub.DatahubServiceClient,
	execution *entities.ExecutionClusterAutoscalerMachineset, fields []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), consts.DatahubTimeout)
	defer cancel()

	entityRow := execution.Row(execution, fields)

	req := &data.WriteDataRequest{
		SchemaMeta: &schemas.SchemaMeta{
			Scope:    schemas.Scope_SCOPE_EXECUTION,
			Category: consts.CategoryClusterAutoscaler,
			Type:     consts.TypeMachineSet,
		},
		WriteData: []*data.WriteData{
			{
				Measurement: consts.MeasurementMachineset,
				Columns:     entityRow.Columns,
				Rows: []*common.Row{
					{
						Time:   entityRow.Time,
						Values: entityRow.Values,
					},
				},
			},
		},
	}
	_, err := datahubClient.WriteData(ctx, req)
	return err
}
