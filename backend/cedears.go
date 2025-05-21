package main

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"log"
	"net/http"
)

func Cedears() {
	favsCedears := []string{"AAPL", "AMZN", "GOOGL", "MELI", "META", "MSFT", "NVDA", "TSLA", "QQQ", "SPY"}

	cedears, err := getCedearsExtApi()
	if err != nil {
		log.Fatal("Error obteniendo datos de la API: ", err)
	}

	filterCedears := filterCedears(cedears, favsCedears)

	fmt.Printf("=== CEDEARS ===\n")
	for _, cedear := range filterCedears {
		fmt.Printf("%-6s $%.2f\n", cedear.Symbol, cedear.LastPrice)
	}

	fmt.Printf("\nTotal: %d acciones\n", len(filterCedears))
}

func getCedearsExtApi() ([]model.Symbol, error) {
	resp, err := http.Get("https://data912.com/live/arg_cedears")
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la API: %w", err)
	}
	defer resp.Body.Close()

	var cedears []model.Symbol
	if err := json.NewDecoder(resp.Body).Decode(&cedears); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %w", err)
	}

	return cedears, nil
}

func filterCedears(allCedears []model.Symbol, filter []string) []model.Symbol {
	var result []model.Symbol

	filterMap := make(map[string]bool)
	for _, s := range filter {
		filterMap[s] = true
	}

	for _, cedear := range allCedears {
		if filterMap[cedear.Symbol] {
			result = append(result, cedear)
		}
	}

	return result
}
