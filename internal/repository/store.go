package repository

import (
	"os"
	"testing"

	"database/sql"
)

func InitDB(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS tshirts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        size TEXT,
        color TEXT,
        price REAL,
        stock INTEGER
    );
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT,
		email TEXT,
		is_admin INTEGER
	);
    `)
	// Ensure admin exists
	var adminPassword = "pass123"
	adminPassword, err = hashPassword(adminPassword)
	if err != nil {
		return err
	}
	if pass := os.Getenv("TSHIRT_ADMIN_PASS"); pass != "" {
		adminPassword, err = hashPassword(pass)
	}
	var exists int
	_ = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = 'admin'").Scan(&exists)
	if exists == 0 {
		_, _ = db.Exec("INSERT INTO users (username, password, email, is_admin) VALUES (?, ?, ?, ?)", "admin", adminPassword, "admin@mail.com", 1)
	}
	return err
}

func setupTestDB(t *testing.T) *sql.DB {
	os.Remove("test.db")
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatal(err)
	}
	if err := InitDB(db); err != nil {
		t.Fatal(err)
	}
	return db
}
