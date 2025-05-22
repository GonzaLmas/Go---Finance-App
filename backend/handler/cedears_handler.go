package handler

import (
	"encoding/json"
	"finance-app/functions"
	"net/http"
)

func HandlerCedears(w http.ResponseWriter, r *http.Request) {
	favsCedears := []string{"AAPL", "AMZN", "GOOGL", "MELI", "META", "MSFT", "NVDA", "TSLA", "QQQ", "SPY"}

	cedears, err := functions.GetCedearsExtApi()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filtered := functions.FilterCedears(cedears, favsCedears)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}
