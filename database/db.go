package database

import (
	"fmt"
	"github.com/google/wire"
	"github.com/raychongtk/go-web/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	WireSet = wire.NewSet(ProvideDBConnection)
)

func ProvideDBConnection() gorm.DB {
	return *gormConnection()
}

func gormConnection() *gorm.DB {
	var config *Config
	util.Load("dev", ".", &config)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUsername, config.DBPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}
