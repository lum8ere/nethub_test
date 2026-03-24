package service

import (
	"context"
	"nethub-mdm/internal/storage/model"
	"nethub-mdm/internal/storage/query"
)

type AuditService interface {
	List(ctx context.Context, limit, offset int) ([]*model.AuditLog, int64, error)
}

type auditService struct {
	q *query.Query
}

func NewAuditService(q *query.Query) AuditService {
	return &auditService{q: q}
}

func (s *auditService) List(ctx context.Context, limit, offset int) ([]*model.AuditLog, int64, error) {
	a := s.q.AuditLog
	return a.WithContext(ctx).Order(a.CreatedAt.Desc()).FindByPage(offset, limit)
}
