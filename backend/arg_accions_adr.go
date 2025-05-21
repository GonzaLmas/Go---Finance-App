package main

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"log"
	"net/http"
)

func ArgAccionsAdr() {
	favsArgAccions := []string{"BBAR", "BMA", "GGAL", "LOMA", "PAM", "YPF"}

	accions, err := getArgAccionsExtApi()
	if err != nil {
		log.Fatal("Error obteniendo datos de la API: ", err)
	}

	filterAccions := filterArgAccions(accions, favsArgAccions)

	fmt.Printf("=== ACCIONES ===\n")
	for _, accion := range filterAccions {
		fmt.Printf("%-6s $%.2f\n", accion.Symbol, accion.LastPrice)
	}
	fmt.Printf("\nTotal: %d acciones\n", len(filterAccions))
}

func getArgAccionsExtApi() ([]model.Symbol, error) {
	resp, err := http.Get("https://data912.com/live/usa_adrs")
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la API: %w", err)
	}
	defer resp.Body.Close()

	var accions []model.Symbol
	if err := json.NewDecoder(resp.Body).Decode(&accions); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %w", err)
	}

	return accions, nil
}

func filterArgAccions(allAccions []model.Symbol, filter []string) []model.Symbol {
	var result []model.Symbol

	filterMap := make(map[string]bool)
	for _, s := range filter {
		filterMap[s] = true
	}

	for _, accion := range allAccions {
		if filterMap[accion.Symbol] {
			result = append(result, accion)
		}
	}

	return result
}
