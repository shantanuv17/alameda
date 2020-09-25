package metrics

import (
	"github.com/pkg/errors"
	EntityInfluxGpuMetric "prophetstor.com/alameda/datahub/pkg/dao/entities/influxdb/gpu/nvidia/metrics"
	RepoInflux "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	InfluxModels "prophetstor.com/alameda/pkg/database/influxdb/models"
)

type MemoryTotalBytesRepository struct {
	influxDB *InfluxDB.InfluxClient
}

func NewMemoryTotalBytesRepositoryWithConfig(cfg InfluxDB.Config) *MemoryTotalBytesRepository {
	return &MemoryTotalBytesRepository{
		influxDB: InfluxDB.NewClient(&cfg),
	}
}

func (r *MemoryTotalBytesRepository) ListMemoryTotalBytes(host, minorNumber string, condition *DBCommon.QueryCondition) ([]*EntityInfluxGpuMetric.MemoryTotalBytesEntity, error) {
	entities := make([]*EntityInfluxGpuMetric.MemoryTotalBytesEntity, 0)

	influxdbStatement := InfluxDB.Statement{
		QueryCondition: condition,
		Measurement:    MemoryTotalBytes,
		GroupByTags:    []string{"host", "uuid"},
	}

	influxdbStatement.AppendWhereClause("AND", "host", "=", host)
	influxdbStatement.AppendWhereClause("AND", "minor_number", "=", minorNumber)
	influxdbStatement.AppendWhereClauseFromTimeCondition()
	influxdbStatement.SetOrderClauseFromQueryCondition()
	influxdbStatement.SetLimitClauseFromQueryCondition()
	cmd := influxdbStatement.BuildQueryCmd()

	response, err := r.influxDB.QueryDB(cmd, string(RepoInflux.Gpu))
	if err != nil {
		return entities, errors.Wrap(err, "failed to list nvidia gpu memory total bytes")
	}

	results := InfluxModels.NewInfluxResults(response)
	for _, result := range results {
		for i := 0; i < result.GetGroupNum(); i++ {
			group := result.GetGroup(i)
			for j := 0; j < group.GetRowNum(); j++ {
				entity := EntityInfluxGpuMetric.NewMemoryTotalBytesEntityFromMap(group.GetRow(j))
				entities = append(entities, &entity)
			}
		}
	}

	return entities, nil
}
