package database

import (
	"log"
	"race_condition/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDB() {
	dsn := "host=	user=postgres password= dbname=bank port=5432 sslmode=disable"
	var err error

	log.Print("Veritabanına bağlanılıyor")
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("HATA: Veritabanı bağlantısı gerçekleştirilemedi.")
	}
	DBConn.AutoMigrate(&model.Account{})

}
