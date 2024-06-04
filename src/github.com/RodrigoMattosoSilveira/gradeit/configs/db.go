package configs

import (
	"fmt"
	"log/slog"
	"os"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
  )

func DBInit () {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		panic(fmt.Sprintf("DBInit: Invalid DB_NAME environment%svariable", " "))
	}

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to SQLite %s database", dbName))
	}
	
	slog.Info(fmt.Sprintf("DBInit: connected successfully to %s", dbName))
	DB = db
}