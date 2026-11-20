package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// (A) Same env-var pattern as the writer, plus the port.
	filePath := os.Getenv("LOG_FILE")
	if filePath == "" {
		filePath = "/usr/src/app/files/output.log"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// (B) Register a handler for GET requests.
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// (C) Read the WHOLE file on EVERY request.
		data, err := os.ReadFile(filePath)
		if err != nil {
			// (D) The file may not exist yet.
			log.Printf("read error: %v", err)
			http.Error(w, "log file not available yet", http.StatusServiceUnavailable)
			return
		}
		// (E) Send the content back as plain text.
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write(data)
	})

	log.Printf("log-reader listening on :%s, reading %s", port, filePath)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}