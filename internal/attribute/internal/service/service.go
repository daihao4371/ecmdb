package service

import (
	"context"
	"github.com/Duke1616/ecmdb/internal/attribute/internal/domain"
	"github.com/Duke1616/ecmdb/internal/attribute/internal/repository"
)

type Service interface {
	CreateAttribute(ctx context.Context, req domain.Attribute) (int64, error)
	SearchAttributeByModelIdentifies(ctx context.Context, identifies string) (domain.AttributeProjection, error)
}

type service struct {
	repo repository.AttributeRepository
}

func NewService(repo repository.AttributeRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateAttribute(ctx context.Context, req domain.Attribute) (int64, error) {
	return s.repo.CreateAttribute(ctx, req)
}

func (s *service) SearchAttributeByModelIdentifies(ctx context.Context, identifies string) (domain.AttributeProjection, error) {
	return s.repo.SearchAttributeByModelIdentifies(ctx, identifies)
}