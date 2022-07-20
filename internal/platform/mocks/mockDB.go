package mocks

import (
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/utils"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"time"
)

var (
	logger, _   = zap.NewProduction()
	sugarLogger = logger.Sugar()
)

func MockDB() *gorm.DB {
	var config gorm.Config

	connectionString := "file::memory:?cache=shared"
	dialect := sqlite.Open(connectionString)
	config.Logger = gormLog.Default.LogMode(gormLog.Warn)
	connection, err := gorm.Open(dialect, &config)
	if err != nil {
		sugarLogger.Panic("[DATABASE][GET_DATABASE_CONNECTION] Error opening gorm connection", err)
	}

	_, err = connection.DB()
	if err != nil {
		sugarLogger.Panic("[DATABASE][GET_DATABASE_CONNECTION] Error configuring connection", err)
	}

	CreateRole(connection)
	CreateUser(connection)
	CreateDrivers(connection)
	CreateTrips(connection)

	return connection
}

func CreateRole(db *gorm.DB) {
	db.SkipDefaultTransaction = true
	tx := db.Exec("DROP TABLE IF EXISTS ROLES;")
	if tx.Error != nil {
		sugarLogger.Panic("Error Creating Table", tx.Error)
	}

	db.AutoMigrate(&entity.Role{})
	db.Create(&entity.Role{Name: "admin", Users: nil})
	db.Create(&entity.Role{Name: "driver", Users: nil})
}

func CreateUser(db *gorm.DB) {
	db.SkipDefaultTransaction = true
	tx := db.Exec("DROP TABLE IF EXISTS USERS;")
	if tx.Error != nil {
		sugarLogger.Panic("Error Creating Table", tx.Error)
	}
	db.AutoMigrate(&entity.User{})
	db.Create(&entity.User{Name: "postman", Email: "postman@gmail.com", Password: utils.HashAndSalt([]byte("monchi")), Token: "", RoleID: 1})
	db.Create(&entity.User{Name: "driver", Email: "driver@gmail.com", Password: utils.HashAndSalt([]byte("cuca")), Token: "", RoleID: 2})
	db.Create(&entity.User{Name: "driver junior", Email: "driverjunior@gmail.com", Password: utils.HashAndSalt([]byte("chango")), Token: "", RoleID: 2})
}

func CreateDrivers(db *gorm.DB) {
	db.SkipDefaultTransaction = true
	tx := db.Exec("DROP TABLE IF EXISTS DRIVERS;")
	if tx.Error != nil {
		sugarLogger.Panic("Error Creating Table", tx.Error)
	}
	db.AutoMigrate(&entity.Driver{})
	db.Create(&entity.Driver{DriverFile: 123456, Description: "test of driver", UserID: 2})
	db.Create(&entity.Driver{DriverFile: 654321, Description: "test2 of driver", UserID: 3})
}

func CreateTrips(db *gorm.DB) {
	db.SkipDefaultTransaction = true
	tx := db.Exec("DROP TABLE IF EXISTS TRIPS;")
	if tx.Error != nil {
		sugarLogger.Panic("Error Creating Table", tx.Error)
	}

	db.AutoMigrate(&entity.Trip{})
	db.Create(&entity.Trip{StartTime: time.Now(), EndTime: time.Time{}, DriverID: 1, Finished: false})
	db.Create(&entity.Trip{StartTime: time.Now(), EndTime: time.Now(), DriverID: 2, Finished: true})

}
