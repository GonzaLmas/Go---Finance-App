package main

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"log"
	"net/http"
)

func ArgLetras() {
	letras, err := getArgLetrasExtApi()
	if err != nil {
		log.Fatal("Error obteniendo datos de la API: ", err)
	}

	fmt.Printf("=== LETRAS ===\n")
	for _, letra := range letras {
		fmt.Printf("%-6s $%.2f\n", letra.Symbol, letra.LastPrice)
	}
	fmt.Printf("\nTotal: %d letras\n", len(letras))
}

func getArgLetrasExtApi() ([]model.Symbol, error) {
	resp, err := http.Get("https://data912.com/live/arg_notes")
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la API: %w", err)
	}
	defer resp.Body.Close()

	var letras []model.Symbol
	if err := json.NewDecoder(resp.Body).Decode(&letras); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %w", err)
	}

	return letras, nil
}
