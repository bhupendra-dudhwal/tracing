package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhupendra-dudhwal/tracing"
)

func main() {
	http.Handle("/step", tracing.LoggingMiddleware("service3", http.HandlerFunc(handleStep)))
	log.Println("[service3] ✅ Listening on :4003")
	http.ListenAndServe(":4003", nil)
}

func handleStep(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Service3 step complete")
}
