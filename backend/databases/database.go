package databases

import (
	"backend/config"
	"backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

// var DATABASE_URI string = config.GetEnvValue("DBUser") + ":" + config.GetEnvValue("DBPassword") + "@tcp(" + config.GetEnvValue("DBHost") + ":" + config.GetEnvValue("DBPort") + ")/" + config.GetEnvValue("DBName") + "?charset=utf8mb4&parseTime=True&loc=Local"
var DATABASE_URI string = "storage.db"

func Connect() error {
	var err error
	switch config.GetEnvValue("DBConnection") {
	case "mysql":
		Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})
	case "postgres":
		// Database, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{
		// 	SkipDefaultTransaction: true,
		// 	PrepareStmt:            true,
		// })
	case "sqlite":
		Database, err = gorm.Open(sqlite.Open(DATABASE_URI), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})
	}

	if err != nil {
		panic(err)
	}

	Database.Logger = logger.Default.LogMode(logger.Info)

	Database.AutoMigrate(&models.Todo{})

	return nil
}
