package database

import (
    "log"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    var err error
    DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect database:", err)
    }

    // TODO: Auto-migrate models here
}
