package filters

import (
	"regexp"
	"time"

	"github.com/mevansam/cf-cli-api/cfapi"
)

// appEventFilter -
type appEventFilter struct {
	session  cfapi.CfSession
	appGUIDs []string

	logger *cfapi.Logger
}

// NewAppEventFilter -
func NewAppEventFilter(session cfapi.CfSession) (filter EventFilter) {

	filter = &appEventFilter{
		session: session,
		logger:  session.GetSessionLogger(),
	}
	return
}

// GetEventsForAllAppsInSpace -
func (f appEventFilter) GetEventsForAllAppsInSpace(from time.Time, inclusive bool) (events []AppEvent, err error) {

	allEvents, err := f.session.GetAllEventsInSpace(from, inclusive)
	if err == nil {
		for _, cfEvent := range allEvents {
			if cfEvent.Type == "app" {
				events = append(events, f.processEvents(cfEvent)...)
			}
		}
	}
	return
}

// GetEventsForApp -
func (f appEventFilter) GetEventsForApp(appGUID string, from time.Time, inclusive bool) (events []AppEvent, err error) {

	cfEvent, err := f.session.GetAllEventsForApp(appGUID, from, inclusive)
	if err == nil {
		if cfEvent.Type == "app" {
			events = append(events, f.processEvents(cfEvent)...)
		}
	}
	return
}

// GetEventsForAllAppsInSpace -
func (f appEventFilter) processEvents(cfEvent cfapi.CfEvent) (appEvents []AppEvent) {

	var (
		eventStateMap  map[int][]eventState
		eventStateList []eventState
		lastEventState eventState
		ok, match      bool
	)

	lastEventState.stateID = -1
	lastEventState.eventType = EtUnknown

	for _, e := range cfEvent.EventList {

		f.logger.DebugMessage("Processing event: %s - %s: %s -> %s",
			e.Timestamp.Format(time.RFC3339), e.Name, cfEvent.Name, e.Description)

		if eventStateMap, ok = stateMap[e.Name]; ok {
			if eventStateList, ok = eventStateMap[lastEventState.stateID]; ok {
				for _, s := range eventStateList {
					if match, _ = regexp.MatchString(s.pattern, e.Description); match {

						// Initialize new state maintaining prev state's
						// event type if match does not have an event type
						if s.eventType == EtUnknown {
							s.eventType = lastEventState.eventType
						}
						if s.trigger {

							appEvents = append(appEvents, AppEvent{
								SourceGUID: cfEvent.GUID,
								SourceName: cfEvent.Name,
								SourceType: cfEvent.Type,
								EventType:  s.eventType,
								Timestamp:  e.Timestamp,
							})

							f.logger.DebugMessage("*** Triggering app event: %s - %s %s",
								e.Timestamp.Format(time.RFC3339), s.eventType, cfEvent.Name)

							// Reset state if triggered
							lastEventState.stateID = -1
							lastEventState.eventType = EtUnknown
						} else {
							lastEventState = s
						}
					}
				}
			}
		}
	}
	return
}

// Event State Table

type eventState struct {
	stateID   int
	last      string
	pattern   string
	eventType EventType
	trigger   bool
}

var stateMap = map[string]map[int][]eventState{

	"audit.app.create": {
		-1: []eventState{{
			stateID:   0,
			pattern:   "^instances: \\d*, memory: \\d*",
			eventType: EtCreated,
			trigger:   false,
		}},
	},
	"audit.app.update": {
		-1: []eventState{
			{
				stateID:   1,
				pattern:   "^instances: \\d*, memory: \\d*",
				eventType: EtUnknown,
				trigger:   false,
			},
			{
				stateID:   4,
				pattern:   "^instances: \\d*$",
				eventType: EtScaled,
				trigger:   true,
			},
			{
				stateID:   5,
				pattern:   "^memory: \\d*$",
				eventType: EtScaled,
				trigger:   true,
			},
		},
		0: []eventState{{
			stateID:   2,
			pattern:   "",
			eventType: EtUnknown,
			trigger:   false,
		}},
		1: []eventState{{
			stateID:   2,
			pattern:   "state: STOPPED",
			eventType: EtModified,
			trigger:   false,
		}},
		2: []eventState{{
			stateID:   3,
			pattern:   "state: STARTED",
			eventType: EtUnknown,
			trigger:   true,
		}},
	},
	"audit.app.restage": {
		-1: []eventState{{
			stateID:   6,
			pattern:   "",
			eventType: EtModified,
			trigger:   true,
		}},
	},
	"audit.app.map-route": {
		-1: []eventState{{
			stateID:   7,
			pattern:   "",
			eventType: EtRouteAdded,
			trigger:   true,
		}},
	},
	"audit.app.unmap-route": {
		-1: []eventState{{
			stateID:   8,
			pattern:   "",
			eventType: EtRouteDeleted,
			trigger:   true,
		}},
	},
	"audit.app.delete-request": {
		-1: []eventState{{
			stateID:   9,
			pattern:   "",
			eventType: EtDeleted,
			trigger:   true,
		}},
	},
}
