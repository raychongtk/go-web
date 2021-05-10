package database

import (
	"github.com/google/wire"
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
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
