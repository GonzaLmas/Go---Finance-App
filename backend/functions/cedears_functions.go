package functions

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"net/http"
	"strings"
)

func GetCedearsExtApi() ([]model.Symbol, error) {
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

func FilterCedears(allCedears []model.Symbol, filter []string) []model.Symbol {
	var result []model.Symbol

	filterMap := make(map[string]bool)
	for _, s := range filter {
		filterMap[strings.ToUpper(s)] = true
	}

	for _, cedear := range allCedears {
		if filterMap[strings.ToUpper(cedear.Symbol)] {
			result = append(result, cedear)
		}
	}

	return result
}
