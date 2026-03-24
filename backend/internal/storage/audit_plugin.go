package storage

import (
	"fmt"
	"nethub-mdm/internal/storage/model"
	"nethub-mdm/pkg/db_manager"
	"nethub-mdm/pkg/logger"
	"nethub-mdm/pkg/types"
	"reflect"

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
			After:     "gorm:after_create",
			Name:      "audit:db_save_create",
			Fn:        p.saveToDB("CREATE"),
		},
		{
			Operation: "Update",
			After:     "gorm:after_update",
			Name:      "audit:db_save_update",
			Fn:        p.saveToDB("UPDATE"),
		},
	}
	return db_manager.SetupGormCallbacks(db, callbacks)
}

func (p *AuditPlugin) saveToDB(op string) func(*gorm.DB) {
	return func(db *gorm.DB) {
		if db.Error != nil || db.Statement.Schema == nil || db.Statement.Table == "audit_logs" {
			return
		}

		recordID := p.extractID(db)

		auditEntry := &model.AuditLog{
			Operation:  types.StrPtr(op),
			TableName_: types.StrPtr(db.Statement.Table),
			RecordID:   recordID,
		}

		err := db.Session(&gorm.Session{NewDB: true, SkipDefaultTransaction: true}).
			Table("audit_logs").
			Create(auditEntry).Error

		if err != nil {
			p.log.Errorf("failed to save audit log: %v", err)
		} else {
			idStr := "NULL"
			if recordID != nil {
				idStr = *recordID
			}
			p.log.Infof("Audit recorded: [%s] table=%s id=%s", op, db.Statement.Table, idStr)
		}
	}
}

func (p *AuditPlugin) extractID(db *gorm.DB) *string {
	rv := reflect.Indirect(db.Statement.ReflectValue)

	if rv.Kind() == reflect.Slice && rv.Len() > 0 {
		rv = reflect.Indirect(rv.Index(0))
	}

	if rv.Kind() != reflect.Struct {
		return nil
	}

	pkField := db.Statement.Schema.PrioritizedPrimaryField
	if pkField == nil && len(db.Statement.Schema.PrimaryFields) > 0 {
		pkField = db.Statement.Schema.PrimaryFields[0]
	}

	if pkField != nil {
		val, isZero := pkField.ValueOf(db.Statement.Context, rv)
		if !isZero {
			return p.formatValue(val)
		}
	}

	f := rv.FieldByName("ID")
	if f.IsValid() && !f.IsZero() {
		return p.formatValue(f.Interface())
	}

	return nil
}

func (p *AuditPlugin) formatValue(val interface{}) *string {
	v := reflect.ValueOf(val)
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	s := fmt.Sprintf("%v", v.Interface())
	if s == "" || s == "<nil>" || s == "0x0" {
		return nil
	}
	return &s
}
