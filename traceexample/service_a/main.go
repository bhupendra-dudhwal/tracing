package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bhupendra-dudhwal/tracing"
)

func main() {
	http.Handle("/start", tracing.LoggingMiddleware("service1", http.HandlerFunc(handleStart)))
	log.Println("[service1] ✅ Listening on :4001")
	http.ListenAndServe(":4001", nil)
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest("GET", "http://localhost:4002/work", nil)
	req.Header.Set("X-Trace-ID", r.Header.Get("X-Trace-ID"))
	req.Header.Set("X-Span-ID", r.Header.Get("X-Span-ID"))

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[service1] ❌ Error calling service2: %v\n", err)
		http.Error(w, "service2 failed", 500)
		return
	}
	defer resp.Body.Close()

	fmt.Fprintln(w, "Service1 completed call to Service2")
}
