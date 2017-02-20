package main

import (
	"fmt"
	"os"

	"github.com/mevansam/cf-cli-api/cli"
)

func main() {

	sessionProvider := cli.NewCfCliSessionProvider()
	logger := cli.NewLogger(true, "true")

	session, err := sessionProvider.NewCfSession(
		"https://api.local.pcfdev.io",
		"admin", "admin",
		"pcfdev-org", "pcfdev-space",
		true, logger)

	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}

	logger.DebugMessage("Session: %# v", session)
}
