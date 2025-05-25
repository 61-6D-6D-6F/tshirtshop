package repository

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/61-6D-6D-6F/tshirtshop/internal/model"
)

func TestOrderRepository_CRUD(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	tShirtRepo := NewTShirtRepository(db)
	cartRepo := NewCartRepository(db)
	orderRepo := NewOrderRepository(db)

	userID := 2
	tShirtRepo.Save(&model.TShirt{
		ID:    1,
		Name:  "newModel",
		Size:  "M",
		Color: "white",
		Price: 5.9,
		Stock: 10,
	})
	cartItem := &model.CartItem{TShirtID: 1, Quantity: 5}
	cartRepo.Add(userID, cartItem)
	// Checkout
	got, err := orderRepo.Checkout(userID)
	tShirt, _ := tShirtRepo.Get(1)
	if err != nil || got.UserID != userID || got.Total != 5.9*5 ||
		got.Items[0].TShirtID != 1 || got.Items[0].Quantity != 5 || tShirt.Stock != 5 {
		t.Fatalf("checkout failed: %v, got: %+v", err, got)
	}
	// List
	orders, err := orderRepo.List(userID)
	if err != nil || orders[0].UserID != 2 || orders[0].Total != 5.9*5 ||
		orders[0].Items[0].TShirtID != 1 || orders[0].Items[0].Quantity != 5 {
		t.Fatalf("list failed: %v, got: %+v", err, orders)
	}
}
