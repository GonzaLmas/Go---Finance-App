package main

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"log"
	"net/http"
)

func UsaAccions() {
	favsAccions := []string{"AAPL", "GOOGL", "MELI", "META", "QQQ", "SPY"}

	accions, err := getUsaAccionsExtApi()
	if err != nil {
		log.Fatal("Error obteniendo datos de la API: ", err)
	}

	filterAccions := filterUsaAccions(accions, favsAccions)

	fmt.Println("=== USA ACCIONS ===")
	for _, accion := range filterAccions {
		fmt.Printf("%-6s $%.2f\n", accion.Symbol, accion.LastPrice)
	}

	fmt.Printf("\nTotal: %d acciones\n", len(filterAccions))
}

func getUsaAccionsExtApi() ([]model.Symbol, error) {
	resp, err := http.Get("https://data912.com/live/usa_stocks")
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

func filterUsaAccions(allAccions []model.Symbol, filter []string) []model.Symbol {
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
