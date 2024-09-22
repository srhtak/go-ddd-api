package persistence

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/srhtak/go-ddd-api/internal/domain"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(connectionString string) (*PostgresUserRepository, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresUserRepository{db: db}, nil
}

func (r *PostgresUserRepository) Create(user *domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, password, created_at, updated_at) VALUES ($1, $2, $3, $4)",
		user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *PostgresUserRepository) GetByUsername(username string) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.QueryRow("SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}