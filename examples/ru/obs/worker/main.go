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
	// Метрики SDK → OpenTelemetry → Prometheus на :9464/metrics.
	exporter, err := otelprom.New()
	if err != nil {
		log.Fatalln(err)
	}
	meterProvider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(exporter))
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println(http.ListenAndServe(":9464", nil))
	}()

	// Трейсинг: спаны на старт и выполнение workflow и activities.
	// В проде провайдеру добавляют экспортёр в Jaeger, Tempo или облачный APM.
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
		Interceptors: []interceptor.ClientInterceptor{tracing}, // воркеры из этого клиента тоже получат трейсинг
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
