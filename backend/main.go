package main

import (
	"finance-app/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/cedears", handler.HandlerCedears)
	http.HandleFunc("/adr", handler.HandlerAdr)
	http.HandleFunc("/letras", handler.HandlerLetras)
	http.HandleFunc("/usaacciones", handler.HandlerUsaActions)
	http.HandleFunc("/cripto", handler.HandlerCripto)
	http.HandleFunc("/dolar", handler.HandlerDolar)

	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback local
	}

	log.Println("Servidor iniciado en http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
