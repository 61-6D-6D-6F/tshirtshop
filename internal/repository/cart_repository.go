package repository

import (
	"database/sql"
	"github.com/61-6D-6D-6F/tshirtshop/internal/model"
)

type CartRepository interface {
	Get(int) (*model.Cart, error)
	Add(int, *model.CartItem) error
	Remove(int, *model.CartItem) error
}

type cartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) Get(id int) (*model.Cart, error) {
	rows, err := r.db.Query("SELECT tshirt_id, quantity FROM cart WHERE user_id = ?", id)
	if err != nil {
		return &model.Cart{}, err
	}
	defer rows.Close()
	var items []model.CartItem
	for rows.Next() {
		var cartItem model.CartItem
		if err := rows.Scan(&cartItem.TShirtID, &cartItem.Quantity); err != nil {
			return &model.Cart{}, err
		}
		items = append(items, cartItem)
	}
	return &model.Cart{UserID: id, Items: items}, nil
}

func (r *cartRepository) Add(id int, cartItem *model.CartItem) error {
	var exists int
	_ = r.db.QueryRow("SELECT COUNT(*) FROM cart WHERE user_id = ? AND tshirt_id = ?", id, cartItem.TShirtID).Scan(&exists)
	if exists == 0 {
		_, err := r.db.Exec("INSERT INTO cart (user_id, tshirt_id, quantity) VALUES (?, ?, ?)", id, cartItem.TShirtID, cartItem.Quantity)
		return err
	}
	_, err := r.db.Exec("UPDATE cart SET quantity = quantity + ? WHERE user_id = ? AND tshirt_id = ?", cartItem.Quantity, id, cartItem.TShirtID)
	return err
}

func (r *cartRepository) Remove(id int, cartItem *model.CartItem) error {
	var qty int
	_ = r.db.QueryRow("SELECT quantity FROM cart WHERE user_id = ? AND tshirt_id = ?", id, cartItem.TShirtID).Scan(&qty)
	if qty <= cartItem.Quantity {
		_, err := r.db.Exec("DELETE FROM cart WHERE user_id = ? AND tshirt_id = ?", id, cartItem.TShirtID)
		return err
	}
	_, err := r.db.Exec("UPDATE cart SET quantity = quantity - ? WHERE user_id = ? AND tshirt_id = ?", cartItem.Quantity, id, cartItem.TShirtID)
	return err
}
