package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/johancuervo/usersGO/internal/domain"
)

type PostgresUserRepository struct {
	db *sql.DB
}

// constructor
func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}
func (r *PostgresUserRepository) GetAllUsers() ([]domain.User, error) {
	rows, err := r.db.Query("SELECT id, name, is_active FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.Id, &user.Name, &user.IsActive); err != nil {
			log.Println("Error escaneando usuario:", err)
			continue
		}
		users = append(users, user)
	}

	return users, nil
}
func (r *PostgresUserRepository) FindByID(id string) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id, name, is_active FROM users WHERE id = $1", id)
	user := &domain.User{}
	err := row.Scan(&user.Id, &user.Name, &user.IsActive)
	return user, err
}

func (r *PostgresUserRepository) Save(user *domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (name, is_active) VALUES ($1, $2)", user.Name, user.IsActive)
	return err
}

func (r *PostgresUserRepository) UpdateUser(user *domain.User, userId string) error {
	res, err := r.db.Exec("UPDATE users SET name = $1, is_active = $2 WHERE id = $3",
		user.Name, user.IsActive, userId)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *PostgresUserRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
