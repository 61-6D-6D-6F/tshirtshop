package model

type TShirt struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Size  string  `json:"size"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
