package estate

import (
	"context"

	"github.com/SawitProRecruitment/UserService/repository/estate"
)

type service struct {
	estateRepository estate.RepositoryInterface
}

type Service interface {
	CreateNewEstate(ctx context.Context, width, length int) (string, error)
}

func NewEstateService(r estate.RepositoryInterface) Service {
	return &service{
		estateRepository: r,
	}
}

func (s *service) CreateNewEstate(ctx context.Context, width, length int) (string, error) {
	return s.estateRepository.Insert(ctx, &estate.Estate{
		Width:  width,
		Length: length,
	})
}
