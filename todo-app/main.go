package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	// 1. Leer el puerto desde la variable de entorno, con valor por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// 2. Definir las rutas
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Todo app running")
	})

	// 3. Abrir el puerto ANTES de anunciar que el servidor arrancó
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("could not listen on port %s: %v", port, err)
	}

	fmt.Printf("Server started in port %s\n", port)

	// 4. Servir peticiones (bloquea para siempre)
	if err := http.Serve(listener, mux); err != nil {
		log.Fatal(err)
	}
}