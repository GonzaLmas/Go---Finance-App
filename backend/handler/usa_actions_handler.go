package handler

import (
	"encoding/json"
	"finance-app/functions"
	"log"
	"net/http"
)

func HandlerUsaActions(w http.ResponseWriter, r *http.Request) {
	favsAccions := []string{"AAPL", "AMZN", "GOOGL", "MELI", "META", "MSFT", "NVDA", "TSLA", "QQQ", "SPY"}

	accions, err := functions.GetUsaAccionsExtApi()
	if err != nil {
		log.Fatal("Error obteniendo datos de la API: ", err)
	}

	filterAccions := functions.FilterUsaAccions(accions, favsAccions)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filterAccions)
}
