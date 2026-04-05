package auth

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
	_ "github.com/spo-iitk/ras-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var logf *os.File

func openConnection() {
	dsn := viper.GetString("DATABASE_URL") + "&search_path=auth,public"

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		logrus.Fatal("Failed to connect to auth database: ", err)
		panic(err)
	}

	db = database

	err = db.AutoMigrate(&User{}, &OTP{}, &CompanySignUpRequest{})
	if err != nil {
		logrus.Fatal("Failed to migrate auth database: ", err)
		panic(err)
	}

	logrus.Info("Connected to auth database")
}

func init() {
	openConnection()
	logf = file()
	go cleanupOTP()
}
