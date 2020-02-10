package influxdb

import (
	"fmt"
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/models"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	InfluxDB "github.com/influxdata/influxdb/client/v2"
	"strconv"
	"time"
)

type InfluxMeasurement struct {
	*schemas.Measurement
	Database string
	Client   *InfluxClient
}

func NewMeasurement(database string, measurement *schemas.Measurement, config Config) *InfluxMeasurement {
	m := InfluxMeasurement{
		Measurement: measurement,
		Database:    database,
		Client:      NewClient(&config),
	}
	return &m
}

func (p *InfluxMeasurement) Read(query *InfluxQuery) ([]*common.Group, error) {
	groups := make([]*common.Group, 0)
	cmd := query.BuildQueryCmd()
	
	response, err := p.Client.QueryDB(cmd, p.Database)
	if err != nil {
		scope.Errorf("failed to read from InfluxDB: %v", err)
		return make([]*common.Group, 0), err
	}

	results := models.NewInfluxResults(response)
	for _, result := range results {
		for i := 0; i < result.GetGroupNum(); i++ {
			g := result.GetGroup(i)
			group := common.Group{}
			for name:= range g.Tags {
				group.Columns = append(group.Columns, name)
			}
			for _, name := range g.Columns {
				if name != "time" {
					group.Columns = append(group.Columns, name)
				}
			}
			for j := 0; j < g.GetRowNum(); j++ {
				r := g.GetRow(j)
				row := common.Row{}
				for _, name := range group.Columns {
					row.Values = append(row.Values, r[name])
				}
				timestamp, _ := time.Parse(time.RFC3339, r["time"])
				row.Time = &timestamp
				group.Rows = append(group.Rows, &row)
			}
			groups = append(groups, &group)
		}
	}

	return groups, nil
}

func (p *InfluxMeasurement) Write(columns []string, rows []*common.Row) error {
	columnTypes := make([]schemas.ColumnType, 0)
	dataTypes := make([]common.DataType, 0)

	// Generate column & data types
	for _, name := range columns {
		for _, column := range p.Columns {
			if column.Name == name {
				columnTypes = append(columnTypes, column.ColumnType)
				dataTypes = append(dataTypes, column.DataType)
				break
			}
		}
	}

	// Generate influx data points
	points := p.buildPoints(columnTypes, dataTypes, columns, rows)

	// Batch write influx data points
	err := p.Client.WritePoints(points, InfluxDB.BatchPointsConfig{Database: p.Database})
	if err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (p *InfluxMeasurement) Drop(query *InfluxQuery) error {
	cmd := query.BuildDropCmd()
	_, err := p.Client.QueryDB(cmd, p.Database)
	if err != nil {
		scope.Errorf("failed to drop data from InfluxDB: %v", err)
		return err
	}
	return nil
}

func (p *InfluxMeasurement) buildPoints(columnTypes []schemas.ColumnType, dataTypes []common.DataType, columns []string, rows []*common.Row) []*InfluxDB.Point {
	points := make([]*InfluxDB.Point, 0)

	for _, row := range rows {
		index := 0
		tags := make(map[string]string)
		fields := make(map[string]interface{})

		// Pack influx tags & fields
		for _, value := range row.Values {
			switch columnTypes[index] {
			case schemas.Tag:
				tags[columns[index]] = value
			case schemas.Field:
				fields[columns[index]] = p.format(value, dataTypes[index])
			default:
				break
			}
			index += 1
		}

		// Add time field depends on row
		if row.Time == nil {
			pt, err := InfluxDB.NewPoint(p.Name, tags, fields, time.Unix(0, 0))
			if err == nil {
				points = append(points, pt)
			} else {
				fmt.Println(err.Error())
			}
		} else {
			pt, err := InfluxDB.NewPoint(p.Name, tags, fields, *row.Time)
			if err == nil {
				points = append(points, pt)
			} else {
				fmt.Println(err.Error())
			}
		}
	}

	return points
}

func (p *InfluxMeasurement) format(value string, dataType common.DataType) interface{} {
	var v interface{}
	var err error

	switch dataType {
	case common.Bool:
		v, err = strconv.ParseBool(value)
	case common.Int:
		v, err = strconv.ParseInt(value, 10, 32)
	case common.Int8:
		v, err = strconv.ParseInt(value, 10, 8)
	case common.Int16:
		v, err = strconv.ParseInt(value, 10, 16)
	case common.Int32:
		v, err = strconv.ParseInt(value, 10, 32)
	case common.Int64:
		v, err = strconv.ParseInt(value, 10, 64)
	case common.Uint:
		v, err = strconv.ParseUint(value, 10, 32)
	case common.Uint8:
		v, err = strconv.ParseUint(value, 10, 8)
	case common.Uint16:
		v, err = strconv.ParseUint(value, 10, 16)
	case common.Uint32:
		v, err = strconv.ParseUint(value, 10, 32)
	case common.Uint64:
		v, err = strconv.ParseUint(value, 10, 64)
	case common.Float32:
		v, err = strconv.ParseFloat(value, 32)
	case common.Float64:
		v, err = strconv.ParseFloat(value, 64)
	case common.String:
		v = value
	default:
		fmt.Println("not support")
		return value
	}

	if err != nil {
		scope.Error(fmt.Sprintf("failed to format string(%s) to type(%d)", value, dataType))
		return nil
	}

	return v
}
