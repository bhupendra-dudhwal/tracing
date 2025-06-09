package middleware

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

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
