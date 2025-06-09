package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/bhupendra-dudhwal/tracing"
)

func main() {
	http.Handle("/work", tracing.LoggingMiddleware("service2", http.HandlerFunc(handleWork)))
	log.Println("[service2] ✅ Listening on :4002")
	http.ListenAndServe(":4002", nil)
}

func callService(url, traceID, parentSpanID string, wg *sync.WaitGroup, serviceName string) {
	defer wg.Done()
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Trace-ID", traceID)
	req.Header.Set("X-Span-ID", parentSpanID)

	start := time.Now()
	log.Printf("[service2 -> %s] --> Request | TraceID=%s | ParentSpanID=%s | URL=%s", serviceName, traceID, parentSpanID, url)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[service2 -> %s] ❌ Error: %v", serviceName, err)
		return
	}
	defer resp.Body.Close()
	io.ReadAll(resp.Body) // consume body
	duration := time.Since(start)
	log.Printf("[service2 -> %s] <-- Response | TraceID=%s | Duration=%s", serviceName, traceID, duration)
}

func handleWork(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	traceID := r.Header.Get("X-Trace-ID")
	parentSpanID := r.Header.Get("X-Span-ID")

	wg.Add(2)
	go callService("http://localhost:4003/step", traceID, parentSpanID, &wg, "service3")
	go callService("http://localhost:4004/step", traceID, parentSpanID, &wg, "service4")

	wg.Wait()
	fmt.Fprintln(w, "Service2 finished parallel calls to Service3 and Service4")
}
