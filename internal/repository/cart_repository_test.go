package repository

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/61-6D-6D-6F/tshirtshop/internal/model"
)

func TestCartRepository_CRUD(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	repo := NewCartRepository(db)

	userID := 2
	cartItem := &model.CartItem{TShirtID: 1, Quantity: 5}
	// Create
	if err := repo.Add(userID, cartItem); err != nil {
		t.Fatalf("save failed: %v", err)
	}
	// Read
	got, err := repo.Get(userID)
	if err != nil || got.UserID != userID {
		t.Fatalf("find failed: %v, got: %+v", err, got)
	}
	// Update - add quantity
	cartItem.Quantity = 1
	if err := repo.Add(userID, cartItem); err != nil {
		t.Fatalf("update add failed: %v", err)
	}
	got, _ = repo.Get(userID)
	if got.Items[0].Quantity != 6 {
		t.Fatalf("update add did not persist")
	}
	// Update - remove quantity
	cartItem.Quantity = 1
	if err := repo.Remove(userID, cartItem); err != nil {
		t.Fatalf("update remove failed: %v", err)
	}
	got, err = repo.Get(userID)
	if got.Items[0].Quantity != 5 {
		t.Fatalf("update remove did not persist")
	}
	// Delete
	cartItem.Quantity = 5
	if err := repo.Remove(userID, cartItem); err != nil {
		t.Fatalf("delete failed: %v", err)
	}
	_, err = repo.Get(userID)
	if err != nil {
		t.Fatalf("should be not found after delete")
	}
}
