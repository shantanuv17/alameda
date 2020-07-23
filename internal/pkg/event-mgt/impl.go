package eventmgt

import (
	"encoding/json"
	EntityEvent "github.com/containers-ai/alameda/datahub/pkg/dao/entities/influxdb/events"
	RepoInflux "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb"
	RepoEvent "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/events"
	Rabbitmq "github.com/containers-ai/alameda/internal/pkg/message-queue/rabbitmq"
	DBCommon "github.com/containers-ai/alameda/pkg/database/common"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	DatahubLog "github.com/containers-ai/alameda/pkg/utils/log"
	ApiEvents "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/events"
	"github.com/golang/protobuf/ptypes"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"time"
)

var (
	scope = DatahubLog.RegisterScope("event_db_measurement", "event DB measurement", 0)
)

func (e *EventMgt) PostEvents(in *ApiEvents.CreateEventsRequest) error {
	points := make([]*InfluxClient.Point, 0)

	for _, event := range in.GetEvents() {
		tags := map[string]string{
			string(EntityEvent.EventClusterId):         event.GetClusterId(),
			string(EntityEvent.EventSourceHost):        event.GetSource().GetHost(),
			string(EntityEvent.EventSourceComponent):   event.GetSource().GetComponent(),
			string(EntityEvent.EventType):              event.GetType().String(),
			string(EntityEvent.EventVersion):           event.GetVersion().String(),
			string(EntityEvent.EventLevel):             event.GetLevel().String(),
			string(EntityEvent.EventSubjectKind):       event.GetSubject().GetKind(),
			string(EntityEvent.EventSubjectNamespace):  event.GetSubject().GetNamespace(),
			string(EntityEvent.EventSubjectName):       event.GetSubject().GetName(),
			string(EntityEvent.EventSubjectApiVersion): event.GetSubject().GetApiVersion(),
		}

		fields := map[string]interface{}{
			string(EntityEvent.EventId):      event.GetId(),
			string(EntityEvent.EventMessage): event.GetMessage(),
			string(EntityEvent.EventData):    event.GetData(),
		}

		tempTime, _ := ptypes.Timestamp(event.GetTime())
		pt, err := InfluxClient.NewPoint(string(RepoEvent.Event), tags, fields, tempTime)
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

	//send to rabbitmq
	err = e.sendEventsToMsgQueue(in)
	if err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (e *EventMgt) ListEvents(in *ApiEvents.ListEventsRequest) ([]*ApiEvents.Event, error) {
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
		Measurement:    RepoEvent.Event,
		QueryCondition: DBCommon.BuildQueryConditionV1(in.GetQueryCondition()),
	}

	influxdbStatement.AppendWhereClauseByList(string(EntityEvent.EventId), "=", "OR", idList)
	influxdbStatement.AppendWhereClauseByList(string(EntityEvent.EventClusterId), "=", "OR", clusterIdList)
	influxdbStatement.AppendWhereClauseByList(string(EntityEvent.EventType), "=", "OR", eventTypeList)
	influxdbStatement.AppendWhereClauseByList(string(EntityEvent.EventVersion), "=", "OR", eventVersionList)
	influxdbStatement.AppendWhereClauseByList(string(EntityEvent.EventLevel), "=", "OR", eventLevelList)

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

func (e *EventMgt) getEventsFromInfluxRows(rows []*InfluxDB.InfluxRow) []*ApiEvents.Event {
	events := make([]*ApiEvents.Event, 0)

	for _, influxdbRow := range rows {
		for _, data := range influxdbRow.Data {
			t, _ := time.Parse(time.RFC3339Nano, data[string(EntityEvent.EventTime)])
			tempTime, _ := ptypes.TimestampProto(t)

			clusterId := data[string(EntityEvent.EventClusterId)]
			sourceHost := data[string(EntityEvent.EventSourceHost)]
			sourceComponent := data[string(EntityEvent.EventSourceComponent)]
			subjectKind := data[string(EntityEvent.EventSubjectKind)]
			subjectNamespace := data[string(EntityEvent.EventSubjectNamespace)]
			subjectName := data[string(EntityEvent.EventSubjectName)]
			subjectApiVersion := data[string(EntityEvent.EventSubjectApiVersion)]

			id := data[string(EntityEvent.EventId)]
			message := data[string(EntityEvent.EventMessage)]
			eventData := data[string(EntityEvent.EventData)]

			eventType := ApiEvents.EventType_EVENT_TYPE_UNDEFINED
			if tempType, exist := data[string(EntityEvent.EventType)]; exist {
				if value, ok := ApiEvents.EventType_value[tempType]; ok {
					eventType = ApiEvents.EventType(value)
				}
			}

			eventVersion := ApiEvents.EventVersion_EVENT_VERSION_UNDEFINED
			if tempVersion, exist := data[string(EntityEvent.EventVersion)]; exist {
				if value, ok := ApiEvents.EventVersion_value[tempVersion]; ok {
					eventVersion = ApiEvents.EventVersion(value)
				}
			}

			eventLevel := ApiEvents.EventLevel_EVENT_LEVEL_UNDEFINED
			if tempLevel, exist := data[string(EntityEvent.EventLevel)]; exist {
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

func (e *EventMgt) sendEventsToMsgQueue(in *ApiEvents.CreateEventsRequest) error {
	messageQueue, err := Rabbitmq.NewRabbitMQSender(e.RabbitMQConfig)
	if err != nil {
		return err
	}
	defer messageQueue.Close()

	events, err := json.Marshal(in.GetEvents())
	if err != nil {
		return err
	}

	err = messageQueue.SendJsonString("event", string(events))
	return err
}
