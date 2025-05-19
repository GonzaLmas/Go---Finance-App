package model

type Symbol struct {
	Symbol    string  `json:"symbol"` // name field
	LastPrice float64 `json:"c"`      // price field
}
