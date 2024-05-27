package tree

import (
	"context"

	"github.com/SawitProRecruitment/UserService/repository/tree"
)

type Service interface {
	CreateNewTree(ctx context.Context, estateId string, xCoordinate, yCoordinate, height int) (string, error)
	RetrievesByEstateID(ctx context.Context, id string) ([]*tree.Tree, error)
	GetStats(ctx context.Context, id string) (*tree.TreeStats, error)
}
