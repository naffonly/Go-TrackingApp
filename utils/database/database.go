package database

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"trackingApp/config"
	companyModel "trackingApp/features/company/model"
	userModel "trackingApp/features/user/model"
	"trackingApp/utils/password"
)

func InitDB(config *config.Config) *gorm.DB {

	usersql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName)

	db, err := gorm.Open(mysql.Open(usersql), &gorm.Config{})
	if err != nil {
		logrus.Error("Model : cannot connect to database, ", err.Error())
		return nil
	}
	Migration(db)
	return db
}

func Migration(db *gorm.DB) {
	logrus.Info("Miggration DB")
	errs := db.AutoMigrate(&userModel.User{}, &companyModel.Company{})
	if errs != nil {
		logrus.Fatal(errs.Error())
		return
	}
	logrus.Info("SeederingSeedering Data")
	seederUser(db)
}

func seederUser(db *gorm.DB) {
	pass, _ := password.HashPassword("superadmin")
	id, _ := uuid.NewRandom()

	data := userModel.User{
		ID:       id.String(),
		Username: "superadmin",
		Name:     "SuperAdmin",
		Email:    "superadmin@gmail.com",
		Password: pass,
		Role:     1,
	}
	db.FirstOrCreate(&data)
}
