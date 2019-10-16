package detective

import (
	"net/http"
	"os"
	"time"

	"github.com/conplementAG/k8s-semantic-detective/pkg/common/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "detective_first_metric",
		Help: "...",
	})
)

func Detect() {
	var port = "2112"
	recordMetrics()

	logging.LogSuccess("Kubernetes Semantic Detective started")
	logging.Log("Listening on " + port)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Kubernetes Semantic Detective Exporter</title></head>
             <body>
             <h1>Kubernetes Semantic Detective Exporter</h1>
             <p><a href='/metrics'>Metrics</a></p>
             </body>
             </html>`))
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logging.Logf("Error starting HTTP server %s", err)
		os.Exit(1)
	}
}
