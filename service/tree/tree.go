package tree

import (
	"context"
	"log"

	"github.com/SawitProRecruitment/UserService/repository/estate"
	"github.com/SawitProRecruitment/UserService/repository/tree"
	"github.com/SawitProRecruitment/UserService/types"
)

type service struct {
	estateRepository estate.RepositoryInterface
	treeRepository   tree.RepositoryInterface
}

type Service interface {
	CreateNewTree(ctx context.Context, estateId string, xCoordinate, yCoordinate, height int) (string, error)

	IsTreePlanted(ctx context.Context, estateID string, xCoordinate, yCoordinate int) bool
	IsOutOfBound(ctx context.Context, estate *estate.Estate, xCoordinate, yCoordinate int) bool
}

func NewTreeService(tree tree.RepositoryInterface, estate estate.RepositoryInterface) Service {
	return &service{
		treeRepository:   tree,
		estateRepository: estate,
	}
}

// CreateNewTree creates a tree in specific estate
// xCoordinate, yCoordinate, height must be given
func (s *service) CreateNewTree(ctx context.Context, estateId string, xCoordinate, yCoordinate, height int) (string, error) {
	// Check if estate id is exist
	estate, err := s.estateRepository.FindByID(ctx, estateId)
	if err != nil || estate == nil {
		log.Println(err)
		return "", types.ErrorEstateNotFound
	}
	// Check if tree is out of bound
	if s.IsOutOfBound(ctx, estate, xCoordinate, yCoordinate) {
		return "", types.ErorrTreeOutOfBound
	}
	// Check if tree already planted in same coordinate
	if s.IsTreePlanted(ctx, estateId, xCoordinate, yCoordinate) {
		return "", types.ErrorTreeAlreadyPlanted
	}

	return s.treeRepository.Insert(ctx, &tree.Tree{
		EstateID:    estateId,
		XCoordinate: xCoordinate,
		YCoordinate: yCoordinate,
		Height:      height,
	})
}

func (s *service) IsOutOfBound(ctx context.Context, estate *estate.Estate, xCoordinate, yCoordinate int) bool {
	return xCoordinate < 0 || xCoordinate > estate.Width || yCoordinate < 0 || yCoordinate > estate.Length
}

func (s *service) IsTreePlanted(ctx context.Context, estateId string, xCoordinate, yCoordinate int) bool {
	if exist, err := s.treeRepository.IsExistInEstate(ctx, estateId, xCoordinate, yCoordinate); err != nil {
		return false
	} else {
		return exist
	}
}