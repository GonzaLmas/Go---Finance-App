package handler

import (
	"encoding/json"
	"finance-app/functions"
	"log"
	"net/http"
)

func HandlerAdr(w http.ResponseWriter, r *http.Request) {
	favsArgAccions := []string{"BBAR", "BMA", "GGAL", "LOMA", "PAM", "YPF"}

	accions, err := functions.GetArgAccionsExtApi()
	if err != nil {
		log.Fatal("Error obteniendo datos de la API: ", err)
	}

	filteredAdr := functions.FilterArgAccions(accions, favsArgAccions)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredAdr)
}
