package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mevansam/cf-cli-api/cfapi"
	"github.com/mevansam/cf-cli-api/filters"
	"github.com/mevansam/cf-cli-api/utils"
)

func main() {

	var (
		err error
	)

	sessionProvider := cfapi.NewCfCliSessionProvider()
	logger := cfapi.NewLogger(true, "true")

	session, err := sessionProvider.NewCfSession(
		"https://api.local.pcfdev.io",
		"admin", "admin",
		"pcfdev-org", "dev1",
		true, logger)

	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}

	from, err := time.Parse("2006-Jan-02", "2010-Jan-01")
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}

	apps, err := session.AppSummary().GetSummariesInCurrentSpace()
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}

	if app, exists := utils.ContainsApp("app1", apps); exists {

		events, err := session.GetAllEventsForApp(app.GUID, from, false)
		if err != nil {
			fmt.Printf("ERROR: %s", err.Error())
			os.Exit(1)
		}

		logger.DebugMessage("Events for app '%s/%s': %# v", app.Name, app.GUID, events)

		filter := filters.NewAppEventFilter(session)
		appEvents, err := filter.GetEventsForApp(app.GUID, from, false)
		if err != nil {
			fmt.Printf("ERROR: %s", err.Error())
			os.Exit(1)
		}

		logger.DebugMessage("Application Events for app '%s/%s': %# v", app.Name, app.GUID, appEvents)
	}
}
