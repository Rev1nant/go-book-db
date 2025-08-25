package model

type Book struct {
	Title  string  `json:"title"`
	Author Author  `json:"author"`
	Genre  []Genre `json:"genre"`
	Price  float64 `json:"price"`
	Amount int     `json:"amount"`
}
