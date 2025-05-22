package functions

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"net/http"
)

func GetArgDolarHistorico() ([]model.ArgDolar, error) {
	resp, err := http.Get("https://api.argentinadatos.com/v1/cotizaciones/dolares/")
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la API: %w", err)
	}
	defer resp.Body.Close()

	var argDolarHist []model.ArgDolar

	if err := json.NewDecoder(resp.Body).Decode(&argDolarHist); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %w", err)
	}

	return argDolarHist, nil
}
