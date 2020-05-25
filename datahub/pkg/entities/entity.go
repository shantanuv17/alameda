package entities

import (
	"github.com/containers-ai/alameda/pkg/utils/log"
	"reflect"
	"strconv"
)

var (
	scope = log.RegisterScope("entity", "datahub entity library", 0)
)

type Entity interface {
	Populate(entity interface{}, columns, values []string)
	Row(entity interface{}, columns []string) []string
}

type DatahubEntity struct {
}

func (p *DatahubEntity) Populate(entity interface{}, columns, values []string) {
	fields := reflect.TypeOf(entity).Elem()
	for index, column := range columns {
		for i := 1; i < fields.NumField(); i++ {
			if column == fields.Field(i).Tag.Get("json") {
				fieldValue := reflect.ValueOf(entity).Elem().Field(i)
				switch fieldValue.Kind() {
				case reflect.Bool:
					valueBool, _ := strconv.ParseBool(values[index])
					fieldValue.SetBool(valueBool)
				case reflect.Int:
					valueInt, _ := strconv.ParseInt(values[index], 10, 32)
					fieldValue.SetInt(valueInt)
				case reflect.Int8:
					valueInt, _ := strconv.ParseInt(values[index], 10, 32)
					fieldValue.SetInt(valueInt)
				case reflect.Int16:
					valueInt, _ := strconv.ParseInt(values[index], 10, 32)
					fieldValue.SetInt(valueInt)
				case reflect.Int32:
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

func (p *DatahubEntity) Row(entity interface{}, columns []string) []string {
	row := make([]string, 0)
	fields := reflect.TypeOf(entity).Elem()
	for _, column := range columns {
		for i := 1; i < fields.NumField(); i++ {
			if column == fields.Field(i).Tag.Get("json") {
				fieldValue := reflect.ValueOf(entity).Elem().Field(i)
				switch fieldValue.Kind() {
				case reflect.Bool:
					row = append(row, strconv.FormatBool(fieldValue.Bool()))
				case reflect.Int:
					row = append(row, strconv.FormatInt(fieldValue.Int(), 10))
				case reflect.Int8:
					row = append(row, strconv.FormatInt(fieldValue.Int(), 10))
				case reflect.Int16:
					row = append(row, strconv.FormatInt(fieldValue.Int(), 10))
				case reflect.Int32:
					row = append(row, strconv.FormatInt(fieldValue.Int(), 10))
				case reflect.Int64:
					row = append(row, strconv.FormatInt(fieldValue.Int(), 10))
				case reflect.Float32:
					row = append(row, strconv.FormatFloat(fieldValue.Float(), 'f', -1, 32))
				case reflect.Float64:
					row = append(row, strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64))
				case reflect.String:
					row = append(row, fieldValue.String())
				default:
					scope.Errorf("field type(%s) not supported", fieldValue.Kind().String())
				}
				break
			}
		}
	}
	return row
}
