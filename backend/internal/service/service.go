package service

import "github.com/SYSU-ECNC/ecnc-oa/backend/internal/repository"

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}
