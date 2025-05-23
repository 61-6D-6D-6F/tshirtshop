package model

type TShirt struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Size  string  `json:"size"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CartItem struct {
	TShirtID int `json:"tshirt_id"`
	Quantity int `json:"quantity"`
}

type Cart struct {
	UserID int        `json:"user_id"`
	Items  []CartItem `json:"items"`
}
