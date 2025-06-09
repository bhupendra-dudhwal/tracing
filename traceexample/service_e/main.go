package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhupendra-dudhwal/tracing"
)

func main() {
	http.Handle("/final", tracing.LoggingMiddleware("service5", http.HandlerFunc(handleFinal)))
	log.Println("[service5] ✅ Listening on :4005")
	http.ListenAndServe(":4005", nil)
}

func handleFinal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Service5 final step complete")
}
