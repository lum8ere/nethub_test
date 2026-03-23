package db_manager

import (
	"fmt"

	"gorm.io/gorm"
)

type Callback struct {
	Operation string
	Before    string
	After     string
	Name      string
	Fn        func(*gorm.DB)
	Match     func(*gorm.DB) bool
}

func SetupGormCallbacks(db *gorm.DB, cbs []Callback) error {
	cb := db.Callback()
	var err error
	for _, c := range cbs {
		switch c.Operation {
		case "Create":
			processor := cb.Create()
			err = processor.Match(c.Match).Before(c.Before).After(c.After).Register(c.Name, c.Fn)

		case "Update":
			processor := cb.Update()
			err = processor.Match(c.Match).Before(c.Before).After(c.After).Register(c.Name, c.Fn)

		case "Delete":
			processor := cb.Delete()
			err = processor.Match(c.Match).Before(c.Before).After(c.After).Register(c.Name, c.Fn)

		case "Query":
			processor := cb.Query()
			err = processor.Match(c.Match).Before(c.Before).After(c.After).Register(c.Name, c.Fn)

		default:
			return fmt.Errorf("invalid operation type: %s", c.Operation)
		}

		if err != nil {
			return fmt.Errorf("failed to register %s callback %s: %w", c.Operation, c.Name, err)
		}
	}
	return nil
}
