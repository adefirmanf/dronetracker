package tree

import (
	"context"

	"github.com/SawitProRecruitment/UserService/repository"
)

type TreeRepositoryImplementation struct {
	r repository.Repository
}

func NewTreeRepositoryImplementation(r repository.Repository) *TreeRepositoryImplementation {
	return &TreeRepositoryImplementation{r: r}
}

func (t *TreeRepositoryImplementation) CreateTree(ctx context.Context, tree *Tree) error {
	return t.r.Db.QueryRowContext(ctx, "INSERT INTO tree (id, height) VALUES ($1, $2) RETURNING id", tree.Id, tree.Height).Scan(&tree.Id)
}
