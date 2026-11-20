package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func generateUUID() string { // same function you already have
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func main() {
	// (A) Where to write. Read from an env var, with a default.
	filePath := os.Getenv("LOG_FILE")
	if filePath == "" {
		filePath = "/usr/src/app/files/output.log"
	}

	// (B) Make sure the folder exists.
	if err := os.MkdirAll(filepath.Dir(filePath), 0o755); err != nil {
		log.Fatalf("cannot create directory: %v", err)
	}

	// (C) Open the file in append mode.
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatalf("cannot open file: %v", err)
	}
	defer f.Close()

	// (D) Random string generated ONCE, at startup.
	randomString := generateUUID()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		now := time.Now().UTC().Format("2006-01-02T15:04:05.000Z07:00")
		line := fmt.Sprintf("%s: %s\n", now, randomString)

		// (E) Write the line to the file AND print it to stdout.
		if _, err := f.WriteString(line); err != nil {
			log.Printf("write error: %v", err)
		}
		fmt.Print(line)

		<-ticker.C
	}
}