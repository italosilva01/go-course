package main

type Product struct {
	ID string `json:"id"` // Essa terceira instrução é indicando que o campo ID, quando for convertido para JSon
	// será representado pelo campo id
	Name     string `json:"name"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}