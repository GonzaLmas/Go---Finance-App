package functions

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"net/http"
	"strings"
)

func GetUsaAccionsExtApi() ([]model.Symbol, error) {
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

func FilterUsaAccions(allAccions []model.Symbol, filter []string) []model.Symbol {
	var result []model.Symbol

	filterMap := make(map[string]bool)
	for _, s := range filter {
		filterMap[strings.ToUpper(s)] = true
	}

	for _, accion := range allAccions {
		if filterMap[strings.ToUpper(accion.Symbol)] {
			result = append(result, accion)
		}
	}

	return result
}
