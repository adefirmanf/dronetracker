package tree

import (
	"context"

	"github.com/google/uuid"
)

type TreeRepositoryInterface interface {
	CreateTree(ctx context.Context, tree *Tree) error
	GetTree(ctx context.Context, id uuid.UUID) (*Tree, error)
}
