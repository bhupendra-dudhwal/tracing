package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhupendra-dudhwal/tracing"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	shutdown := tracing.InitTracer("service-c")
	defer shutdown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, span := tracing.Tracer.Start(ctx, "service-c-handler")
		defer span.End()

		log.Printf("[Service C] TraceID: %s", span.SpanContext().TraceID())

		fmt.Fprintln(w, "Handled by Service C")
	}

	mux := http.NewServeMux()
	mux.Handle("/", otelhttp.NewHandler(http.HandlerFunc(handler), "service-c-root"))

	log.Println("Service C running on :8003")
	log.Fatal(http.ListenAndServe(":8003", mux))
}
