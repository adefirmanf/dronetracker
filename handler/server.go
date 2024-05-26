package handler

import (
	"github.com/SawitProRecruitment/UserService/service/estate"
	"github.com/SawitProRecruitment/UserService/service/tree"
)

type Server struct {
	Validator *Validator

	EstateService estate.Service
	TreeService   tree.Service
}

type NewServerOptions struct {
	Validator *Validator

	EstateService estate.Service
	TreeService   tree.Service
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		Validator: opts.Validator,

		EstateService: opts.EstateService,
		TreeService:   opts.TreeService,
	}
}
