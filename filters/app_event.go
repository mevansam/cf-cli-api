package filters

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// EventType -
type EventType string

const (
	// EtCreated -
	EtCreated = EventType("created")
	// EtDeleted -
	EtDeleted = EventType("deleted")
	// EtModified -
	EtModified = EventType("modified")
	// EtScaled -
	EtScaled = EventType("scaled")
	// EtRouteAdded -
	EtRouteAdded = EventType("routed-added")
	// EtRouteDeleted -
	EtRouteDeleted = EventType("routed-deleted")

	// EtUnknown -
	EtUnknown = ""
)

// validEvents -
var validEvents = map[string]EventType{
	string(EtCreated):      EtCreated,
	string(EtDeleted):      EtDeleted,
	string(EtModified):     EtModified,
	string(EtScaled):       EtScaled,
	string(EtRouteAdded):   EtRouteAdded,
	string(EtRouteDeleted): EtRouteDeleted,
	string(EtUnknown):      EtUnknown,
}

// AppEvent -
type AppEvent struct {
	SourceGUID string
	SourceName string
	SourceType string
	EventType  EventType
	Timestamp  time.Time
}

// NewAppEvent -
func NewAppEvent(data string) (ae AppEvent, err error) {

	var (
		ok bool
	)

	fields := strings.Split(data, "|")
	if len(fields) < 5 {
		err = fmt.Errorf("The string data should have 5 or more fields. '%d' fields were extracted.", len(fields))
		return
	}

	ae.SourceGUID = fields[0]
	if ok, err = regexp.MatchString("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}", ae.SourceGUID); err != nil {
		return
	}
	if !ok {
		err = fmt.Errorf("The app GUID '%s' is not a valid GUID.", ae.SourceGUID)
		return
	}
	if ae.EventType, ok = validEvents[fields[3]]; !ok {
		err = fmt.Errorf("Event type '%s' is not valid.", fields[3])
		return
	}
	if ae.Timestamp, err = time.Parse(time.RFC3339, fields[4]); err != nil {
		return
	}

	ae.SourceName = fields[1]
	ae.SourceType = fields[2]

	return
}

// String -
func (ae AppEvent) String() string {

	return fmt.Sprintf("%s|%s|%s|%s|%s",
		ae.SourceGUID, ae.SourceName, ae.SourceType, ae.EventType,
		ae.Timestamp.Format(time.RFC3339))
}
