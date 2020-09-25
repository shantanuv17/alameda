package keycodes

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	EventMgt "prophetstor.com/alameda/internal/pkg/event-mgt"
	AlamedaUtils "prophetstor.com/alameda/pkg/utils"
	K8SUtils "prophetstor.com/alameda/pkg/utils/kubernetes"
	ApiEvents "prophetstor.com/api/datahub/events"
	"time"
)

func NewKeycodeEvent(level ApiEvents.EventLevel, message string) *ApiEvents.Event {
	namespace := K8SUtils.GetRunningNamespace()

	clusterId, err := K8SUtils.GetClusterUID(K8SClient)
	if err != nil {
		scope.Errorf("failed to get cluster id: %s", err.Error())
	}

	source := &ApiEvents.EventSource{
		Host:      "",
		Component: fmt.Sprintf("%s-datahub", namespace),
	}

	subject := &ApiEvents.K8SObjectReference{
		Kind:       "Pod",
		Namespace:  namespace,
		Name:       "Federator.ai",
		ApiVersion: "v1",
	}

	event := &ApiEvents.Event{
		Time:      &timestamp.Timestamp{Seconds: time.Now().Unix()},
		Id:        AlamedaUtils.GenerateUUID(),
		ClusterId: clusterId,
		Source:    source,
		Type:      ApiEvents.EventType_EVENT_TYPE_LICENSE,
		Version:   ApiEvents.EventVersion_EVENT_VERSION_V1,
		Level:     level,
		Subject:   subject,
		Message:   message,
	}

	return event
}

func PostEvent(level ApiEvents.EventLevel, message string) error {
	if level == ApiEvents.EventLevel_EVENT_LEVEL_INFO {
		scope.Info(message)
	} else {
		scope.Error(message)
	}

	request := &ApiEvents.CreateEventsRequest{}
	request.Events = append(request.Events, NewKeycodeEvent(level, message))

	return EventMgt.PostEvents(request)
}
