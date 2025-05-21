package main

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"log"
	"net/http"
)

func ArgDolar() {
	argDolars, err := getArgDolar()
	if err != nil {
		log.Fatal("Error obteniendo datos de la API: ", err)
	}

	fmt.Printf("=== DÓLARES ===\n")
	for _, dolar := range argDolars {
		fmt.Printf("%s: Compra: $%.2f  Venta: $%.2f\n", dolar.Nombre, dolar.Compra, dolar.Venta)
	}
}

func getArgDolar() ([]model.ArgDolar, error) {
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
