package user

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]User, error) {
	rows, err := r.db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *Repository) GetByID(id int) (*User, error) {
	var user User
	err := r.db.QueryRow("SELECT id, name, age FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) Create(user *User) error {
	err := r.db.QueryRow(
		"INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id",
		user.Name, user.Age).Scan(&user.ID)
	return err
}

func (r *Repository) Update(user *User) error {
	_, err := r.db.Exec("UPDATE users SET name=$1, age=$2 WHERE id=$3", user.Name, user.Age, user.ID)
	return err
}

func (r *Repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
