package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Estructura para la respuesta de la API externa
type ArgAccionsExtApi struct {
	Symbol    string  `json:"symbol"`
	LastPrice float64 `json:"c"` // Usamos el campo "c" como precio
}

func ArgAccions() {
	// 1. Lista de símbolos que nos interesan (definida en código)
	favsArgAccions := []string{"GGAL", "LOMA", "PAMP", "YPFD"}

	// 2. Obtener todos los datos de la API externa
	accions, err := getArgAccionsExtApi()
	if err != nil {
		log.Fatal("Error obteniendo datos de la API:", err)
	}

	// 3. Filtrar las acciones que nos interesan
	filterAccions := filterArgAccions(accions, favsArgAccions)

	// 4. Mostrar resultados en consola
	fmt.Println("=== ACCIONES ===")
	for _, accion := range filterAccions {
		fmt.Printf("%-6s $%.2f\n", accion.Symbol, accion.LastPrice)
	}
	fmt.Printf("\nTotal: %d acciones\n", len(filterAccions))
}

func getArgAccionsExtApi() ([]ArgAccionsExtApi, error) {
	resp, err := http.Get("https://data912.com/live/arg_stocks")
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la API: %w", err)
	}
	defer resp.Body.Close()

	var accions []ArgAccionsExtApi
	if err := json.NewDecoder(resp.Body).Decode(&accions); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %w", err)
	}

	return accions, nil
}

func filterArgAccions(allAccions []ArgAccionsExtApi, filter []string) []ArgAccionsExtApi {
	var result []ArgAccionsExtApi

	// Convertir el filtro a mapa para búsqueda rápida
	filterMap := make(map[string]bool)
	for _, s := range filter {
		filterMap[s] = true
	}

	// Filtrar acciones
	for _, accion := range allAccions {
		if filterMap[accion.Symbol] {
			result = append(result, accion)
		}
	}

	return result
}
