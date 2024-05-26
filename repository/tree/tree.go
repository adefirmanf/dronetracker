package tree

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
	return "tree"
}

func (r *Repository) Insert(ctx context.Context, tree *Tree) (string, error) {
	q := fmt.Sprintf("INSERT INTO %s (estate_id, x_coordinate, y_coordinate, height) VALUES ($1, $2, $3, $4) RETURNING id", r.tableName())
	err := r.dbHandler.Db.QueryRowContext(ctx, q, tree.EstateID, tree.XCoordinate, tree.YCoordinate, tree.Height).Scan(&tree.ID)
	if err != nil {
		return "", err
	}
	return tree.ID, nil
}

func (r *Repository) IsExistInEstate(ctx context.Context, estateId string, xCoordinate, yCoordinate int) (bool, error) {
	q := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE estate_id = $1 AND x_coordinate = $2 AND y_coordinate = $3)", r.tableName())
	var exist bool
	err := r.dbHandler.Db.QueryRowContext(ctx, q, estateId, xCoordinate, yCoordinate).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (r *Repository) FindByEstateID(ctx context.Context, estateId string) ([]*Tree, error) {
	q := fmt.Sprintf("SELECT id, estate_id, x_coordinate, y_coordinate, height, created_at, updated_at FROM %s WHERE estate_id = $1", r.tableName())
	rows, err := r.dbHandler.Db.QueryContext(ctx, q, estateId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trees []*Tree
	for rows.Next() {
		tree := Tree{}
		if err := rows.Scan(&tree.ID, &tree.EstateID, &tree.XCoordinate, &tree.YCoordinate, &tree.Height, &tree.CreatedAt, &tree.UpdatedAt); err != nil {
			return nil, err
		}
		trees = append(trees, &tree)
	}

	return trees, nil
}

func (r *Repository) GetStats(ctx context.Context, estateID string) (*TreeStats, error) {
	q := fmt.Sprintf("SELECT median(height), count(id), min(height), max(height) FROM %s WHERE estate_id = $1", r.tableName())
	row := r.dbHandler.Db.QueryRowContext(ctx, q, estateID)

	var stats TreeStats
	if err := row.Scan(&stats.Median, &stats.Count, &stats.Min, &stats.Max); err != nil {
		return nil, err
	}
	return &stats, nil

}
