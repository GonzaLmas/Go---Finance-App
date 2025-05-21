package main

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"log"
	"net/http"
)

func ArgDolarHistorico() {
	argDolarHist, err := getArgDolarHistorico()
	if err != nil {
		log.Fatal("Error obteniendo datos de la API: ", err)
	}

	fmt.Printf("=== HISTÓRICO DÓLAR ===\n")
	for _, dolar := range argDolarHist {
		fmt.Printf("%s: Compra: $%.2f  Venta: $%.2f\n", dolar.Casa, dolar.Compra, dolar.Venta)
	}
}

func getArgDolarHistorico() ([]model.ArgDolar, error) {
	resp, err := http.Get("https://api.argentinadatos.com/v1/cotizaciones/dolares/")
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la API: %w", err)
	}
	defer resp.Body.Close()

	var argDolarHist []model.ArgDolar

	if err := json.NewDecoder(resp.Body).Decode(&argDolarHist); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %w", err)
	}

	return argDolarHist, nil
}
