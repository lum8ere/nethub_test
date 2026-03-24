package service

import (
	"context"
	"nethub-mdm/internal/storage/model"
	"nethub-mdm/internal/storage/query"
)

type LocationService interface {
	Create(ctx context.Context, loc *model.Location) error
	List(ctx context.Context, name string, isActive *bool, limit, offset int) ([]*model.Location, int64, error)
	Update(ctx context.Context, id string, loc *model.Location) error
	Delete(ctx context.Context, id string) error
}

type locationService struct {
	q *query.Query
}

func NewLocationService(q *query.Query) LocationService {
	return &locationService{q: q}
}

func (s *locationService) Create(ctx context.Context, loc *model.Location) error {
	return s.q.Location.WithContext(ctx).Create(loc)
}

func (s *locationService) List(ctx context.Context, name string, isActive *bool, limit, offset int) ([]*model.Location, int64, error) {
	l := s.q.Location
	db := l.WithContext(ctx)
	if name != "" {
		db = db.Where(l.Name.Like("%" + name + "%"))
	}
	if isActive != nil {
		db = db.Where(l.IsActive.Is(*isActive))
	}
	return db.FindByPage(offset, limit)
}

func (s *locationService) Update(ctx context.Context, id string, loc *model.Location) error {
	loc.ID = &id
	_, err := s.q.Location.WithContext(ctx).Where(s.q.Location.ID.Eq(id)).Updates(loc)
	return err
}

func (s *locationService) Delete(ctx context.Context, id string) error {
	_, err := s.q.Location.WithContext(ctx).Where(s.q.Location.ID.Eq(id)).Delete()
	return err
}
