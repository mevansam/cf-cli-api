package filters_test

import (
	"time"

	"code.cloudfoundry.org/cli/cf/models"

	"github.com/mevansam/cf-cli-api/cfapi"
	. "github.com/mevansam/cf-cli-api/cfapi/mocks"
	"github.com/mevansam/cf-cli-api/filters"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Application Event Collection Tests", func() {

	var (
		logger *cfapi.Logger

		session *MockSession
		filter  filters.EventFilter

		to time.Time
	)

	BeforeEach(func() {
		logger = cfapi.NewLogger(true, "true")

		session = &MockSession{Logger: logger}
		filter = filters.NewAppEventFilter(session)

		session.MockGetAllEventsInSpace = func(from time.Time, inclusive bool) (events map[string]cfapi.CfEvent, err error) {
			events = make(map[string]cfapi.CfEvent)
			for guid, cfEvent := range testEvents {
				eventList := []models.EventFields{}
				for _, event := range cfEvent.EventList {
					if event.Timestamp.After(from) && event.Timestamp.Before(to) {
						eventList = append(eventList, event)
					}
				}
				cfEvent.EventList = eventList
				events[guid] = cfEvent
			}
			return
		}
		session.MockGetAllEventsForApp = func(appGUID string, from time.Time, inclusive bool) (cfEvent cfapi.CfEvent, err error) {
			cfEvent, ok := testEvents[appGUID]
			if ok {
				eventList := []models.EventFields{}
				for _, event := range cfEvent.EventList {
					if event.Timestamp.After(from) && event.Timestamp.Before(to) {
						eventList = append(eventList, event)
					}
				}
				cfEvent.EventList = eventList
			}
			return
		}
	})

	Context("Reading and triggering events for one or more applications", func() {

		It("Should return events for an app between a certain range", func() {

			from, _ := time.Parse(time.RFC3339, "2017-03-01T00:00:00+04:00")
			to, _ = time.Parse(time.RFC3339, "2017-03-01T11:11:20+04:00")

			appEvents, err := filter.GetEventsForApp("d9d8b1c8-42a7-4bdf-b337-512232c653ca", from, false)
			Expect(err).Should(BeNil())
			Expect(len(appEvents)).To(Equal(0))

			to, _ = time.Parse(time.RFC3339, "2017-03-01T11:15:55+04:00")
			appEvents, err = filter.GetEventsForAllAppsInSpace(from, false)
			Expect(err).Should(BeNil())
			Expect(len(appEvents)).To(Equal(1))

			Expect(appEvents[0].SourceName).To(Equal("spring-music"))
			Expect(appEvents[0].SourceType).To(Equal("app"))
			Expect(appEvents[0].EventType).To(Equal(filters.EtCreated))

			from = appEvents[0].Timestamp
			to, _ = time.Parse(time.RFC3339, "2017-03-01T19:48:15+04:00")
			appEvents, err = filter.GetEventsForAllAppsInSpace(from, false)
			Expect(err).Should(BeNil())
			Expect(len(appEvents)).To(Equal(1))

			Expect(appEvents[0].SourceName).To(Equal("spring-music"))
			Expect(appEvents[0].SourceType).To(Equal("app"))
			Expect(appEvents[0].EventType).To(Equal(filters.EtModified))

			from = appEvents[0].Timestamp
			to, _ = time.Parse(time.RFC3339, "2017-03-01T19:48:35+04:00")
			appEvents, err = filter.GetEventsForAllAppsInSpace(from, false)
			Expect(err).Should(BeNil())
			Expect(len(appEvents)).To(Equal(0))

			to, _ = time.Parse(time.RFC3339, "2017-03-01T21:20:00+04:00")
			appEvents, err = filter.GetEventsForAllAppsInSpace(from, false)
			Expect(err).Should(BeNil())
			Expect(len(appEvents)).To(Equal(2))

			Expect(appEvents[0].EventType).To(Equal(filters.EtModified))
			Expect(appEvents[1].EventType).To(Equal(filters.EtRouteAdded))

			from = appEvents[1].Timestamp
			to, _ = time.Parse(time.RFC3339, "2017-03-01T22:32:00+04:00")
			appEvents, err = filter.GetEventsForAllAppsInSpace(from, false)
			Expect(err).Should(BeNil())
			Expect(len(appEvents)).To(Equal(2))

			Expect(appEvents[0].EventType).To(Equal(filters.EtScaled))
			Expect(appEvents[1].EventType).To(Equal(filters.EtScaled))

			from = appEvents[1].Timestamp
			to, _ = time.Parse(time.RFC3339, "2017-03-01T22:43:13+04:00")
			appEvents, err = filter.GetEventsForAllAppsInSpace(from, false)
			Expect(err).Should(BeNil())
			Expect(len(appEvents)).To(Equal(2))

			Expect(appEvents[0].EventType).To(Equal(filters.EtModified))
			Expect(appEvents[1].EventType).To(Equal(filters.EtScaled))

			from = appEvents[1].Timestamp
			to, _ = time.Parse(time.RFC3339, "2017-03-01T23:59:00+04:00")
			appEvents, err = filter.GetEventsForAllAppsInSpace(from, false)
			Expect(err).Should(BeNil())
			Expect(len(appEvents)).To(Equal(3))

			Expect(appEvents[0].EventType).To(Equal(filters.EtRouteDeleted))
			Expect(appEvents[1].EventType).To(Equal(filters.EtRouteDeleted))
			Expect(appEvents[2].EventType).To(Equal(filters.EtDeleted))
		})
	})
})

var testEvents = map[string]cfapi.CfEvent{
	"11efe132-b21f-4471-b625-8fee4561a296": {
		GUID: "11efe132-b21f-4471-b625-8fee4561a296",
		Name: "terraform",
		Type: "route",
		EventList: []models.EventFields{
			{
				GUID:        "03853794-d29d-4206-9eca-a4c9286599ee",
				Name:        "audit.route.create",
				Timestamp:   time.Unix(1483474464, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
		},
	},
	"7a3c27d8-c7ee-4e77-b327-961622e8b48c": {
		GUID: "7a3c27d8-c7ee-4e77-b327-961622e8b48c",
		Name: "spring-music-dopy-regma",
		Type: "route",
		EventList: []models.EventFields{
			{
				GUID:        "1983c818-fe25-41a1-bfb5-d80938157ffb",
				Name:        "audit.route.create",
				Timestamp:   time.Unix(1488204489, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
		},
	},
	"d9d8b1c8-42a7-4bdf-b337-512232c653ca": {
		GUID: "d9d8b1c8-42a7-4bdf-b337-512232c653ca",
		Name: "spring-music",
		Type: "app",
		EventList: []models.EventFields{
			{
				GUID:        "90568ca7-4856-4997-93f3-ad8000a6afc9",
				Name:        "audit.app.create",
				Timestamp:   time.Unix(1488352272, 0),
				Description: "instances: 2, memory: 512, state: STOPPED, environment_json: PRIVATE DATA HIDDEN",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "cfd5c30d-12d2-4790-9f3d-cb608ca2409a",
				Name:        "audit.app.map-route",
				Timestamp:   time.Unix(1488352278, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "0efd14ef-b2ca-41f4-99b8-37d515dcd2af",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488352278, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "2dc90133-8c85-4cc9-a895-85778a75cab7",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488352302, 0),
				Description: "state: STARTED",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "0d658986-9d24-4a0f-a864-5598df13e487",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488352524, 0),
				Description: "instances: 1, memory: 1024, environment_json: PRIVATE DATA HIDDEN",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "b3d9942a-bd3f-46b8-a665-d420d8f6bd61",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488352551, 0),
				Description: "state: STOPPED",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "fe390d28-4f85-48af-acc0-7918977b3852",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488352556, 0),
				Description: "state: STARTED",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "cf29271a-a6fd-48f0-b275-cf15a23f9e68",
				Name:        "audit.app.process.crash",
				Timestamp:   time.Unix(1488352681, 0),
				Description: "index: 0, reason: CRASHED, exit_description: Downloading failed",
				Actor:       "d9d8b1c8-42a7-4bdf-b337-512232c653ca",
				ActorName:   "web",
			},
			{
				GUID:        "a01f015b-34e4-419d-af72-7fd241cdd0b9",
				Name:        "app.crash",
				Timestamp:   time.Unix(1488352681, 0),
				Description: "index: 0, reason: CRASHED, exit_description: Downloading failed",
				Actor:       "d9d8b1c8-42a7-4bdf-b337-512232c653ca",
				ActorName:   "spring-music",
			},
			{
				GUID:        "ca43ce1f-a42a-430c-a003-d67fb8487751",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488383293, 0),
				Description: "instances: 1, memory: 512, environment_json: PRIVATE DATA HIDDEN",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "6836b96d-d207-4e24-9122-76e003e16184",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488383315, 0),
				Description: "state: STOPPED",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "08a582fd-8497-4b4c-bd5b-66e708989e31",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488383319, 0),
				Description: "state: STARTED",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "95c17f05-69a0-407a-9029-73c1edea6719",
				Name:        "audit.app.map-route",
				Timestamp:   time.Unix(1488388434, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "be6533b7-503b-4b61-b306-fcd5369ad6fa",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488388434, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "88e3de37-2f40-4c64-b671-2bd463ddbb5b",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488388898, 0),
				Description: "instances: 3",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "c2dba054-59c3-4dfc-a664-af8c4f8024d9",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488388992, 0),
				Description: "memory: 1024",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "b1bf4caa-8ef2-40a7-afbd-6ae7a2ed8990",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488388994, 0),
				Description: "state: STOPPED",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "b8d12b43-6a68-479d-8a26-f67962d8ab4a",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488388996, 0),
				Description: "state: STARTED",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "887fb14c-8874-4574-8979-fb12a38c5b1e",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488393111, 0),
				Description: "instances: 1, memory: 512, environment_json: PRIVATE DATA HIDDEN",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "e0c809f0-6be3-4d2d-823a-955816844f21",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488393129, 0),
				Description: "state: STOPPED",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "11ed0b4a-8c25-4e22-9fbd-670e24faf4c8",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488393132, 0),
				Description: "state: STARTED",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "e156bff0-2cc1-48a2-b4b2-c9d28b6be4a7",
				Name:        "audit.app.process.crash",
				Timestamp:   time.Unix(1488393215, 0),
				Description: "index: 0, reason: CRASHED, exit_description: Downloading failed",
				Actor:       "d9d8b1c8-42a7-4bdf-b337-512232c653ca",
				ActorName:   "web",
			},
			{
				GUID:        "fb746642-414b-4dc6-9964-72a4b35b9b26",
				Name:        "app.crash",
				Timestamp:   time.Unix(1488393215, 0),
				Description: "index: 0, reason: CRASHED, exit_description: Downloading failed",
				Actor:       "d9d8b1c8-42a7-4bdf-b337-512232c653ca",
				ActorName:   "spring-music",
			},
			{
				GUID:        "1d0e3dff-7525-43ab-a851-f4fbc2f94bb4",
				Name:        "audit.app.update",
				Timestamp:   time.Unix(1488393693, 0),
				Description: "instances: 1",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "63c8a61f-ebe1-46d6-8f22-8dc13a804c1e",
				Name:        "audit.app.package.delete",
				Timestamp:   time.Unix(1488393793, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "1ee7cf46-8dd3-43fe-9041-c80cce8cd797",
				Name:        "audit.app.package.delete",
				Timestamp:   time.Unix(1488393793, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "a1611017-7bff-4aa2-9388-7d50493a06e3",
				Name:        "audit.app.package.delete",
				Timestamp:   time.Unix(1488393793, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "cab729d1-c7ab-44b0-afa9-c7e9bcb6d7f0",
				Name:        "audit.app.package.delete",
				Timestamp:   time.Unix(1488393793, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "2c5a297e-7348-4df0-bd9e-2289fdd3da15",
				Name:        "audit.app.droplet.delete",
				Timestamp:   time.Unix(1488393793, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "e913de60-6342-4029-944e-f20a7773aad3",
				Name:        "audit.app.droplet.delete",
				Timestamp:   time.Unix(1488393793, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "5ec15786-b92f-41f3-a18e-2a0e0038bdb0",
				Name:        "audit.app.droplet.delete",
				Timestamp:   time.Unix(1488393793, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "a1949121-3cb9-48cc-b12d-29bf9664d842",
				Name:        "audit.app.droplet.delete",
				Timestamp:   time.Unix(1488393793, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "af6b86a2-e47d-45c3-a120-1335212e92dc",
				Name:        "audit.app.process.delete",
				Timestamp:   time.Unix(1488393793, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "8bbe98cc-da7d-4e8d-9d89-5e4ca6e7fa93",
				Name:        "audit.app.unmap-route",
				Timestamp:   time.Unix(1488393793, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "04beee77-62a9-4b53-a629-c1dba66f8958",
				Name:        "audit.app.unmap-route",
				Timestamp:   time.Unix(1488393793, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
			{
				GUID:        "3132f90f-93d8-45b3-99c5-6928df317ff8",
				Name:        "audit.app.delete-request",
				Timestamp:   time.Unix(1488393793, 0),
				Description: "recursive: true",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
		},
	},
	"c5c0237b-e4a1-4622-8e83-7dcc2455e7e6": {
		GUID: "c5c0237b-e4a1-4622-8e83-7dcc2455e7e6",
		Name: "spring-music-telegonic-pregeneration",
		Type: "route",
		EventList: []models.EventFields{
			{
				GUID:        "e575e1e4-8a38-4b54-aa1f-ee499fab7f91",
				Name:        "audit.route.create",
				Timestamp:   time.Unix(1488352276, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
		},
	},
	"111406c8-fda4-4556-b7b5-267dddcdd90f": {
		GUID: "111406c8-fda4-4556-b7b5-267dddcdd90f",
		Name: "smmks",
		Type: "route",
		EventList: []models.EventFields{
			{
				GUID:        "e6e2321d-835e-4e88-9adf-3e1af9524e51",
				Name:        "audit.route.create",
				Timestamp:   time.Unix(1488388433, 0),
				Description: "",
				Actor:       "5b3fedc4-22e9-4276-85e1-f16d60330adc",
				ActorName:   "msamaratunga@pivotal.io",
			},
		},
	},
}
