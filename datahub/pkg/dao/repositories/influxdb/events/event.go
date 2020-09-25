package events

import (
	"github.com/golang/protobuf/ptypes"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	EntityInfluxEvent "prophetstor.com/alameda/datahub/pkg/dao/entities/influxdb/events"
	RepoInflux "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	Log "prophetstor.com/alameda/pkg/utils/log"
	ApiEvents "prophetstor.com/api/datahub/events"
	"time"
)

var (
	scope = Log.RegisterScope("event_db_measurement", "event DB measurement", 0)
)

type EventRepository struct {
	influxDB *InfluxDB.InfluxClient
}

func NewEventRepository(influxDBCfg *InfluxDB.Config) *EventRepository {
	return &EventRepository{
		influxDB: &InfluxDB.InfluxClient{
			Address:  influxDBCfg.Address,
			Username: influxDBCfg.Username,
			Password: influxDBCfg.Password,
		},
	}
}

func (e *EventRepository) CreateEvents(in *ApiEvents.CreateEventsRequest) error {
	points := make([]*InfluxClient.Point, 0)

	for _, event := range in.GetEvents() {
		tags := map[string]string{
			string(EntityInfluxEvent.EventClusterId):         event.GetClusterId(),
			string(EntityInfluxEvent.EventSourceHost):        event.GetSource().GetHost(),
			string(EntityInfluxEvent.EventSourceComponent):   event.GetSource().GetComponent(),
			string(EntityInfluxEvent.EventType):              event.GetType().String(),
			string(EntityInfluxEvent.EventVersion):           event.GetVersion().String(),
			string(EntityInfluxEvent.EventLevel):             event.GetLevel().String(),
			string(EntityInfluxEvent.EventSubjectKind):       event.GetSubject().GetKind(),
			string(EntityInfluxEvent.EventSubjectNamespace):  event.GetSubject().GetNamespace(),
			string(EntityInfluxEvent.EventSubjectName):       event.GetSubject().GetName(),
			string(EntityInfluxEvent.EventSubjectApiVersion): event.GetSubject().GetApiVersion(),
		}

		fields := map[string]interface{}{
			string(EntityInfluxEvent.EventId):      event.GetId(),
			string(EntityInfluxEvent.EventMessage): event.GetMessage(),
			string(EntityInfluxEvent.EventData):    event.GetData(),
		}

		tempTime, _ := ptypes.Timestamp(event.GetTime())
		pt, err := InfluxClient.NewPoint(string(Event), tags, fields, tempTime)
		if err != nil {
			scope.Error(err.Error())
		}

		points = append(points, pt)
	}

	err := e.influxDB.WritePoints(points, InfluxClient.BatchPointsConfig{
		Database: string(RepoInflux.Event),
	})

	if err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (e *EventRepository) ListEvents(in *ApiEvents.ListEventsRequest) ([]*ApiEvents.Event, error) {
	idList := in.GetId()
	clusterIdList := in.GetClusterId()

	eventTypeList := make([]string, 0)
	for _, eventType := range in.GetType() {
		eventTypeList = append(eventTypeList, eventType.String())
	}

	eventVersionList := make([]string, 0)
	for _, eventVersion := range in.GetVersion() {
		eventVersionList = append(eventVersionList, eventVersion.String())
	}

	eventLevelList := make([]string, 0)
	for _, eventLevel := range in.GetLevel() {
		eventLevelList = append(eventLevelList, eventLevel.String())
	}

	influxdbStatement := InfluxDB.Statement{
		Measurement:    Event,
		QueryCondition: DBCommon.BuildQueryConditionV1(in.GetQueryCondition()),
	}

	influxdbStatement.AppendWhereClauseByList(string(EntityInfluxEvent.EventId), "=", "OR", idList)
	influxdbStatement.AppendWhereClauseByList(string(EntityInfluxEvent.EventClusterId), "=", "OR", clusterIdList)
	influxdbStatement.AppendWhereClauseByList(string(EntityInfluxEvent.EventType), "=", "OR", eventTypeList)
	influxdbStatement.AppendWhereClauseByList(string(EntityInfluxEvent.EventVersion), "=", "OR", eventVersionList)
	influxdbStatement.AppendWhereClauseByList(string(EntityInfluxEvent.EventLevel), "=", "OR", eventLevelList)

	influxdbStatement.AppendWhereClauseFromTimeCondition()
	influxdbStatement.SetOrderClauseFromQueryCondition()
	influxdbStatement.SetLimitClauseFromQueryCondition()
	cmd := influxdbStatement.BuildQueryCmd()

	results, err := e.influxDB.QueryDB(cmd, string(RepoInflux.Event))
	if err != nil {
		return make([]*ApiEvents.Event, 0), err
	}

	influxdbRows := InfluxDB.PackMap(results)
	events := e.getEventsFromInfluxRows(influxdbRows)

	return events, nil
}

func (e *EventRepository) getEventsFromInfluxRows(rows []*InfluxDB.InfluxRow) []*ApiEvents.Event {
	events := make([]*ApiEvents.Event, 0)

	for _, influxdbRow := range rows {
		for _, data := range influxdbRow.Data {
			t, _ := time.Parse(time.RFC3339Nano, data[string(EntityInfluxEvent.EventTime)])
			tempTime, _ := ptypes.TimestampProto(t)

			clusterId := data[string(EntityInfluxEvent.EventClusterId)]
			sourceHost := data[string(EntityInfluxEvent.EventSourceHost)]
			sourceComponent := data[string(EntityInfluxEvent.EventSourceComponent)]
			subjectKind := data[string(EntityInfluxEvent.EventSubjectKind)]
			subjectNamespace := data[string(EntityInfluxEvent.EventSubjectNamespace)]
			subjectName := data[string(EntityInfluxEvent.EventSubjectName)]
			subjectApiVersion := data[string(EntityInfluxEvent.EventSubjectApiVersion)]

			id := data[string(EntityInfluxEvent.EventId)]
			message := data[string(EntityInfluxEvent.EventMessage)]
			eventData := data[string(EntityInfluxEvent.EventData)]

			eventType := ApiEvents.EventType_EVENT_TYPE_UNDEFINED
			if tempType, exist := data[string(EntityInfluxEvent.EventType)]; exist {
				if value, ok := ApiEvents.EventType_value[tempType]; ok {
					eventType = ApiEvents.EventType(value)
				}
			}

			eventVersion := ApiEvents.EventVersion_EVENT_VERSION_UNDEFINED
			if tempVersion, exist := data[string(EntityInfluxEvent.EventVersion)]; exist {
				if value, ok := ApiEvents.EventVersion_value[tempVersion]; ok {
					eventVersion = ApiEvents.EventVersion(value)
				}
			}

			eventLevel := ApiEvents.EventLevel_EVENT_LEVEL_UNDEFINED
			if tempLevel, exist := data[string(EntityInfluxEvent.EventLevel)]; exist {
				if value, ok := ApiEvents.EventLevel_value[tempLevel]; ok {
					eventLevel = ApiEvents.EventLevel(value)
				}
			}

			event := ApiEvents.Event{
				Time:      tempTime,
				Id:        id,
				ClusterId: clusterId,
				Source: &ApiEvents.EventSource{
					Host:      sourceHost,
					Component: sourceComponent,
				},
				Type:    eventType,
				Version: eventVersion,
				Level:   eventLevel,
				Subject: &ApiEvents.K8SObjectReference{
					Kind:       subjectKind,
					Namespace:  subjectNamespace,
					Name:       subjectName,
					ApiVersion: subjectApiVersion,
				},
				Message: message,
				Data:    eventData,
			}

			events = append(events, &event)
		}
	}

	return events
}
