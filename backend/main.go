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
	http.HandleFunc("/usaacciones", handler.HandlerUsaActions)
	http.HandleFunc("/cripto", handler.HandlerCripto)
	http.HandleFunc("/dolar", handler.HandlerDolar)
	http.HandleFunc("/dolar_historico", handler.HandlerDolarHistorico)

	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)

	log.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
