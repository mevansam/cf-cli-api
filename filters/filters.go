package filters

import "time"

// EventFilter -
type EventFilter interface {
	GetEventsForAllAppsInSpace(from time.Time) ([]AppEvent, error)
	GetEventsForApp(appGUID string, from time.Time) ([]AppEvent, error)
}
