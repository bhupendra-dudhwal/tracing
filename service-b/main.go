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

		url := "http://service-c:8003"
		fmt.Println("url - ", url)
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		fmt.Printf("\n req - %+v\nerr- %+v", req, err)
		res, err := http.DefaultClient.Do(req)
		fmt.Printf("\n res - %+v\nerr- %+v", res, err)
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
