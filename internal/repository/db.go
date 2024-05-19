package repository

import (
	"Varian_v2/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(ctx context.Context, cfg config.DBConfig) (*Repository, error) {
	connStr := fmt.Sprintf(`user=%s password=%s
   host=%s port=%d dbname=%s`,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
	conn, err := pgxpool.New(ctx, connStr)
	return &Repository{pool: conn}, err
}

func (r *Repository) GetContact() ([]*Contact, error) {
	rows, err := r.pool.Query(context.Background(), "SELECT * FROM public.contacts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []*Contact
	for rows.Next() {
		contact := &Contact{}
		err := rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName,
			&contact.Email, &contact.Phone)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact) // Добавить contact в массив
	}
	return contacts, nil
}

func (r *Repository) Ping(ctx context.Context) error {
	return r.pool.Ping(ctx)
}
