package model

type ArgDolar struct {
	Compra             float64
	Venta              float64
	Casa               string
	Nombre             string
	Moneda             string
	FechaActualizacion string
}

type ArgDolarResponse struct {
	Compra             float64 `json:"compra"`
	Venta              float64 `json:"venta"`
	Casa               string  `json:"casa"`
	Nombre             string  `json:"nombre"`
	Moneda             string  `json:"moneda"`
	FechaActualizacion string  `json:"fecha_actualizacion"`
}
