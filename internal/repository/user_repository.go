package repository

import (
	"context"

	"project/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(ctx, query, user.Name, user.Email).Scan(&user.ID)
}

func (r *UserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, email FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}
	return users, nil
}