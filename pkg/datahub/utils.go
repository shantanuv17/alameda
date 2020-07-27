package datahub

import (
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	"reflect"
)

func ExtractField(entities interface{}, fieldName string) reflect.StructField {
	var structField reflect.StructField

	fieldType := reflect.TypeOf(entities).Elem()
	switch fieldType.Kind() {
	case reflect.Struct:
		structField, _ = fieldType.FieldByName(fieldName)
	case reflect.Slice:
		structField, _ = fieldType.Elem().FieldByName(fieldName)
	default:
		scope.Errorf("data type(%d) of entity is not supported", fieldType.Kind())
	}

	return structField
}

func DeepCopyEntity(entities interface{}, results *data.Data) {
	entityType := reflect.TypeOf(entities).Elem().Elem()

	entityPtr := reflect.ValueOf(entities)
	entityValue := entityPtr.Elem()

	for _, rawdata := range results.Rawdata {
		for _, group := range rawdata.Groups {
			for _, row := range group.Rows {
				obj := reflect.New(entityType)
				inputs := make([]reflect.Value, 0)
				inputs = append(inputs, obj)
				inputs = append(inputs, reflect.ValueOf(row.Time))
				inputs = append(inputs, reflect.ValueOf(group.Columns))
				inputs = append(inputs, reflect.ValueOf(row.Values))
				obj.MethodByName("Populate").Call(inputs)

				entityValue.Set(reflect.Append(entityValue, reflect.Indirect(obj)))
			}
		}
	}
}
