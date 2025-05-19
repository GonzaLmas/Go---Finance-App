package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CedearsExtApi struct {
	Symbol    string  `json:"symbol"`
	LastPrice float64 `json:"c"` // Usamos el campo "c" como precio
}

func Cedears() {
	favsCedears := []string{"AAPL", "GOOGL", "MELI", "META", "QQQ", "SPY"}

	cedears, err := getCedearsExtApi()
	if err != nil {
		log.Fatal("Error obteniendo datos de la API: ", err)
	}

	filterCedears := filterCedears(cedears, favsCedears)

	fmt.Println("=== CEDEARS ===")
	for _, cedear := range filterCedears {
		fmt.Printf("%-6s $%.2f\n", cedear.Symbol, cedear.LastPrice)
	}

	fmt.Printf("\nTotal: %d acciones\n", len(filterCedears))
}

func getCedearsExtApi() ([]CedearsExtApi, error) {
	resp, err := http.Get("https://data912.com/live/arg_cedears")
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la API: %w", err)
	}
	defer resp.Body.Close()

	var cedears []CedearsExtApi
	if err := json.NewDecoder(resp.Body).Decode(&cedears); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %w", err)
	}

	return cedears, nil
}

func filterCedears(allCedears []CedearsExtApi, filter []string) []CedearsExtApi {
	var result []CedearsExtApi

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
