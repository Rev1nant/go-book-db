package domain

type Book struct {
	BookID int
	Title  string
	Author Author
	Genre  []Genre
	Price  float64
	Amount int
}
