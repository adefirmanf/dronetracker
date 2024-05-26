package tree

import "context"

type RepositoryInterface interface {
	Insert(ctx context.Context, tree *Tree) (string, error)
	FindByEstateID(ctx context.Context, estateId string) ([]*Tree, error)

	IsExistInEstate(ctx context.Context, estateId string, xCoordinate, yCoordinate int) (bool, error)
}
