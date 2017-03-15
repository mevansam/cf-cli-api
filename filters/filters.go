package filters

import "time"

// EventFilter -
type EventFilter interface {
	GetEventsForAllAppsInSpace(from time.Time, inclusive bool) ([]AppEvent, error)
	GetEventsForApp(appGUID string, from time.Time, inclusive bool) ([]AppEvent, error)
}
