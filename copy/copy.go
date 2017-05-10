package copy

import (
	"code.cloudfoundry.org/cli/cf/models"
	"github.com/mevansam/cf-cli-api/cfapi"
)

// ApplicationsManager -
type ApplicationsManager interface {
	Init(srcCCSession cfapi.CfSession,
		destCCSession cfapi.CfSession,
		logger *cfapi.Logger) error

	ApplicationsToBeCopied(appNames []string, copyAsDroplet bool) (ApplicationCollection, error)
	DoCopy(applications ApplicationCollection, services ServiceCollection, appHostFormat string, appRouteDomain string) error
	Close()
}

// ApplicationCollection -
type ApplicationCollection interface {
}

// ServicesManager -
type ServicesManager interface {
	Init(srcCCSession cfapi.CfSession,
		destCCSession cfapi.CfSession,
		serviceKeyFormat string,
		logger *cfapi.Logger) error

	ServicesToBeCopied(appNames []string, siToCopyAsUpsServices []string, stToCopyAsUpsServices []string) (ServiceCollection, error)
	DoCopy(services ServiceCollection, recreate bool) error
	Close()
}

// ServiceCollection -
type ServiceCollection interface {
	AppBindings(appName string) ([]string, bool)
}

// ApplicationContentProvider -
type ApplicationContentProvider interface {
	NewApplication(srcApp *models.Application, downloadPath string, copyAsDroplet bool) ApplicationContent
}

// ApplicationContent -
type ApplicationContent interface {
	App() *models.Application
	Download(session cfapi.CfSession) error
	Upload(session cfapi.CfSession, params models.AppParams) (models.Application, error)
}
