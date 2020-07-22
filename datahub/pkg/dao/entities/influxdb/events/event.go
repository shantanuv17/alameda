package events

import (
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb"
)

const (
	EventTime              influxdb.Tag = "time"
	EventClusterId         influxdb.Tag = "cluster_id"
	EventSourceComponent   influxdb.Tag = "source_component"
	EventSourceHost        influxdb.Tag = "source_host"
	EventType              influxdb.Tag = "type"
	EventVersion           influxdb.Tag = "version"
	EventLevel             influxdb.Tag = "level"
	EventSubjectKind       influxdb.Tag = "subject_kind"
	EventSubjectNamespace  influxdb.Tag = "subject_namespace"
	EventSubjectName       influxdb.Tag = "subject_name"
	EventSubjectApiVersion influxdb.Tag = "subject_api_version"

	EventId      influxdb.Field = "id"
	EventMessage influxdb.Field = "message"
	EventData    influxdb.Field = "data"
)

var (
	// ControllerTags is list of tags of alameda_controller_recommendation measurement
	EventTags = []influxdb.Tag{
		EventTime,
		EventClusterId,
		EventSourceComponent,
		EventSourceHost,
		EventType,
		EventVersion,
		EventLevel,
		EventSubjectKind,
		EventSubjectNamespace,
		EventSubjectName,
		EventSubjectApiVersion,
	}
	// ControllerFields is list of fields of alameda_controller_recommendation measurement
	EventFields = []influxdb.Field{
		EventId,
		EventMessage,
		EventData,
	}
)
