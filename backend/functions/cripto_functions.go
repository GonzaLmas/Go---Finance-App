package functions

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"net/http"
)

func GetCriptoExtApi(favsCripto []string) ([]model.Symbol, error) {
	var criptos []model.Symbol

	for _, cripto := range favsCripto {
		url := "https://api.binance.com/api/v3/ticker/price?symbol=" + cripto
		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error al conectar con la API: %w", err)
		}
		defer resp.Body.Close()

		var criptoResp model.Symbol
		if err := json.NewDecoder(resp.Body).Decode(&criptoResp); err != nil {
			return nil, fmt.Errorf("error decodificando respuesta de %s: %w", cripto, err)
		}

		criptos = append(criptos, criptoResp)
	}

	return criptos, nil
}
