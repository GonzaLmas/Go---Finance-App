package model

type Symbol struct {
	Symbol    string  `json:"symbol"`   // name field
	LastPrice float64 `json:"c"`        // price field
	Price     string  `json:"price"`    // binance price field
	Base      string  `json:"base"`     // coinbase
	Currency  string  `json:"currency"` // coinbase
	Amount    string  `json:"amount"`   // coinbase
}
