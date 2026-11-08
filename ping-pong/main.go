package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Contador en memoria: se reinicia si el pod se reinicia.
	var counter atomic.Int64

	mux := http.NewServeMux()
	mux.HandleFunc("GET /pingpong", func(w http.ResponseWriter, r *http.Request) {
		// Add devuelve el valor nuevo; restamos 1 para que la primera respuesta sea "pong 0".
		n := counter.Add(1) - 1
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "pong %d", n)
	})

	log.Printf("Server started in port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}