package models

type Transaction struct {
    ID        string  `json:"id"`
    ProductID string  `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Total     float64 `json:"total"`
}

var Transactions = []Transaction{}