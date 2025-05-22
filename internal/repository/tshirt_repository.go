package repository

import (
	"errors"

	"database/sql"
	"github.com/61-6D-6D-6F/tshirtshop/internal/model"
)

type TShirtRepository interface {
	List() ([]model.TShirt, error)
	Save(model.TShirt) error
	Update(int, model.TShirt) error
	Delete(int) error
}

type tShirtRepository struct {
	db *sql.DB
}

func NewTShirtRepository(db *sql.DB) TShirtRepository {
	return &tShirtRepository{db: db}
}

func (r *tShirtRepository) List() ([]model.TShirt, error) {
	rows, err := r.db.Query("SELECT id, name, size, color, price, stock FROM tshirts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tshirts []model.TShirt
	for rows.Next() {
		var tshirt model.TShirt
		if err := rows.Scan(&tshirt.ID, &tshirt.Name, &tshirt.Size,
			&tshirt.Color, &tshirt.Price, &tshirt.Stock); err != nil {
			return nil, err
		}
		tshirts = append(tshirts, tshirt)

	}
	return tshirts, nil
}

func (r *tShirtRepository) Save(tShirt model.TShirt) error {
	_, err := r.db.Exec("INSERT INTO tshirts (name, size, color, price, stock) VALUES (?, ?, ?, ?, ?)",
		tShirt.Name, tShirt.Size, tShirt.Color, tShirt.Price, tShirt.Stock)
	return err
}

func InitDB(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS tshirts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        size TEXT,
        color TEXT,
        price REAL,
        stock INTEGER
    )
    `)
	return err
}

func (r *tShirtRepository) Update(id int, tShirt model.TShirt) error {
	res, err := r.db.Exec("UPDATE tshirts SET name=?, size=?, color=?, price=?, stock=? WHERE id=?",
		tShirt.Name, tShirt.Size, tShirt.Color, tShirt.Price, tShirt.Stock, id)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	if aff == 0 {
		return errors.New("error: Not found")
	}
	return nil
}

func (r *tShirtRepository) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM tshirts WHERE id=?", id)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	if aff == 0 {
		return errors.New("error: Not found")
	}
	return nil
}
