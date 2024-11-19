package main

import (
	"gorm.io/gen"

	"server/internal/dal/mysql/querygen"
	"server/internal/types/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../internal/dal/mysql/query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.ApplyBasic(
		model.User{},
	)
	g.ApplyInterface(func(querygen.UserQuery) {}, model.User{})
	g.Execute()
}
