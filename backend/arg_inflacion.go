package main

import (
	"encoding/json"
	"finance-app/model"
	"fmt"
	"log"
	"net/http"
)

func ArgInflacion() {
	inflacion, err := getArgInflacion()
	if err != nil {
		log.Fatal("Error obteniendo datos de la API: ", err)
	}

	fmt.Printf("=== HISTÓRICO INFLACIÓN ===\n")
	for _, inf := range inflacion {
		fmt.Printf("Fecha: %s  Inflación: %.1f \n", inf.Fecha, inf.Valor)
	}
}

func getArgInflacion() ([]model.Inflacion, error) {
	resp, err := http.Get("https://api.argentinadatos.com/v1/finanzas/indices/inflacion/")
	if err != nil {
		return nil, fmt.Errorf("error al conectar con la API: %w", err)
	}
	defer resp.Body.Close()

	var inflacion []model.Inflacion

	if err := json.NewDecoder(resp.Body).Decode(&inflacion); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta: %w", err)
	}

	return inflacion, nil
}
