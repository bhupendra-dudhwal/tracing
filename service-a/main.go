package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhupendra-dudhwal/tracing"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	shutdown := tracing.InitTracer("service-a")
	defer shutdown()

	handler := func(w http.ResponseWriter, r *http.Request) {

		// inside handler
		ctx := r.Context()
		_, span := tracing.Tracer.Start(ctx, "service-a-handler")
		defer span.End()

		log.Printf("Service A - TraceID: %s", span.SpanContext().TraceID().String())

		client := http.Client{
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		}

		url := "http://service-b:8002"
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			http.Error(w, "Service A Request failed", http.StatusInternalServerError)
			return
		}
		// otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

		res, err := client.Do(req)
		if err != nil {
			http.Error(w, "Service B failed", http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		fmt.Fprintln(w, "Handled by Service A")
	}

	mux := http.NewServeMux()
	mux.Handle("/", otelhttp.NewHandler(http.HandlerFunc(handler), "service-a-root"))

	log.Println("Service A running on :8001")
	log.Fatal(http.ListenAndServe(":8001", mux))
}
