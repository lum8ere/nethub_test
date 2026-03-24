package service

import (
	"context"
	"testing"
	"time"

	"nethub-mdm/internal/storage/model"
	"nethub-mdm/internal/storage/query"
	"nethub-mdm/pkg/types"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Device struct {
	ID           string `gorm:"primaryKey"`
	CreatedAt    *time.Time
	CreatedBy    *string
	DeletedAt    gorm.DeletedAt
	Name         *string
	UserID       *string
	PlatformCode string
	IP           *string
	Hostname     *string
	Location     *string
	IsActive     *bool
}

func (Device) TableName() string {
	return "devices"
}

func setupTestDB(t *testing.T) *query.Query {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}

	err = db.AutoMigrate(&Device{})
	if err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	return query.Use(db)
}
func TestDeviceService_List(t *testing.T) {
	q := setupTestDB(t)
	svc := NewDeviceService(q)
	ctx := context.Background()

	q.Device.WithContext(ctx).Create(&model.Device{
		ID:           types.NewObjectIdStrRef(),
		Hostname:     types.StrPtr("web-server"),
		IsActive:     types.BoolPtr(true),
		PlatformCode: "LINUX",
	})
	q.Device.WithContext(ctx).Create(&model.Device{
		ID:           types.NewObjectIdStrRef(),
		Hostname:     types.StrPtr("db-server"),
		IsActive:     types.BoolPtr(false),
		PlatformCode: "LINUX",
	})

	t.Run("filter by active", func(t *testing.T) {
		isActive := true
		devices, count, err := svc.List(ctx, "", &isActive, 10, 0)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if count != 1 {
			t.Errorf("expected 1 active device, got %d", count)
		}
		if *devices[0].Hostname != "web-server" {
			t.Errorf("wrong device returned")
		}
	})

	t.Run("search by hostname", func(t *testing.T) {
		devices, count, err := svc.List(ctx, "db", nil, 10, 0)
		if err != nil || count != 1 {
			t.Fatalf("search failed")
		}
		if *devices[0].Hostname != "db-server" {
			t.Errorf("wrong search result")
		}
	})
}

func TestDeviceService_CreateAndGet(t *testing.T) {
	q := setupTestDB(t)
	svc := NewDeviceService(q)
	ctx := context.Background()

	t.Run("create and find device", func(t *testing.T) {
		id := types.NewObjectIdStr()
		err := svc.Create(ctx, &model.Device{
			ID:           types.StrPtr(id),
			Hostname:     types.StrPtr("new-device"),
			PlatformCode: "LINUX",
		})
		if err != nil {
			t.Fatalf("failed to create: %v", err)
		}

		found, err := svc.GetByID(ctx, id)
		if err != nil {
			t.Fatalf("failed to find: %v", err)
		}
		if *found.Hostname != "new-device" {
			t.Errorf("expected new-device, got %s", *found.Hostname)
		}
	})
}
