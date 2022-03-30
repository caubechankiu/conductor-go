package metrics

import (
	"net/http"
	"strconv"

	"github.com/netflix/conductor/client/go/settings"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ProvideMetrics(metricsSettings *settings.MetricsSettings) {
	if metricsSettings == nil {
		metricsSettings = settings.NewDefaultMetricsSettings()
	}
	http.Handle(
		metricsSettings.ApiEndpoint,
		promhttp.Handler(),
	)
	portString := strconv.Itoa(metricsSettings.Port)
	http.ListenAndServe(":"+portString, nil)
}