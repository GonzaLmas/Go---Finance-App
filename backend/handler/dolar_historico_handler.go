package handler

import (
	"encoding/json"
	"finance-app/functions"
	"log"
	"net/http"
)

func HandlerDolarHistorico(w http.ResponseWriter, r *http.Request) {
	argDolarHist, err := functions.GetArgDolarHistorico()
	if err != nil {
		log.Fatal("Error obteniendo datos de la API: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(argDolarHist)
}
