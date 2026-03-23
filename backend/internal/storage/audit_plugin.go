package storage

import (
	"nethub-mdm/pkg/db_manager"
	"nethub-mdm/pkg/logger"

	"gorm.io/gorm"
)

type AuditPlugin struct {
	log logger.Logger
}

func NewAuditPlugin(log logger.Logger) *AuditPlugin {
	return &AuditPlugin{log: log}
}

func (p *AuditPlugin) Name() string {
	return "audit_plugin"
}

func (p *AuditPlugin) Initialize(db *gorm.DB) error {
	callbacks := []db_manager.Callback{
		{
			Operation: "Create",
			After:     "gorm:create",
			Name:      "audit:after_create",
			Fn:        p.afterCreate,
		},
		{
			Operation: "Update",
			After:     "gorm:update",
			Name:      "audit:after_update",
			Fn:        p.afterUpdate,
		},
	}

	return db_manager.SetupGormCallbacks(db, callbacks)
}

func (p *AuditPlugin) afterCreate(db *gorm.DB) {
	if db.Error == nil {
		p.log.Infof("[AUDIT] Created record in table: %s", db.Statement.Table)
	}
}

func (p *AuditPlugin) afterUpdate(db *gorm.DB) {
	if db.Error == nil {
		p.log.Infof("[AUDIT] Updated record in table: %s", db.Statement.Table)
	}
}
