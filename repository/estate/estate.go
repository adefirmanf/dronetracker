package estate

import (
	"context"
	"fmt"

	"github.com/SawitProRecruitment/UserService/repository"
)

type Repository struct {
	dbHandler repository.Repository
}

func NewRepository(r *repository.Repository) RepositoryInterface {
	return &Repository{
		dbHandler: *r,
	}
}

func (r *Repository) tableName() string {
	return "estate"
}

func (r *Repository) FindByID(ctx context.Context, id string) (*Estate, error) {
	q := fmt.Sprintf("SELECT id, width, length, created_at, updated_at FROM %s WHERE id = $1", r.tableName())
	estate := &Estate{}
	err := r.dbHandler.Db.QueryRowContext(ctx, q, id).Scan(&estate.ID, &estate.Width, &estate.Length, &estate.CreatedAt, &estate.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return estate, nil
}

func (r *Repository) Insert(ctx context.Context, estate *Estate) (string, error) {
	q := fmt.Sprintf("INSERT INTO %s (width, length) VALUES ($1, $2) RETURNING id", r.tableName())
	err := r.dbHandler.Db.QueryRowContext(ctx, q, estate.Width, estate.Length).Scan(&estate.ID)
	if err != nil {
		return "", err
	}
	return estate.ID, nil
}
