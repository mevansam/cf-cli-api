package cfapi

import (
	"os"

	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/api/applicationbits"
	"code.cloudfoundry.org/cli/cf/api/applications"
	"code.cloudfoundry.org/cli/cf/api/organizations"
	"code.cloudfoundry.org/cli/cf/api/spaces"
	"code.cloudfoundry.org/cli/cf/models"
)

// CfSessionProvider -
type CfSessionProvider interface {
	NewCfSession(
		apiEndPoint string,
		userName string,
		password string,
		orgName string,
		spaceName string,
		sslDisabled bool,
		logger *Logger) (cfSession CfSession, err error)

	NewCfSessionFromFilepath(
		configPath string,
		sslDisabled bool,
		logger *Logger) (cfSession CfSession, err error)
}

// CfSession -
type CfSession interface {
	Close()

	HasTarget() bool

	GetSessionUsername() string
	GetSessionOrg() models.OrganizationFields
	SetSessionOrg(models.OrganizationFields)
	GetSessionSpace() models.SpaceFields
	SetSessionSpace(models.SpaceFields)

	// Cloud Countroller APIs

	Organizations() organizations.OrganizationRepository
	Spaces() spaces.SpaceRepository

	Services() api.ServiceRepository
	ServicePlans() api.ServicePlanRepository
	ServiceSummary() api.ServiceSummaryRepository
	UserProvidedServices() api.UserProvidedServiceInstanceRepository
	ServiceKeys() api.ServiceKeyRepository
	ServiceBindings() api.ServiceBindingRepository

	AppSummary() api.AppSummaryRepository
	Applications() applications.Repository
	ApplicationBits() applicationbits.Repository
	Routes() api.RouteRepository
	Domains() api.DomainRepository

	GetServiceCredentials(models.ServiceBindingFields) (*ServiceBindingDetail, error)

	DownloadAppContent(ppGUID string, outputFile *os.File, asDroplet bool) error
	UploadDroplet(appGUID string, contentType string, dropletUploadRequest *os.File) error
}
