package application

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/spo-iitk/ras-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func openConnection() {
	dsn := viper.GetString("DATABASE_URL") + "&search_path=application,public"

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Fatal("Failed to connect to application database: ", err)
		panic(err)
	}

	db = database

	err = db.AutoMigrate(&Proforma{}, &ApplicationQuestion{}, &ApplicationQuestionAnswer{},
		&ProformaEvent{}, &EventCoordinator{}, &EventStudent{}, &ApplicationResume{})
	if err != nil {
		logrus.Fatal("Failed to migrate application database: ", err)
		panic(err)
	}

	logrus.Info("Connected to application database")
}

func init() {
	openConnection()
	gCalendarConnect()
}

type EventType string

const (
	ApplicationSubmitted EventType = "Application"
	Recruited            EventType = "Recruited"
	PIOPPOACCEPTED       EventType = "PIO-PPO"
)
