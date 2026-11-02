package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/contrib/opentelemetry"
	"go.temporal.io/sdk/interceptor"
	tlog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"

	"tcourse/hello"
)

func main() {
	// SDK metrics → OpenTelemetry → Prometheus on :9464/metrics.
	exporter, err := otelprom.New()
	if err != nil {
		log.Fatalln(err)
	}
	meterProvider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(exporter))
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println(http.ListenAndServe(":9464", nil))
	}()

	// Tracing: spans for starting and running workflows and activities.
	// In production the provider gets an exporter to Jaeger, Tempo or a cloud APM.
	tracerProvider := sdktrace.NewTracerProvider()
	tracing, err := opentelemetry.NewTracingInterceptor(opentelemetry.TracerOptions{
		Tracer: tracerProvider.Tracer("tcourse"),
	})
	if err != nil {
		log.Fatalln(err)
	}

	c, err := client.Dial(client.Options{
		Logger: tlog.NewStructuredLogger(slog.New(slog.NewJSONHandler(os.Stdout, nil))),
		MetricsHandler: opentelemetry.NewMetricsHandler(opentelemetry.MetricsHandlerOptions{
			Meter: meterProvider.Meter("temporal-sdk"),
		}),
		Interceptors: []interceptor.ClientInterceptor{tracing}, // workers created from this client get tracing too
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	w := worker.New(c, hello.TaskQueue, worker.Options{})
	w.RegisterWorkflow(hello.GreetingWorkflow)
	w.RegisterActivity(hello.Greet)
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln(err)
	}
}
