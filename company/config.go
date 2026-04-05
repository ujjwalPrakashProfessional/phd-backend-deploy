package company

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func openConnection() {
	dsn := viper.GetString("DATABASE_URL") + "&search_path=company,public"

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Fatal("Failed to connect to company database: ", err)
		panic(err)
	}

	db = database

	err = db.AutoMigrate(&Company{}, &CompanyHR{})
	if err != nil {
		logrus.Fatal("Failed to migrate company database: ", err)
		panic(err)
	}

	logrus.Info("Connected to company database")
}

func init() {
	openConnection()
}
