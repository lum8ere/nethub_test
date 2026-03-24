package service

import (
	"context"
	"nethub-mdm/internal/storage/model"
	"nethub-mdm/internal/storage/query"
	"nethub-mdm/pkg/types"
)

type DeviceService interface {
	Create(ctx context.Context, dev *model.Device) error
	List(ctx context.Context, hostname string, isActive *bool, limit, offset int) ([]*model.Device, int64, error)
	GetByID(ctx context.Context, id string) (*model.Device, error)
	Update(ctx context.Context, id string, dev *model.Device) error
	Delete(ctx context.Context, id string) error
}

type deviceService struct {
	q *query.Query
}

func NewDeviceService(q *query.Query) DeviceService {
	return &deviceService{q: q}
}

func (s *deviceService) Create(ctx context.Context, dev *model.Device) error {
	return s.q.Device.WithContext(ctx).Create(dev)
}

func (s *deviceService) List(ctx context.Context, hostname string, isActive *bool, limit, offset int) ([]*model.Device, int64, error) {
	d := s.q.Device
	q := d.WithContext(ctx)

	if hostname != "" {
		q = q.Where(d.Hostname.Like("%" + hostname + "%"))
	}
	if isActive != nil {
		q = q.Where(d.IsActive.Is(*isActive))
	}

	return q.FindByPage(offset, limit)
}

func (s *deviceService) GetByID(ctx context.Context, id string) (*model.Device, error) {
	return s.q.Device.WithContext(ctx).Where(s.q.Device.ID.Eq(id)).First()
}

func (s *deviceService) Update(ctx context.Context, id string, dev *model.Device) error {
	dev.ID = types.StrPtr(id)
	_, err := s.q.Device.WithContext(ctx).Where(s.q.Device.ID.Eq(id)).Updates(dev)
	return err
}

func (s *deviceService) Delete(ctx context.Context, id string) error {
	_, err := s.q.Device.WithContext(ctx).Where(s.q.Device.ID.Eq(id)).Delete()
	return err
}
