package mock_test

import (
	"os"
	"time"

	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/api/appevents"
	"code.cloudfoundry.org/cli/cf/api/applicationbits"
	"code.cloudfoundry.org/cli/cf/api/applications"
	"code.cloudfoundry.org/cli/cf/api/organizations"
	"code.cloudfoundry.org/cli/cf/api/spaces"
	"code.cloudfoundry.org/cli/cf/i18n"
	"code.cloudfoundry.org/cli/cf/models"
	"github.com/mevansam/cf-cli-api/cfapi"
)

// MockSessionProvider -
type MockSessionProvider struct {
	MockSessionMap map[string]cfapi.CfSession
}

// MockSession -
type MockSession struct {
	Logger *cfapi.Logger

	MockHasTarget            func() bool
	MockSetSessionTarget     func(string, string) error
	MockGetSessionUsername   func() string
	MockGetSessionOrg        func() models.OrganizationFields
	MockSetSessionOrg        func(models.OrganizationFields)
	MockGetSessionSpace      func() models.SpaceFields
	MockSetSessionSpace      func(models.SpaceFields)
	MockOrganizations        func() organizations.OrganizationRepository
	MockSpaces               func() spaces.SpaceRepository
	MockServices             func() api.ServiceRepository
	MockServicePlans         func() api.ServicePlanRepository
	MockServiceSummary       func() api.ServiceSummaryRepository
	MockUserProvidedServices func() api.UserProvidedServiceInstanceRepository
	MockServiceKeys          func() api.ServiceKeyRepository
	MockServiceBindings      func() api.ServiceBindingRepository
	MockAppSummary           func() api.AppSummaryRepository
	MockApplications         func() applications.Repository
	MockApplicationBits      func() applicationbits.Repository
	MockAppEvents            func() appevents.Repository
	MockRoutes               func() api.RouteRepository
	MockDomains              func() api.DomainRepository

	MockGetAllEventsInSpace func(time.Time, bool) (map[string]cfapi.CfEvent, error)
	MockGetAllEventsForApp  func(string, time.Time, bool) (cfapi.CfEvent, error)

	MockGetServiceCredentials func(models.ServiceBindingFields) (*cfapi.ServiceBindingDetail, error)
	MockDownloadAppContent    func(string, *os.File, bool) error
	MockUploadDroplet         func(string, string, *os.File) error
}

// mockLocale -
type mockLocale struct{}

// Locale -
func (l *mockLocale) Locale() string {
	return "en_us"
}

// NewCfSession -
func (p *MockSessionProvider) NewCfSession(
	apiEndPoint string,
	userName string,
	password string,
	orgName string,
	spaceName string,
	sslDisabled bool,
	logger *cfapi.Logger) (cfSession cfapi.CfSession, err error) {

	if i18n.T == nil {
		i18n.T = i18n.Init(&mockLocale{})
	}

	return &MockSession{Logger: logger}, nil
}

// NewCfSessionFromFilepath -
func (p *MockSessionProvider) NewCfSessionFromFilepath(
	configPath string,
	sslDisabled bool,
	logger *cfapi.Logger) (cfapi.CfSession, error) {

	if i18n.T == nil {
		i18n.T = i18n.Init(&mockLocale{})
	}

	return p.MockSessionMap[configPath], nil
}

// Close -
func (m *MockSession) Close() {
}

// GetSessionLogger -
func (m *MockSession) GetSessionLogger() *cfapi.Logger {
	return m.Logger
}

// HasTarget -
func (m *MockSession) HasTarget() bool {
	return m.MockHasTarget()
}

// SetSessionTarget -
func (m *MockSession) SetSessionTarget(orgName, spaceName string) error {
	return m.MockSetSessionTarget(orgName, spaceName)
}

// GetSessionUsername -
func (m *MockSession) GetSessionUsername() string {
	return m.MockGetSessionUsername()
}

// GetSessionOrg -
func (m *MockSession) GetSessionOrg() models.OrganizationFields {
	return m.MockGetSessionOrg()
}

// SetSessionOrg -
func (m *MockSession) SetSessionOrg(org models.OrganizationFields) {
	m.MockSetSessionOrg(org)
}

// GetSessionSpace -
func (m *MockSession) GetSessionSpace() models.SpaceFields {
	return m.MockGetSessionSpace()
}

// SetSessionSpace -
func (m *MockSession) SetSessionSpace(space models.SpaceFields) {
	m.MockSetSessionSpace(space)
}

// Organizations -
func (m *MockSession) Organizations() organizations.OrganizationRepository {
	return m.MockOrganizations()
}

// Spaces -
func (m *MockSession) Spaces() spaces.SpaceRepository {
	return m.MockSpaces()
}

// Services -
func (m *MockSession) Services() api.ServiceRepository {
	return m.MockServices()
}

// ServicePlans -
func (m *MockSession) ServicePlans() api.ServicePlanRepository {
	return m.MockServicePlans()
}

// ServiceSummary -
func (m *MockSession) ServiceSummary() api.ServiceSummaryRepository {
	return m.MockServiceSummary()
}

// UserProvidedServices -
func (m *MockSession) UserProvidedServices() api.UserProvidedServiceInstanceRepository {
	return m.MockUserProvidedServices()
}

// ServiceKeys -
func (m *MockSession) ServiceKeys() api.ServiceKeyRepository {
	return m.MockServiceKeys()
}

// AppSummary -
func (m *MockSession) AppSummary() api.AppSummaryRepository {
	return m.MockAppSummary()
}

// Applications -
func (m *MockSession) Applications() applications.Repository {
	return m.MockApplications()
}

// ApplicationBits -
func (m *MockSession) ApplicationBits() applicationbits.Repository {
	return m.MockApplicationBits()
}

// AppEvents -
func (m *MockSession) AppEvents() appevents.Repository {
	return m.MockAppEvents()
}

// Routes -
func (m *MockSession) Routes() api.RouteRepository {
	return m.MockRoutes()
}

// Domains -
func (m *MockSession) Domains() api.DomainRepository {
	return m.MockDomains()
}

// ServiceBindings -
func (m *MockSession) ServiceBindings() api.ServiceBindingRepository {
	return m.MockServiceBindings()
}

// GetAllEventsInSpace -
func (m *MockSession) GetAllEventsInSpace(from time.Time, inclusive bool) (events map[string]cfapi.CfEvent, err error) {
	events, err = m.MockGetAllEventsInSpace(from, inclusive)
	return
}

// GetAllEventsForApp -
func (m *MockSession) GetAllEventsForApp(appGUID string, from time.Time, inclusive bool) (cfEvent cfapi.CfEvent, err error) {
	cfEvent, err = m.MockGetAllEventsForApp(appGUID, from, inclusive)
	return
}

// GetServiceCredentials -
func (m *MockSession) GetServiceCredentials(serviceBinding models.ServiceBindingFields) (*cfapi.ServiceBindingDetail, error) {
	return m.MockGetServiceCredentials(serviceBinding)
}

// DownloadAppContent -
func (m *MockSession) DownloadAppContent(appGUID string, outputFile *os.File, asDroplet bool) error {
	return m.MockDownloadAppContent(appGUID, outputFile, asDroplet)
}

// UploadDroplet -
func (m *MockSession) UploadDroplet(appGUID string, contentType string, dropletUploadRequest *os.File) error {
	return m.MockUploadDroplet(appGUID, contentType, dropletUploadRequest)
}
