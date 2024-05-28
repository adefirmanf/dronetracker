package drone

import (
	"github.com/SawitProRecruitment/UserService/repository/estate"
	"github.com/SawitProRecruitment/UserService/repository/tree"
)

type Service interface {
	GetDronePlane(estate *estate.Estate, tree []*tree.Tree, maxDistance int) (int, Coordinate)
}
