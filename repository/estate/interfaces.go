package estate

import "context"

type RepositoryInterface interface {
	FindByID(ctx context.Context, id string) (*Estate, error)
	Insert(ctx context.Context, estate *Estate) (string, error)
}
