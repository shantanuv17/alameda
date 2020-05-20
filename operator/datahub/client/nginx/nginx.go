package nginx

import (
	"context"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/code"

	// autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	// "github.com/containers-ai/alameda/operator/datahub/client"
	"github.com/containers-ai/alameda/operator/datahub/client/nginx/entity"
	"github.com/containers-ai/alameda/operator/pkg/nginx"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

const (
	tagDatahubColumnType = "datahubcolumntype"
	tagDatahubColumn     = "datahubcolumn"
	tagDatahubDataType   = "datahubdatatype"
)

type NginxRepository struct {
	datahubClient datahub.DatahubServiceClient
	schemaConfig  config
	logger        *log.Scope
}

func NewNginxRepository(datahubClient datahub.DatahubServiceClient, logger *log.Scope) NginxRepository {
	if logger == nil {
		logger = log.RegisterScope("datahub-client", "", 0)
	}
	return NginxRepository{
		datahubClient: datahubClient,
		// TODO:
		schemaConfig: defaultConfig,
		logger:       logger,
	}
}

func (n NginxRepository) CreateNginxs(ctx context.Context, nginxs []nginx.Nginx) error {
	if len(nginxs) == 0 {
		return nil
	}
	req, err := n.newWriteDataRequestByNginxs(nginxs)
	if err != nil {
		return errors.Wrap(err, "new WriteDataRequest by nginxs failed")
	}
	if err := n.sendWriteDataRequest(ctx, req); err != nil {
		return err
	}
	return nil
}

func (n NginxRepository) ListNginxs(ctx context.Context, option ListNginxsOption) ([]nginx.Nginx, error) {
	req, err := n.newReadDataRequestForNginxs(option)
	if err != nil {
		return nil, errors.Wrap(err, "new ReadDataRequest failed")
	}
	data, err := n.sendReadDataRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	entities := make([]entity.Nginx, 0)
	if err := decodeSlice(data, &entities); err != nil {
		return nil, errors.Wrap(err, "decode data failed")
	}
	nginxs := make([]nginx.Nginx, 0, len(entities))
	for _, e := range entities {
		nginxs = append(nginxs, nginx.Nginx{
			ClusterName:            e.ClusterName,
			ExporterNamespace:      e.ExporterNamespace,
			AlamedaScalerName:      e.AlamedaScalerName,
			AlamedaScalerNamespace: e.AlamedaScalerNamespace,
			Policy:                 e.Policy,
			EnableExecution:        e.EnableExecution,
			ExporterPods:           e.ExporterPods,
			HTTPResponseTime:       e.HTTPResponseTime,
			ResourceMeta: nginx.ResourceMeta{
				KubernetesMeta: nginx.KubernetesMeta{
					Namespace:        e.ResourceK8SNamespace,
					Name:             e.ResourceK8SName,
					ServiceNamespace: e.ResourceK8SServiceNamespace,
					ServiceName:      e.ResourceK8SServiceName,
					RouteNamespace:   e.ResourceK8SRouteNamespace,
					RouteName:        e.ResourceK8SRouteName,
					Kind:             e.ResourceK8SKind,
					ReadyReplicas:    e.ReadyReplicas,
					SpecReplicas:     e.SpecReplicas,
					CPULimit:         e.ResourceK8SCPULimit,
					CPURequest:       e.ResourceK8SCPURequest,
					MemoryLimit:      e.ResourceK8SMemoryLimit,
					MemoryRequest:    e.ResourceK8SMemoryRequest,
				},
			},
			MaxReplicas: e.MaxReplicas,
			MinReplicas: e.MinReplicas,
		})
	}
	return nginxs, nil
}

func (n NginxRepository) DeleteNginxs(ctx context.Context, nginxs []nginx.Nginx) error {
	if len(nginxs) == 0 {
		return nil
	}
	req, err := n.newDeleteDataRequestByNginxs(nginxs)
	if err != nil {
		return errors.Wrap(err, "new DeleteDataRequeset by nginxs failed")
	}
	if err := n.sendDeleteDataRequest(ctx, req); err != nil {
		return err
	}
	return nil
}

func (n NginxRepository) DeleteNginxsByOption(ctx context.Context, option DeleteNginxsOption) error {
	deleteData, err := newDeleteData(n.schemaConfig.nginx.nginx.delete.measurements[0], option)
	if err != nil {
		return errors.Wrap(err, "new DeleteData failed")
	}
	schemaMeta := n.schemaConfig.nginx.nginx.delete.schemaMeta
	req := data.DeleteDataRequest{
		SchemaMeta: &schemas.SchemaMeta{
			Scope:    schemas.Scope(schemas.Scope_value[schemaMeta.scope]),
			Category: schemaMeta.category,
			Type:     schemaMeta.type_,
		},
		DeleteData: []*data.DeleteData{
			&deleteData,
		},
	}
	if err := n.sendDeleteDataRequest(ctx, req); err != nil {
		return err
	}
	return nil
}

func (n NginxRepository) sendWriteDataRequest(ctx context.Context, req data.WriteDataRequest) error {
	n.logger.Debugf("Write data to Datahub. Request: %+v", req)
	status, err := n.datahubClient.WriteData(ctx, &req)
	if err != nil {
		return errors.Wrap(err, "send WriteDataRequest failed")
	} else if status == nil {
		return errors.New("receive nil status")
	} else if status.Code != int32(code.Code_OK) {
		return errors.Errorf("send WriteDataRequest failed: statuscode: %d, message: %s", status.Code, status.Message)
	}
	return nil
}

func (n NginxRepository) sendReadDataRequest(ctx context.Context, req data.ReadDataRequest) (data.Data, error) {
	n.logger.Debugf("Read data from Datahub. Request: %+v", req)
	resp, err := n.datahubClient.ReadData(ctx, &req)
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
	n.logger.Debugf("Read data from Datahub. Response: %+v", resp)
	return *resp.Data, nil
}

func (n NginxRepository) sendDeleteDataRequest(ctx context.Context, req data.DeleteDataRequest) error {
	n.logger.Debugf("Delete data from Datahub. Request: %+v", req)
	status, err := n.datahubClient.DeleteData(ctx, &req)
	if err != nil {
		return errors.Wrap(err, "send DeleteDataRequest failed")
	} else if status == nil {
		return errors.New("receive nil status")
	} else if status.Code != int32(code.Code_OK) {
		return errors.Errorf("send DeleteDataRequest failed: statuscode: %d, message: %s", status.Code, status.Message)
	}
	return nil
}

func (n NginxRepository) newWriteDataRequestByNginxs(nginxs []nginx.Nginx) (data.WriteDataRequest, error) {
	entities := make([]entity.Nginx, 0, len(nginxs))
	for _, nginx := range nginxs {
		entity := entity.NewNginx(nginx)
		entities = append(entities, entity)
	}

	writeData, err := newWriteData(n.schemaConfig.nginx.nginx.create.measurements[0], entities)
	if err != nil {
		return data.WriteDataRequest{}, errors.Wrap(err, "new WriteData failed")
	}
	req := data.WriteDataRequest{
		SchemaMeta: &schemas.SchemaMeta{
			Scope:    schemas.Scope(schemas.Scope_value[n.schemaConfig.nginx.nginx.create.schemaMeta.scope]),
			Category: n.schemaConfig.nginx.nginx.create.schemaMeta.category,
			Type:     n.schemaConfig.nginx.nginx.create.schemaMeta.type_,
		},
		WriteData: []*data.WriteData{&writeData},
	}
	return req, nil
}

func (n NginxRepository) newDeleteDataRequestByNginxs(nginxs []nginx.Nginx) (data.DeleteDataRequest, error) {
	entities := make([]entity.Nginx, 0, len(nginxs))
	for _, nginx := range nginxs {
		entity := entity.NewNginx(nginx)
		entities = append(entities, entity)
	}
	deleteData, err := newDeleteData(n.schemaConfig.nginx.nginx.delete.measurements[0], entities)
	if err != nil {
		return data.DeleteDataRequest{}, errors.Wrap(err, "new DeleteData failed")
	}
	req := data.DeleteDataRequest{
		SchemaMeta: &schemas.SchemaMeta{
			Scope:    schemas.Scope(schemas.Scope_value[n.schemaConfig.nginx.nginx.delete.schemaMeta.scope]),
			Category: n.schemaConfig.nginx.nginx.delete.schemaMeta.category,
			Type:     n.schemaConfig.nginx.nginx.delete.schemaMeta.type_,
		},
		DeleteData: []*data.DeleteData{&deleteData},
	}
	return req, nil
}

func (n NginxRepository) newReadDataRequestForNginxs(option ListNginxsOption) (data.ReadDataRequest, error) {
	condition, err := newCondition(option)
	if err != nil {
		return data.ReadDataRequest{}, errors.Wrap(err, "new condition failed")
	}
	readData := make([]*data.ReadData, 0, len(n.schemaConfig.nginx.nginx.list.measurements))
	for _, measurement := range n.schemaConfig.nginx.nginx.list.measurements {
		readData = append(readData, &data.ReadData{
			Measurement: measurement.name,
			QueryCondition: &common.QueryCondition{
				WhereCondition: []*common.Condition{&condition},
			},
		})
	}

	req := data.ReadDataRequest{
		SchemaMeta: &schemas.SchemaMeta{
			Scope:    schemas.Scope(schemas.Scope_value[n.schemaConfig.nginx.nginx.list.schemaMeta.scope]),
			Category: n.schemaConfig.nginx.nginx.list.schemaMeta.category,
			Type:     n.schemaConfig.nginx.nginx.list.schemaMeta.type_,
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

	slice := reflect.ValueOf(dataRows)
	if slice.Len() == 0 {
		return data.WriteData{}, nil
	}

	columns, err := listDatahubColumns(slice.Index(0))
	if err != nil {
		return data.WriteData{}, errors.Wrap(err, "list Datahub columns failed")
	}

	rows := make([]*common.Row, 0, slice.Len())
	for i := 0; i < slice.Len(); i++ {
		element := slice.Index(i)

		values := make([]string, 0, len(columns))
		for j := 0; j < element.NumField(); j++ {
			fieldValue := element.Field(j)
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
		Columns:     columns,
		Rows:        rows,
	}
	return w, nil
}

// newDeleteData returns DeleteData containing with measuremnt and whereConditions
// Currently type of argument "dataRows" must be slice of struct, and each type of the field must in the second swith cases.
func newDeleteData(measurement measurement, dataRows interface{}) (data.DeleteData, error) {
	dataV := reflect.ValueOf(dataRows)
	dataT := dataV.Type()
	switch dataV.Kind() {
	case reflect.Slice:
	case reflect.Struct:
		tmpSlice := reflect.MakeSlice(reflect.SliceOf(dataT), 0, 1)
		tmpSlice = reflect.Append(tmpSlice, dataV)
		dataRows = tmpSlice.Interface()
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
			Types:     []common.DataType{},
		}
		for j := 0; j < rV.NumField(); j++ {
			field := rT.Field(j)
			columnType, exist := field.Tag.Lookup(tagDatahubColumnType)
			if !exist {
				return data.DeleteData{}, errors.Errorf(`tag("%s") not found`, tagDatahubColumnType)
			} else if columnType == "" {
				return data.DeleteData{}, errors.Errorf(`tag("%s") value empty`, tagDatahubColumnType)
			} else if columnType != entity.Tag {
				continue
			}

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

			datahubDataType, exist := rT.Field(j).Tag.Lookup(tagDatahubDataType)
			if !exist {
				return data.DeleteData{}, errors.Errorf(`tag("%s") not found`, tagDatahubDataType)
			} else if datahubDataType == "" {
				return data.DeleteData{}, errors.Errorf(`tag("%s") value empty`, tagDatahubDataType)
			}
			dataType, exist := common.DataType_value[datahubDataType]
			if !exist {
				return data.DeleteData{}, errors.Errorf(`datatype("%s") not supported`, datahubDataType)
			}

			cond.Keys = append(cond.Keys, datahubColumn)
			cond.Values = append(cond.Values, value)
			cond.Operators = append(cond.Operators, "=")
			cond.Types = append(cond.Types, common.DataType(dataType))
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
		Types:     make([]common.DataType, 0, rV.NumField()),
	}
	for i := 0; i < rV.NumField(); i++ {
		field := rT.Field(i)
		datahubColumn, exist := field.Tag.Lookup(tagDatahubColumn)
		if !exist || datahubColumn == "" {
			continue
		}

		datahubDataType, exist := field.Tag.Lookup(tagDatahubDataType)
		if !exist || datahubDataType == "" {
			continue
		}
		dataType, exist := common.DataType_value[datahubDataType]
		if !exist {
			return common.Condition{}, errors.Errorf(`datatype("%s") not supported`, datahubDataType)
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
		cond.Types = append(cond.Types, common.DataType(dataType))
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

func listDatahubColumns(entity reflect.Value) ([]string, error) {
	switch entity.Kind() {
	case reflect.Struct:
	default:
		return nil, errors.Errorf(`type("%s") not supported`, entity.Kind().String())
	}

	columns := make([]string, 0)
	rT := entity.Type()
	for i := 0; i < entity.NumField(); i++ {
		field := rT.Field(i)
		column, exist := field.Tag.Lookup(tagDatahubColumn)
		if !exist {
			return nil, errors.Errorf(`tag("%s") not exist`, tagDatahubColumn)
		} else if column == "" {
			return nil, errors.Errorf(`tag("%s") empty`, tagDatahubColumn)
		}
		columns = append(columns, column)
	}
	return columns, nil
}
