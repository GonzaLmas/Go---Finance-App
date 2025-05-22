package functions

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"net/http"
)

func GetArgLetrasExtApi() ([]model.Symbol, error) {
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
