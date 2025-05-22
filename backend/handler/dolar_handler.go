package handler

import (
	"encoding/json"
	"finance-app/functions"
	"log"
	"net/http"
)

func HandlerDolar(w http.ResponseWriter, r *http.Request) {
	argDolars, err := functions.GetArgDolar()
	if err != nil {
		log.Fatal("Error obteniendo datos de la API: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(argDolars)
}
