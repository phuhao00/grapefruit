package main

import (
	"gorm.io/gen"
	"grapefruit/config"
	"grapefruit/internal/adapter/psql"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/adapter/psql/query",
		ModelPkgPath: "../../../internal/domain/po",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	_, err = config.GormGenerateLoadConfig(dir + "/config/dev.toml")
	if err != nil {
		return
	}
	psql.InitGormDB()
	// reuse your gorm db
	g.UseDB(psql.GetGormDB())

	g.ApplyBasic(
		g.GenerateModel("user"),
		g.GenerateModel("company"),
		g.GenerateModel("resume"),
		g.GenerateModel("job"),
	)
	// Generate the code
	g.Execute()
}
