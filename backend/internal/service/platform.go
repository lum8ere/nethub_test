package service

import (
	"context"
	"nethub-mdm/internal/storage/model"
	"nethub-mdm/internal/storage/query"
)

type PlatformService interface {
	Create(ctx context.Context, p *model.Platform) error
	List(ctx context.Context) ([]*model.Platform, error)
}

type platformService struct {
	q *query.Query
}

func NewPlatformService(q *query.Query) PlatformService {
	return &platformService{q: q}
}

func (s *platformService) Create(ctx context.Context, p *model.Platform) error {
	return s.q.Platform.WithContext(ctx).Create(p)
}

func (s *platformService) List(ctx context.Context) ([]*model.Platform, error) {
	return s.q.Platform.WithContext(ctx).Find()
}
