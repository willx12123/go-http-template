package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"server/internal/dal/db/query"
	"server/internal/pkg/config"
)

var (
	DB *gorm.DB
)

func Init() {
	conf := config.Default().Database
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s search_path=public host=%s port=%d TimeZone=Asia/Shanghai",
		conf.User, conf.Password, conf.DB, conf.Host, conf.Port,
	)
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	query.SetDefault(DB)
}
