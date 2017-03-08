package cfapi

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"code.cloudfoundry.org/cli/cf/api/resources"
	"code.cloudfoundry.org/cli/cf/models"
	"code.cloudfoundry.org/cli/util/generic"
)

// CfEvent -
type CfEvent struct {
	GUID string
	Name string
	Type string

	EventList []models.EventFields
}

type eventResource struct {
	resources.Resource
	Entity struct {
		Timestamp time.Time
		Type      string `json:"type"`
		Actor     string `json:"actor"`
		ActorType string `json:"actor_type"`
		ActorName string `json:"actor_name"`
		Actee     string `json:"actee"`
		ActeeType string `json:"actee_type"`
		ActeeName string `json:"actee_name"`
		Metadata  map[string]interface{}
	}
}

// GetAllEventsInSpace -
func (s *CfCliSession) GetAllEventsInSpace(from time.Time) (events map[string]CfEvent, err error) {

	events = make(map[string]CfEvent)
	spaceGUID := s.GetSessionSpace().GUID
	timeFilter := url.QueryEscape(fmt.Sprintf("timestamp>%s", from.Format("2006-01-02 15:04:05-07:00")))

	count := 0

	err = s.ccGateway.ListPaginatedResources(
		s.config.APIEndpoint(),
		fmt.Sprintf("/v2/events?results-per-page=100&order-direction=asc&q=space_guid:%s&q=%s", spaceGUID, timeFilter),
		eventResource{},

		func(resource interface{}) bool {

			eventResource := resource.(eventResource)

			metadata := generic.NewMap(eventResource.Entity.Metadata)
			if metadata.Has("request") {
				metadata = generic.NewMap(metadata.Get("request"))
			}

			eventFields := models.EventFields{
				GUID:        eventResource.Metadata.GUID,
				Name:        eventResource.Entity.Type,
				Timestamp:   eventResource.Entity.Timestamp,
				Actor:       eventResource.Entity.Actor,
				ActorName:   eventResource.Entity.ActorName,
				Description: formatDescription(metadata, knownMetadataKeys),
			}

			sourceGUID := eventResource.Entity.Actee
			event, exists := events[sourceGUID]
			if exists {
				event.EventList = append(event.EventList, eventFields)
			} else {
				event = CfEvent{
					GUID:      eventResource.Entity.Actee,
					Name:      eventResource.Entity.ActeeName,
					Type:      eventResource.Entity.ActeeType,
					EventList: []models.EventFields{eventFields},
				}
			}
			events[sourceGUID] = event

			// Do not paginate as there is an error
			// in the next url returned in the response
			count++
			return count < 100
		})

	return
}

// GetAllEventsForApp -
func (s *CfCliSession) GetAllEventsForApp(appGUID string, from time.Time) (cfEvent CfEvent, err error) {

	timeFilter := url.QueryEscape(fmt.Sprintf("timestamp>%s", from.Format("2006-01-02 15:04:05-07:00")))
	count := 0

	err = s.ccGateway.ListPaginatedResources(
		s.config.APIEndpoint(),
		fmt.Sprintf("/v2/events?results-per-page=100&order-direction=asc&q=actee:%s&q=%s", appGUID, timeFilter),
		eventResource{},

		func(resource interface{}) bool {

			eventResource := resource.(eventResource)

			metadata := generic.NewMap(eventResource.Entity.Metadata)
			if metadata.Has("request") {
				metadata = generic.NewMap(metadata.Get("request"))
			}

			eventFields := models.EventFields{
				GUID:        eventResource.Metadata.GUID,
				Name:        eventResource.Entity.Type,
				Timestamp:   eventResource.Entity.Timestamp,
				Actor:       eventResource.Entity.Actor,
				ActorName:   eventResource.Entity.ActorName,
				Description: formatDescription(metadata, knownMetadataKeys),
			}

			if len(cfEvent.GUID) == 0 {
				cfEvent.GUID = eventResource.Entity.Actee
				cfEvent.Name = eventResource.Entity.ActeeName
				cfEvent.Type = eventResource.Entity.ActeeType
			}
			cfEvent.EventList = append(cfEvent.EventList, eventFields)

			// Do not paginate as there is an error
			// in the next url returned in the response
			count++
			return count < 100
		})

	return
}

// Event description formatting

var knownMetadataKeys = []string{
	"index",
	"reason",
	"exit_description",
	"exit_status",
	"recursive",
	"disk_quota",
	"instances",
	"memory",
	"state",
	"command",
	"environment_json",
}

func formatDescription(metadata generic.Map, keys []string) string {
	parts := []string{}
	for _, key := range keys {
		value := metadata.Get(key)
		if value != nil {
			parts = append(parts, fmt.Sprintf("%s: %s", key, formatDescriptionPart(value)))
		}
	}
	return strings.Join(parts, ", ")
}

func formatDescriptionPart(val interface{}) string {
	switch val := val.(type) {
	case string:
		return val
	case float64:
		return strconv.FormatFloat(val, byte('f'), -1, 64)
	case bool:
		if val {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%s", val)
	}
}
