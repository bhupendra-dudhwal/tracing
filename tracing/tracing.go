package tracing

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"

	"math/rand"
)

var Tracer trace.Tracer

func InitTracer(serviceName string) func() {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://jaeger:14268/api/traces")))
	if err != nil {
		log.Fatalf("Failed to create Jaeger exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		)),
	)

	otel.SetTracerProvider(tp)
	Tracer = tp.Tracer(serviceName)

	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}
}

func generateID() string {
	return fmt.Sprintf("%x", rand.Int63())
}

func LoggingMiddleware(serviceName string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get("X-Trace-ID")
		if traceID == "" {
			traceID = generateID()
		}

		parentSpanID := r.Header.Get("X-Span-ID")
		spanID := generateID()

		start := time.Now()
		r.Header.Set("X-Trace-ID", traceID)
		r.Header.Set("X-Parent-Span-ID", parentSpanID)
		r.Header.Set("X-Span-ID", spanID)

		log.Printf("[%s] --> Request | TraceID=%s | ParentSpanID=%s | SpanID=%s | Method=%s | URL=%s",
			serviceName, traceID, parentSpanID, spanID, r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("[%s] <-- Response | TraceID=%s | SpanID=%s | Duration=%s",
			serviceName, traceID, spanID, duration)
	})
}
