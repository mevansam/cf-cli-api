package filters_test

import (
	"fmt"
	"time"

	"github.com/mevansam/cf-cli-api/filters"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Application Event Type Tests", func() {

	Context("Application Events", func() {

		It("Should serialize an AppEvent type to a string", func() {

			timestamp, _ := time.Parse(time.RFC3339, "2017-03-01T00:00:00+04:00")

			ae := filters.AppEvent{
				SourceGUID: "19b9d70b-6ebe-47d7-9313-f0c213445036",
				SourceName: "some_test_app",
				SourceType: "app",
				EventType:  filters.EtCreated,
				Timestamp:  timestamp,
			}
			Expect(fmt.Sprintf("%s", ae)).To(Equal("19b9d70b-6ebe-47d7-9313-f0c213445036|some_test_app|app|created|2017-03-01T00:00:00+04:00"))
		})
		It("Should deserialize an AppEvent from a string", func() {

			timestamp, _ := time.Parse(time.RFC3339, "2017-03-01T00:00:00+04:00")

			ae, err := filters.NewAppEvent("19b9d70b-6ebe-47d7-9313-f0c213445036|some_test_app|app|created|2017-03-01T00:00:00+04:00")
			Expect(err).Should(BeNil())
			Expect(ae.SourceGUID).To(Equal("19b9d70b-6ebe-47d7-9313-f0c213445036"))
			Expect(ae.SourceName).To(Equal("some_test_app"))
			Expect(ae.SourceType).To(Equal("app"))
			Expect(ae.EventType).To(Equal(filters.EtCreated))
			Expect(ae.Timestamp).To(Equal(timestamp))
		})
		It("Should fail to deserialize an AppEvent with an number of fields", func() {
			_, err := filters.NewAppEvent("19b9d70b-6ebe-47d7-9313-f0c213445036|some_test_app|app|2017-03-01T00:00:00+04:00")
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).To(Equal("The string data should have 5 or more fields. '4' fields were extracted."))
		})
		It("Should fail to deserialize an AppEvent with an invalid AppID", func() {
			_, err := filters.NewAppEvent("1234|some_test_app|app|created|2017-03-01T00:00:00+04:00")
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).To(Equal("The app GUID '1234' is not a valid GUID."))
		})
		It("Should fail to deserialize an AppEvent with an invalid Event type", func() {
			_, err := filters.NewAppEvent("19b9d70b-6ebe-47d7-9313-f0c213445036|some_test_app|app|nonevent|2017-03-01T00:00:00+04:00")
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).To(Equal("Event type 'nonevent' is not valid."))
		})
		It("Should fail to deserialize an AppEvent with an invalid Timestamp", func() {
			_, err := filters.NewAppEvent("19b9d70b-6ebe-47d7-9313-f0c213445036|some_test_app|app|created|01/03/2017")
			Expect(err).ShouldNot(BeNil())
		})
	})
})
