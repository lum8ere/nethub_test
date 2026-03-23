package db_manager

import (
	"context"
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Manager interface {
	DB() *gorm.DB
	Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error
	Close() error
	UsePlugins(plugins ...gorm.Plugin) error
}

type gormManager struct {
	db    *gorm.DB
	sqlDB *sql.DB
}

func NewDbManager(dsn string) (Manager, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)

	return &gormManager{db: db, sqlDB: sqlDB}, nil
}

func (m *gormManager) DB() *gorm.DB { return m.db }

func (m *gormManager) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return m.db.WithContext(ctx).Transaction(fn)
}

func (m *gormManager) Close() error { return m.sqlDB.Close() }

func (m *gormManager) UsePlugins(plugins ...gorm.Plugin) error {
	for _, p := range plugins {
		if err := m.db.Use(p); err != nil {
			return fmt.Errorf("failed to use plugin %s: %w", p.Name(), err)
		}
	}
	return nil
}
