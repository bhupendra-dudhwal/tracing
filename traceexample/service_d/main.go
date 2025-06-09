package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bhupendra-dudhwal/tracing"
)

func main() {
	http.Handle("/step", tracing.LoggingMiddleware("service4", http.HandlerFunc(handleStep)))
	log.Println("[service4] ✅ Listening on :4004")
	http.ListenAndServe(":4004", nil)
}

func handleStep(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest("GET", "http://localhost:4005/final", nil)
	req.Header.Set("X-Trace-ID", r.Header.Get("X-Trace-ID"))
	req.Header.Set("X-Span-ID", r.Header.Get("X-Span-ID"))
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[service4] ❌ Error calling service5: %v\n", err)
		http.Error(w, "service5 failed", 500)
		return
	}
	defer resp.Body.Close()

	fmt.Fprintln(w, "Service4 step (and Service5 call) complete")
}
