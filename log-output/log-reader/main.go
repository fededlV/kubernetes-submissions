package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// (A) Same env-var pattern as the writer, plus the port.
	filePath := os.Getenv("LOG_FILE")
	if filePath == "" {
		filePath = "/usr/src/app/files/output.log"
	}
	countFile := os.Getenv("COUNT_FILE")
	if countFile == "" {
		countFile = "/usr/src/app/files/pingpong.txt"
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

		lines := strings.Split(strings.TrimSpace(string(data)), "\n")
		last := lines[len(lines)-1]

		count := "0" // if ping-pong hasn't received requests yet, the file doesn't exist
		if c, err := os.ReadFile(countFile); err == nil {
			count = strings.TrimSpace(string(c))
		}

		// (E) Send the content back as plain text.
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "%s.\nPing / Pongs: %s\n", last, count)
	})

	log.Printf("log-reader listening on :%s, reading %s", port, filePath)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}