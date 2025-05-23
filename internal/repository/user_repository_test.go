package repository

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/61-6D-6D-6F/tshirtshop/internal/model"
)

func TestUserRepository_CRUD(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
	repo := NewUserRepository(db)

	user := &model.User{Username: "bob", Password: "pass123", Email: "bob@mail.com"}
	// Create
	if err := repo.Save(user); err != nil {
		t.Fatalf("save failed: %v", err)
	}
	// Read
	// admin has ID 1
	got, err := repo.Get(2)
	if err != nil || got.Username != "bob" {
		t.Fatalf("find failed: %v, got: %+v", err, got)
	}
	// Login
	got, err = repo.TryLogin(user.Username, user.Password)
	if err != nil || got.Username != "bob" {
		t.Fatalf("login failed: %v, %+v", err, got)
	}
	// Register same username
	err = repo.TryRegister(user)
	if err == nil {
		t.Fatalf("register existed failed: %v", err)
	}
	// Update
	user.Username = "john"
	if err := repo.Update(2, user); err != nil {
		t.Fatalf("update failed: %v", err)
	}
	got, _ = repo.Get(2)
	if got.Username != "john" {
		t.Fatalf("update did not persist")
	}
	// List
	users, err := repo.List()
	if err != nil || len(users) != 2 {
		t.Fatalf("list failed: %v, users: %+v", err, users)
	}
	// Delete
	if err := repo.Delete(2); err != nil {
		t.Fatalf("delete failed: %v", err)
	}
	_, err = repo.Get(2)
	if err == nil {
		t.Fatalf("should be not found after delete")
	}
	// Register
	err = repo.TryRegister(user)
	if err != nil {
		t.Fatalf("register new failed: %v", err)
	}
}
