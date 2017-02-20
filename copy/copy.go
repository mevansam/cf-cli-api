package copy

import (
	"code.cloudfoundry.org/cli/cf/models"
	"github.com/mevansam/cf-cli-api/cli"
)

// ApplicationsManager -
type ApplicationsManager interface {
	Init(srcCCSession cli.CfSession,
		destCCSession cli.CfSession,
		logger *cli.Logger) error

	ApplicationsToBeCopied(appNames []string, copyAsDroplet bool) (ApplicationCollection, error)
	DoCopy(applications ApplicationCollection, services ServiceCollection, appHostFormat string, appRouteDomain string) error
	Close()
}

// ApplicationCollection -
type ApplicationCollection interface {
}

// ServicesManager -
type ServicesManager interface {
	Init(srcCCSession cli.CfSession,
		destCCSession cli.CfSession,
		destTarget, destOrg, destSpace string,
		logger *cli.Logger) error

	ServicesToBeCopied(appNames []string, upsServices []string) (ServiceCollection, error)
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
	Download(session cli.CfSession) error
	Upload(session cli.CfSession, params models.AppParams) (models.Application, error)
}
