package database

import (
	"log"

	"github.com/dhuharohim/golang-crud-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() error {
    var err error
    
    // Configure GORM logger
    newLogger := logger.New(
        log.New(log.Writer(), "\r\n", log.LstdFlags),
        logger.Config{
            LogLevel: logger.Info,
        },
    )

    // Open database connection
    DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{
        Logger: newLogger,
    })
    if err != nil {
        return err
    }

    // Auto migrate the schema
    err = DB.AutoMigrate(&models.Book{})
    if err != nil {
        return err
    }

    log.Println("Successfully connected to database")
    return nil
}
