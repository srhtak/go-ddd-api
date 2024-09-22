package persistence

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/srhtak/go-ddd-api/internal/domain"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(connectionString string) (*PostgresUserRepository, error) {
    log.Println("Attempting to connect to the database")
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        log.Printf("Error opening database connection: %v", err)
        return nil, err
    }

    log.Println("Pinging the database")
    if err = db.Ping(); err != nil {
        log.Printf("Error pinging database: %v", err)
        return nil, err
    }

    log.Println("Running migrations")
    if err = RunMigrations(db); err != nil {
        log.Printf("Error running migrations: %v", err)
        return nil, err
    }

    log.Println("PostgresUserRepository initialized successfully")
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