package storage

import (
	"nethub-mdm/internal/storage/model"
	"nethub-mdm/pkg/db_manager"
	"nethub-mdm/pkg/logger"
	"nethub-mdm/pkg/types"

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
			Name:      "audit:db_save_create",
			Fn:        p.saveToDB("CREATE"),
		},
		{
			Operation: "Update",
			After:     "gorm:update",
			Name:      "audit:db_save_update",
			Fn:        p.saveToDB("UPDATE"),
		},
	}
	return db_manager.SetupGormCallbacks(db, callbacks)
}

func (p *AuditPlugin) saveToDB(op string) func(*gorm.DB) {
	return func(db *gorm.DB) {
		if db.Error != nil {
			return
		}

		var recordID string
		if dev, ok := db.Statement.Dest.(*model.Device); ok {
			recordID = *dev.ID
		}

		auditEntry := &model.AuditLog{
			Operation:  types.StrPtr(op),
			TableName_: types.StrPtr(db.Statement.Table),
			RecordID:   types.StrPtr(recordID),
		}

		if err := db.Session(&gorm.Session{NewDB: true}).Create(auditEntry).Error; err != nil {
			p.log.Errorf("failed to save audit log to DB: %v", err)
		} else {
			p.log.Infof("Audit: [%s] recorded for table %s (ID: %d)", op, db.Statement.Table, recordID)
		}
	}
}
