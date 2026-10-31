package main

import (
    "crypto/rand"
    "fmt"
    "time"
)

func generateUUID() string {
    b := make([]byte, 16)
    if _, err := rand.Read(b); err != nil {
        panic(err)
    }
    // configurar versió 4 y variante RFC 4122
    b[6] = (b[6] & 0x0f) | 0x40
    b[8] = (b[8] & 0x3f) | 0x80
    return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func main() {
    instanceID := generateUUID()
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for {
        now := time.Now().UTC().Format("2006-01-02T15:04:05.000Z07:00")
        fmt.Printf("%s: %s\n", now, instanceID)
        <-ticker.C
    }

    
}