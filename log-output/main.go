package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.000Z07:00"

type status struct {
	Timestamp string `json:"timestamp"`
	ID        string `json:"id"`
}

func generateUUID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	// configurar versión 4 y variante RFC 4122
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func now() string {
	return time.Now().UTC().Format(timeLayout)
}

func main() {
	// El string aleatorio se guarda en memoria durante toda la vida del proceso.
	instanceID := generateUUID()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Goroutine que sigue imprimiendo en los logs cada 5 segundos.
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			fmt.Printf("%s: %s\n", now(), instanceID)
			<-ticker.C
		}
	}()

	statusHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status{Timestamp: now(), ID: instanceID})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/status", statusHandler)
	mux.HandleFunc("/", statusHandler)

	log.Printf("Server started in port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}