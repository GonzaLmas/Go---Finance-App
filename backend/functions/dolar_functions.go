package functions

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"net/http"
)

func GetArgDolar() ([]model.ArgDolar, error) {
	resp, err := http.Get("https://dolarapi.com/v1/dolares")
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la API: %w", err)
	}
	defer resp.Body.Close()

	var argDolar []model.ArgDolar

	if err := json.NewDecoder(resp.Body).Decode(&argDolar); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %w", err)
	}

	return argDolar, nil
}
