package main

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Cripto() {
	favsCripto := []string{"BTCUSDT", "ETHUSDT"}

	criptos, err := getCriptoExtApi(favsCripto)
	if err != nil {
		log.Fatal("Error obteniendo datos de la API: ", err)
	}

	fmt.Printf("=== CRIPTOMONEDAS ===\n")
	for _, cripto := range criptos {
		priceFloat, err := strconv.ParseFloat(cripto.Price, 64)
		if err != nil {
			fmt.Printf("Error al convertir el precio de %s: %v\n", cripto.Symbol, err)
			continue
		}
		fmt.Printf("%-8s $%.2f\n", cripto.Symbol, priceFloat)
	}
	fmt.Printf("\nTotal: %d criptomonedas\n", len(criptos))
}

func getCriptoExtApi(favsCripto []string) ([]model.Symbol, error) {
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
