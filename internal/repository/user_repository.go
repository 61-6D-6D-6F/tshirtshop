package repository

import (
	"errors"

	"database/sql"

	"github.com/61-6D-6D-6F/tshirtshop/internal/model"
)

type UserRepository interface {
	List() ([]*model.User, error)
	Get(int) (*model.User, error)
	Save(*model.User) error
	Update(int, *model.User) error
	Delete(int) error
	TryLogin(string, string) (*model.User, error)
	TryRegister(*model.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) List() ([]*model.User, error) {
	rows, err := r.db.Query("SELECT id, username, email, is_admin FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.IsAdmin); err != nil {
			return nil, err
		}
		users = append(users, &user)

	}
	return users, nil
}

func (r *userRepository) Get(id int) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, username, password, email, is_admin FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.IsAdmin)
	if err == sql.ErrNoRows {
		return nil, errors.New("error: Not found")
	}
	return &user, err
}

func (r *userRepository) Save(user *model.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, password, email, is_admin) VALUES (?, ?, ?, ?)",
		user.Username, user.Password, user.Email, 0)
	return err
}

func (r *userRepository) Update(id int, user *model.User) error {
	res, err := r.db.Exec("UPDATE users SET username=?, password=?, email=? WHERE id=?",
		user.Username, user.Password, user.Email, id)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	if aff == 0 {
		return errors.New("error: Not found")
	}
	return nil
}

func (r *userRepository) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	if aff == 0 {
		return errors.New("error: Not found")
	}
	return nil
}

func (r *userRepository) TryLogin(username string, password string) (*model.User, error) {
	var user model.User
	var isAdmin int
	row := r.db.QueryRow("SELECT username, password, email, is_admin FROM users WHERE username = ?", username)
	if err := row.Scan(&user.Username, &user.Password, &user.Email, &isAdmin); err != nil {
		return &user, errors.New("error: Not found")
	}
	if user.Password != password {
		return &user, errors.New("error: Password mismatch")
	}
	user.IsAdmin = isAdmin == 1
	return &user, nil
}

func (r *userRepository) TryRegister(user *model.User) error {
	var exists int
	_ = r.db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", user.Username).Scan(&exists)
	if exists != 0 {
		return errors.New("error: Username exist")
	}
	_, err := r.db.Exec("INSERT INTO users (username, password, email, is_admin) VALUES (?, ?, ?, ?)",
		user.Username, user.Password, user.Email, 0)
	if err != nil {
		return err
	}
	return nil
}
