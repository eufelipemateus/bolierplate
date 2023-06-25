package database

import (
	"github.com/eufelipemateus/bolierplate-go/models"
	"gorm.io/gen"
)

func GenerateDB() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./database/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(DB)

	g.ApplyBasic(
		models.ExampleModel{},
	)

	g.Execute()

	DB.AutoMigrate(
		models.ExampleModel{},
	)
}
