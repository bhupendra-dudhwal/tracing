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
		ctx := r.Context()
		_, span := tracing.Tracer.Start(ctx, "service-a-handler")
		defer span.End()

		url := "http://service-b:8002"
		fmt.Println("url - ", url)
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		fmt.Printf("\n req - %+v\nerr- %+v", req, err)
		res, err := http.DefaultClient.Do(req)
		fmt.Printf("\n res - %+v\nerr- %+v", res, err)
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
