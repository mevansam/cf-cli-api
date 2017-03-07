package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mevansam/cf-cli-api/cfapi"
)

func main() {

	sessionProvider := cfapi.NewCfCliSessionProvider()
	logger := cfapi.NewLogger(true, "true")

	session, err := sessionProvider.NewCfSession(
		"https://api.run.pivotal.io",
		"msamaratunga@pivotal.io", "!1mksKLD",
		"pcfp", "mevan-dev-2",
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

	events, err := session.GetAllEventsForApp("19b9d70b-6ebe-47d7-9313-f0c213445036", from)
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}

	logger.DebugMessage("Events: %# v", events)
}
