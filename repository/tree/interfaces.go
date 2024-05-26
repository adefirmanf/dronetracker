package tree

import "context"

type RepositoryInterface interface {
	Insert(ctx context.Context, tree *Tree) (string, error)
	GetStats(ctx context.Context, estateID string) (*TreeStats, error)
	FindByEstateID(ctx context.Context, estateId string) ([]*Tree, error)
	IsExistInEstate(ctx context.Context, estateId string, xCoordinate, yCoordinate int) (bool, error)
}
