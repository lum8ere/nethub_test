package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	// dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database for codegen")
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "../../internal/storage/query",
		ModelPkgPath: "../../internal/storage/model",

		// WithUnitTest: true, // прикольная вещь, потом надо будет потыкать

		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,

		FieldNullable:     true,
		FieldCoverable:    true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db)

	softDeleteField := gen.FieldType("deleted_at", "gorm.DeletedAt")

	allModels := g.GenerateAllTable(softDeleteField)

	g.ApplyBasic(allModels...)

	g.Execute()
}
