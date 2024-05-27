package estate

import (
	"context"

	"github.com/SawitProRecruitment/UserService/repository/estate"
)

type Service interface {
	RetrieveEstate(ctx context.Context, id string) (*estate.Estate, error)
	CreateNewEstate(ctx context.Context, width, length int) (string, error)
}
