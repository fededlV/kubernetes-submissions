package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	countFile := os.Getenv("COUNT_FILE")
	if countFile == "" {
		countFile = "/usr/src/app/files/pingpong.txt"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	var mu sync.Mutex
	counter := 0
	if data, err := os.ReadFile(countFile); err == nil {
		counter, _ = strconv.Atoi(strings.TrimSpace(string(data)))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /pingpong", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		n := counter
		counter++
		if err := os.WriteFile(countFile, []byte(strconv.Itoa(counter)), 0o644); err != nil {
			log.Printf("failed to write count file: %v", err)
		}
		mu.Unlock()
		fmt.Fprintf(w, "pong %d", n)
	})

	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
