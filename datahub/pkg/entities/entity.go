package entities

import (
	"github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"reflect"
	"strconv"
	"time"
)

var (
	scope = log.RegisterScope("entity", "datahub entity library", 0)
)

type Row struct {
	Time    *timestamp.Timestamp
	Columns []string
	Values  []string
}

type Entity interface {
	Populate(entity interface{}, timestamp *timestamp.Timestamp, columns, values []string)
	Row(entity interface{}, fields []string) *Row
}

type DatahubEntity struct {
	Time *time.Time
}

func (p *DatahubEntity) Populate(entity interface{}, timestamp *timestamp.Timestamp, columns, values []string) {
	p.Time = p.timestamp(timestamp)

	fields := reflect.TypeOf(entity).Elem()
	for index, column := range columns {
		for i := 1; i < fields.NumField(); i++ {
			if column == fields.Field(i).Tag.Get("json") {
				fieldValue := reflect.ValueOf(entity).Elem().Field(i)
				switch fieldValue.Kind() {
				case reflect.Bool:
					valueBool, _ := strconv.ParseBool(values[index])
					fieldValue.SetBool(valueBool)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
					valueInt, _ := strconv.ParseInt(values[index], 10, 32)
					fieldValue.SetInt(valueInt)
				case reflect.Int64:
					valueInt, _ := strconv.ParseInt(values[index], 10, 64)
					fieldValue.SetInt(valueInt)
				case reflect.Float32:
					valueFloat, _ := strconv.ParseFloat(values[index], 32)
					fieldValue.SetFloat(valueFloat)
				case reflect.Float64:
					valueFloat, _ := strconv.ParseFloat(values[index], 64)
					fieldValue.SetFloat(valueFloat)
				case reflect.String:
					fieldValue.SetString(values[index])
				default:
					scope.Errorf("field type(%s) not supported", fieldValue.Kind().String())
				}
				break
			}
		}
	}
}

func (p *DatahubEntity) Row(entity interface{}, fields []string) *Row {
	row := Row{Time: p.timestampProto()}

	// If fields is empty which means to iterate through all the fields of the entity
	values := reflect.TypeOf(entity).Elem()
	fieldNames := fields
	if len(fieldNames) == 0 {
		for i := 1; i < values.NumField(); i++ {
			fieldNames = append(fieldNames, values.Field(i).Name)
		}
	}

	for _, name := range fieldNames {
		for i := 1; i < values.NumField(); i++ {
			if name == values.Field(i).Name {
				fieldValue := reflect.ValueOf(entity).Elem().Field(i)
				switch fieldValue.Kind() {
				case reflect.Bool:
					row.Columns = append(row.Columns, values.Field(i).Tag.Get("json"))
					row.Values = append(row.Values, strconv.FormatBool(fieldValue.Bool()))
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					row.Columns = append(row.Columns, values.Field(i).Tag.Get("json"))
					row.Values = append(row.Values, strconv.FormatInt(fieldValue.Int(), 10))
				case reflect.Float32:
					row.Columns = append(row.Columns, values.Field(i).Tag.Get("json"))
					row.Values = append(row.Values, strconv.FormatFloat(fieldValue.Float(), 'f', -1, 32))
				case reflect.Float64:
					row.Columns = append(row.Columns, values.Field(i).Tag.Get("json"))
					row.Values = append(row.Values, strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64))
				case reflect.String:
					row.Columns = append(row.Columns, values.Field(i).Tag.Get("json"))
					row.Values = append(row.Values, fieldValue.String())
				default:
					scope.Errorf("field type(%s) not supported", fieldValue.Kind().String())
				}
			}
		}
	}

	return &row
}

func (p *DatahubEntity) timestamp(timestamp *timestamp.Timestamp) *time.Time {
	ts, _ := ptypes.Timestamp(timestamp)
	return &ts
}

func (p *DatahubEntity) timestampProto() *timestamp.Timestamp {
	ts := &timestamp.Timestamp{}
	if p.Time == nil {
		ts, _ = ptypes.TimestampProto(time.Unix(0, 0).UTC())
	} else {
		ts, _ = ptypes.TimestampProto(*p.Time)
	}
	return ts
}
