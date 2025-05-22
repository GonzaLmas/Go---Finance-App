package main

import (
	"finance-app/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cedears", handler.HandlerCedears)
	http.HandleFunc("/adr", handler.HandlerAdr)
	http.HandleFunc("/letras", handler.HandlerLetras)

	// Servir archivos estáticos (HTML, JS, CSS desde carpeta frontend)
	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)

	// Levantar el servidor
	log.Println("Servidor iniciado en http://localhost:5500")
	log.Fatal(http.ListenAndServe(":5500", nil))
}
