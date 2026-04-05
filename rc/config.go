package rc // will be reanamed later

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func openConnection() {
	dsn := viper.GetString("DATABASE_URL") + "&search_path=rc,public"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Fatal("Failed to connect to cycle database: ", err)
		panic(err)
	}

	db = database

	err = db.AutoMigrate(&RecruitmentCycle{}, &RecruitmentCycleQuestion{},
		&RecruitmentCycleQuestionsAnswer{}, &CompanyRecruitmentCycle{}, &Notice{},
		&StudentRecruitmentCycle{}, &StudentRecruitmentCycleResume{})
	if err != nil {
		logrus.Fatal("Failed to migrate cycle database: ", err)
		panic(err)
	}

	logrus.Info("Connected to cycle database")
}

func init() {
	openConnection()
}
