package repository

import (
	"database/sql"
	"errors"

	"github.com/61-6D-6D-6F/tshirtshop/internal/model"
)

type OrderRepository interface {
	List(int) ([]*model.Order, error)
	Checkout(int) (*model.Order, error)
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) List(id int) ([]*model.Order, error) {
	rows, err := r.db.Query("SELECT order_id, total FROM orders WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []*model.Order
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.OrderID, &order.Total); err != nil {
			return nil, err
		}
		order.UserID = id
		itemRows, _ := r.db.Query("SELECT tshirt_id, quantity FROM order_items WHERE order_id = ?", order.OrderID)
		for itemRows.Next() {
			var cartItem model.CartItem
			itemRows.Scan(&cartItem.TShirtID, &cartItem.Quantity)
			order.Items = append(order.Items, cartItem)
		}
		itemRows.Close()
		orders = append(orders, &order)
	}
	return orders, nil
}

func (r *orderRepository) Checkout(id int) (*model.Order, error) {
	cartRepo := NewCartRepository(r.db)
	tShirtRepo := NewTShirtRepository(r.db)

	cart, err := cartRepo.Get(id)
	if err != nil || len(cart.Items) == 0 {
		return &model.Order{}, errors.New("error: Cart empty")
	}
	var total float64
	for _, item := range cart.Items {
		tShirt, err := tShirtRepo.Get(item.TShirtID)
		if err != nil {
			return &model.Order{}, errors.New("error: Bad tshirt in cart")
		}
		if tShirt.Stock < item.Quantity {
			return &model.Order{}, errors.New("error: Not enough stock for " + tShirt.Name)

		}
		total += float64(item.Quantity) * tShirt.Price
	}
	// Payment here
	// Success simulated
	res, err := r.db.Exec("INSERT INTO orders (user_id, total) VALUES (?, ?)", id, total)
	if err != nil {
		return &model.Order{}, err
	}
	orderID, _ := res.LastInsertId()
	for _, item := range cart.Items {
		_, err = r.db.Exec("INSERT INTO order_items (order_id, tshirt_id, quantity) VALUES (?, ?, ?)",
			orderID, item.TShirtID, item.Quantity)
		if err != nil {
			return &model.Order{}, err
		}
		_, err = r.db.Exec("UPDATE tshirts SET stock = stock - ? WHERE id = ?",
			item.Quantity, item.TShirtID)
		if err != nil {
			return &model.Order{}, err
		}
	}
	_, err = r.db.Exec("DELETE FROM cart WHERE user_id = ?", id)
	if err != nil {
		return &model.Order{}, err
	}
	return &model.Order{
		OrderID: int(orderID),
		UserID:  id,
		Items:   cart.Items,
		Total:   total,
	}, nil
}
