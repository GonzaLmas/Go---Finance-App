package functions

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"net/http"
	"strings"
)

func GetCriptoExtApi(favsCripto []string) ([]model.Symbol, error) {
	var criptos []model.Symbol

	for _, cripto := range favsCripto {
		pair := strings.Replace(cripto, "USDT", "USD", 1)
		url := fmt.Sprintf("https://api.coinbase.com/v2/prices/%s/spot", pair)

		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error al conectar con Coinbase: %w", err)
		}
		defer resp.Body.Close()

		var apiResponse struct {
			Data model.Symbol `json:"data"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
			return nil, fmt.Errorf("error decodificando respuesta de %s: %w", cripto, err)
		}

		criptos = append(criptos, apiResponse.Data)
	}

	return criptos, nil
}
