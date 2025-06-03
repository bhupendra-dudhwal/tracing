package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhupendra-dudhwal/tracing"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	shutdown := tracing.InitTracer("service-b")
	defer shutdown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, span := tracing.Tracer.Start(ctx, "service-b-handler")
		defer span.End()

		log.Printf("[Service B] TraceID: %s", span.SpanContext().TraceID())

		client := http.Client{
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		}

		url := "http://service-c:8003"
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

		if err != nil {
			http.Error(w, "Service B Request failed", http.StatusInternalServerError)
			return
		}
		// otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

		res, err := client.Do(req)
		if err != nil {
			http.Error(w, "Service C failed", http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		fmt.Fprintln(w, "Handled by Service B")
	}

	mux := http.NewServeMux()
	mux.Handle("/", otelhttp.NewHandler(http.HandlerFunc(handler), "service-b-root"))

	log.Println("Service B running on :8002")
	log.Fatal(http.ListenAndServe(":8002", mux))
}
