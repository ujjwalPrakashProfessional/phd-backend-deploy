package student

import (
	"log"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	_ "github.com/spo-iitk/ras-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func openConnection() {
	dsn := viper.GetString("DATABASE_URL") + "&search_path=student,public"
	log.Println(dsn)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Fatal("Failed to connect to student database: ", err)
		panic(err)
	}

	db = database

	err = db.AutoMigrate(&Student{})
	if err != nil {
		logrus.Fatal("Failed to migrate student database: ", err)
		panic(err)
	}

	logrus.Info("Connected to student database")
}

func init() {
	openConnection()
}
