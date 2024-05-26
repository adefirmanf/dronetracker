package drone

import (
	"github.com/SawitProRecruitment/UserService/repository/estate"
	"github.com/SawitProRecruitment/UserService/repository/tree"
)

type Drone struct {
	Estate estate.Estate
	Tree   []tree.Tree
}

type DroneOpts struct {
	Logging bool
}
