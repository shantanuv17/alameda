package kafka

import (
	"context"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/code"

	// autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	// "github.com/containers-ai/alameda/operator/datahub/client"
	"github.com/containers-ai/alameda/operator/datahub/client/kafka/entity"
	"github.com/containers-ai/alameda/operator/pkg/kafka"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

const (
	tagDatahubColumn = "datahubcolumn"
)

type KafkaRepository struct {
	datahubClient datahub.DatahubServiceClient
	schemaConfig  config
	logger        *log.Scope
}

func NewKafkaRepository(datahubClient datahub.DatahubServiceClient, logger *log.Scope) KafkaRepository {
	if logger == nil {
		logger = log.RegisterScope("datahub-client", "", 0)
	}
	return KafkaRepository{
		datahubClient: datahubClient,
		// TODO:
		schemaConfig: defaultConfig,
		logger:       logger,
	}
}

func (k KafkaRepository) CreateTopics(ctx context.Context, topics []kafka.Topic) error {
	req, err := k.newWriteDataRequestByTopics(topics)
	if err != nil {
		return errors.Wrap(err, "new WriteDataRequest by topics failed")
	}
	if err := k.sendWriteDataRequest(ctx, req); err != nil {
		return err
	}
	return nil
}

func (k KafkaRepository) CreateConsumerGroups(ctx context.Context, consumerGroups []kafka.ConsumerGroup) error {
	req, err := k.newWriteDataRequesByConsumerGroups(consumerGroups)
	if err != nil {
		return errors.Wrap(err, "new WriteDataRequest by consumerGroups failed")
	}
	if err := k.sendWriteDataRequest(ctx, req); err != nil {
		return err
	}
	return nil
}

func (k KafkaRepository) ListTopics(ctx context.Context, option ListTopicsOption) ([]kafka.Topic, error) {
	req, err := k.newReadDataRequestForTopics(option)
	if err != nil {
		return nil, errors.Wrap(err, "new ReadDataRequest failed")
	}
	data, err := k.sendReadDataRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	entities := make([]entity.KafkaTopic, 0)
	if err := decodeSlice(data, entities); err != nil {
		return nil, errors.Wrap(err, "decode data failed")
	}
	topics := make([]kafka.Topic, 0, len(entities))
	for _, e := range entities {
		topics = append(topics, kafka.Topic{
			Name:              e.Name,
			ExporterNamespace: e.ExporterNamespace,
			ClusterName:       e.ClusterName,
			AlamedaScalerName: e.AlamedaScalerName,
		})
	}
	return topics, nil
}

func (k KafkaRepository) ListConsumerGroups(ctx context.Context, option ListConsumerGroupsOption) ([]kafka.ConsumerGroup, error) {
	req, err := k.newReadDataRequestForConsumerGroups(option)
	if err != nil {
		return nil, errors.Wrap(err, "new ReadDataRequest failed")
	}
	data, err := k.sendReadDataRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	entities := make([]entity.KafkaConsumerGroup, 0)
	if err := decodeSlice(data, entities); err != nil {
		return nil, errors.Wrap(err, "decode data failed")
	}
	consumerGroups := make([]kafka.ConsumerGroup, 0, len(entities))
	for _, e := range entities {
		consumerGroups = append(consumerGroups, kafka.ConsumerGroup{
			Name:              e.Name,
			ExporterNamespace: e.ExporterNamespace,
			ClusterName:       e.ClusterName,
			AlamedaScalerName: e.AlamedaScalerName,
			Policy:            e.Policy,
			EnableExecution:   e.EnableExecution,
			ConsumeTopic:      e.ConsumeTopic,
			ResourceMeta: kafka.ResourceMeta{
				CustomName: e.ResourceCustomName,
				KubernetesMeta: kafka.KubernetesMeta{
					Namespace:     e.ResourceK8SNamespace,
					Name:          e.ResourceK8SName,
					Kind:          e.ResourceK8SKind,
					ReadyReplicas: e.ReadyReplicas,
					SpecReplicas:  e.SpecReplicas,
				},
			},
		})
	}
	return consumerGroups, nil
}

func (k KafkaRepository) DeleteTopics(ctx context.Context, topics []kafka.Topic) error {
	req, err := k.newDeleteDataRequestByTopics(topics)
	if err != nil {
		return errors.Wrap(err, "new DeleteDataRequeset by topics failed")
	}
	if err := k.sendDeleteDataRequest(ctx, req); err != nil {
		return err
	}
	return nil
}

func (k KafkaRepository) DeleteConsumerGroups(ctx context.Context, consumerGroups []kafka.ConsumerGroup) error {
	req, err := k.newDeleteDataRequestByConsumerGroups(consumerGroups)
	if err != nil {
		return errors.Wrap(err, "new DeleteDataRequeset by consumergroups failed")
	}
	if err := k.sendDeleteDataRequest(ctx, req); err != nil {
		return err
	}
	return nil
}

func (k KafkaRepository) sendWriteDataRequest(ctx context.Context, req data.WriteDataRequest) error {
	k.logger.Infof("Write data to Datahub. Request: %+v", req)
	status, err := k.datahubClient.WriteData(ctx, &req)
	if err != nil {
		return errors.Wrap(err, "send WriteDataRequest failed")
	} else if status == nil {
		return errors.New("receive nil status")
	} else if status.Code != int32(code.Code_OK) {
		return errors.Errorf("send WriteDataRequest failed: statuscode: %d, message: %s", status.Code, status.Message)
	}
	return nil
}

func (k KafkaRepository) sendReadDataRequest(ctx context.Context, req data.ReadDataRequest) (data.Data, error) {
	k.logger.Infof("Read data from Datahub. Request: %+v", req)
	resp, err := k.datahubClient.ReadData(ctx, &req)
	if err != nil {
		return data.Data{}, errors.Wrap(err, "send ReadDataRequest failed")
	} else if resp == nil {
		return data.Data{}, errors.New("receive nil response")
	} else if resp.Status == nil {
		return data.Data{}, errors.New("receive nil status")
	} else if resp.Status.Code != int32(code.Code_OK) {
		return data.Data{}, errors.Errorf("send ReadDataRequest failed: statuscode: %d, message: %s", resp.Status.Code, resp.Status.Message)
	} else if resp.Data == nil {
		return data.Data{}, errors.New("receive nil responce data")
	}
	return *resp.Data, nil
}

func (k KafkaRepository) sendDeleteDataRequest(ctx context.Context, req data.DeleteDataRequest) error {
	k.logger.Infof("Delete data from Datahub. Request: %+v", req)
	status, err := k.datahubClient.DeleteData(ctx, &req)
	if err != nil {
		return errors.Wrap(err, "send DeleteDataRequest failed")
	} else if status == nil {
		return errors.New("receive nil status")
	} else if status.Code != int32(code.Code_OK) {
		return errors.Errorf("send DeleteDataRequest failed: statuscode: %d, message: %s", status.Code, status.Message)
	}
	return nil
}

func (k KafkaRepository) newWriteDataRequestByTopics(topics []kafka.Topic) (data.WriteDataRequest, error) {
	entities := make([]entity.KafkaTopic, 0, len(topics))
	for _, topic := range topics {
		entity := entity.NewKafkaTopic(topic)
		entities = append(entities, entity)
	}

	writeData, err := newWriteData(k.schemaConfig.kafka.topic.create.measurements[0], entities)
	if err != nil {
		return data.WriteDataRequest{}, errors.Wrap(err, "new WriteData failed")
	}
	req := data.WriteDataRequest{
		SchemaMeta: &schemas.SchemaMeta{
			Scope:    schemas.Scope(schemas.Scope_value[k.schemaConfig.kafka.topic.create.schemaMeta.scope]),
			Category: k.schemaConfig.kafka.topic.create.schemaMeta.category,
			Type:     k.schemaConfig.kafka.topic.create.schemaMeta.type_,
		},
		WriteData: []*data.WriteData{&writeData},
	}
	return req, nil
}

func (k KafkaRepository) newWriteDataRequesByConsumerGroups(consumerGroups []kafka.ConsumerGroup) (data.WriteDataRequest, error) {
	entities := make([]entity.KafkaConsumerGroup, 0, len(consumerGroups))
	for _, consumerGroup := range consumerGroups {
		entity := entity.NewKafkaConsumerGroup(consumerGroup)
		entities = append(entities, entity)
	}

	writeData, err := newWriteData(k.schemaConfig.kafka.consumerGroup.create.measurements[0], entities)
	if err != nil {
		return data.WriteDataRequest{}, errors.Wrap(err, "new WriteData failed")
	}
	req := data.WriteDataRequest{
		SchemaMeta: &schemas.SchemaMeta{
			Scope:    schemas.Scope(schemas.Scope_value[k.schemaConfig.kafka.consumerGroup.create.schemaMeta.scope]),
			Category: k.schemaConfig.kafka.consumerGroup.create.schemaMeta.category,
			Type:     k.schemaConfig.kafka.consumerGroup.create.schemaMeta.type_,
		},
		WriteData: []*data.WriteData{&writeData},
	}
	return req, nil
}

func (k KafkaRepository) newDeleteDataRequestByTopics(topics []kafka.Topic) (data.DeleteDataRequest, error) {
	entities := make([]entity.KafkaTopic, 0, len(topics))
	for _, topic := range topics {
		entity := entity.NewKafkaTopic(topic)
		entities = append(entities, entity)
	}

	deleteData, err := newDeleteData(k.schemaConfig.kafka.topic.delete.measurements[0], topics)
	if err != nil {
		return data.DeleteDataRequest{}, errors.Wrap(err, "new DeleteData failed")
	}
	req := data.DeleteDataRequest{
		SchemaMeta: &schemas.SchemaMeta{
			Scope:    schemas.Scope(schemas.Scope_value[k.schemaConfig.kafka.topic.delete.schemaMeta.scope]),
			Category: k.schemaConfig.kafka.topic.delete.schemaMeta.category,
			Type:     k.schemaConfig.kafka.topic.delete.schemaMeta.type_,
		},
		DeleteData: []*data.DeleteData{&deleteData},
	}
	return req, nil
}

func (k KafkaRepository) newDeleteDataRequestByConsumerGroups(consumerGroups []kafka.ConsumerGroup) (data.DeleteDataRequest, error) {
	entities := make([]entity.KafkaConsumerGroup, 0, len(consumerGroups))
	for _, consumerGroup := range consumerGroups {
		entity := entity.NewKafkaConsumerGroup(consumerGroup)
		entities = append(entities, entity)
	}
	deleteData, err := newDeleteData(k.schemaConfig.kafka.consumerGroup.delete.measurements[0], entities)
	if err != nil {
		return data.DeleteDataRequest{}, errors.Wrap(err, "new DeleteData failed")
	}
	req := data.DeleteDataRequest{
		SchemaMeta: &schemas.SchemaMeta{
			Scope:    schemas.Scope(schemas.Scope_value[k.schemaConfig.kafka.consumerGroup.delete.schemaMeta.scope]),
			Category: k.schemaConfig.kafka.consumerGroup.delete.schemaMeta.category,
			Type:     k.schemaConfig.kafka.consumerGroup.delete.schemaMeta.type_,
		},
		DeleteData: []*data.DeleteData{&deleteData},
	}
	return req, nil
}

func (k KafkaRepository) newReadDataRequestForTopics(option ListTopicsOption) (data.ReadDataRequest, error) {
	condition, err := newCondition(option)
	if err != nil {
		return data.ReadDataRequest{}, errors.Wrap(err, "new condition failed")
	}
	readData := make([]*data.ReadData, 0, len(k.schemaConfig.kafka.topic.list.measurements))
	for _, measurement := range k.schemaConfig.kafka.topic.list.measurements {
		readData = append(readData, &data.ReadData{
			Measurement: measurement.name,
			QueryCondition: &common.QueryCondition{
				WhereCondition: []*common.Condition{
					&condition,
				},
			},
		})
	}

	req := data.ReadDataRequest{
		SchemaMeta: &schemas.SchemaMeta{
			Scope:    schemas.Scope(schemas.Scope_value[k.schemaConfig.kafka.topic.list.schemaMeta.scope]),
			Category: k.schemaConfig.kafka.topic.list.schemaMeta.category,
			Type:     k.schemaConfig.kafka.topic.list.schemaMeta.type_,
		},
		ReadData: readData,
	}
	return req, nil
}

func (k KafkaRepository) newReadDataRequestForConsumerGroups(option ListConsumerGroupsOption) (data.ReadDataRequest, error) {
	condition, err := newCondition(option)
	if err != nil {
		return data.ReadDataRequest{}, errors.Wrap(err, "new condition failed")
	}
	readData := make([]*data.ReadData, 0, len(k.schemaConfig.kafka.consumerGroup.list.measurements))
	for _, measurement := range k.schemaConfig.kafka.consumerGroup.list.measurements {
		readData = append(readData, &data.ReadData{
			Measurement: measurement.name,
			QueryCondition: &common.QueryCondition{
				WhereCondition: []*common.Condition{&condition},
			},
		})
	}

	req := data.ReadDataRequest{
		SchemaMeta: &schemas.SchemaMeta{
			Scope:    schemas.Scope(schemas.Scope_value[k.schemaConfig.kafka.consumerGroup.list.schemaMeta.scope]),
			Category: k.schemaConfig.kafka.consumerGroup.list.schemaMeta.category,
			Type:     k.schemaConfig.kafka.consumerGroup.list.schemaMeta.type_,
		},
		ReadData: readData,
	}
	return req, nil
}

func newWriteData(measurement measurement, dataRows interface{}) (data.WriteData, error) {
	switch reflect.ValueOf(dataRows).Kind() {
	case reflect.Slice:
	// TODO:
	// case reflect.Struct:
	default:
		return data.WriteData{}, errors.Errorf("not supported value(%s)", reflect.TypeOf(dataRows).Kind())
	}

	rv := reflect.ValueOf(dataRows)
	rows := make([]*common.Row, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		v := rv.Index(i)
		m := datahubColumnToFieldMap(v.Interface())

		values := make([]string, 0, len(measurement.columns))
		for _, column := range measurement.columns {
			fieldValue, exist := m[column]
			if !exist {
				return data.WriteData{}, errors.Errorf(`field for column("%s") not found`, column)
			}

			switch fieldValue.Kind() {
			case reflect.Int:
				values = append(values, strconv.FormatInt(fieldValue.Int(), 10))
			case reflect.Int8:
				values = append(values, strconv.FormatInt(fieldValue.Int(), 10))
			case reflect.Int16:
				values = append(values, strconv.FormatInt(fieldValue.Int(), 10))
			case reflect.Int32:
				values = append(values, strconv.FormatInt(fieldValue.Int(), 10))
			case reflect.Int64:
				values = append(values, strconv.FormatInt(fieldValue.Int(), 10))
			case reflect.Float32:
				values = append(values, strconv.FormatFloat(fieldValue.Float(), 'f', -1, 32))
			case reflect.Float64:
				values = append(values, strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64))
			case reflect.String:
				values = append(values, fieldValue.String())
			case reflect.Bool:
				values = append(values, strconv.FormatBool(fieldValue.Bool()))
			default:
				return data.WriteData{}, errors.Errorf("field type(%s) not supported", fieldValue.Kind().String())
			}
		}
		rows = append(rows, &common.Row{
			Values: values,
		})
	}

	w := data.WriteData{
		Measurement: measurement.name,
		Columns:     measurement.columns,
		Rows:        rows,
	}
	return w, nil
}

// newDeleteData returns DeleteData containing with measuremnt and whereConditions
// Currently type of argument "dataRows" must be slice of struct, and each type of the field must in the second swith cases.
func newDeleteData(measurement measurement, dataRows interface{}) (data.DeleteData, error) {
	switch reflect.ValueOf(dataRows).Kind() {
	case reflect.Slice:
	// TODO:
	// case reflect.Struct:
	default:
		return data.DeleteData{}, errors.Errorf("not supported value(%s)", reflect.TypeOf(dataRows).Kind())
	}

	whereConditions := make([]*common.Condition, 0)
	for i := 0; i < reflect.ValueOf(dataRows).Len(); i++ {
		rV := reflect.ValueOf(dataRows).Index(i)
		rT := rV.Type()

		cond := common.Condition{
			Keys:      []string{},
			Values:    []string{},
			Operators: []string{},
		}
		for j := 0; j < rV.NumField(); j++ {
			f := rV.Field(j)
			value := ""
			switch f.Kind() {
			case reflect.Int:
				value = strconv.FormatInt(f.Int(), 10)
			case reflect.Int8:
				value = strconv.FormatInt(f.Int(), 10)
			case reflect.Int16:
				value = strconv.FormatInt(f.Int(), 10)
			case reflect.Int32:
				value = strconv.FormatInt(f.Int(), 10)
			case reflect.Int64:
				value = strconv.FormatInt(f.Int(), 10)
			case reflect.Float32:
				value = strconv.FormatFloat(f.Float(), 'f', -1, 10)
			case reflect.Float64:
				value = strconv.FormatFloat(f.Float(), 'f', -1, 10)
			case reflect.String:
				value = f.String()
			case reflect.Bool:
				value = strconv.FormatBool(f.Bool())
			default:
				return data.DeleteData{}, errors.Errorf(`not supported "%s"`, f.Kind().String())
			}
			if value == "" {
				continue
			}

			datahubColumn, exist := rT.Field(j).Tag.Lookup(tagDatahubColumn)
			if !exist || datahubColumn == "" {
				continue
			}
			cond.Keys = append(cond.Keys, datahubColumn)
			cond.Values = append(cond.Values, value)
			cond.Operators = append(cond.Operators, "=")
		}
		whereConditions = append(whereConditions, &cond)
	}

	w := data.DeleteData{
		Measurement: measurement.name,
		QueryCondition: &common.QueryCondition{
			WhereCondition: whereConditions,
		},
	}
	return w, nil
}

func datahubColumnToFieldMap(s interface{}) map[string]reflect.Value {
	m := make(map[string]reflect.Value)

	rT := reflect.TypeOf(s)
	rV := reflect.ValueOf(s)
	for i := 0; i < rT.NumField(); i++ {
		switch rT.Field(i).Type.Kind() {
		case reflect.Struct:
			rm := datahubColumnToFieldMap(rV.Field(i).Interface())
			for tag, value := range rm {
				m[tag] = value
			}
		default:
			tag, ok := rT.Field(i).Tag.Lookup(tagDatahubColumn)
			if !ok {

			}
			m[tag] = rV.Field(i)
		}
	}

	return m
}

func newCondition(option interface{}) (common.Condition, error) {
	rV := reflect.ValueOf(option)
	rT := rV.Type()
	switch rT.Kind() {
	case reflect.Struct:
	default:
		return common.Condition{}, errors.Errorf(`not supported type("%s")`, rT.Kind().String())
	}

	cond := common.Condition{
		Keys:      make([]string, 0, rV.NumField()),
		Values:    make([]string, 0, rV.NumField()),
		Operators: make([]string, 0, rV.NumField()),
	}
	for i := 0; i < rV.NumField(); i++ {
		field := rT.Field(i)
		datahubColumn, exist := field.Tag.Lookup(tagDatahubColumn)
		if !exist || datahubColumn == "" {
			continue
		}

		value := ""
		fieldValue := rV.Field(i)
		fieldKind := fieldValue.Kind()
		switch fieldKind {
		case reflect.Int:
			value = strconv.FormatInt(fieldValue.Int(), 10)
		case reflect.Int8:
			value = strconv.FormatInt(fieldValue.Int(), 10)
		case reflect.Int16:
			value = strconv.FormatInt(fieldValue.Int(), 10)
		case reflect.Int32:
			value = strconv.FormatInt(fieldValue.Int(), 10)
		case reflect.Int64:
			value = strconv.FormatInt(fieldValue.Int(), 10)
		case reflect.Float32:
			value = strconv.FormatFloat(fieldValue.Float(), 'f', -1, 32)
		case reflect.Float64:
			value = strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64)
		case reflect.String:
			value = fieldValue.String()
		case reflect.Bool:
			value = strconv.FormatBool(fieldValue.Bool())
		default:
			return common.Condition{}, errors.Errorf(`field type("%s") not supported`, fieldKind)
		}
		if value == "" {
			continue
		}

		cond.Keys = append(cond.Keys, datahubColumn)
		cond.Values = append(cond.Values, value)
		cond.Operators = append(cond.Operators, "=")
	}

	return cond, nil
}

// decodeSlice
// type of argument "items" must be pointer to slice
func decodeSlice(data data.Data, items interface{}) error {
	slicePtr := reflect.ValueOf(items)
	rV := slicePtr.Elem()
	rT := rV.Type()
	switch rT.Kind() {
	case reflect.Slice:
	default:
		return errors.Errorf(`items type("%s") not supported`, rT.Kind().String())
	}

	ln := rV.Len()
	if ln > 0 {
	}

	elementType := rV.Type().Elem()
	for _, rawData := range data.Rawdata {
		if rawData == nil {
			continue
		}

		for _, group := range rawData.Groups {
			if group == nil {
				continue
			}

			columnToIndexMap := make(map[string]int, len(group.Columns))
			for i, column := range group.Columns {
				columnToIndexMap[column] = i
			}

			for _, row := range group.Rows {
				if row == nil {
					continue
				}

				var element reflect.Value
				if elementType.Kind() == reflect.Ptr {
					element = reflect.New(elementType.Elem())
				}
				if elementType.Kind() == reflect.Struct {
					element = reflect.New(elementType).Elem()
				}

				elementType := element.Type()
				for i := 0; i < element.NumField(); i++ {
					datahubColumn, exist := elementType.Field(i).Tag.Lookup(tagDatahubColumn)
					if !exist || datahubColumn == "" {
						continue
					}

					index, exist := columnToIndexMap[datahubColumn]
					if !exist {
						continue
					}
					rawValue := row.Values[index]

					fieldValue := element.Field(i)
					kind := element.Field(i).Kind()
					switch kind {
					case reflect.Int:
						value, err := strconv.ParseInt(rawValue, 10, 64)
						if err != nil {
							return errors.Wrap(err, "parse int failed")
						}
						fieldValue.SetInt(value)
					case reflect.Int8:
						value, err := strconv.ParseInt(rawValue, 10, 8)
						if err != nil {
							return errors.Wrap(err, "parse int failed")
						}
						fieldValue.SetInt(value)
					case reflect.Int16:
						value, err := strconv.ParseInt(rawValue, 10, 16)
						if err != nil {
							return errors.Wrap(err, "parse int failed")
						}
						fieldValue.SetInt(value)
					case reflect.Int32:
						value, err := strconv.ParseInt(rawValue, 10, 32)
						if err != nil {
							return errors.Wrap(err, "parse int failed")
						}
						fieldValue.SetInt(value)
					case reflect.Int64:
						value, err := strconv.ParseInt(rawValue, 10, 64)
						if err != nil {
							return errors.Wrap(err, "parse int failed")
						}
						fieldValue.SetInt(value)
					case reflect.Float32:
						value, err := strconv.ParseFloat(rawValue, 32)
						if err != nil {
							return errors.Wrap(err, "parse float failed")
						}
						fieldValue.SetFloat(value)
					case reflect.Float64:
						value, err := strconv.ParseFloat(rawValue, 64)
						if err != nil {
							return errors.Wrap(err, "parse float failed")
						}
						fieldValue.SetFloat(value)
					case reflect.String:
						fieldValue.SetString(rawValue)
					case reflect.Bool:
						value, err := strconv.ParseBool(rawValue)
						if err != nil {
							return errors.Wrap(err, "parse bool failed")
						}
						fieldValue.SetBool(value)
					default:
						return errors.Errorf(`not supported "%s"`, kind.String())
					}
				}
				rV.Set(reflect.Append(rV, element))
			}
		}
	}
	return nil
}
