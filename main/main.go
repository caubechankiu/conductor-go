package main

import (
	"os"

	"github.com/netflix/conductor/client/go/example"
	"github.com/netflix/conductor/client/go/metrics"
	"github.com/netflix/conductor/client/go/orkestrator"
	"github.com/netflix/conductor/client/go/settings"
	log "github.com/sirupsen/logrus"
)

//Example init function that shows how to configure logging
//Using json formatter and changing level to Debug
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	//Stdout, change to a file for production use case
	log.SetOutput(os.Stdout)

	// Set to debug for demonstration.  Change to Info for production use cases.
	log.SetLevel(log.DebugLevel)
}

func main() {
	authenticationSettings := settings.NewAuthenticationSettings(
		"keyId",
		"keySecret",
	)
	httpSettings := settings.NewHttpSettingsWithBaseUrlAndDebug(
		"https://play.orkes.io/api",
		true,
	)

	metricsCollector := metrics.NewMetricsCollector()
	go metrics.ProvideDefaultMetrics()

	workerOrkestrator := orkestrator.NewWorkerOrkestrator(
		authenticationSettings,
		httpSettings,
		metricsCollector,
		1,
		5000,
	)
	workerOrkestrator.StartWorker(
		"go_task_example",
		"",
		example.TaskExecuteFunctionExample1,
		true,
	)
	workerOrkestrator.StartWorker(
		"go_task_example",
		"",
		example.TaskExecuteFunctionExample1,
		true,
	)
}
