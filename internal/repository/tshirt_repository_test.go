package repository

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/61-6D-6D-6F/tshirtshop/internal/model"
)

func TestTShirtRepository_CRUD(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	repo := NewTShirtRepository(db)

	tShirt := &model.TShirt{Name: "model1", Size: "M", Color: "white", Price: 9.99, Stock: 10}
	// Create
	if err := repo.Save(tShirt); err != nil {
		t.Fatalf("save failed: %v", err)
	}
	// Read
	got, err := repo.Get(1)
	if err != nil || got.Name != "model1" {
		t.Fatalf("find failed: %v, got: %+v", err, got)
	}
	// Update
	tShirt.Name = "model2"
	if err := repo.Update(1, tShirt); err != nil {
		t.Fatalf("update failed: %v", err)
	}
	got, _ = repo.Get(1)
	if got.Name != "model2" {
		t.Fatalf("update did not persist")
	}
	// List
	tShirts, err := repo.List()
	if err != nil || len(tShirts) != 1 {
		t.Fatalf("list failed: %v, tShirts: %+v", err, tShirts)
	}
	// Delete
	if err := repo.Delete(1); err != nil {
		t.Fatalf("delete failed: %v", err)
	}
	_, err = repo.Get(1)
	if err == nil {
		t.Fatalf("should be not found after delete")
	}
}
