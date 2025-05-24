package handler

import (
	"encoding/json"
	"finance-app/functions"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func HandlerCripto(w http.ResponseWriter, r *http.Request) {
	favsCripto := []string{"BTC-USD", "ETH-USD"}

	criptos, err := functions.GetCriptoExtApi(favsCripto)
	if err != nil {
		http.Error(w, "Error obteniendo datos de criptomonedas", http.StatusInternalServerError)
		log.Printf("Error obteniendo datos de criptomonedas: %v", err)
		return
	}

	var response []map[string]interface{}
	for _, cripto := range criptos {
		priceFloat, err := strconv.ParseFloat(cripto.Amount, 64)
		if err != nil {
			log.Printf("Error al convertir el precio de %s: %v", cripto.Base, err)
			continue
		}

		response = append(response, map[string]interface{}{
			"symbol": fmt.Sprintf("%s", cripto.Base), // Ej: "BTC-USD"
			"price":  priceFloat,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error formateando respuesta", http.StatusInternalServerError)
		log.Printf("Error formateando respuesta: %v", err)
	}
}
